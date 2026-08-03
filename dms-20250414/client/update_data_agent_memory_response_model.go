// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataAgentMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataAgentMemoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataAgentMemoryResponseBody) *UpdateDataAgentMemoryResponse
	GetBody() *UpdateDataAgentMemoryResponseBody
}

type UpdateDataAgentMemoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataAgentMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataAgentMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentMemoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataAgentMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataAgentMemoryResponse) GetBody() *UpdateDataAgentMemoryResponseBody {
	return s.Body
}

func (s *UpdateDataAgentMemoryResponse) SetHeaders(v map[string]*string) *UpdateDataAgentMemoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataAgentMemoryResponse) SetStatusCode(v int32) *UpdateDataAgentMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataAgentMemoryResponse) SetBody(v *UpdateDataAgentMemoryResponseBody) *UpdateDataAgentMemoryResponse {
	s.Body = v
	return s
}

func (s *UpdateDataAgentMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
