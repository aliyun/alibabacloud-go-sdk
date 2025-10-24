// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSM2CertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSM2CertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSM2CertResponse
	GetStatusCode() *int32
	SetBody(v *CreateSM2CertResponseBody) *CreateSM2CertResponse
	GetBody() *CreateSM2CertResponseBody
}

type CreateSM2CertResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSM2CertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSM2CertResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSM2CertResponse) GoString() string {
	return s.String()
}

func (s *CreateSM2CertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSM2CertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSM2CertResponse) GetBody() *CreateSM2CertResponseBody {
	return s.Body
}

func (s *CreateSM2CertResponse) SetHeaders(v map[string]*string) *CreateSM2CertResponse {
	s.Headers = v
	return s
}

func (s *CreateSM2CertResponse) SetStatusCode(v int32) *CreateSM2CertResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSM2CertResponse) SetBody(v *CreateSM2CertResponseBody) *CreateSM2CertResponse {
	s.Body = v
	return s
}

func (s *CreateSM2CertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
