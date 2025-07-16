// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslateReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTranslateReportResponseBody
	GetCode() *int32
	SetData(v string) *GetTranslateReportResponseBody
	GetData() *string
	SetMessage(v string) *GetTranslateReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTranslateReportResponseBody
	GetRequestId() *string
}

type GetTranslateReportResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"result":[{"time":"2021-10-21 00:00:00","total":100}]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTranslateReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranslateReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslateReportResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTranslateReportResponseBody) GetData() *string {
	return s.Data
}

func (s *GetTranslateReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTranslateReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranslateReportResponseBody) SetCode(v int32) *GetTranslateReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranslateReportResponseBody) SetData(v string) *GetTranslateReportResponseBody {
	s.Data = &v
	return s
}

func (s *GetTranslateReportResponseBody) SetMessage(v string) *GetTranslateReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranslateReportResponseBody) SetRequestId(v string) *GetTranslateReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranslateReportResponseBody) Validate() error {
	return dara.Validate(s)
}
