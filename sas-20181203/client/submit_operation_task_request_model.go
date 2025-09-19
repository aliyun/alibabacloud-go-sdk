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
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the ID of the check item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 132
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The dimension of the task that you want to submit. Valid values:
	//
	// 	- Instance dimension: INSTANCE
	//
	// 	- Check item dimension: CHECK_ID
	//
	// example:
	//
	// CHECK_ID
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	// The asset information required to submit the tasks for instances.
	OperationTaskInstances []*SubmitOperationTaskRequestOperationTaskInstances `json:"OperationTaskInstances,omitempty" xml:"OperationTaskInstances,omitempty" type:"Repeated"`
	// The key linked to cross-page selections during task submission.
	//
	// >  You can call the [CreateAssetSelectionConfig](~~CreateAssetSelectionConfig~~) operation to query the associated key from the BusinessType field.
	//
	// example:
	//
	// CSPM_OPERATION_RELATION_KEY_173***
	RelationKey *string `json:"RelationKey,omitempty" xml:"RelationKey,omitempty"`
	// The temporary parameters required for the repair task.
	RepairTempParam []*SubmitOperationTaskRequestRepairTempParam `json:"RepairTempParam,omitempty" xml:"RepairTempParam,omitempty" type:"Repeated"`
	// The type of the task that you want to submit. Valid values:
	//
	// 	- Repair task: REPAIR
	//
	// 	- Rollback task: ROLLBACK
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
	return dara.Validate(s)
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
	// The ID of the task that you want to roll back
	//
	// example:
	//
	// 7d0b10e35e80c9e5ebac5f1054****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The service provider of the asset. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud.
	//
	// 	- **1**: an asset outside Alibaba Cloud.
	//
	// 	- **2**: an asset in a data center.
	//
	// 	- **3**, **4**, **5**, and **7**: an asset from a third-party cloud service provider.
	//
	// 	- **8**: a lightweight asset.
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
	// The name of the temporary repair parameter.
	//
	// example:
	//
	// IPPort
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the temporary repair parameter.
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
