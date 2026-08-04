// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealNameCertificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRealNameCertificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRealNameCertificationResponse
	GetStatusCode() *int32
	SetBody(v *CreateRealNameCertificationResponseBody) *CreateRealNameCertificationResponse
	GetBody() *CreateRealNameCertificationResponseBody
}

type CreateRealNameCertificationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRealNameCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRealNameCertificationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRealNameCertificationResponse) GoString() string {
	return s.String()
}

func (s *CreateRealNameCertificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRealNameCertificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRealNameCertificationResponse) GetBody() *CreateRealNameCertificationResponseBody {
	return s.Body
}

func (s *CreateRealNameCertificationResponse) SetHeaders(v map[string]*string) *CreateRealNameCertificationResponse {
	s.Headers = v
	return s
}

func (s *CreateRealNameCertificationResponse) SetStatusCode(v int32) *CreateRealNameCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRealNameCertificationResponse) SetBody(v *CreateRealNameCertificationResponseBody) *CreateRealNameCertificationResponse {
	s.Body = v
	return s
}

func (s *CreateRealNameCertificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
