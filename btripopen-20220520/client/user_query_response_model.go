// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UserQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UserQueryResponse
	GetStatusCode() *int32
	SetBody(v *UserQueryResponseBody) *UserQueryResponse
	GetBody() *UserQueryResponseBody
}

type UserQueryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UserQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UserQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s UserQueryResponse) GoString() string {
	return s.String()
}

func (s *UserQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UserQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UserQueryResponse) GetBody() *UserQueryResponseBody {
	return s.Body
}

func (s *UserQueryResponse) SetHeaders(v map[string]*string) *UserQueryResponse {
	s.Headers = v
	return s
}

func (s *UserQueryResponse) SetStatusCode(v int32) *UserQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *UserQueryResponse) SetBody(v *UserQueryResponseBody) *UserQueryResponse {
	s.Body = v
	return s
}

func (s *UserQueryResponse) Validate() error {
	return dara.Validate(s)
}
