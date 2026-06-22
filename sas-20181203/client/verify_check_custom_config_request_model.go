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
	// The ID of the check item.
	//
	// example:
	//
	// 76
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The input parameters for custom check item validation.
	CustomCheckConfig *VerifyCheckCustomConfigRequestCustomCheckConfig `json:"CustomCheckConfig,omitempty" xml:"CustomCheckConfig,omitempty" type:"Struct"`
	// The list of custom parameter configuration items for the check item.
	CustomConfigs []*VerifyCheckCustomConfigRequestCustomConfigs `json:"CustomConfigs,omitempty" xml:"CustomConfigs,omitempty" type:"Repeated"`
	// The repair parameters supported by the repair feature of the check item.
	RepairConfigs []*VerifyCheckCustomConfigRequestRepairConfigs `json:"RepairConfigs,omitempty" xml:"RepairConfigs,omitempty" type:"Repeated"`
	// The validation type for Threat Detection Service parameters. Valid values:
	//
	// - **REPAIR_CONFIG**: repair and custom parameter validation (default).
	//
	// - **CHECK_ITEM_CONFIG**: custom check item validation.
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
	// The definition rule of the custom check item.
	//
	// example:
	//
	// {"AssociatedData":{"ToDataList":[{"DataName":"ACS_ECS_Instance","PropertyPath":"InstanceId","FromPropertyPath":"InstanceId"}]},"MatchProperty":{"Operator":"AND","MatchProperties":[{"DataName":"ACS_ECS_Disk","PropertyPath":"DiskId","MatchOperator":"EQ","MatchPropertyValue":"testId"}]}}
	CheckRule *string `json:"CheckRule,omitempty" xml:"CheckRule,omitempty"`
	// The asset instance on which you want to test the rule.
	CloudAssetInstance *VerifyCheckCustomConfigRequestCustomCheckConfigCloudAssetInstance `json:"CloudAssetInstance,omitempty" xml:"CloudAssetInstance,omitempty" type:"Struct"`
	// The asset subtype of the cloud service.
	//
	// example:
	//
	// DISK
	InstanceSubType *string `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// The asset type of the cloud service. Valid values:
	//
	// - **ECS**: server
	//
	// - **SLB**: load balancing
	//
	// - **RDS**: ApsaraDB RDS database
	//
	// - **MONGODB**: ApsaraDB for MongoDB database
	//
	// - **KVSTORE**: ApsaraDB for Redis database
	//
	// - **ACR**: ACR
	//
	// - **CSK**: CSK
	//
	// - **VPC**: VPC
	//
	// - **ACTIONTRAIL**: ActionTrail
	//
	// - **CDN**: CDN
	//
	// - **CAS**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// - **RDC**: Apsara Devops
	//
	// - **RAM**: RAM
	//
	// - **DDOS**: distributed deny of service
	//
	// - **WAF**: WAF
	//
	// - **OSS**: access control
	//
	// - **POLARDB**: POLARDB
	//
	// - **POSTGRESQL**: PostgreSQL
	//
	// - **MSE**: MSE
	//
	// - **NAS**: NAS
	//
	// - **SDDP**: SDDP
	//
	// - **EIP**: EIP.
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The cloud asset vendor. Valid values:
	//
	// - **ALIYUN**: Alibaba Cloud
	//
	// - **Tencent**: Tencent Cloud
	//
	// - **HUAWEICLOUD**: Huawei Cloud
	//
	// - **Azure**: Microsoft Azure
	//
	// - **AWS**: Amazon Web Services (AWS).
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
	// The instance ID of the asset.
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
	// The name of the custom configuration item for the check item. The name is unique within the check item.
	//
	// example:
	//
	// IPList
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type of the custom configuration item for the check item. Set this parameter to DELETE only for deletion operations. You do not need to specify this parameter for creation or update operations.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The user-configured value string of the custom configuration item for the check item.
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
	// The ID of the repair flow that corresponds to the repair operation.
	//
	// example:
	//
	// 7fec0a3395b345c18f108ffc9fc0****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the repair parameter for the check item. The name is unique within the check item.
	//
	// example:
	//
	// IPLists
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type of the custom configuration item for the check item. Set this parameter to DELETE only for deletion operations. You do not need to specify this parameter for creation or update operations.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The user-configured value string of the repair configuration item for the check item.
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
