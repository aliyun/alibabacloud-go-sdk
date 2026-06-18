// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTaskDetailResponseBody
	GetCode() *string
	SetData(v *ListTaskDetailResponseBodyData) *ListTaskDetailResponseBody
	GetData() *ListTaskDetailResponseBodyData
	SetMessage(v string) *ListTaskDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTaskDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskDetailResponseBody
	GetSuccess() *bool
}

type ListTaskDetailResponseBody struct {
	// The request status code. A return value of OK indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Detailed job information.
	Data *ListTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API invocation succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTaskDetailResponseBody) GetData() *ListTaskDetailResponseBodyData {
	return s.Data
}

func (s *ListTaskDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskDetailResponseBody) SetCode(v string) *ListTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskDetailResponseBody) SetData(v *ListTaskDetailResponseBodyData) *ListTaskDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskDetailResponseBody) SetMessage(v string) *ListTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskDetailResponseBody) SetRequestId(v string) *ListTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskDetailResponseBody) SetSuccess(v bool) *ListTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskDetailResponseBodyData struct {
	// Number of items per page.
	//
	// example:
	//
	// 20
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of job information.
	Record []*ListTaskDetailResponseBodyDataRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
	// Total number of jobs.
	//
	// example:
	//
	// 50
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTaskDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTaskDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskDetailResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListTaskDetailResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTaskDetailResponseBodyData) GetRecord() []*ListTaskDetailResponseBodyDataRecord {
	return s.Record
}

func (s *ListTaskDetailResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListTaskDetailResponseBodyData) SetPageNo(v int64) *ListTaskDetailResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListTaskDetailResponseBodyData) SetPageSize(v int64) *ListTaskDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTaskDetailResponseBodyData) SetRecord(v []*ListTaskDetailResponseBodyDataRecord) *ListTaskDetailResponseBodyData {
	s.Record = v
	return s
}

func (s *ListTaskDetailResponseBodyData) SetTotal(v int64) *ListTaskDetailResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTaskDetailResponseBodyData) Validate() error {
	if s.Record != nil {
		for _, item := range s.Record {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskDetailResponseBodyDataRecord struct {
	// Called number.
	//
	// example:
	//
	// 186****0000
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
	// Calling number.
	//
	// example:
	//
	// 136****0000
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// Hang-up direction. Valid values:
	//
	// - **User**.
	//
	// - **Machine**.
	//
	// example:
	//
	// 用户
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Call duration. Unit: seconds.
	//
	// example:
	//
	// 30
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// End time.
	//
	// example:
	//
	// 2021-05-20 00:03:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Record ID.
	//
	// example:
	//
	// 12****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Current retry count.
	//
	// example:
	//
	// 1
	RetryCurTimes *int32 `json:"RetryCurTimes,omitempty" xml:"RetryCurTimes,omitempty"`
	// Retry Count.
	//
	// example:
	//
	// 1
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// Start Time.
	//
	// example:
	//
	// 2021-05-20 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Detail status.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Call status code.
	//
	// example:
	//
	// 200100
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// Call status code description.
	//
	// example:
	//
	// 呼叫成功
	StatusCodeDesc *string `json:"StatusCodeDesc,omitempty" xml:"StatusCodeDesc,omitempty"`
	// Intent tags.
	//
	// example:
	//
	// 肯定
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTaskDetailResponseBodyDataRecord) String() string {
	return dara.Prettify(s)
}

func (s ListTaskDetailResponseBodyDataRecord) GoString() string {
	return s.String()
}

func (s *ListTaskDetailResponseBodyDataRecord) GetCalled() *string {
	return s.Called
}

func (s *ListTaskDetailResponseBodyDataRecord) GetCaller() *string {
	return s.Caller
}

func (s *ListTaskDetailResponseBodyDataRecord) GetDirection() *string {
	return s.Direction
}

func (s *ListTaskDetailResponseBodyDataRecord) GetDuration() *int32 {
	return s.Duration
}

func (s *ListTaskDetailResponseBodyDataRecord) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTaskDetailResponseBodyDataRecord) GetId() *int64 {
	return s.Id
}

func (s *ListTaskDetailResponseBodyDataRecord) GetRetryCurTimes() *int32 {
	return s.RetryCurTimes
}

func (s *ListTaskDetailResponseBodyDataRecord) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *ListTaskDetailResponseBodyDataRecord) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTaskDetailResponseBodyDataRecord) GetStatus() *string {
	return s.Status
}

func (s *ListTaskDetailResponseBodyDataRecord) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListTaskDetailResponseBodyDataRecord) GetStatusCodeDesc() *string {
	return s.StatusCodeDesc
}

func (s *ListTaskDetailResponseBodyDataRecord) GetTags() *string {
	return s.Tags
}

func (s *ListTaskDetailResponseBodyDataRecord) SetCalled(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Called = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetCaller(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Caller = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetDirection(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Direction = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetDuration(v int32) *ListTaskDetailResponseBodyDataRecord {
	s.Duration = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetEndTime(v string) *ListTaskDetailResponseBodyDataRecord {
	s.EndTime = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetId(v int64) *ListTaskDetailResponseBodyDataRecord {
	s.Id = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetRetryCurTimes(v int32) *ListTaskDetailResponseBodyDataRecord {
	s.RetryCurTimes = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetRetryTimes(v int32) *ListTaskDetailResponseBodyDataRecord {
	s.RetryTimes = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetStartTime(v string) *ListTaskDetailResponseBodyDataRecord {
	s.StartTime = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetStatus(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Status = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetStatusCode(v string) *ListTaskDetailResponseBodyDataRecord {
	s.StatusCode = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetStatusCodeDesc(v string) *ListTaskDetailResponseBodyDataRecord {
	s.StatusCodeDesc = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) SetTags(v string) *ListTaskDetailResponseBodyDataRecord {
	s.Tags = &v
	return s
}

func (s *ListTaskDetailResponseBodyDataRecord) Validate() error {
	return dara.Validate(s)
}
