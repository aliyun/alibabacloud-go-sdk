// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretVersionIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecretVersionIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecretVersionIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListSecretVersionIdsResponseBody) *ListSecretVersionIdsResponse
	GetBody() *ListSecretVersionIdsResponseBody
}

type ListSecretVersionIdsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretVersionIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretVersionIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecretVersionIdsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecretVersionIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecretVersionIdsResponse) GetBody() *ListSecretVersionIdsResponseBody {
	return s.Body
}

func (s *ListSecretVersionIdsResponse) SetHeaders(v map[string]*string) *ListSecretVersionIdsResponse {
	s.Headers = v
	return s
}

func (s *ListSecretVersionIdsResponse) SetStatusCode(v int32) *ListSecretVersionIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretVersionIdsResponse) SetBody(v *ListSecretVersionIdsResponseBody) *ListSecretVersionIdsResponse {
	s.Body = v
	return s
}

func (s *ListSecretVersionIdsResponse) Validate() error {
	return dara.Validate(s)
}
