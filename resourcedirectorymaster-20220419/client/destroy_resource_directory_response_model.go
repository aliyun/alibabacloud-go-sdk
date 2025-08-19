// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyResourceDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DestroyResourceDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DestroyResourceDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *DestroyResourceDirectoryResponseBody) *DestroyResourceDirectoryResponse
	GetBody() *DestroyResourceDirectoryResponseBody
}

type DestroyResourceDirectoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DestroyResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DestroyResourceDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DestroyResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DestroyResourceDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DestroyResourceDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DestroyResourceDirectoryResponse) GetBody() *DestroyResourceDirectoryResponseBody {
	return s.Body
}

func (s *DestroyResourceDirectoryResponse) SetHeaders(v map[string]*string) *DestroyResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DestroyResourceDirectoryResponse) SetStatusCode(v int32) *DestroyResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DestroyResourceDirectoryResponse) SetBody(v *DestroyResourceDirectoryResponseBody) *DestroyResourceDirectoryResponse {
	s.Body = v
	return s
}

func (s *DestroyResourceDirectoryResponse) Validate() error {
	return dara.Validate(s)
}
