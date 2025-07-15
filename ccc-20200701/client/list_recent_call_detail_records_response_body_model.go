// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentCallDetailRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRecentCallDetailRecordsResponseBody
	GetCode() *string
	SetData(v *ListRecentCallDetailRecordsResponseBodyData) *ListRecentCallDetailRecordsResponseBody
	GetData() *ListRecentCallDetailRecordsResponseBodyData
	SetHttpStatusCode(v int32) *ListRecentCallDetailRecordsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRecentCallDetailRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRecentCallDetailRecordsResponseBody
	GetRequestId() *string
}

type ListRecentCallDetailRecordsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListRecentCallDetailRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListRecentCallDetailRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecentCallDetailRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRecentCallDetailRecordsResponseBody) GetData() *ListRecentCallDetailRecordsResponseBodyData {
	return s.Data
}

func (s *ListRecentCallDetailRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRecentCallDetailRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRecentCallDetailRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecentCallDetailRecordsResponseBody) SetCode(v string) *ListRecentCallDetailRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) SetData(v *ListRecentCallDetailRecordsResponseBodyData) *ListRecentCallDetailRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) SetHttpStatusCode(v int32) *ListRecentCallDetailRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) SetMessage(v string) *ListRecentCallDetailRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) SetRequestId(v string) *ListRecentCallDetailRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRecentCallDetailRecordsResponseBodyData struct {
	List []*ListRecentCallDetailRecordsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListRecentCallDetailRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRecentCallDetailRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsResponseBodyData) GetList() []*ListRecentCallDetailRecordsResponseBodyDataList {
	return s.List
}

func (s *ListRecentCallDetailRecordsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecentCallDetailRecordsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecentCallDetailRecordsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRecentCallDetailRecordsResponseBodyData) SetList(v []*ListRecentCallDetailRecordsResponseBodyDataList) *ListRecentCallDetailRecordsResponseBodyData {
	s.List = v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyData) SetPageNumber(v int32) *ListRecentCallDetailRecordsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyData) SetPageSize(v int32) *ListRecentCallDetailRecordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyData) SetTotalCount(v int32) *ListRecentCallDetailRecordsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListRecentCallDetailRecordsResponseBodyDataList struct {
	// example:
	//
	// agent@ccc-test
	AgentIds *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// example:
	//
	// 16
	CallDuration *string `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	// example:
	//
	// 1332315****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 0533128****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// Success
	ContactDisposition *string `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	// example:
	//
	// job-7660472242845****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// Outbound
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
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
	// skillgroup@ccc-test
	SkillGroupIds *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRecentCallDetailRecordsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListRecentCallDetailRecordsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetAgentIds() *string {
	return s.AgentIds
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetCallDuration() *string {
	return s.CallDuration
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetContactDisposition() *string {
	return s.ContactDisposition
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetContactId() *string {
	return s.ContactId
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetContactType() *string {
	return s.ContactType
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetSkillGroupIds() *string {
	return s.SkillGroupIds
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetAgentIds(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.AgentIds = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetCallDuration(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.CallDuration = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetCalledNumber(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.CalledNumber = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetCallingNumber(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.CallingNumber = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetContactDisposition(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.ContactDisposition = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetContactId(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetContactType(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.ContactType = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetDuration(v int64) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.Duration = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetInstanceId(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetSkillGroupIds(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.SkillGroupIds = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetStartTime(v int64) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
