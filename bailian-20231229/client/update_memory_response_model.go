// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMemoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMemoryResponseBody) *UpdateMemoryResponse
	GetBody() *UpdateMemoryResponseBody
}

type UpdateMemoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMemoryResponse) GetBody() *UpdateMemoryResponseBody {
	return s.Body
}

func (s *UpdateMemoryResponse) SetHeaders(v map[string]*string) *UpdateMemoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemoryResponse) SetStatusCode(v int32) *UpdateMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemoryResponse) SetBody(v *UpdateMemoryResponseBody) *UpdateMemoryResponse {
	s.Body = v
	return s
}

func (s *UpdateMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
