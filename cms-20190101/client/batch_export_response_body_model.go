// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchExportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnchor(v int64) *BatchExportResponseBody
	GetAnchor() *int64
	SetCode(v int32) *BatchExportResponseBody
	GetCode() *int32
	SetCursor(v string) *BatchExportResponseBody
	GetCursor() *string
	SetDataResults(v []*MetricStat) *BatchExportResponseBody
	GetDataResults() []*MetricStat
	SetHasNext(v bool) *BatchExportResponseBody
	GetHasNext() *bool
	SetLength(v int32) *BatchExportResponseBody
	GetLength() *int32
	SetMessage(v string) *BatchExportResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchExportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchExportResponseBody
	GetSuccess() *bool
}

type BatchExportResponseBody struct {
	// The timestamp of the data requested by the backend. A larger timestamp indicates that the data export time is closer to the current time.
	//
	// example:
	//
	// 1678781819000
	Anchor *int64 `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The Cursor information that is used to call this operation again.
	//
	// >  If `null` is returned, the monitoring data is exported.
	//
	// example:
	//
	// v2.5eyJidWNrZXRzIjo0LCJjdXJzb3IiOiIxNjQxNDU0ODAwMDAwMWUxY2YxNWY0NTU0MTliZjllYTY4OWQ2ODI1OTU1Yzc1NmZjMDQ2OTMxMzczMzM2MzUzMTMxMzEzMzM0MzMzODM5MzEzMTMwMjQyYzY5MmQzMjdhNjU2MjY3N2E2NjZhNzczOTY2NmM3Mjc0NjM3MzY5Njg3NDcyMjQyYyIsImN1cnNvclZlcnNpb24iOiJxdWVyeSIsImVuZFRpbWUiOjE2NDE0NTUyMzYxMTIsImV4cG9ydEVuZFRpbWUiOjE2NDE0NTUyMzYxMTIsImV4cG9ydFN0YXJ0VGltZSI6MTY0MTQ1NDYzNjExMiwiZXhwcmVzc1JhbmdlIjpmYWxzZSwiaGFzTmV4dCI6dHJ1ZSwiaW5wdXRNZXRyaWMiOiJDUFVVdGlsaXphdGlvbiIsImlucHV0TmFtZXNwYWNlIjoiYWNzX2Vjc19kYXNoYm9hcmQiLCJsaW1pdCI6MTAwMCwibG9nVGltZU1vZGUiOnRydWUsIm1hdGNoZXJzIjp7ImNoYWluIjpbeyJsYWJlbCI6InVzZXJJZCIsIm9wZXJhdG9yIjoiRVFVQUxTIiwidmFsdWUiOiIxNzM2NTExMTM0Mzg5MTEwIn1dfSwibWV0cmljIjoiQ1BVVXRpbGl6YXRpb24iLCJtZXRyaWNUeXBlIjoiTUVUUklDIiwibmFtZXNwYWNlIjoiYWNzX2Vjc19kYXNoYm9hcmQiLCJuZXh0UGtBZGFwdGVyIjp7fSwib2Zmc2V0IjowLCJwYXJlbnRVaWQiOjEyNzA2NzY2Nzk1NDY3MDQsInN0YXJ0VGltZSI6MTY0MTQ1NDYzNjExMiwic3RlcCI6LTEsInRpbWVvdXQiOjEyMCwid2luZG93Ijo2MH0***
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// The data returned in this call.
	DataResults []*MetricStat `json:"DataResults,omitempty" xml:"DataResults,omitempty" type:"Repeated"`
	// Indicates whether the data has been exported. Valid values:
	//
	// 	- true: Some data is not exported.
	//
	// 	- false: All the data is exported.
	//
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// The number of data entries returned in this call.
	//
	// example:
	//
	// 1000
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 251402CD-305C-1617-808E-D8C11FC8138D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchExportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchExportResponseBody) GoString() string {
	return s.String()
}

func (s *BatchExportResponseBody) GetAnchor() *int64 {
	return s.Anchor
}

func (s *BatchExportResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BatchExportResponseBody) GetCursor() *string {
	return s.Cursor
}

func (s *BatchExportResponseBody) GetDataResults() []*MetricStat {
	return s.DataResults
}

func (s *BatchExportResponseBody) GetHasNext() *bool {
	return s.HasNext
}

func (s *BatchExportResponseBody) GetLength() *int32 {
	return s.Length
}

func (s *BatchExportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchExportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchExportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchExportResponseBody) SetAnchor(v int64) *BatchExportResponseBody {
	s.Anchor = &v
	return s
}

func (s *BatchExportResponseBody) SetCode(v int32) *BatchExportResponseBody {
	s.Code = &v
	return s
}

func (s *BatchExportResponseBody) SetCursor(v string) *BatchExportResponseBody {
	s.Cursor = &v
	return s
}

func (s *BatchExportResponseBody) SetDataResults(v []*MetricStat) *BatchExportResponseBody {
	s.DataResults = v
	return s
}

func (s *BatchExportResponseBody) SetHasNext(v bool) *BatchExportResponseBody {
	s.HasNext = &v
	return s
}

func (s *BatchExportResponseBody) SetLength(v int32) *BatchExportResponseBody {
	s.Length = &v
	return s
}

func (s *BatchExportResponseBody) SetMessage(v string) *BatchExportResponseBody {
	s.Message = &v
	return s
}

func (s *BatchExportResponseBody) SetRequestId(v string) *BatchExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchExportResponseBody) SetSuccess(v bool) *BatchExportResponseBody {
	s.Success = &v
	return s
}

func (s *BatchExportResponseBody) Validate() error {
	return dara.Validate(s)
}
