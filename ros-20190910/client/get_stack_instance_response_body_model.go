// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackInstanceResponseBody
	GetRequestId() *string
	SetStackInstance(v *GetStackInstanceResponseBodyStackInstance) *GetStackInstanceResponseBody
	GetStackInstance() *GetStackInstanceResponseBodyStackInstance
}

type GetStackInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B8A6B693-82C8-419D-8796-DE99EC33CFF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the stack.
	StackInstance *GetStackInstanceResponseBodyStackInstance `json:"StackInstance,omitempty" xml:"StackInstance,omitempty" type:"Struct"`
}

func (s GetStackInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackInstanceResponseBody) GetStackInstance() *GetStackInstanceResponseBodyStackInstance {
	return s.StackInstance
}

func (s *GetStackInstanceResponseBody) SetRequestId(v string) *GetStackInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackInstanceResponseBody) SetStackInstance(v *GetStackInstanceResponseBodyStackInstance) *GetStackInstanceResponseBody {
	s.StackInstance = v
	return s
}

func (s *GetStackInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStackInstanceResponseBodyStackInstance struct {
	// The ID of the destination account to which the stack belongs.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the most recent successful drift detection was performed on the stack group.
	//
	// > This parameter is returned only if drift detection is performed on the stack group.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The outputs of the stack.
	//
	// >  This parameter is returned if OutputOption is set to Enabled.
	Outputs []map[string]interface{} `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*GetStackInstanceResponseBodyStackInstanceParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	// The ID of the folder in the resource directory.
	//
	// > This parameter is returned only if the stack group is granted service-managed permissions.
	//
	// example:
	//
	// fd-4PvlVLOL8v
	RdFolderId *string `json:"RdFolderId,omitempty" xml:"RdFolderId,omitempty"`
	// The region ID of the stack.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the stack when the most recent successful drift detection was performed on the stack group.
	//
	// Valid values:
	//
	// 	- DRIFTED: The stack has drifted.
	//
	// 	- NOT_CHECKED: No successful drift detection is performed on the stack.
	//
	// 	- IN_SYNC: The stack is being synchronized.
	//
	// > This parameter is returned only if drift detection is performed on the stack group.
	//
	// example:
	//
	// IN_SYNC
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The ID of the stack group.
	//
	// example:
	//
	// fd0ddef9-9540-4b42-a464-94f77835****
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The stack ID.
	//
	// > This parameter is returned only if the stack is in the CURRENT state.
	//
	// example:
	//
	// 35ad60e3-6a92-42d8-8812-f0700d45****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The state of the stack.
	//
	// Valid values:
	//
	// 	- CURRENT: The stack is up-to-date with the stack group.
	//
	// 	- OUTDATED: The stack is not up-to-date with the stack group. Stacks are in the OUTDATED state due to the following possible reasons:
	//
	//     	- When the CreateStackInstances operation is called to create stacks, the stacks fail to be created.
	//
	//     	- When the UpdateStackInstances or UpdateStackGroup operation is called to update stacks, the stacks fail to be updated, or only specific stacks are updated.
	//
	//     	- The creation or update operation is not complete.
	//
	// example:
	//
	// CURRENT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the stack instance is in the OUTDATED state.
	//
	// > This parameter is returned only if the stack instance is in the OUTDATED state.
	//
	// example:
	//
	// User initiated stop
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s GetStackInstanceResponseBodyStackInstance) String() string {
	return dara.Prettify(s)
}

func (s GetStackInstanceResponseBodyStackInstance) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBodyStackInstance) GetAccountId() *string {
	return s.AccountId
}

func (s *GetStackInstanceResponseBodyStackInstance) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *GetStackInstanceResponseBodyStackInstance) GetOutputs() []map[string]interface{} {
	return s.Outputs
}

func (s *GetStackInstanceResponseBodyStackInstance) GetParameterOverrides() []*GetStackInstanceResponseBodyStackInstanceParameterOverrides {
	return s.ParameterOverrides
}

func (s *GetStackInstanceResponseBodyStackInstance) GetRdFolderId() *string {
	return s.RdFolderId
}

func (s *GetStackInstanceResponseBodyStackInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackInstanceResponseBodyStackInstance) GetStackDriftStatus() *string {
	return s.StackDriftStatus
}

func (s *GetStackInstanceResponseBodyStackInstance) GetStackGroupId() *string {
	return s.StackGroupId
}

func (s *GetStackInstanceResponseBodyStackInstance) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *GetStackInstanceResponseBodyStackInstance) GetStackId() *string {
	return s.StackId
}

func (s *GetStackInstanceResponseBodyStackInstance) GetStatus() *string {
	return s.Status
}

func (s *GetStackInstanceResponseBodyStackInstance) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetStackInstanceResponseBodyStackInstance) SetAccountId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.AccountId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetDriftDetectionTime(v string) *GetStackInstanceResponseBodyStackInstance {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetOutputs(v []map[string]interface{}) *GetStackInstanceResponseBodyStackInstance {
	s.Outputs = v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetParameterOverrides(v []*GetStackInstanceResponseBodyStackInstanceParameterOverrides) *GetStackInstanceResponseBodyStackInstance {
	s.ParameterOverrides = v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetRdFolderId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.RdFolderId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetRegionId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.RegionId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackDriftStatus(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackGroupId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackGroupId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackGroupName(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackGroupName = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStatus(v string) *GetStackInstanceResponseBodyStackInstance {
	s.Status = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStatusReason(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StatusReason = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) Validate() error {
	return dara.Validate(s)
}

type GetStackInstanceResponseBodyStackInstanceParameterOverrides struct {
	// The name of the parameter that is used to override a specific parameter.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter that is used to override a specific parameter.
	//
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackInstanceResponseBodyStackInstanceParameterOverrides) String() string {
	return dara.Prettify(s)
}

func (s GetStackInstanceResponseBodyStackInstanceParameterOverrides) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) SetParameterKey(v string) *GetStackInstanceResponseBodyStackInstanceParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) SetParameterValue(v string) *GetStackInstanceResponseBodyStackInstanceParameterOverrides {
	s.ParameterValue = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) Validate() error {
	return dara.Validate(s)
}
