// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarTaskAndActionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetails(v *DescribeSoarTaskAndActionsResponseBodyDetails) *DescribeSoarTaskAndActionsResponseBody
	GetDetails() *DescribeSoarTaskAndActionsResponseBodyDetails
	SetPage(v *DescribeSoarTaskAndActionsResponseBodyPage) *DescribeSoarTaskAndActionsResponseBody
	GetPage() *DescribeSoarTaskAndActionsResponseBodyPage
	SetRequestId(v string) *DescribeSoarTaskAndActionsResponseBody
	GetRequestId() *string
}

type DescribeSoarTaskAndActionsResponseBody struct {
	// The execution details of each task.
	Details *DescribeSoarTaskAndActionsResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	Page    *DescribeSoarTaskAndActionsResponseBodyPage    `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 18017A93-3D5D-503A-8308-914543F1CBA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBody) GetDetails() *DescribeSoarTaskAndActionsResponseBodyDetails {
	return s.Details
}

func (s *DescribeSoarTaskAndActionsResponseBody) GetPage() *DescribeSoarTaskAndActionsResponseBodyPage {
	return s.Page
}

func (s *DescribeSoarTaskAndActionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarTaskAndActionsResponseBody) SetDetails(v *DescribeSoarTaskAndActionsResponseBodyDetails) *DescribeSoarTaskAndActionsResponseBody {
	s.Details = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBody) SetPage(v *DescribeSoarTaskAndActionsResponseBodyPage) *DescribeSoarTaskAndActionsResponseBody {
	s.Page = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBody) SetRequestId(v string) *DescribeSoarTaskAndActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBody) Validate() error {
	if s.Details != nil {
		if err := s.Details.Validate(); err != nil {
			return err
		}
	}
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSoarTaskAndActionsResponseBodyDetails struct {
	ActionLogNum *int32 `json:"ActionLogNum,omitempty" xml:"ActionLogNum,omitempty"`
	// The list of component actions during the running of the playbook.
	Actions []*DescribeSoarTaskAndActionsResponseBodyDetailsActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The end of the time range during which the playbook is run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1699868848767
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error message of the task. If the task is successful, this field is empty.
	//
	// example:
	//
	// stime not match
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The request parameters of the task.
	//
	// example:
	//
	// {
	//
	//     "input1": "xx.xx.xx.xx",
	//
	//     "input2": "7d"
	//
	// }
	RawEventReq *string `json:"RawEventReq,omitempty" xml:"RawEventReq,omitempty"`
	// The request ID of the task. The value is unique.
	//
	// example:
	//
	// 17f75844-75cc-4174-86da-cec07a690142
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	// The beginning of the time range during which the playbook is run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1699868848645
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// 	- **running**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The MD5 value of the playbook.
	//
	// example:
	//
	// ed127287-6699-4e4d-b986-9f770879xxx
	TaskFlowMd5 *string `json:"TaskFlowMd5,omitempty" xml:"TaskFlowMd5,omitempty"`
	// The name of the task. The value is the same as the playbook UUID.
	//
	// example:
	//
	// 92af3c79-1754-4646-9366-9ddbd1e45536
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task type. Valid values:
	//
	// 	- **debug**: a debugging task
	//
	// 	- **manual**: a manual task
	//
	// 	- **siem**: an event-triggered task
	//
	// example:
	//
	// siem
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The ID of the Alibaba Cloud account that triggers the task.
	//
	// example:
	//
	// 127xxxx4392
	TriggerUser *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBodyDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetActionLogNum() *int32 {
	return s.ActionLogNum
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetActions() []*DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	return s.Actions
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetRawEventReq() *string {
	return s.RawEventReq
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetRequestUuid() *string {
	return s.RequestUuid
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetStatus() *string {
	return s.Status
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetTaskFlowMd5() *string {
	return s.TaskFlowMd5
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetTriggerType() *string {
	return s.TriggerType
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) GetTriggerUser() *string {
	return s.TriggerUser
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetActionLogNum(v int32) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.ActionLogNum = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetActions(v []*DescribeSoarTaskAndActionsResponseBodyDetailsActions) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.Actions = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetEndTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.EndTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetErrorMsg(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetRawEventReq(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.RawEventReq = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetRequestUuid(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetStartTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.StartTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetStatus(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.Status = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTaskFlowMd5(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TaskFlowMd5 = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTaskName(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTriggerType(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TriggerType = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTriggerUser(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TriggerUser = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSoarTaskAndActionsResponseBodyDetailsActions struct {
	// The action name of the component.
	//
	// example:
	//
	// formatdata
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The UUID of the component execution record.
	//
	// example:
	//
	// 091be399-a937-4276-af78-xxxxxxxx
	ActionUuid *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	// The name of the asset that is used by the component.
	//
	// example:
	//
	// SLS Asset
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The component name.
	//
	// example:
	//
	// DataFormat
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	// The end of the time range during which the component is run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1699868848766
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The custom name of the node in the component.
	//
	// example:
	//
	// DataFormat_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The beginning of the time range during which the component is run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1699868848731
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The running result of the component. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBodyDetailsActions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBodyDetailsActions) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) GetAction() *string {
	return s.Action
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) GetActionUuid() *string {
	return s.ActionUuid
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) GetAssetName() *string {
	return s.AssetName
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) GetComponent() *string {
	return s.Component
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) GetStatus() *string {
	return s.Status
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetAction(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Action = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetActionUuid(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetAssetName(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.AssetName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetComponent(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Component = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetEndTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.EndTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetNodeName(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.NodeName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetStartTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.StartTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetStatus(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Status = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) Validate() error {
	return dara.Validate(s)
}

type DescribeSoarTaskAndActionsResponseBodyPage struct {
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBodyPage) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSoarTaskAndActionsResponseBodyPage) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSoarTaskAndActionsResponseBodyPage) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSoarTaskAndActionsResponseBodyPage) SetPageNumber(v string) *DescribeSoarTaskAndActionsResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyPage) SetPageSize(v string) *DescribeSoarTaskAndActionsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyPage) SetTotalCount(v string) *DescribeSoarTaskAndActionsResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
