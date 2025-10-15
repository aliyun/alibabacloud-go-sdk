// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationClientSecretsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationClientSecretsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationClientSecretsResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationClientSecretsResponseBody) *ListApplicationClientSecretsResponse
	GetBody() *ListApplicationClientSecretsResponseBody
}

type ListApplicationClientSecretsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationClientSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationClientSecretsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationClientSecretsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationClientSecretsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationClientSecretsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationClientSecretsResponse) GetBody() *ListApplicationClientSecretsResponseBody {
	return s.Body
}

func (s *ListApplicationClientSecretsResponse) SetHeaders(v map[string]*string) *ListApplicationClientSecretsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationClientSecretsResponse) SetStatusCode(v int32) *ListApplicationClientSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationClientSecretsResponse) SetBody(v *ListApplicationClientSecretsResponseBody) *ListApplicationClientSecretsResponse {
	s.Body = v
	return s
}

func (s *ListApplicationClientSecretsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
