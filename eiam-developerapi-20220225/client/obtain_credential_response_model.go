// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ObtainCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ObtainCredentialResponse
	GetStatusCode() *int32
	SetBody(v *ObtainCredentialResponseBody) *ObtainCredentialResponse
	GetBody() *ObtainCredentialResponseBody
}

type ObtainCredentialResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialResponse) GoString() string {
	return s.String()
}

func (s *ObtainCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ObtainCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ObtainCredentialResponse) GetBody() *ObtainCredentialResponseBody {
	return s.Body
}

func (s *ObtainCredentialResponse) SetHeaders(v map[string]*string) *ObtainCredentialResponse {
	s.Headers = v
	return s
}

func (s *ObtainCredentialResponse) SetStatusCode(v int32) *ObtainCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainCredentialResponse) SetBody(v *ObtainCredentialResponseBody) *ObtainCredentialResponse {
	s.Body = v
	return s
}

func (s *ObtainCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
