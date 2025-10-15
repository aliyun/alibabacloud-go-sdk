// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationClientSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationClientSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationClientSecretResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationClientSecretResponseBody) *DisableApplicationClientSecretResponse
	GetBody() *DisableApplicationClientSecretResponseBody
}

type DisableApplicationClientSecretResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationClientSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationClientSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationClientSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationClientSecretResponse) GetBody() *DisableApplicationClientSecretResponseBody {
	return s.Body
}

func (s *DisableApplicationClientSecretResponse) SetHeaders(v map[string]*string) *DisableApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationClientSecretResponse) SetStatusCode(v int32) *DisableApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationClientSecretResponse) SetBody(v *DisableApplicationClientSecretResponseBody) *DisableApplicationClientSecretResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationClientSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
