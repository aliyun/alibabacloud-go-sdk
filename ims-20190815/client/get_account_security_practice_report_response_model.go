// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountSecurityPracticeReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountSecurityPracticeReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountSecurityPracticeReportResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountSecurityPracticeReportResponseBody) *GetAccountSecurityPracticeReportResponse
	GetBody() *GetAccountSecurityPracticeReportResponseBody
}

type GetAccountSecurityPracticeReportResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountSecurityPracticeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountSecurityPracticeReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponse) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountSecurityPracticeReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountSecurityPracticeReportResponse) GetBody() *GetAccountSecurityPracticeReportResponseBody {
	return s.Body
}

func (s *GetAccountSecurityPracticeReportResponse) SetHeaders(v map[string]*string) *GetAccountSecurityPracticeReportResponse {
	s.Headers = v
	return s
}

func (s *GetAccountSecurityPracticeReportResponse) SetStatusCode(v int32) *GetAccountSecurityPracticeReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponse) SetBody(v *GetAccountSecurityPracticeReportResponseBody) *GetAccountSecurityPracticeReportResponse {
	s.Body = v
	return s
}

func (s *GetAccountSecurityPracticeReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
