// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportBizAlertInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportBizAlertInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportBizAlertInfoResponse
	GetStatusCode() *int32
	SetBody(v *ReportBizAlertInfoResponseBody) *ReportBizAlertInfoResponse
	GetBody() *ReportBizAlertInfoResponseBody
}

type ReportBizAlertInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportBizAlertInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportBizAlertInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportBizAlertInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportBizAlertInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportBizAlertInfoResponse) GetBody() *ReportBizAlertInfoResponseBody {
	return s.Body
}

func (s *ReportBizAlertInfoResponse) SetHeaders(v map[string]*string) *ReportBizAlertInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportBizAlertInfoResponse) SetStatusCode(v int32) *ReportBizAlertInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportBizAlertInfoResponse) SetBody(v *ReportBizAlertInfoResponseBody) *ReportBizAlertInfoResponse {
	s.Body = v
	return s
}

func (s *ReportBizAlertInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
