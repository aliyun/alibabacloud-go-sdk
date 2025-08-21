// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *UnbindDirectoryResponseBody) *UnbindDirectoryResponse
	GetBody() *UnbindDirectoryResponseBody
}

type UnbindDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UnbindDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindDirectoryResponse) GetBody() *UnbindDirectoryResponseBody {
	return s.Body
}

func (s *UnbindDirectoryResponse) SetHeaders(v map[string]*string) *UnbindDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UnbindDirectoryResponse) SetStatusCode(v int32) *UnbindDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDirectoryResponse) SetBody(v *UnbindDirectoryResponseBody) *UnbindDirectoryResponse {
	s.Body = v
	return s
}

func (s *UnbindDirectoryResponse) Validate() error {
	return dara.Validate(s)
}
