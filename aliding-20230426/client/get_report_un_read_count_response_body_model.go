// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportUnReadCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *GetReportUnReadCountResponseBody
	GetCount() *int64
	SetRequestId(v string) *GetReportUnReadCountResponseBody
	GetRequestId() *string
}

type GetReportUnReadCountResponseBody struct {
	// example:
	//
	// 1
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetReportUnReadCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReportUnReadCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportUnReadCountResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *GetReportUnReadCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetReportUnReadCountResponseBody) SetCount(v int64) *GetReportUnReadCountResponseBody {
	s.Count = &v
	return s
}

func (s *GetReportUnReadCountResponseBody) SetRequestId(v string) *GetReportUnReadCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReportUnReadCountResponseBody) Validate() error {
	return dara.Validate(s)
}
