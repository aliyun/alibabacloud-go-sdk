// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWorkflowsResponseBody
	GetCode() *int32
	SetData(v *ListWorkflowsResponseBodyData) *ListWorkflowsResponseBody
	GetData() *ListWorkflowsResponseBodyData
	SetMessage(v string) *ListWorkflowsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWorkflowsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkflowsResponseBody
	GetSuccess() *bool
}

type ListWorkflowsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListWorkflowsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39938688-0BAB-5AD8-BF02-F4910FAC7589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWorkflowsResponseBody) GetData() *ListWorkflowsResponseBodyData {
	return s.Data
}

func (s *ListWorkflowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkflowsResponseBody) SetCode(v int32) *ListWorkflowsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowsResponseBody) SetData(v *ListWorkflowsResponseBodyData) *ListWorkflowsResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowsResponseBody) SetMessage(v string) *ListWorkflowsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowsResponseBody) SetRequestId(v string) *ListWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowsResponseBody) SetSuccess(v bool) *ListWorkflowsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkflowsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowsResponseBodyData struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListWorkflowsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 64
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListWorkflowsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowsResponseBodyData) GetRecords() []*ListWorkflowsResponseBodyDataRecords {
	return s.Records
}

func (s *ListWorkflowsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListWorkflowsResponseBodyData) SetMaxResults(v int32) *ListWorkflowsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetNextToken(v string) *ListWorkflowsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetPageNumber(v int32) *ListWorkflowsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetPageSize(v int32) *ListWorkflowsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetRecords(v []*ListWorkflowsResponseBodyDataRecords) *ListWorkflowsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetTotal(v int32) *ListWorkflowsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) Validate() error {
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

type ListWorkflowsResponseBodyDataRecords struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// work-day
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// example:
	//
	// 1963096506470832
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 3
	CurrentExecuteStatus *int32 `json:"CurrentExecuteStatus,omitempty" xml:"CurrentExecuteStatus,omitempty"`
	// example:
	//
	// my first workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2025-06-29 15:56:36
	LastExecuteEndTime *string `json:"LastExecuteEndTime,omitempty" xml:"LastExecuteEndTime,omitempty"`
	// example:
	//
	// 4
	LastExecuteStatus *int32 `json:"LastExecuteStatus,omitempty" xml:"LastExecuteStatus,omitempty"`
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// myWorkflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0 0 12 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// HangKong
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// example:
	//
	// 1963096506470832
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
	// example:
	//
	// 10
	WorkflowId *int64  `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	Xattrs     *string `json:"Xattrs,omitempty" xml:"Xattrs,omitempty"`
}

func (s ListWorkflowsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *ListWorkflowsResponseBodyDataRecords) GetAppType() *int32 {
	return s.AppType
}

func (s *ListWorkflowsResponseBodyDataRecords) GetCalendar() *string {
	return s.Calendar
}

func (s *ListWorkflowsResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkflowsResponseBodyDataRecords) GetCurrentExecuteStatus() *int32 {
	return s.CurrentExecuteStatus
}

func (s *ListWorkflowsResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *ListWorkflowsResponseBodyDataRecords) GetLastExecuteEndTime() *string {
	return s.LastExecuteEndTime
}

func (s *ListWorkflowsResponseBodyDataRecords) GetLastExecuteStatus() *int32 {
	return s.LastExecuteStatus
}

func (s *ListWorkflowsResponseBodyDataRecords) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *ListWorkflowsResponseBodyDataRecords) GetName() *string {
	return s.Name
}

func (s *ListWorkflowsResponseBodyDataRecords) GetStatus() *int32 {
	return s.Status
}

func (s *ListWorkflowsResponseBodyDataRecords) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *ListWorkflowsResponseBodyDataRecords) GetTimeType() *int32 {
	return s.TimeType
}

func (s *ListWorkflowsResponseBodyDataRecords) GetTimezone() *string {
	return s.Timezone
}

func (s *ListWorkflowsResponseBodyDataRecords) GetUpdater() *string {
	return s.Updater
}

func (s *ListWorkflowsResponseBodyDataRecords) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkflowsResponseBodyDataRecords) GetXattrs() *string {
	return s.Xattrs
}

func (s *ListWorkflowsResponseBodyDataRecords) SetAppName(v string) *ListWorkflowsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetAppType(v int32) *ListWorkflowsResponseBodyDataRecords {
	s.AppType = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetCalendar(v string) *ListWorkflowsResponseBodyDataRecords {
	s.Calendar = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetCreator(v string) *ListWorkflowsResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetCurrentExecuteStatus(v int32) *ListWorkflowsResponseBodyDataRecords {
	s.CurrentExecuteStatus = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetDescription(v string) *ListWorkflowsResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetLastExecuteEndTime(v string) *ListWorkflowsResponseBodyDataRecords {
	s.LastExecuteEndTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetLastExecuteStatus(v int32) *ListWorkflowsResponseBodyDataRecords {
	s.LastExecuteStatus = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetMaxConcurrency(v int32) *ListWorkflowsResponseBodyDataRecords {
	s.MaxConcurrency = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetName(v string) *ListWorkflowsResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetStatus(v int32) *ListWorkflowsResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetTimeExpression(v string) *ListWorkflowsResponseBodyDataRecords {
	s.TimeExpression = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetTimeType(v int32) *ListWorkflowsResponseBodyDataRecords {
	s.TimeType = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetTimezone(v string) *ListWorkflowsResponseBodyDataRecords {
	s.Timezone = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetUpdater(v string) *ListWorkflowsResponseBodyDataRecords {
	s.Updater = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetWorkflowId(v int64) *ListWorkflowsResponseBodyDataRecords {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) SetXattrs(v string) *ListWorkflowsResponseBodyDataRecords {
	s.Xattrs = &v
	return s
}

func (s *ListWorkflowsResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
