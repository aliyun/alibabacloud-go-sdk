// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppSecretIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppSecretIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppSecretIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppSecretIdsResponseBody) *ListAppSecretIdsResponse
	GetBody() *ListAppSecretIdsResponseBody
}

type ListAppSecretIdsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppSecretIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppSecretIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppSecretIdsResponse) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppSecretIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppSecretIdsResponse) GetBody() *ListAppSecretIdsResponseBody {
	return s.Body
}

func (s *ListAppSecretIdsResponse) SetHeaders(v map[string]*string) *ListAppSecretIdsResponse {
	s.Headers = v
	return s
}

func (s *ListAppSecretIdsResponse) SetStatusCode(v int32) *ListAppSecretIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppSecretIdsResponse) SetBody(v *ListAppSecretIdsResponseBody) *ListAppSecretIdsResponse {
	s.Body = v
	return s
}

func (s *ListAppSecretIdsResponse) Validate() error {
	return dara.Validate(s)
}
