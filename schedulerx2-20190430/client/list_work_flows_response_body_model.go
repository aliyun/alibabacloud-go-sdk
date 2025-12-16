// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWorkFlowsResponseBody
	GetCode() *int32
	SetData(v *ListWorkFlowsResponseBodyData) *ListWorkFlowsResponseBody
	GetData() *ListWorkFlowsResponseBodyData
	SetMessage(v string) *ListWorkFlowsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWorkFlowsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkFlowsResponseBody
	GetSuccess() *bool
}

type ListWorkFlowsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the workflow.
	Data *ListWorkFlowsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// workflow is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 883AFE93-FB03-4FA9-A958-E750C6DE120C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkFlowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkFlowsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWorkFlowsResponseBody) GetData() *ListWorkFlowsResponseBodyData {
	return s.Data
}

func (s *ListWorkFlowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkFlowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkFlowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkFlowsResponseBody) SetCode(v int32) *ListWorkFlowsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkFlowsResponseBody) SetData(v *ListWorkFlowsResponseBodyData) *ListWorkFlowsResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkFlowsResponseBody) SetMessage(v string) *ListWorkFlowsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkFlowsResponseBody) SetRequestId(v string) *ListWorkFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkFlowsResponseBody) SetSuccess(v bool) *ListWorkFlowsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkFlowsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkFlowsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The row of data.
	Records []*ListWorkFlowsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListWorkFlowsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkFlowsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkFlowsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkFlowsResponseBodyData) GetRecords() []*ListWorkFlowsResponseBodyDataRecords {
	return s.Records
}

func (s *ListWorkFlowsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListWorkFlowsResponseBodyData) SetPageNumber(v int32) *ListWorkFlowsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListWorkFlowsResponseBodyData) SetPageSize(v int32) *ListWorkFlowsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListWorkFlowsResponseBodyData) SetRecords(v []*ListWorkFlowsResponseBodyDataRecords) *ListWorkFlowsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListWorkFlowsResponseBodyData) SetTotal(v int32) *ListWorkFlowsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListWorkFlowsResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkFlowsResponseBodyDataRecords struct {
	// The calendar.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The creator.
	//
	// example:
	//
	// 1144881807903942
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The job description.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application ID.
	//
	// example:
	//
	// hxm.test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum concurrency.
	//
	// example:
	//
	// 1
	MaxConcurrency *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// test3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// example:
	//
	// 4a06d5ea-f576-4326-842c-fb14ea043d8d
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The time expression.
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type of the workflow.
	//
	// example:
	//
	// cron
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The updater.
	//
	// example:
	//
	// 1144881807903942
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 1339
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkFlowsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetCalendar() *string {
	return s.Calendar
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetGroupId() *string {
	return s.GroupId
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetMaxConcurrency() *string {
	return s.MaxConcurrency
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetName() *string {
	return s.Name
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetNamespace() *string {
	return s.Namespace
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetTimeType() *string {
	return s.TimeType
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetUpdater() *string {
	return s.Updater
}

func (s *ListWorkFlowsResponseBodyDataRecords) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetCalendar(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.Calendar = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetCreator(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetDescription(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetGroupId(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.GroupId = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetMaxConcurrency(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.MaxConcurrency = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetName(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetNamespace(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.Namespace = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetTimeExpression(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.TimeExpression = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetTimeType(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.TimeType = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetUpdater(v string) *ListWorkFlowsResponseBodyDataRecords {
	s.Updater = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) SetWorkflowId(v int64) *ListWorkFlowsResponseBodyDataRecords {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkFlowsResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
