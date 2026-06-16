// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0MemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMem0MemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMem0MemoryResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *DeleteMem0MemoryResponse
	GetBody() map[string]interface{}
}

type DeleteMem0MemoryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMem0MemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0MemoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteMem0MemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMem0MemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMem0MemoryResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *DeleteMem0MemoryResponse) SetHeaders(v map[string]*string) *DeleteMem0MemoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteMem0MemoryResponse) SetStatusCode(v int32) *DeleteMem0MemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMem0MemoryResponse) SetBody(v map[string]interface{}) *DeleteMem0MemoryResponse {
	s.Body = v
	return s
}

func (s *DeleteMem0MemoryResponse) Validate() error {
	return dara.Validate(s)
}
