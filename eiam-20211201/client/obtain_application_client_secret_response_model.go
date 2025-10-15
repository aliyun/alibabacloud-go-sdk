// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainApplicationClientSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ObtainApplicationClientSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ObtainApplicationClientSecretResponse
	GetStatusCode() *int32
	SetBody(v *ObtainApplicationClientSecretResponseBody) *ObtainApplicationClientSecretResponse
	GetBody() *ObtainApplicationClientSecretResponseBody
}

type ObtainApplicationClientSecretResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainApplicationClientSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s ObtainApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *ObtainApplicationClientSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ObtainApplicationClientSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ObtainApplicationClientSecretResponse) GetBody() *ObtainApplicationClientSecretResponseBody {
	return s.Body
}

func (s *ObtainApplicationClientSecretResponse) SetHeaders(v map[string]*string) *ObtainApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *ObtainApplicationClientSecretResponse) SetStatusCode(v int32) *ObtainApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainApplicationClientSecretResponse) SetBody(v *ObtainApplicationClientSecretResponseBody) *ObtainApplicationClientSecretResponse {
	s.Body = v
	return s
}

func (s *ObtainApplicationClientSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
