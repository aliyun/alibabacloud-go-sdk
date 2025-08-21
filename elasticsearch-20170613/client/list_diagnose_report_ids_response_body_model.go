// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseReportIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListDiagnoseReportIdsResponseBodyHeaders) *ListDiagnoseReportIdsResponseBody
	GetHeaders() *ListDiagnoseReportIdsResponseBodyHeaders
	SetRequestId(v string) *ListDiagnoseReportIdsResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ListDiagnoseReportIdsResponseBody
	GetResult() []*string
}

type ListDiagnoseReportIdsResponseBody struct {
	// The header of the response.
	Headers *ListDiagnoseReportIdsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiagnoseReportIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportIdsResponseBody) GetHeaders() *ListDiagnoseReportIdsResponseBodyHeaders {
	return s.Headers
}

func (s *ListDiagnoseReportIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnoseReportIdsResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ListDiagnoseReportIdsResponseBody) SetHeaders(v *ListDiagnoseReportIdsResponseBodyHeaders) *ListDiagnoseReportIdsResponseBody {
	s.Headers = v
	return s
}

func (s *ListDiagnoseReportIdsResponseBody) SetRequestId(v string) *ListDiagnoseReportIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnoseReportIdsResponseBody) SetResult(v []*string) *ListDiagnoseReportIdsResponseBody {
	s.Result = v
	return s
}

func (s *ListDiagnoseReportIdsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDiagnoseReportIdsResponseBodyHeaders struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListDiagnoseReportIdsResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportIdsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportIdsResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListDiagnoseReportIdsResponseBodyHeaders) SetXTotalCount(v int32) *ListDiagnoseReportIdsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListDiagnoseReportIdsResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}
