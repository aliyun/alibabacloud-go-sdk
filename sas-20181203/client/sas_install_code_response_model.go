// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSasInstallCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SasInstallCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SasInstallCodeResponse
	GetStatusCode() *int32
	SetBody(v *SasInstallCodeResponseBody) *SasInstallCodeResponse
	GetBody() *SasInstallCodeResponseBody
}

type SasInstallCodeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SasInstallCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SasInstallCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SasInstallCodeResponse) GoString() string {
	return s.String()
}

func (s *SasInstallCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SasInstallCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SasInstallCodeResponse) GetBody() *SasInstallCodeResponseBody {
	return s.Body
}

func (s *SasInstallCodeResponse) SetHeaders(v map[string]*string) *SasInstallCodeResponse {
	s.Headers = v
	return s
}

func (s *SasInstallCodeResponse) SetStatusCode(v int32) *SasInstallCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SasInstallCodeResponse) SetBody(v *SasInstallCodeResponseBody) *SasInstallCodeResponse {
	s.Body = v
	return s
}

func (s *SasInstallCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
