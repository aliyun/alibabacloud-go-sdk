// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMem0MemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMem0MemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMem0MemoryResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *UpdateMem0MemoryResponse
	GetBody() map[string]interface{}
}

type UpdateMem0MemoryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMem0MemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMem0MemoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateMem0MemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMem0MemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMem0MemoryResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *UpdateMem0MemoryResponse) SetHeaders(v map[string]*string) *UpdateMem0MemoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateMem0MemoryResponse) SetStatusCode(v int32) *UpdateMem0MemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMem0MemoryResponse) SetBody(v map[string]interface{}) *UpdateMem0MemoryResponse {
	s.Body = v
	return s
}

func (s *UpdateMem0MemoryResponse) Validate() error {
	return dara.Validate(s)
}
