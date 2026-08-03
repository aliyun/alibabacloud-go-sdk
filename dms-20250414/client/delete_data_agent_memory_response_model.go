// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataAgentMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataAgentMemoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataAgentMemoryResponseBody) *DeleteDataAgentMemoryResponse
	GetBody() *DeleteDataAgentMemoryResponseBody
}

type DeleteDataAgentMemoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataAgentMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataAgentMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentMemoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataAgentMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataAgentMemoryResponse) GetBody() *DeleteDataAgentMemoryResponseBody {
	return s.Body
}

func (s *DeleteDataAgentMemoryResponse) SetHeaders(v map[string]*string) *DeleteDataAgentMemoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataAgentMemoryResponse) SetStatusCode(v int32) *DeleteDataAgentMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataAgentMemoryResponse) SetBody(v *DeleteDataAgentMemoryResponseBody) *DeleteDataAgentMemoryResponse {
	s.Body = v
	return s
}

func (s *DeleteDataAgentMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
