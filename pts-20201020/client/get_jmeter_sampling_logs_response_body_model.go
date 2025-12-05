// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSamplingLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetJMeterSamplingLogsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetJMeterSamplingLogsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetJMeterSamplingLogsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetJMeterSamplingLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetJMeterSamplingLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetJMeterSamplingLogsResponseBody
	GetRequestId() *string
	SetSampleResults(v []*string) *GetJMeterSamplingLogsResponseBody
	GetSampleResults() []*string
	SetSuccess(v bool) *GetJMeterSamplingLogsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetJMeterSamplingLogsResponseBody
	GetTotalCount() *int64
}

type GetJMeterSamplingLogsResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sampling results of the sampler.
	SampleResults []*string `json:"SampleResults,omitempty" xml:"SampleResults,omitempty" type:"Repeated"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of log entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetJMeterSamplingLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSamplingLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJMeterSamplingLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetJMeterSamplingLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetJMeterSamplingLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJMeterSamplingLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetJMeterSamplingLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetJMeterSamplingLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJMeterSamplingLogsResponseBody) GetSampleResults() []*string {
	return s.SampleResults
}

func (s *GetJMeterSamplingLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJMeterSamplingLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetJMeterSamplingLogsResponseBody) SetCode(v string) *GetJMeterSamplingLogsResponseBody {
	s.Code = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetHttpStatusCode(v int32) *GetJMeterSamplingLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetMessage(v string) *GetJMeterSamplingLogsResponseBody {
	s.Message = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetPageNumber(v int32) *GetJMeterSamplingLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetPageSize(v int32) *GetJMeterSamplingLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetRequestId(v string) *GetJMeterSamplingLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetSampleResults(v []*string) *GetJMeterSamplingLogsResponseBody {
	s.SampleResults = v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetSuccess(v bool) *GetJMeterSamplingLogsResponseBody {
	s.Success = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetTotalCount(v int64) *GetJMeterSamplingLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
