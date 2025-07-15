// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoicemailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVoicemailsResponseBody
	GetCode() *string
	SetData(v *ListVoicemailsResponseBodyData) *ListVoicemailsResponseBody
	GetData() *ListVoicemailsResponseBodyData
	SetHttpStatusCode(v int32) *ListVoicemailsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVoicemailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVoicemailsResponseBody
	GetRequestId() *string
}

type ListVoicemailsResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListVoicemailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVoicemailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVoicemailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoicemailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVoicemailsResponseBody) GetData() *ListVoicemailsResponseBodyData {
	return s.Data
}

func (s *ListVoicemailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVoicemailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVoicemailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVoicemailsResponseBody) SetCode(v string) *ListVoicemailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoicemailsResponseBody) SetData(v *ListVoicemailsResponseBodyData) *ListVoicemailsResponseBody {
	s.Data = v
	return s
}

func (s *ListVoicemailsResponseBody) SetHttpStatusCode(v int32) *ListVoicemailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVoicemailsResponseBody) SetMessage(v string) *ListVoicemailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoicemailsResponseBody) SetRequestId(v string) *ListVoicemailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoicemailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVoicemailsResponseBodyData struct {
	List []*ListVoicemailsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVoicemailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVoicemailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVoicemailsResponseBodyData) GetList() []*ListVoicemailsResponseBodyDataList {
	return s.List
}

func (s *ListVoicemailsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoicemailsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoicemailsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVoicemailsResponseBodyData) SetList(v []*ListVoicemailsResponseBodyDataList) *ListVoicemailsResponseBodyData {
	s.List = v
	return s
}

func (s *ListVoicemailsResponseBodyData) SetPageNumber(v int32) *ListVoicemailsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListVoicemailsResponseBodyData) SetPageSize(v int32) *ListVoicemailsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListVoicemailsResponseBodyData) SetTotalCount(v int32) *ListVoicemailsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListVoicemailsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListVoicemailsResponseBodyDataList struct {
	// example:
	//
	// 0533128****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 073xxxx7539
	Caller       *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CdrStartTime *int64  `json:"CdrStartTime,omitempty" xml:"CdrStartTime,omitempty"`
	// example:
	//
	// job-125152394144124921
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// 16
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// voicemail-test
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecordingDuration *int64  `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// example:
	//
	// 1631440860000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListVoicemailsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListVoicemailsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListVoicemailsResponseBodyDataList) GetCallee() *string {
	return s.Callee
}

func (s *ListVoicemailsResponseBodyDataList) GetCaller() *string {
	return s.Caller
}

func (s *ListVoicemailsResponseBodyDataList) GetCdrStartTime() *int64 {
	return s.CdrStartTime
}

func (s *ListVoicemailsResponseBodyDataList) GetContactId() *string {
	return s.ContactId
}

func (s *ListVoicemailsResponseBodyDataList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListVoicemailsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoicemailsResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListVoicemailsResponseBodyDataList) GetRecordingDuration() *int64 {
	return s.RecordingDuration
}

func (s *ListVoicemailsResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListVoicemailsResponseBodyDataList) SetCallee(v string) *ListVoicemailsResponseBodyDataList {
	s.Callee = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) SetCaller(v string) *ListVoicemailsResponseBodyDataList {
	s.Caller = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) SetCdrStartTime(v int64) *ListVoicemailsResponseBodyDataList {
	s.CdrStartTime = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) SetContactId(v string) *ListVoicemailsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) SetDuration(v int64) *ListVoicemailsResponseBodyDataList {
	s.Duration = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) SetInstanceId(v string) *ListVoicemailsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) SetName(v string) *ListVoicemailsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) SetRecordingDuration(v int64) *ListVoicemailsResponseBodyDataList {
	s.RecordingDuration = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) SetStartTime(v string) *ListVoicemailsResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListVoicemailsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
