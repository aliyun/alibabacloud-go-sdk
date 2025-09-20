// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileCompressionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileCompressionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileCompressionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileCompressionTaskResponseBody) *CreateFileCompressionTaskResponse
	GetBody() *CreateFileCompressionTaskResponseBody
}

type CreateFileCompressionTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileCompressionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileCompressionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileCompressionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileCompressionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileCompressionTaskResponse) GetBody() *CreateFileCompressionTaskResponseBody {
	return s.Body
}

func (s *CreateFileCompressionTaskResponse) SetHeaders(v map[string]*string) *CreateFileCompressionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFileCompressionTaskResponse) SetStatusCode(v int32) *CreateFileCompressionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileCompressionTaskResponse) SetBody(v *CreateFileCompressionTaskResponseBody) *CreateFileCompressionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateFileCompressionTaskResponse) Validate() error {
	return dara.Validate(s)
}
