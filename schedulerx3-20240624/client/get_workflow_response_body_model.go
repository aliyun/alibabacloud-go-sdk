// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWorkflowResponseBody
	GetCode() *int32
	SetData(v *GetWorkflowResponseBodyData) *GetWorkflowResponseBody
	GetData() *GetWorkflowResponseBodyData
	SetMessage(v string) *GetWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkflowResponseBody
	GetSuccess() *bool
}

type GetWorkflowResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 27B1345D-5F71-5972-8E4C-AABA6C6232F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWorkflowResponseBody) GetData() *GetWorkflowResponseBodyData {
	return s.Data
}

func (s *GetWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkflowResponseBody) SetCode(v int32) *GetWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowResponseBody) SetData(v *GetWorkflowResponseBodyData) *GetWorkflowResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkflowResponseBody) SetMessage(v string) *GetWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowResponseBody) SetRequestId(v string) *GetWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowResponseBody) SetSuccess(v bool) *GetWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkflowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowResponseBodyData struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// example:
	//
	// 18582193685027xx
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// my first workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 2
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
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// example:
	//
	// 18582193685027xx
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
	// example:
	//
	// 10
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// example:
	//
	// null
	Xattrs *string `json:"Xattrs,omitempty" xml:"Xattrs,omitempty"`
}

func (s GetWorkflowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetWorkflowResponseBodyData) GetCalendar() *string {
	return s.Calendar
}

func (s *GetWorkflowResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetWorkflowResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetWorkflowResponseBodyData) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *GetWorkflowResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetWorkflowResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetWorkflowResponseBodyData) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *GetWorkflowResponseBodyData) GetTimeType() *int32 {
	return s.TimeType
}

func (s *GetWorkflowResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *GetWorkflowResponseBodyData) GetUpdater() *string {
	return s.Updater
}

func (s *GetWorkflowResponseBodyData) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkflowResponseBodyData) GetXattrs() *string {
	return s.Xattrs
}

func (s *GetWorkflowResponseBodyData) SetAppName(v string) *GetWorkflowResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetCalendar(v string) *GetWorkflowResponseBodyData {
	s.Calendar = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetCreator(v string) *GetWorkflowResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetDescription(v string) *GetWorkflowResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetMaxConcurrency(v int32) *GetWorkflowResponseBodyData {
	s.MaxConcurrency = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetName(v string) *GetWorkflowResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetStatus(v int32) *GetWorkflowResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetTimeExpression(v string) *GetWorkflowResponseBodyData {
	s.TimeExpression = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetTimeType(v int32) *GetWorkflowResponseBodyData {
	s.TimeType = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetTimezone(v string) *GetWorkflowResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetUpdater(v string) *GetWorkflowResponseBodyData {
	s.Updater = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetWorkflowId(v int64) *GetWorkflowResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowResponseBodyData) SetXattrs(v string) *GetWorkflowResponseBodyData {
	s.Xattrs = &v
	return s
}

func (s *GetWorkflowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
