// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslateReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranslateReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranslateReportResponse
	GetStatusCode() *int32
	SetBody(v *GetTranslateReportResponseBody) *GetTranslateReportResponse
	GetBody() *GetTranslateReportResponseBody
}

type GetTranslateReportResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranslateReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranslateReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranslateReportResponse) GoString() string {
	return s.String()
}

func (s *GetTranslateReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranslateReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranslateReportResponse) GetBody() *GetTranslateReportResponseBody {
	return s.Body
}

func (s *GetTranslateReportResponse) SetHeaders(v map[string]*string) *GetTranslateReportResponse {
	s.Headers = v
	return s
}

func (s *GetTranslateReportResponse) SetStatusCode(v int32) *GetTranslateReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranslateReportResponse) SetBody(v *GetTranslateReportResponseBody) *GetTranslateReportResponse {
	s.Body = v
	return s
}

func (s *GetTranslateReportResponse) Validate() error {
	return dara.Validate(s)
}
