// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateMemoryResponseBody) *CreateMemoryResponse
	GetBody() *CreateMemoryResponseBody
}

type CreateMemoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponse) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemoryResponse) GetBody() *CreateMemoryResponseBody {
	return s.Body
}

func (s *CreateMemoryResponse) SetHeaders(v map[string]*string) *CreateMemoryResponse {
	s.Headers = v
	return s
}

func (s *CreateMemoryResponse) SetStatusCode(v int32) *CreateMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemoryResponse) SetBody(v *CreateMemoryResponseBody) *CreateMemoryResponse {
	s.Body = v
	return s
}

func (s *CreateMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
