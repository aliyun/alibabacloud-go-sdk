// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBindAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBindAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListBindAccountResponseBody) *ListBindAccountResponse
	GetBody() *ListBindAccountResponseBody
}

type ListBindAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBindAccountResponse) GoString() string {
	return s.String()
}

func (s *ListBindAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBindAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBindAccountResponse) GetBody() *ListBindAccountResponseBody {
	return s.Body
}

func (s *ListBindAccountResponse) SetHeaders(v map[string]*string) *ListBindAccountResponse {
	s.Headers = v
	return s
}

func (s *ListBindAccountResponse) SetStatusCode(v int32) *ListBindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindAccountResponse) SetBody(v *ListBindAccountResponseBody) *ListBindAccountResponse {
	s.Body = v
	return s
}

func (s *ListBindAccountResponse) Validate() error {
	return dara.Validate(s)
}
