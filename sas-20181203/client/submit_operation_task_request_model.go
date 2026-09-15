// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *SubmitOperationTaskRequest
	GetCheckId() *int64
	SetDimensionType(v string) *SubmitOperationTaskRequest
	GetDimensionType() *string
	SetOperationTaskInstances(v []*SubmitOperationTaskRequestOperationTaskInstances) *SubmitOperationTaskRequest
	GetOperationTaskInstances() []*SubmitOperationTaskRequestOperationTaskInstances
	SetRelationKey(v string) *SubmitOperationTaskRequest
	GetRelationKey() *string
	SetRepairTempParam(v []*SubmitOperationTaskRequestRepairTempParam) *SubmitOperationTaskRequest
	GetRepairTempParam() []*SubmitOperationTaskRequestRepairTempParam
	SetType(v string) *SubmitOperationTaskRequest
	GetType() *string
}

type SubmitOperationTaskRequest struct {
	// The ID of the check item.
	//
	// > Call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the check item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 132
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The task dimension for the submitted operation task. Valid values:
	//
	// - INSTANCE: instance dimension
	//
	// - CHECK_ID: check item dimension
	//
	// example:
	//
	// CHECK_ID
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	// The asset information required to submit instance tasks.
	OperationTaskInstances []*SubmitOperationTaskRequestOperationTaskInstances `json:"OperationTaskInstances,omitempty" xml:"OperationTaskInstances,omitempty" type:"Repeated"`
	// The relation key associated with cross-page selection when submitting the operation.
	//
	// > Call the [CreateAssetSelectionConfig](~~CreateAssetSelectionConfig~~) operation and use the BusinessType field to obtain the relation key.
	//
	// example:
	//
	// CSPM_OPERATION_RELATION_KEY_173***
	RelationKey *string `json:"RelationKey,omitempty" xml:"RelationKey,omitempty"`
	// The temporary parameters required for the remediation task.
	RepairTempParam []*SubmitOperationTaskRequestRepairTempParam `json:"RepairTempParam,omitempty" xml:"RepairTempParam,omitempty" type:"Repeated"`
	// The task type for the submitted task. Valid values:
	//
	// - REPAIR: remediation task
	//
	// - ROLLBACK: rollback task
	//
	// This parameter is required.
	//
	// example:
	//
	// REPAIR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitOperationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitOperationTaskRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *SubmitOperationTaskRequest) GetDimensionType() *string {
	return s.DimensionType
}

func (s *SubmitOperationTaskRequest) GetOperationTaskInstances() []*SubmitOperationTaskRequestOperationTaskInstances {
	return s.OperationTaskInstances
}

func (s *SubmitOperationTaskRequest) GetRelationKey() *string {
	return s.RelationKey
}

func (s *SubmitOperationTaskRequest) GetRepairTempParam() []*SubmitOperationTaskRequestRepairTempParam {
	return s.RepairTempParam
}

func (s *SubmitOperationTaskRequest) GetType() *string {
	return s.Type
}

func (s *SubmitOperationTaskRequest) SetCheckId(v int64) *SubmitOperationTaskRequest {
	s.CheckId = &v
	return s
}

func (s *SubmitOperationTaskRequest) SetDimensionType(v string) *SubmitOperationTaskRequest {
	s.DimensionType = &v
	return s
}

func (s *SubmitOperationTaskRequest) SetOperationTaskInstances(v []*SubmitOperationTaskRequestOperationTaskInstances) *SubmitOperationTaskRequest {
	s.OperationTaskInstances = v
	return s
}

func (s *SubmitOperationTaskRequest) SetRelationKey(v string) *SubmitOperationTaskRequest {
	s.RelationKey = &v
	return s
}

func (s *SubmitOperationTaskRequest) SetRepairTempParam(v []*SubmitOperationTaskRequestRepairTempParam) *SubmitOperationTaskRequest {
	s.RepairTempParam = v
	return s
}

func (s *SubmitOperationTaskRequest) SetType(v string) *SubmitOperationTaskRequest {
	s.Type = &v
	return s
}

func (s *SubmitOperationTaskRequest) Validate() error {
	if s.OperationTaskInstances != nil {
		for _, item := range s.OperationTaskInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RepairTempParam != nil {
		for _, item := range s.RepairTempParam {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitOperationTaskRequestOperationTaskInstances struct {
	// The instance ID of the server.
	//
	// example:
	//
	// i-uf6533m4vuo3oa33****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the server.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID to roll back when performing a rollback task.
	//
	// example:
	//
	// 7d0b10e35e80c9e5ebac5f1054****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The asset vendor. Valid values:
	//
	// - **0**: Alibaba Cloud asset
	//
	// - **1**: asset outside the cloud
	//
	// - **2**: IDC asset
	//
	// - **3**, **4**, **5**, **7**: asset from another cloud provider
	//
	// - **8**: lightweight asset
	//
	// example:
	//
	// 7
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s SubmitOperationTaskRequestOperationTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationTaskRequestOperationTaskInstances) GoString() string {
	return s.String()
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) GetVendor() *string {
	return s.Vendor
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) SetInstanceId(v string) *SubmitOperationTaskRequestOperationTaskInstances {
	s.InstanceId = &v
	return s
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) SetRegionId(v string) *SubmitOperationTaskRequestOperationTaskInstances {
	s.RegionId = &v
	return s
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) SetTaskId(v string) *SubmitOperationTaskRequestOperationTaskInstances {
	s.TaskId = &v
	return s
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) SetVendor(v string) *SubmitOperationTaskRequestOperationTaskInstances {
	s.Vendor = &v
	return s
}

func (s *SubmitOperationTaskRequestOperationTaskInstances) Validate() error {
	return dara.Validate(s)
}

type SubmitOperationTaskRequestRepairTempParam struct {
	// The name of the temporary remediation parameter.
	//
	// example:
	//
	// IPPort
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the temporary remediation parameter.
	//
	// example:
	//
	// 192.168.1XX.1XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SubmitOperationTaskRequestRepairTempParam) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationTaskRequestRepairTempParam) GoString() string {
	return s.String()
}

func (s *SubmitOperationTaskRequestRepairTempParam) GetName() *string {
	return s.Name
}

func (s *SubmitOperationTaskRequestRepairTempParam) GetValue() *string {
	return s.Value
}

func (s *SubmitOperationTaskRequestRepairTempParam) SetName(v string) *SubmitOperationTaskRequestRepairTempParam {
	s.Name = &v
	return s
}

func (s *SubmitOperationTaskRequestRepairTempParam) SetValue(v string) *SubmitOperationTaskRequestRepairTempParam {
	s.Value = &v
	return s
}

func (s *SubmitOperationTaskRequestRepairTempParam) Validate() error {
	return dara.Validate(s)
}
