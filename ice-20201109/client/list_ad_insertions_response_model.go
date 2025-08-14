// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdInsertionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAdInsertionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAdInsertionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAdInsertionsResponseBody) *ListAdInsertionsResponse
	GetBody() *ListAdInsertionsResponseBody
}

type ListAdInsertionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAdInsertionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAdInsertionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAdInsertionsResponse) GoString() string {
	return s.String()
}

func (s *ListAdInsertionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAdInsertionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAdInsertionsResponse) GetBody() *ListAdInsertionsResponseBody {
	return s.Body
}

func (s *ListAdInsertionsResponse) SetHeaders(v map[string]*string) *ListAdInsertionsResponse {
	s.Headers = v
	return s
}

func (s *ListAdInsertionsResponse) SetStatusCode(v int32) *ListAdInsertionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdInsertionsResponse) SetBody(v *ListAdInsertionsResponseBody) *ListAdInsertionsResponse {
	s.Body = v
	return s
}

func (s *ListAdInsertionsResponse) Validate() error {
	return dara.Validate(s)
}
