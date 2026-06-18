// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineRecordDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHotlineRecordDetailResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListHotlineRecordDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHotlineRecordDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotlineRecordDetailResponseBody
	GetRequestId() *string
	SetResultData(v *ListHotlineRecordDetailResponseBodyResultData) *ListHotlineRecordDetailResponseBody
	GetResultData() *ListHotlineRecordDetailResponseBodyResultData
	SetSuccess(v bool) *ListHotlineRecordDetailResponseBody
	GetSuccess() *bool
}

type ListHotlineRecordDetailResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned data.
	ResultData *ListHotlineRecordDetailResponseBodyResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Struct"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotlineRecordDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHotlineRecordDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHotlineRecordDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotlineRecordDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotlineRecordDetailResponseBody) GetResultData() *ListHotlineRecordDetailResponseBodyResultData {
	return s.ResultData
}

func (s *ListHotlineRecordDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHotlineRecordDetailResponseBody) SetCode(v string) *ListHotlineRecordDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetHttpStatusCode(v int32) *ListHotlineRecordDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetMessage(v string) *ListHotlineRecordDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetRequestId(v string) *ListHotlineRecordDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetResultData(v *ListHotlineRecordDetailResponseBodyResultData) *ListHotlineRecordDetailResponseBody {
	s.ResultData = v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) SetSuccess(v bool) *ListHotlineRecordDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBody) Validate() error {
	if s.ResultData != nil {
		if err := s.ResultData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotlineRecordDetailResponseBodyResultData struct {
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Query result data.
	Data []*ListHotlineRecordDetailResponseBodyResultDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Page size.
	//
	// example:
	//
	// 100
	OnePageSize *int64 `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 10
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 945
	TotalResults *int64 `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s ListHotlineRecordDetailResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordDetailResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailResponseBodyResultData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListHotlineRecordDetailResponseBodyResultData) GetData() []*ListHotlineRecordDetailResponseBodyResultDataData {
	return s.Data
}

func (s *ListHotlineRecordDetailResponseBodyResultData) GetOnePageSize() *int64 {
	return s.OnePageSize
}

func (s *ListHotlineRecordDetailResponseBodyResultData) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListHotlineRecordDetailResponseBodyResultData) GetTotalResults() *int64 {
	return s.TotalResults
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetCurrentPage(v int64) *ListHotlineRecordDetailResponseBodyResultData {
	s.CurrentPage = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetData(v []*ListHotlineRecordDetailResponseBodyResultDataData) *ListHotlineRecordDetailResponseBodyResultData {
	s.Data = v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetOnePageSize(v int64) *ListHotlineRecordDetailResponseBodyResultData {
	s.OnePageSize = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetTotalPage(v int64) *ListHotlineRecordDetailResponseBodyResultData {
	s.TotalPage = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) SetTotalResults(v int64) *ListHotlineRecordDetailResponseBodyResultData {
	s.TotalResults = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotlineRecordDetailResponseBodyResultDataData struct {
	// Hotline End Time.
	//
	// example:
	//
	// 1614578410000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// OSS URL of the hotline call recording.
	//
	// example:
	//
	// http://xxx.xxxxx/xx.wav
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// Servicer Account.
	//
	// example:
	//
	// 123@123.com
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	// Start Time of the hotline call.
	//
	// example:
	//
	// 1614578400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHotlineRecordDetailResponseBodyResultDataData) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordDetailResponseBodyResultDataData) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) GetOssUrl() *string {
	return s.OssUrl
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) GetServicerName() *string {
	return s.ServicerName
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) SetEndTime(v int64) *ListHotlineRecordDetailResponseBodyResultDataData {
	s.EndTime = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) SetOssUrl(v string) *ListHotlineRecordDetailResponseBodyResultDataData {
	s.OssUrl = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) SetServicerName(v string) *ListHotlineRecordDetailResponseBodyResultDataData {
	s.ServicerName = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) SetStartTime(v int64) *ListHotlineRecordDetailResponseBodyResultDataData {
	s.StartTime = &v
	return s
}

func (s *ListHotlineRecordDetailResponseBodyResultDataData) Validate() error {
	return dara.Validate(s)
}
