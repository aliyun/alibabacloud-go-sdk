// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCredentialReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCredentialReportResponse
	GetStatusCode() *int32
	SetBody(v *GetCredentialReportResponseBody) *GetCredentialReportResponse
	GetBody() *GetCredentialReportResponseBody
}

type GetCredentialReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCredentialReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCredentialReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialReportResponse) GoString() string {
	return s.String()
}

func (s *GetCredentialReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCredentialReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCredentialReportResponse) GetBody() *GetCredentialReportResponseBody {
	return s.Body
}

func (s *GetCredentialReportResponse) SetHeaders(v map[string]*string) *GetCredentialReportResponse {
	s.Headers = v
	return s
}

func (s *GetCredentialReportResponse) SetStatusCode(v int32) *GetCredentialReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCredentialReportResponse) SetBody(v *GetCredentialReportResponseBody) *GetCredentialReportResponse {
	s.Body = v
	return s
}

func (s *GetCredentialReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
