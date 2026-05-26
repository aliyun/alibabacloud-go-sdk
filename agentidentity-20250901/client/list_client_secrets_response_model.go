// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientSecretsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientSecretsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientSecretsResponse
	GetStatusCode() *int32
	SetBody(v *ListClientSecretsResponseBody) *ListClientSecretsResponse
	GetBody() *ListClientSecretsResponseBody
}

type ListClientSecretsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientSecretsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientSecretsResponse) GoString() string {
	return s.String()
}

func (s *ListClientSecretsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientSecretsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientSecretsResponse) GetBody() *ListClientSecretsResponseBody {
	return s.Body
}

func (s *ListClientSecretsResponse) SetHeaders(v map[string]*string) *ListClientSecretsResponse {
	s.Headers = v
	return s
}

func (s *ListClientSecretsResponse) SetStatusCode(v int32) *ListClientSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientSecretsResponse) SetBody(v *ListClientSecretsResponseBody) *ListClientSecretsResponse {
	s.Body = v
	return s
}

func (s *ListClientSecretsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
