// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCardSmsTemplateReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCardSmsTemplateReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCardSmsTemplateReportResponse
	GetStatusCode() *int32
	SetBody(v *QueryCardSmsTemplateReportResponseBody) *QueryCardSmsTemplateReportResponse
	GetBody() *QueryCardSmsTemplateReportResponseBody
}

type QueryCardSmsTemplateReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCardSmsTemplateReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCardSmsTemplateReportResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCardSmsTemplateReportResponse) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCardSmsTemplateReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCardSmsTemplateReportResponse) GetBody() *QueryCardSmsTemplateReportResponseBody {
	return s.Body
}

func (s *QueryCardSmsTemplateReportResponse) SetHeaders(v map[string]*string) *QueryCardSmsTemplateReportResponse {
	s.Headers = v
	return s
}

func (s *QueryCardSmsTemplateReportResponse) SetStatusCode(v int32) *QueryCardSmsTemplateReportResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardSmsTemplateReportResponse) SetBody(v *QueryCardSmsTemplateReportResponseBody) *QueryCardSmsTemplateReportResponse {
	s.Body = v
	return s
}

func (s *QueryCardSmsTemplateReportResponse) Validate() error {
	return dara.Validate(s)
}
