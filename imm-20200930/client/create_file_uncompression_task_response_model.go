// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileUncompressionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileUncompressionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileUncompressionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileUncompressionTaskResponseBody) *CreateFileUncompressionTaskResponse
	GetBody() *CreateFileUncompressionTaskResponseBody
}

type CreateFileUncompressionTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileUncompressionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileUncompressionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileUncompressionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileUncompressionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileUncompressionTaskResponse) GetBody() *CreateFileUncompressionTaskResponseBody {
	return s.Body
}

func (s *CreateFileUncompressionTaskResponse) SetHeaders(v map[string]*string) *CreateFileUncompressionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFileUncompressionTaskResponse) SetStatusCode(v int32) *CreateFileUncompressionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileUncompressionTaskResponse) SetBody(v *CreateFileUncompressionTaskResponseBody) *CreateFileUncompressionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateFileUncompressionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
