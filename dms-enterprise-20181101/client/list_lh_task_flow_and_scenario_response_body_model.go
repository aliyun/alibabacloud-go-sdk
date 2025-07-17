// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLhTaskFlowAndScenarioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListLhTaskFlowAndScenarioResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListLhTaskFlowAndScenarioResponseBody
	GetErrorMessage() *string
	SetRawDAGList(v *ListLhTaskFlowAndScenarioResponseBodyRawDAGList) *ListLhTaskFlowAndScenarioResponseBody
	GetRawDAGList() *ListLhTaskFlowAndScenarioResponseBodyRawDAGList
	SetRequestId(v string) *ListLhTaskFlowAndScenarioResponseBody
	GetRequestId() *string
	SetScenarioDAGList(v *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList) *ListLhTaskFlowAndScenarioResponseBody
	GetScenarioDAGList() *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList
	SetSuccess(v bool) *ListLhTaskFlowAndScenarioResponseBody
	GetSuccess() *bool
}

type ListLhTaskFlowAndScenarioResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task flows in the default business scenario.
	RawDAGList *ListLhTaskFlowAndScenarioResponseBodyRawDAGList `json:"RawDAGList,omitempty" xml:"RawDAGList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 48602B78-0DDF-414C-8688-70CAB6070115
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task flows in other business scenarios.
	ScenarioDAGList *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList `json:"ScenarioDAGList,omitempty" xml:"ScenarioDAGList,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLhTaskFlowAndScenarioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListLhTaskFlowAndScenarioResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListLhTaskFlowAndScenarioResponseBody) GetRawDAGList() *ListLhTaskFlowAndScenarioResponseBodyRawDAGList {
	return s.RawDAGList
}

func (s *ListLhTaskFlowAndScenarioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLhTaskFlowAndScenarioResponseBody) GetScenarioDAGList() *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList {
	return s.ScenarioDAGList
}

func (s *ListLhTaskFlowAndScenarioResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLhTaskFlowAndScenarioResponseBody) SetErrorCode(v string) *ListLhTaskFlowAndScenarioResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBody) SetErrorMessage(v string) *ListLhTaskFlowAndScenarioResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBody) SetRawDAGList(v *ListLhTaskFlowAndScenarioResponseBodyRawDAGList) *ListLhTaskFlowAndScenarioResponseBody {
	s.RawDAGList = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBody) SetRequestId(v string) *ListLhTaskFlowAndScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBody) SetScenarioDAGList(v *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList) *ListLhTaskFlowAndScenarioResponseBody {
	s.ScenarioDAGList = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBody) SetSuccess(v bool) *ListLhTaskFlowAndScenarioResponseBody {
	s.Success = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLhTaskFlowAndScenarioResponseBodyRawDAGList struct {
	Dag []*ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag `json:"Dag,omitempty" xml:"Dag,omitempty" type:"Repeated"`
}

func (s ListLhTaskFlowAndScenarioResponseBodyRawDAGList) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponseBodyRawDAGList) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGList) GetDag() []*ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	return s.Dag
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGList) SetDag(v []*ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) *ListLhTaskFlowAndScenarioResponseBodyRawDAGList {
	s.Dag = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGList) Validate() error {
	return dara.Validate(s)
}

type ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag struct {
	// Indicates whether the task flow can be modified. Valid values:
	//
	// 	- **true**: The task flow can be modified.
	//
	// 	- **false**: The task flow cannot be modified.
	//
	// example:
	//
	// true
	CanEdit *bool `json:"CanEdit,omitempty" xml:"CanEdit,omitempty"`
	// The ID of the user who creates the task flow.
	//
	// example:
	//
	// 51****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the user who creates the workspace.
	//
	// example:
	//
	// Creator_Name
	CreatorNickName *string `json:"CreatorNickName,omitempty" xml:"CreatorNickName,omitempty"`
	// The name of the task flow.
	//
	// example:
	//
	// Dag_Name
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The user ID of the task flow owner.
	//
	// example:
	//
	// 51****
	DagOwnerId *string `json:"DagOwnerId,omitempty" xml:"DagOwnerId,omitempty"`
	// The name of the task flow owner.
	//
	// example:
	//
	// Owner_Name
	DagOwnerNickName *string `json:"DagOwnerNickName,omitempty" xml:"DagOwnerNickName,omitempty"`
	// The extended field. No meaning is specified for this field.
	//
	// example:
	//
	// -
	DataFlowId *int64 `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The extended field. No meaning is specified for this field.
	//
	// example:
	//
	// -
	DemoId *string `json:"DemoId,omitempty" xml:"DemoId,omitempty"`
	// The ID of the latest deployment record.
	//
	// example:
	//
	// 12**
	DeployId *int64 `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	// The ID of the task flow.
	//
	// example:
	//
	// 134137****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the task flow is deleted. Valid values:
	//
	// 	- **true**: deleted
	//
	// 	- **false**: not deleted
	//
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The status of the latest execution. Valid values:
	//
	// 	- **0**: invalid
	//
	// 	- **1**: scheduling disabled
	//
	// 	- **2**: waiting to be scheduled
	//
	// example:
	//
	// 0
	LatestInstanceStatus *int32 `json:"LatestInstanceStatus,omitempty" xml:"LatestInstanceStatus,omitempty"`
	// The time when the latest execution record was generated.
	//
	// example:
	//
	// 2022-04-14
	LatestInstanceTime *int32 `json:"LatestInstanceTime,omitempty" xml:"LatestInstanceTime,omitempty"`
	// The ID of the business scenario.
	//
	// example:
	//
	// 2**
	ScenarioId *int64 `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 1
	SpaceId *int64 `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// The status of the task flow. Valid values:
	//
	// 	- **0**: invalid
	//
	// 	- **1**: scheduling disabled
	//
	// 	- **2**: waiting to be scheduled
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetCanEdit() *bool {
	return s.CanEdit
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetCreatorNickName() *string {
	return s.CreatorNickName
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetDagName() *string {
	return s.DagName
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetDagOwnerId() *string {
	return s.DagOwnerId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetDagOwnerNickName() *string {
	return s.DagOwnerNickName
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetDataFlowId() *int64 {
	return s.DataFlowId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetDemoId() *string {
	return s.DemoId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetDeployId() *int64 {
	return s.DeployId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetId() *int64 {
	return s.Id
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetLatestInstanceStatus() *int32 {
	return s.LatestInstanceStatus
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetLatestInstanceTime() *int32 {
	return s.LatestInstanceTime
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetSpaceId() *int64 {
	return s.SpaceId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) GetStatus() *int32 {
	return s.Status
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetCanEdit(v bool) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.CanEdit = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetCreatorId(v string) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.CreatorId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetCreatorNickName(v string) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.CreatorNickName = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetDagName(v string) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.DagName = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetDagOwnerId(v string) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.DagOwnerId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetDagOwnerNickName(v string) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.DagOwnerNickName = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetDataFlowId(v int64) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.DataFlowId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetDemoId(v string) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.DemoId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetDeployId(v int64) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.DeployId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetId(v int64) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.Id = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetIsDeleted(v bool) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.IsDeleted = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetLatestInstanceStatus(v int32) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.LatestInstanceStatus = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetLatestInstanceTime(v int32) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.LatestInstanceTime = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetScenarioId(v int64) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.ScenarioId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetSpaceId(v int64) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.SpaceId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) SetStatus(v int32) *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag {
	s.Status = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyRawDAGListDag) Validate() error {
	return dara.Validate(s)
}

type ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList struct {
	ScenarioDAG []*ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG `json:"ScenarioDAG,omitempty" xml:"ScenarioDAG,omitempty" type:"Repeated"`
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList) GetScenarioDAG() []*ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG {
	return s.ScenarioDAG
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList) SetScenarioDAG(v []*ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList {
	s.ScenarioDAG = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGList) Validate() error {
	return dara.Validate(s)
}

type ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG struct {
	// The list of task flows.
	DagList *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList `json:"DagList,omitempty" xml:"DagList,omitempty" type:"Struct"`
	// The information about the business scenario.
	Scenario *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario `json:"Scenario,omitempty" xml:"Scenario,omitempty" type:"Struct"`
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG) GetDagList() *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList {
	return s.DagList
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG) GetScenario() *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario {
	return s.Scenario
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG) SetDagList(v *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG {
	s.DagList = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG) SetScenario(v *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG {
	s.Scenario = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAG) Validate() error {
	return dara.Validate(s)
}

type ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList struct {
	Dag []*ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag `json:"Dag,omitempty" xml:"Dag,omitempty" type:"Repeated"`
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList) GetDag() []*ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	return s.Dag
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList) SetDag(v []*ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList {
	s.Dag = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagList) Validate() error {
	return dara.Validate(s)
}

type ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag struct {
	// Indicates whether the task flow can be modified. Valid values:
	//
	// - **true**: The task flow can be modified.
	//
	// - **false**: The task flow cannot be modified.
	//
	// example:
	//
	// true
	CanEdit *bool `json:"CanEdit,omitempty" xml:"CanEdit,omitempty"`
	// The ID of the user who creates the task flow.
	//
	// example:
	//
	// 51****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the user who creates the workspace.
	//
	// example:
	//
	// Creator_Name
	CreatorNickName *string `json:"CreatorNickName,omitempty" xml:"CreatorNickName,omitempty"`
	// The name of the task flow.
	//
	// example:
	//
	// Dag_Name
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The user ID of the task flow owner.
	//
	// example:
	//
	// 51****
	DagOwnerId *string `json:"DagOwnerId,omitempty" xml:"DagOwnerId,omitempty"`
	// The name of the task flow owner.
	//
	// example:
	//
	// Owner_Name
	DagOwnerNickName *string `json:"DagOwnerNickName,omitempty" xml:"DagOwnerNickName,omitempty"`
	// The extended field. No meaning is specified for this field.
	//
	// example:
	//
	// -
	DataFlowId *int64 `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The extended field. No meaning is specified for this field.
	//
	// example:
	//
	// -
	DemoId *string `json:"DemoId,omitempty" xml:"DemoId,omitempty"`
	// The ID of the latest deployment record.
	//
	// example:
	//
	// 12**
	DeployId *int64 `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	// The ID of the task flow.
	//
	// example:
	//
	// 9***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the task flow is deleted. Valid values:
	//
	// - **true**: deleted
	//
	// - **false**: not deleted
	//
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The status of the latest execution. Valid values:
	//
	// - 0: invalid
	//
	// - 1: scheduling disabled
	//
	// - 2: waiting to be scheduled
	//
	// example:
	//
	// 1
	LatestInstanceStatus *int32 `json:"LatestInstanceStatus,omitempty" xml:"LatestInstanceStatus,omitempty"`
	// The time when the latest execution record was generated.
	//
	// example:
	//
	// 2022-04-14
	LatestInstanceTime *int32 `json:"LatestInstanceTime,omitempty" xml:"LatestInstanceTime,omitempty"`
	// The ID of the business scenario.
	//
	// example:
	//
	// 2**
	ScenarioId *int64 `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 1
	SpaceId *int64 `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// The status of the task flow. Valid values:
	//
	// - **0**: invalid
	//
	// - **1**: scheduling disabled
	//
	// - **2**: waiting to be scheduled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetCanEdit() *bool {
	return s.CanEdit
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetCreatorNickName() *string {
	return s.CreatorNickName
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetDagName() *string {
	return s.DagName
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetDagOwnerId() *string {
	return s.DagOwnerId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetDagOwnerNickName() *string {
	return s.DagOwnerNickName
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetDataFlowId() *int64 {
	return s.DataFlowId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetDemoId() *string {
	return s.DemoId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetDeployId() *int64 {
	return s.DeployId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetId() *int64 {
	return s.Id
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetLatestInstanceStatus() *int32 {
	return s.LatestInstanceStatus
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetLatestInstanceTime() *int32 {
	return s.LatestInstanceTime
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetSpaceId() *int64 {
	return s.SpaceId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) GetStatus() *int32 {
	return s.Status
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetCanEdit(v bool) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.CanEdit = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetCreatorId(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.CreatorId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetCreatorNickName(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.CreatorNickName = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetDagName(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.DagName = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetDagOwnerId(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.DagOwnerId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetDagOwnerNickName(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.DagOwnerNickName = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetDataFlowId(v int64) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.DataFlowId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetDemoId(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.DemoId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetDeployId(v int64) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.DeployId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetId(v int64) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.Id = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetIsDeleted(v bool) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.IsDeleted = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetLatestInstanceStatus(v int32) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.LatestInstanceStatus = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetLatestInstanceTime(v int32) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.LatestInstanceTime = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetScenarioId(v int64) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.ScenarioId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetSpaceId(v int64) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.SpaceId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) SetStatus(v int32) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag {
	s.Status = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGDagListDag) Validate() error {
	return dara.Validate(s)
}

type ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario struct {
	// The ID of the user who creates the business scenario.
	//
	// example:
	//
	// 51****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the business scenario.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the business scenario.
	//
	// example:
	//
	// Scenario_2
	ScenarioName *string `json:"ScenarioName,omitempty" xml:"ScenarioName,omitempty"`
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) GetDescription() *string {
	return s.Description
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) GetScenarioName() *string {
	return s.ScenarioName
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) SetCreatorId(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario {
	s.CreatorId = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) SetDescription(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario {
	s.Description = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) SetScenarioName(v string) *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario {
	s.ScenarioName = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponseBodyScenarioDAGListScenarioDAGScenario) Validate() error {
	return dara.Validate(s)
}
