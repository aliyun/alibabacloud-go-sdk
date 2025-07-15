// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyQueueEventLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLegacyQueueEventLogsResponseBody
	GetCode() *string
	SetData(v *ListLegacyQueueEventLogsResponseBodyData) *ListLegacyQueueEventLogsResponseBody
	GetData() *ListLegacyQueueEventLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListLegacyQueueEventLogsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListLegacyQueueEventLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLegacyQueueEventLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLegacyQueueEventLogsResponseBody
	GetSuccess() *bool
}

type ListLegacyQueueEventLogsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListLegacyQueueEventLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1A5A8998-41F9-5F85-BFCF-EB2B6E376812
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLegacyQueueEventLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyQueueEventLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLegacyQueueEventLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLegacyQueueEventLogsResponseBody) GetData() *ListLegacyQueueEventLogsResponseBodyData {
	return s.Data
}

func (s *ListLegacyQueueEventLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLegacyQueueEventLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLegacyQueueEventLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLegacyQueueEventLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLegacyQueueEventLogsResponseBody) SetCode(v string) *ListLegacyQueueEventLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBody) SetData(v *ListLegacyQueueEventLogsResponseBodyData) *ListLegacyQueueEventLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBody) SetHttpStatusCode(v int32) *ListLegacyQueueEventLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBody) SetMessage(v string) *ListLegacyQueueEventLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBody) SetRequestId(v string) *ListLegacyQueueEventLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBody) SetSuccess(v bool) *ListLegacyQueueEventLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLegacyQueueEventLogsResponseBodyData struct {
	List []*ListLegacyQueueEventLogsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLegacyQueueEventLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyQueueEventLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLegacyQueueEventLogsResponseBodyData) GetList() []*ListLegacyQueueEventLogsResponseBodyDataList {
	return s.List
}

func (s *ListLegacyQueueEventLogsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLegacyQueueEventLogsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLegacyQueueEventLogsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLegacyQueueEventLogsResponseBodyData) SetList(v []*ListLegacyQueueEventLogsResponseBodyDataList) *ListLegacyQueueEventLogsResponseBodyData {
	s.List = v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyData) SetPageNumber(v int32) *ListLegacyQueueEventLogsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyData) SetPageSize(v int32) *ListLegacyQueueEventLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyData) SetTotalCount(v int32) *ListLegacyQueueEventLogsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListLegacyQueueEventLogsResponseBodyDataList struct {
	// example:
	//
	// 456328****
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// example:
	//
	// 8012****
	Ani *string `json:"Ani,omitempty" xml:"Ani,omitempty"`
	// example:
	//
	// agent@ccc-test
	AnswerPhone *string `json:"AnswerPhone,omitempty" xml:"AnswerPhone,omitempty"`
	// example:
	//
	// 15
	AnswerTime *int64  `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	Cause      *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// example:
	//
	// 1312211****
	Dnis *string `json:"Dnis,omitempty" xml:"Dnis,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10
	QueueTime *int64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// example:
	//
	// 2021-12-03T10:15:30
	StatisticDate *string `json:"StatisticDate,omitempty" xml:"StatisticDate,omitempty"`
	// example:
	//
	// acc3733
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	Vq *string `json:"Vq,omitempty" xml:"Vq,omitempty"`
}

func (s ListLegacyQueueEventLogsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyQueueEventLogsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetAcid() *string {
	return s.Acid
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetAni() *string {
	return s.Ani
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetAnswerPhone() *string {
	return s.AnswerPhone
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetAnswerTime() *int64 {
	return s.AnswerTime
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetCause() *string {
	return s.Cause
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetDnis() *string {
	return s.Dnis
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetQueueTime() *int64 {
	return s.QueueTime
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetStatisticDate() *string {
	return s.StatisticDate
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) GetVq() *string {
	return s.Vq
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetAcid(v string) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.Acid = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetAni(v string) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.Ani = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetAnswerPhone(v string) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.AnswerPhone = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetAnswerTime(v int64) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.AnswerTime = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetCause(v string) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.Cause = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetDnis(v string) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.Dnis = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetId(v int64) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetQueueTime(v int64) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.QueueTime = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetStatisticDate(v string) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.StatisticDate = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetTenantId(v string) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.TenantId = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) SetVq(v string) *ListLegacyQueueEventLogsResponseBodyDataList {
	s.Vq = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
