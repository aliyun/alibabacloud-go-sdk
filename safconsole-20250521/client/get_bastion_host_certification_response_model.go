// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBastionHostCertificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBastionHostCertificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBastionHostCertificationResponse
	GetStatusCode() *int32
	SetBody(v *GetBastionHostCertificationResponseBody) *GetBastionHostCertificationResponse
	GetBody() *GetBastionHostCertificationResponseBody
}

type GetBastionHostCertificationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBastionHostCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBastionHostCertificationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBastionHostCertificationResponse) GoString() string {
	return s.String()
}

func (s *GetBastionHostCertificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBastionHostCertificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBastionHostCertificationResponse) GetBody() *GetBastionHostCertificationResponseBody {
	return s.Body
}

func (s *GetBastionHostCertificationResponse) SetHeaders(v map[string]*string) *GetBastionHostCertificationResponse {
	s.Headers = v
	return s
}

func (s *GetBastionHostCertificationResponse) SetStatusCode(v int32) *GetBastionHostCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBastionHostCertificationResponse) SetBody(v *GetBastionHostCertificationResponseBody) *GetBastionHostCertificationResponse {
	s.Body = v
	return s
}

func (s *GetBastionHostCertificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
