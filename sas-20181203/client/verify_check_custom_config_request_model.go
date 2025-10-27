// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckCustomConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *VerifyCheckCustomConfigRequest
	GetCheckId() *int64
	SetCustomCheckConfig(v *VerifyCheckCustomConfigRequestCustomCheckConfig) *VerifyCheckCustomConfigRequest
	GetCustomCheckConfig() *VerifyCheckCustomConfigRequestCustomCheckConfig
	SetCustomConfigs(v []*VerifyCheckCustomConfigRequestCustomConfigs) *VerifyCheckCustomConfigRequest
	GetCustomConfigs() []*VerifyCheckCustomConfigRequestCustomConfigs
	SetRepairConfigs(v []*VerifyCheckCustomConfigRequestRepairConfigs) *VerifyCheckCustomConfigRequest
	GetRepairConfigs() []*VerifyCheckCustomConfigRequestRepairConfigs
	SetType(v string) *VerifyCheckCustomConfigRequest
	GetType() *string
}

type VerifyCheckCustomConfigRequest struct {
	// Check item ID.
	//
	// example:
	//
	// 76
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// Custom check item to validate input parameters.
	CustomCheckConfig *VerifyCheckCustomConfigRequestCustomCheckConfig `json:"CustomCheckConfig,omitempty" xml:"CustomCheckConfig,omitempty" type:"Struct"`
	// List of custom configuration items for the check item.
	CustomConfigs []*VerifyCheckCustomConfigRequestCustomConfigs `json:"CustomConfigs,omitempty" xml:"CustomConfigs,omitempty" type:"Repeated"`
	// Repair parameters supported by the check item\\"s repair function.
	RepairConfigs []*VerifyCheckCustomConfigRequestRepairConfigs `json:"RepairConfigs,omitempty" xml:"RepairConfigs,omitempty" type:"Repeated"`
	// Situation Awareness parameter validation types:
	//
	// - **REPAIR_CONFIG**: Repair and custom parameter validation (default)
	//
	// - **CHECK_ITEM_CONFIG**: Custom check item validation
	//
	// example:
	//
	// REPAIR_CONFIG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s VerifyCheckCustomConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigRequest) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *VerifyCheckCustomConfigRequest) GetCustomCheckConfig() *VerifyCheckCustomConfigRequestCustomCheckConfig {
	return s.CustomCheckConfig
}

func (s *VerifyCheckCustomConfigRequest) GetCustomConfigs() []*VerifyCheckCustomConfigRequestCustomConfigs {
	return s.CustomConfigs
}

func (s *VerifyCheckCustomConfigRequest) GetRepairConfigs() []*VerifyCheckCustomConfigRequestRepairConfigs {
	return s.RepairConfigs
}

func (s *VerifyCheckCustomConfigRequest) GetType() *string {
	return s.Type
}

func (s *VerifyCheckCustomConfigRequest) SetCheckId(v int64) *VerifyCheckCustomConfigRequest {
	s.CheckId = &v
	return s
}

func (s *VerifyCheckCustomConfigRequest) SetCustomCheckConfig(v *VerifyCheckCustomConfigRequestCustomCheckConfig) *VerifyCheckCustomConfigRequest {
	s.CustomCheckConfig = v
	return s
}

func (s *VerifyCheckCustomConfigRequest) SetCustomConfigs(v []*VerifyCheckCustomConfigRequestCustomConfigs) *VerifyCheckCustomConfigRequest {
	s.CustomConfigs = v
	return s
}

func (s *VerifyCheckCustomConfigRequest) SetRepairConfigs(v []*VerifyCheckCustomConfigRequestRepairConfigs) *VerifyCheckCustomConfigRequest {
	s.RepairConfigs = v
	return s
}

func (s *VerifyCheckCustomConfigRequest) SetType(v string) *VerifyCheckCustomConfigRequest {
	s.Type = &v
	return s
}

func (s *VerifyCheckCustomConfigRequest) Validate() error {
	if s.CustomCheckConfig != nil {
		if err := s.CustomCheckConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CustomConfigs != nil {
		for _, item := range s.CustomConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RepairConfigs != nil {
		for _, item := range s.RepairConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VerifyCheckCustomConfigRequestCustomCheckConfig struct {
	// Define rules for custom inspection items.
	//
	// example:
	//
	// {"AssociatedData":{"ToDataList":[{"DataName":"ACS_ECS_Instance","PropertyPath":"InstanceId","FromPropertyPath":"InstanceId"}]},"MatchProperty":{"Operator":"AND","MatchProperties":[{"DataName":"ACS_ECS_Disk","PropertyPath":"DiskId","MatchOperator":"EQ","MatchPropertyValue":"testId"}]}}
	CheckRule *string `json:"CheckRule,omitempty" xml:"CheckRule,omitempty"`
	// Asset instance that requires testing rules
	CloudAssetInstance *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance `json:"CloudAssetInstance,omitempty" xml:"CloudAssetInstance,omitempty" type:"Struct"`
	// Asset subtype of the cloud product
	//
	// example:
	//
	// DISK
	InstanceSubType *string `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// Asset types of cloud products. Values:
	//
	// - **ECS**: Elastic Compute Service
	//
	// - **SLB**: Server Load Balancer
	//
	// - **RDS**: Relational Database Service
	//
	// - **MONGODB**: MongoDB Database
	//
	// - **KVSTORE**: Redis Database
	//
	// - **ACR**: Container Registry
	//
	// - **CSK**: CSK
	//
	// - **VPC**: Virtual Private Cloud
	//
	// - **ACTIONTRAIL**: Action Trail
	//
	// - **CDN**: Content Delivery Network
	//
	// - **CAS**: Digital Certificate Management Service [formerly SSL Certificates]
	//
	// - **RDC**: DevOps
	//
	// - **RAM**: Resource Access Management
	//
	// - **DDOS**: Distributed Denial of Service
	//
	// - **WAF**: Web Application Firewall
	//
	// - **OSS**: Object Storage Service
	//
	// - **POLARDB**: POLARDB
	//
	// - **POSTGRESQL**: PostgreSQL
	//
	// - **MSE**: MSE
	//
	// - **NAS**: Network Attached Storage
	//
	// - **SDDP**: Sensitive Data Discovery and Protection
	//
	// - **EIP**: Elastic IP
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Cloud asset vendor. Values:
	//
	// - **ALIYUN**: Alibaba Cloud
	//
	// - **Tencent**: Tencent Cloud
	//
	// - **HUAWEICLOUD**: Huawei Cloud
	//
	// - **Azure**: Microsoft
	//
	// - **AWS**: Amazon Web Services (AWS)
	//
	// example:
	//
	// ALIYUN
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s VerifyCheckCustomConfigRequestCustomCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigRequestCustomCheckConfig) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) GetCheckRule() *string {
	return s.CheckRule
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) GetCloudAssetInstance() *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance {
	return s.CloudAssetInstance
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) GetInstanceSubType() *string {
	return s.InstanceSubType
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) GetInstanceType() *string {
	return s.InstanceType
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) GetVendor() *string {
	return s.Vendor
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) SetCheckRule(v string) *VerifyCheckCustomConfigRequestCustomCheckConfig {
	s.CheckRule = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) SetCloudAssetInstance(v *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance) *VerifyCheckCustomConfigRequestCustomCheckConfig {
	s.CloudAssetInstance = v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) SetInstanceSubType(v string) *VerifyCheckCustomConfigRequestCustomCheckConfig {
	s.InstanceSubType = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) SetInstanceType(v string) *VerifyCheckCustomConfigRequestCustomCheckConfig {
	s.InstanceType = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) SetVendor(v string) *VerifyCheckCustomConfigRequestCustomCheckConfig {
	s.Vendor = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfig) Validate() error {
	if s.CloudAssetInstance != nil {
		if err := s.CloudAssetInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance struct {
	// Instance ID of the asset.
	//
	// example:
	//
	// i-0jl4mjgl261cfrz5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance) SetInstanceId(v string) *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance {
	s.InstanceId = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance) SetRegionId(v string) *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance {
	s.RegionId = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance) Validate() error {
	return dara.Validate(s)
}

type VerifyCheckCustomConfigRequestCustomConfigs struct {
	// Name of the custom configuration item for the check item, unique within the same check item.
	//
	// example:
	//
	// IPList
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Operation type for the custom configuration item of the check item. Only pass DELETE when deleting; no need to pass for creation or update.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// User-configured value string for the custom configuration item of the check item.
	//
	// example:
	//
	// 10.12.4.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s VerifyCheckCustomConfigRequestCustomConfigs) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigRequestCustomConfigs) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigRequestCustomConfigs) GetName() *string {
	return s.Name
}

func (s *VerifyCheckCustomConfigRequestCustomConfigs) GetOperation() *string {
	return s.Operation
}

func (s *VerifyCheckCustomConfigRequestCustomConfigs) GetValue() *string {
	return s.Value
}

func (s *VerifyCheckCustomConfigRequestCustomConfigs) SetName(v string) *VerifyCheckCustomConfigRequestCustomConfigs {
	s.Name = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomConfigs) SetOperation(v string) *VerifyCheckCustomConfigRequestCustomConfigs {
	s.Operation = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomConfigs) SetValue(v string) *VerifyCheckCustomConfigRequestCustomConfigs {
	s.Value = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestCustomConfigs) Validate() error {
	return dara.Validate(s)
}

type VerifyCheckCustomConfigRequestRepairConfigs struct {
	// ID of the repair process during the repair.
	//
	// example:
	//
	// 7fec0a3395b345c18f108ffc9fc0****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// Name of the repair parameter for the check item, unique within the same check item.
	//
	// example:
	//
	// IPLists
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Operation type for the custom configuration item of the check item. Only pass DELETE when deleting; no need to pass for creation or update.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// User-configured value string for the repair parameter of the check item.
	//
	// example:
	//
	// 172.26.49.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s VerifyCheckCustomConfigRequestRepairConfigs) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigRequestRepairConfigs) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) GetFlowId() *string {
	return s.FlowId
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) GetName() *string {
	return s.Name
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) GetOperation() *string {
	return s.Operation
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) GetValue() *string {
	return s.Value
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) SetFlowId(v string) *VerifyCheckCustomConfigRequestRepairConfigs {
	s.FlowId = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) SetName(v string) *VerifyCheckCustomConfigRequestRepairConfigs {
	s.Name = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) SetOperation(v string) *VerifyCheckCustomConfigRequestRepairConfigs {
	s.Operation = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) SetValue(v string) *VerifyCheckCustomConfigRequestRepairConfigs {
	s.Value = &v
	return s
}

func (s *VerifyCheckCustomConfigRequestRepairConfigs) Validate() error {
	return dara.Validate(s)
}
