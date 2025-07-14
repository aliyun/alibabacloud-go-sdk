// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecretsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecretsResponse
	GetStatusCode() *int32
	SetBody(v *ListSecretsResponseBody) *ListSecretsResponse
	GetBody() *ListSecretsResponseBody
}

type ListSecretsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecretsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecretsResponse) GetBody() *ListSecretsResponseBody {
	return s.Body
}

func (s *ListSecretsResponse) SetHeaders(v map[string]*string) *ListSecretsResponse {
	s.Headers = v
	return s
}

func (s *ListSecretsResponse) SetStatusCode(v int32) *ListSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretsResponse) SetBody(v *ListSecretsResponseBody) *ListSecretsResponse {
	s.Body = v
	return s
}

func (s *ListSecretsResponse) Validate() error {
	return dara.Validate(s)
}
