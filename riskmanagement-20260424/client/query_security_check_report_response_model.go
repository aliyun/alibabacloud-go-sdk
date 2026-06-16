// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecurityCheckReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySecurityCheckReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySecurityCheckReportResponse
	GetStatusCode() *int32
	SetBody(v *QuerySecurityCheckReportResponseBody) *QuerySecurityCheckReportResponse
	GetBody() *QuerySecurityCheckReportResponseBody
}

type QuerySecurityCheckReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecurityCheckReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecurityCheckReportResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySecurityCheckReportResponse) GoString() string {
	return s.String()
}

func (s *QuerySecurityCheckReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySecurityCheckReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySecurityCheckReportResponse) GetBody() *QuerySecurityCheckReportResponseBody {
	return s.Body
}

func (s *QuerySecurityCheckReportResponse) SetHeaders(v map[string]*string) *QuerySecurityCheckReportResponse {
	s.Headers = v
	return s
}

func (s *QuerySecurityCheckReportResponse) SetStatusCode(v int32) *QuerySecurityCheckReportResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecurityCheckReportResponse) SetBody(v *QuerySecurityCheckReportResponseBody) *QuerySecurityCheckReportResponse {
	s.Body = v
	return s
}

func (s *QuerySecurityCheckReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
