// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportResponseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetReportResponseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetReportResponseResponseBody
	GetErrorMessage() *string
	SetOutRequestNo(v string) *GetReportResponseResponseBody
	GetOutRequestNo() *string
	SetReportUrl(v string) *GetReportResponseResponseBody
	GetReportUrl() *string
}

type GetReportResponseResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	OutRequestNo *string `json:"outRequestNo,omitempty" xml:"outRequestNo,omitempty"`
	ReportUrl    *string `json:"reportUrl,omitempty" xml:"reportUrl,omitempty"`
}

func (s GetReportResponseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReportResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportResponseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetReportResponseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetReportResponseResponseBody) GetOutRequestNo() *string {
	return s.OutRequestNo
}

func (s *GetReportResponseResponseBody) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *GetReportResponseResponseBody) SetErrorCode(v string) *GetReportResponseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetReportResponseResponseBody) SetErrorMessage(v string) *GetReportResponseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetReportResponseResponseBody) SetOutRequestNo(v string) *GetReportResponseResponseBody {
	s.OutRequestNo = &v
	return s
}

func (s *GetReportResponseResponseBody) SetReportUrl(v string) *GetReportResponseResponseBody {
	s.ReportUrl = &v
	return s
}

func (s *GetReportResponseResponseBody) Validate() error {
	return dara.Validate(s)
}
