package cloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/signers"
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

func TestCreateDescribePriceACSRequest(t *testing.T) {
	node := &SlimK8sNode{
		InstanceType:       "ecs.g6.large",
		RegionID:           "cn-hangzhou",
		PriceUnit:          "Hour",
		MemorySizeInKiB:    "16KiB",
		IsIoOptimized:      true,
		OSType:             "Linux",
		ProviderID:         "Ali-XXX-node-01",
		InstanceTypeFamily: "g6",
	}

	disk := &SlimK8sDisk{
		DiskType:         "data",
		RegionID:         "cn-hangzhou",
		PriceUnit:        "Hour",
		SizeInGiB:        "20",
		DiskCategory:     "diskCategory",
		PerformanceLevel: "cloud_essd",
		ProviderID:       "d-Ali-XXX-01",
		StorageClass:     "testStorageClass",
	}

	cases := []struct {
		name          string
		testStruct    interface{}
		expectedError error
	}{
		{
			name:          "test CreateDescribePriceACSRequest with SlimK8sNode struct Object",
			testStruct:    node,
			expectedError: nil,
		},
		{
			name:          "test CreateDescribePriceACSRequest with SlimK8sDisk struct Object",
			testStruct:    disk,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := createDescribePriceACSRequest(c.testStruct)
			if err != nil && c.expectedError == nil {
				t.Fatalf("Case name %s: Error converting to Alibaba cloud request", c.name)
			}
		})
	}
}

func TestProcessDescribePriceAndCreateAlibabaPricing(t *testing.T) {
	// Skipping this test case since it exposes secret but a good test case to verify when
	// supporting a new family of instances, steps to perform are
	// STEP 1: Comment the t.Skip() line and then replace XXX_KEY_ID with the alibaba key id of your account and XXX_SECRET_ID with alibaba cloud secret of your account.
	// STEP 2: Once you verify describePrice is working and no change needed in processDescribePriceAndCreateAlibabaPricing, you can go ahead and revert the step 1 changes.

	// This test case was use to test all general puprose instances

	t.Skip()

	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", "XXX_KEY_ID", "XXX_SECRET_ID")
	if err != nil {
		t.Errorf("Error connecting to the Alibaba cloud")
	}
	aak := credentials.NewAccessKeyCredential("XXX_KEY_ID", "XXX_SECRET_ID")
	signer := signers.NewAccessKeySigner(aak)

	cases := []struct {
		name          string
		teststruct    interface{}
		expectedError error
	}{
		{
			name: "test Enhanced General Purpose Type g6e instance family",
			teststruct: &SlimK8sNode{
				InstanceType:       "ecs.g6e.xlarge",
				RegionID:           "cn-hangzhou",
				PriceUnit:          "Hour",
				MemorySizeInKiB:    "16777216KiB",
				IsIoOptimized:      true,
				OSType:             "Linux",
				ProviderID:         "cn-hangzhou.i-test-01",
				InstanceTypeFamily: "g6e",
			},
			expectedError: nil,
		},
		{
			name: "test General Purpose Type g6 instance family",
			teststruct: &SlimK8sNode{
				InstanceType:       "ecs.g6.3xlarge",
				RegionID:           "cn-hangzhou",
				PriceUnit:          "Hour",
				MemorySizeInKiB:    "50331648KiB",
				IsIoOptimized:      true,
				OSType:             "Linux",
				ProviderID:         "cn-hangzhou.i-test-02",
				InstanceTypeFamily: "g6",
			},
			expectedError: nil,
		},
		{
			name: "test General Purpose Type g5 instance family",
			teststruct: &SlimK8sNode{
				InstanceType:       "ecs.g5.2xlarge",
				RegionID:           "cn-hangzhou",
				PriceUnit:          "Hour",
				MemorySizeInKiB:    "33554432KiB",
				IsIoOptimized:      true,
				OSType:             "Linux",
				ProviderID:         "cn-hangzhou.i-test-03",
				InstanceTypeFamily: "g5",
			},
			expectedError: nil,
		},
		{
			name: "test General Purpose Type sn2 instance family",
			teststruct: &SlimK8sNode{
				InstanceType:       "ecs.sn2.large",
				RegionID:           "cn-hangzhou",
				PriceUnit:          "Hour",
				MemorySizeInKiB:    "16777216KiB",
				IsIoOptimized:      true,
				OSType:             "Linux",
				ProviderID:         "cn-hangzhou.i-test-04",
				InstanceTypeFamily: "sn2",
			},
			expectedError: nil,
		},
		{
			name: "test General Purpose Type with Enhanced Network Performance sn2ne instance family",
			teststruct: &SlimK8sNode{
				InstanceType:       "ecs.sn2ne.2xlarge",
				RegionID:           "cn-hangzhou",
				PriceUnit:          "Hour",
				MemorySizeInKiB:    "33554432KiB",
				IsIoOptimized:      true,
				OSType:             "Linux",
				ProviderID:         "cn-hangzhou.i-test-05",
				InstanceTypeFamily: "sn2ne",
			},
			expectedError: nil,
		},
		{
			name:          "test for a nil information",
			teststruct:    nil,
			expectedError: fmt.Errorf("unsupported ECS pricing component at this time"),
		},
		{
			name: "test Cloud Disk with Category cloud representing basic disk",
			teststruct: &SlimK8sDisk{
				DiskType:     "data",
				RegionID:     "cn-hangzhou",
				PriceUnit:    "Hour",
				SizeInGiB:    "20",
				DiskCategory: "cloud",
				ProviderID:   "d-Ali-cloud-XXX-01",
				StorageClass: "temp",
			},
			expectedError: nil,
		},
		{
			name: "test Cloud Disk with Category cloud_efficiency representing ultra disk",
			teststruct: &SlimK8sDisk{
				DiskType:     "data",
				RegionID:     "cn-hangzhou",
				PriceUnit:    "Hour",
				SizeInGiB:    "40",
				DiskCategory: "cloud_efficiency",
				ProviderID:   "d-Ali-cloud-XXX-02",
				StorageClass: "temp",
			},
			expectedError: nil,
		},
		{
			name: "test Cloud Disk with Category cloud_ssd representing standard SSD",
			teststruct: &SlimK8sDisk{
				DiskType:     "data",
				RegionID:     "cn-hangzhou",
				PriceUnit:    "Hour",
				SizeInGiB:    "40",
				DiskCategory: "cloud_efficiency",
				ProviderID:   "d-Ali-cloud-XXX-02",
				StorageClass: "temp",
			},
			expectedError: nil,
		},
		{
			name: "test Cloud Disk with Category cloud_essd representing Enhanced SSD with PL2 performance level",
			teststruct: &SlimK8sDisk{
				DiskType:         "data",
				RegionID:         "cn-hangzhou",
				PriceUnit:        "Hour",
				SizeInGiB:        "80",
				DiskCategory:     "cloud_ssd",
				PerformanceLevel: "PL2",
				ProviderID:       "d-Ali-cloud-XXX-04",
				StorageClass:     "temp",
			},
			expectedError: nil,
		},
	}
	custom := &CustomPricing{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pricingObj, err := processDescribePriceAndCreateAlibabaPricing(client, c.teststruct, signer, custom)
			if err != nil && c.expectedError == nil {
				t.Fatalf("Case name %s: got an error %s", c.name, err)
			}
			if c.teststruct != nil {
				if pricingObj == nil {
					t.Fatalf("Case name %s: got a nil pricing object", c.name)
				}
				t.Logf("Case name %s: Pricing Information gathered for instanceType is %v", c.name, pricingObj.PricingTerms.PricingDetails.TradePrice)
			}
		})
	}
}

func TestGetInstanceFamilyFromType(t *testing.T) {
	cases := []struct {
		name                   string
		instanceType           string
		expectedInstanceFamily string
	}{
		{
			name:                   "test if ecs.[instance-family].[different-type] work",
			instanceType:           "ecs.sn2ne.2xlarge",
			expectedInstanceFamily: "sn2ne",
		},
		{
			name:                   "test if random word gives you ALIBABA_UNKNOWN_INSTANCE_FAMILY_TYPE value ",
			instanceType:           "random.value",
			expectedInstanceFamily: ALIBABA_UNKNOWN_INSTANCE_FAMILY_TYPE,
		},
		{
			name:                   "test if random instance family gives you ALIBABA_NOT_SUPPORTED_INSTANCE_FAMILY_TYPE value ",
			instanceType:           "ecs.g7e.2xlarge",
			expectedInstanceFamily: ALIBABA_NOT_SUPPORTED_INSTANCE_FAMILY_TYPE,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			returnValue := getInstanceFamilyFromType(c.instanceType)
			if returnValue != c.expectedInstanceFamily {
				t.Fatalf("Case name %s: expected instance family of type %s but got %s", c.name, c.expectedInstanceFamily, returnValue)
			}
		})
	}
}

func TestDetermineKeyForPricing(t *testing.T) {
	type randomK8sStruct struct {
		name string
	}
	cases := []struct {
		name          string
		testVar       interface{}
		expectedKey   string
		expectedError error
	}{
		{
			name: "test when all RegionID, InstanceType, OSType & ALIBABA_OPTIMIZE_KEYWORD words are used in Node key",
			testVar: &SlimK8sNode{
				InstanceType:       "ecs.sn2.large",
				RegionID:           "cn-hangzhou",
				PriceUnit:          "Hour",
				MemorySizeInKiB:    "16777216KiB",
				IsIoOptimized:      true,
				OSType:             "linux",
				ProviderID:         "cn-hangzhou.i-test-04",
				InstanceTypeFamily: "sn2",
			},
			expectedKey:   "cn-hangzhou::ecs.sn2.large::linux::optimize",
			expectedError: nil,
		},
		{
			name: "test missing InstanceType to create Node key",
			testVar: &SlimK8sNode{
				RegionID:        "cn-hangzhou",
				PriceUnit:       "Hour",
				MemorySizeInKiB: "16777216KiB",
				IsIoOptimized:   true,
				OSType:          "linux",
				ProviderID:      "cn-hangzhou.i-test-04",
			},
			expectedKey:   "cn-hangzhou::linux::optimize",
			expectedError: nil,
		},
		{
			name: "test random k8s struct should return unsupported error",
			testVar: &randomK8sStruct{
				name: "test struct",
			},
			expectedKey:   "",
			expectedError: fmt.Errorf("unsupported ECS type randomK8sStruct for DescribePrice at this time"),
		},
		{
			name:          "test for nil check",
			testVar:       nil,
			expectedKey:   "",
			expectedError: fmt.Errorf("unsupported ECS type randomK8sStruct for DescribePrice at this time"),
		},
		{
			name: "test when all RegionID, InstanceType, OSType & ALIBABA_OPTIMIZE_KEYWORD words are used to key",
			testVar: &SlimK8sDisk{
				DiskType:     "data",
				RegionID:     "cn-hangzhou",
				PriceUnit:    "Hour",
				SizeInGiB:    "40",
				DiskCategory: "cloud_efficiency",
				ProviderID:   "d-Ali-cloud-XXX-02",
				StorageClass: "temp",
			},
			expectedKey:   "cn-hangzhou::data::cloud_efficiency::40",
			expectedError: nil,
		},
		{
			name: "test missing InstanceType to create key",
			testVar: &SlimK8sDisk{
				DiskType:         "data",
				RegionID:         "cn-hangzhou",
				PriceUnit:        "Hour",
				SizeInGiB:        "80",
				DiskCategory:     "cloud_ssd",
				PerformanceLevel: "PL2",
				ProviderID:       "d-Ali-cloud-XXX-04",
				StorageClass:     "temp",
			},
			expectedKey:   "cn-hangzhou::data::cloud_ssd::PL2::80",
			expectedError: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			returnString, returnErr := determineKeyForPricing(c.testVar)
			if c.expectedError == nil && returnErr != nil {
				t.Fatalf("Case name %s: expected error was nil but recieved error %v", c.name, returnErr)
			}
			if returnString != c.expectedKey {
				t.Fatalf("Case name %s: determineKeyForPricing recieved %s but expected %s", c.name, returnString, c.expectedKey)
			}
		})
	}
}

func TestGenerateSlimK8sNodeFromV1Node(t *testing.T) {
	testv1Node := &v1.Node{}
	testv1Node.Labels = make(map[string]string)
	testv1Node.Labels["topology.kubernetes.io/region"] = "us-east-1"
	testv1Node.Labels["beta.kubernetes.io/os"] = "linux"
	testv1Node.Labels["node.kubernetes.io/instance-type"] = "ecs.sn2ne.2xlarge"
	testv1Node.Status.Capacity = v1.ResourceList{
		v1.ResourceMemory: *resource.NewQuantity(16, resource.BinarySI),
	}
	cases := []struct {
		name             string
		testNode         *v1.Node
		expectedSlimNode *SlimK8sNode
	}{
		{
			name:     "test a generic *v1.Node to *SlimK8sNode Conversion",
			testNode: testv1Node,
			expectedSlimNode: &SlimK8sNode{
				InstanceType:       "ecs.sn2ne.2xlarge",
				RegionID:           "us-east-1",
				PriceUnit:          ALIBABA_HOUR_PRICE_UNIT,
				MemorySizeInKiB:    "16",
				IsIoOptimized:      true,
				OSType:             "linux",
				InstanceTypeFamily: "sn2ne",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			returnSlimK8sNode := generateSlimK8sNodeFromV1Node(c.testNode)
			if returnSlimK8sNode.InstanceType != c.expectedSlimNode.InstanceType {
				t.Fatalf("unexpected conversion in function generateSlimK8sNodeFromV1Node expected InstanceType: %s , recieved InstanceType: %s", c.expectedSlimNode.InstanceType, returnSlimK8sNode.InstanceType)
			}
			if returnSlimK8sNode.RegionID != c.expectedSlimNode.RegionID {
				t.Fatalf("unexpected conversion in function generateSlimK8sNodeFromV1Node expected RegionID: %s , recieved RegionID: %s", c.expectedSlimNode.RegionID, returnSlimK8sNode.RegionID)
			}
			if returnSlimK8sNode.PriceUnit != c.expectedSlimNode.PriceUnit {
				t.Fatalf("unexpected conversion in function generateSlimK8sNodeFromV1Node expected PriceUnit: %s , recieved PriceUnit: %s", c.expectedSlimNode.PriceUnit, returnSlimK8sNode.PriceUnit)
			}
			if returnSlimK8sNode.MemorySizeInKiB != c.expectedSlimNode.MemorySizeInKiB {
				t.Fatalf("unexpected conversion in function generateSlimK8sNodeFromV1Node expected MemorySizeInKiB: %s , recieved MemorySizeInKiB: %s", c.expectedSlimNode.MemorySizeInKiB, returnSlimK8sNode.MemorySizeInKiB)
			}
			if returnSlimK8sNode.OSType != c.expectedSlimNode.OSType {
				t.Fatalf("unexpected conversion in function generateSlimK8sNodeFromV1Node expected OSType: %s , recieved OSType: %s", c.expectedSlimNode.OSType, returnSlimK8sNode.OSType)
			}
			if returnSlimK8sNode.InstanceTypeFamily != c.expectedSlimNode.InstanceTypeFamily {
				t.Fatalf("unexpected conversion in function generateSlimK8sNodeFromV1Node expected InstanceTypeFamily: %s , recieved InstanceTypeFamily: %s", c.expectedSlimNode.InstanceTypeFamily, returnSlimK8sNode.InstanceTypeFamily)
			}
		})
	}
}

func TestGenerateSlimK8sDiskFromV1PV(t *testing.T) {
	testv1PV := &v1.PersistentVolume{}
	testv1PV.Spec.Capacity = v1.ResourceList{
		v1.ResourceStorage: *resource.NewQuantity(16*1024*1024*1024, resource.BinarySI),
	}
	testv1PV.Spec.CSI = &v1.CSIPersistentVolumeSource{}
	testv1PV.Spec.CSI.VolumeHandle = "testPV"
	testv1PV.Spec.CSI.VolumeAttributes = map[string]string{
		"performanceLevel": "PL2",
		"type":             "cloud_essd",
	}
	testv1PV.Spec.CSI.VolumeHandle = "testPV"
	testv1PV.Spec.StorageClassName = "testStorageClass"
	cases := []struct {
		name             string
		testPV           *v1.PersistentVolume
		expectedSlimDisk *SlimK8sDisk
		inpRegionID      string
	}{
		{
			name:   "test a generic *v1.Node to *SlimK8sNode Conversion",
			testPV: testv1PV,
			expectedSlimDisk: &SlimK8sDisk{
				DiskType:         ALIBABA_DATA_DISK_CATEGORY,
				RegionID:         "us-east-1",
				PriceUnit:        ALIBABA_HOUR_PRICE_UNIT,
				SizeInGiB:        "16",
				DiskCategory:     "cloud_essd",
				PerformanceLevel: "PL2",
				ProviderID:       "testPV",
				StorageClass:     "testStorageClass",
			},
			inpRegionID: "us-east-1",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			returnSlimK8sDisk := generateSlimK8sDiskFromV1PV(c.testPV, c.inpRegionID)
			if returnSlimK8sDisk.DiskType != c.expectedSlimDisk.DiskType {
				t.Fatalf("unexpected conversion in function generateSlimK8sDiskFromV1PV expected DiskType: %s , recieved DiskType: %s", c.expectedSlimDisk.DiskType, returnSlimK8sDisk.DiskType)
			}
			if returnSlimK8sDisk.RegionID != c.expectedSlimDisk.RegionID {
				t.Fatalf("unexpected conversion in function generateSlimK8sDiskFromV1PV expected RegionID: %s , recieved RegionID Type: %s", c.expectedSlimDisk.RegionID, returnSlimK8sDisk.RegionID)
			}
			if returnSlimK8sDisk.PriceUnit != c.expectedSlimDisk.PriceUnit {
				t.Fatalf("unexpected conversion in function generateSlimK8sDiskFromV1PV expected PriceUnit: %s , recieved PriceUnit Type: %s", c.expectedSlimDisk.PriceUnit, returnSlimK8sDisk.PriceUnit)
			}
			if returnSlimK8sDisk.SizeInGiB != c.expectedSlimDisk.SizeInGiB {
				t.Fatalf("unexpected conversion in function generateSlimK8sDiskFromV1PV expected SizeInGiB: %s , recieved SizeInGiB Type: %s", c.expectedSlimDisk.SizeInGiB, returnSlimK8sDisk.SizeInGiB)
			}
			if returnSlimK8sDisk.DiskCategory != c.expectedSlimDisk.DiskCategory {
				t.Fatalf("unexpected conversion in function generateSlimK8sDiskFromV1PV expected DiskCategory: %s , recieved DiskCategory Type: %s", c.expectedSlimDisk.DiskCategory, returnSlimK8sDisk.DiskCategory)
			}
			if returnSlimK8sDisk.PerformanceLevel != c.expectedSlimDisk.PerformanceLevel {
				t.Fatalf("unexpected conversion in function generateSlimK8sDiskFromV1PV expected PerformanceLevel: %s , recieved PerformanceLevel Type: %s", c.expectedSlimDisk.PerformanceLevel, returnSlimK8sDisk.PerformanceLevel)
			}
			if returnSlimK8sDisk.ProviderID != c.expectedSlimDisk.ProviderID {
				t.Fatalf("unexpected conversion in function generateSlimK8sDiskFromV1PV expected ProviderID: %s , recieved ProviderID Type: %s", c.expectedSlimDisk.ProviderID, returnSlimK8sDisk.ProviderID)
			}
			if returnSlimK8sDisk.StorageClass != c.expectedSlimDisk.StorageClass {
				t.Fatalf("unexpected conversion in function generateSlimK8sDiskFromV1PV expected StorageClass: %s , recieved StorageClass Type: %s", c.expectedSlimDisk.StorageClass, returnSlimK8sDisk.StorageClass)
			}
		})
	}
}

func TestGetNumericalValueFromResourceQuantity(t *testing.T) {
	cases := []struct {
		name                 string
		inputResourceQuanity string
		expectedValue        string
	}{
		{
			name:                 "positive scenario: when inputResourceQuantity is 10Gi",
			inputResourceQuanity: "10Gi",
			expectedValue:        "10",
		},
		{
			name:                 "negative scenario: when inputResourceQuantity is Gi",
			inputResourceQuanity: "Gi",
			expectedValue:        ALIBABA_DEFAULT_DATADISK_SIZE,
		},
		{
			name:                 "negative scenario: when inputResourceQuantity is 10",
			inputResourceQuanity: "10",
			expectedValue:        ALIBABA_DEFAULT_DATADISK_SIZE,
		},
		{
			name:                 "negative scenario: when inputResourceQuantity is empty string",
			inputResourceQuanity: "",
			expectedValue:        ALIBABA_DEFAULT_DATADISK_SIZE,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			returnValue := getNumericalValueFromResourceQuantity(c.inputResourceQuanity)
			if c.expectedValue != returnValue {
				t.Fatalf("Case name %s: getNumericalValueFromResourceQuantity recieved %s but expected %s", c.name, returnValue, c.expectedValue)
			}
		})
	}
}

func TestDeterminePVRegion(t *testing.T) {
	genericNodeAffinityTestStruct := v1.NodeSelectorTerm{
		MatchExpressions: []v1.NodeSelectorRequirement{
			{
				Key:      "topology.diskplugin.csi.alibabacloud.com/zone",
				Operator: v1.NodeSelectorOpIn,
				Values:   []string{"us-east-1a"},
			},
		},
		MatchFields: []v1.NodeSelectorRequirement{},
	}

	// testPV1 contains the Label with region information as well as node affinity in spec
	testPV1 := &v1.PersistentVolume{}
	testPV1.Name = "testPV1"
	testPV1.Labels = make(map[string]string)
	testPV1.Labels[ALIBABA_DISK_TOPOLOGY_REGION_LABEL] = "us-east-1"
	testPV1.Spec.NodeAffinity = &v1.VolumeNodeAffinity{
		Required: &v1.NodeSelector{
			NodeSelectorTerms: []v1.NodeSelectorTerm{genericNodeAffinityTestStruct},
		},
	}

	// testPV2 contains the only zone label
	testPV2 := &v1.PersistentVolume{}
	testPV2.Name = "testPV2"
	testPV2.Labels = make(map[string]string)
	testPV2.Labels[ALIBABA_DISK_TOPOLOGY_ZONE_LABEL] = "us-east-1a"

	// testPV3 contains only node affinity in spec
	testPV3 := &v1.PersistentVolume{}
	testPV3.Name = "testPV3"
	testPV3.Spec.NodeAffinity = &v1.VolumeNodeAffinity{
		Required: &v1.NodeSelector{
			NodeSelectorTerms: []v1.NodeSelectorTerm{genericNodeAffinityTestStruct},
		},
	}

	// testPV4 contains no label/annotation or any node affinity
	testPV4 := &v1.PersistentVolume{}
	testPV4.Name = "testPV4"

	cases := []struct {
		name           string
		inputPV        *v1.PersistentVolume
		expectedRegion string
	}{
		{
			name:           "When Region label topology.diskplugin.csi.alibabacloud.com/region is present along with node affinity details",
			inputPV:        testPV1,
			expectedRegion: "us-east-1",
		},
		{
			name:           "When zone label topology.diskplugin.csi.alibabacloud.com/zone is present function has to determine region",
			inputPV:        testPV2,
			expectedRegion: "us-east-1",
		},
		{
			name:           "When only node affinity detail is present function has to determine the region",
			inputPV:        testPV3,
			expectedRegion: "us-east-1",
		},
		{
			name:           "When no region/zone information is present function returns empty to default to cluster region",
			inputPV:        testPV4,
			expectedRegion: "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			returnRegion := determinePVRegion(c.inputPV)
			if c.expectedRegion != returnRegion {
				t.Fatalf("Case name %s: determinePVRegion recieved region :%s but expected region: %s", c.name, returnRegion, c.expectedRegion)
			}
		})
	}

}
