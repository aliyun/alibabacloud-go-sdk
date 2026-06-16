// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMem0MemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMem0MemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMem0MemoryResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *GetMem0MemoryResponse
	GetBody() map[string]interface{}
}

type GetMem0MemoryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMem0MemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMem0MemoryResponse) GoString() string {
	return s.String()
}

func (s *GetMem0MemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMem0MemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMem0MemoryResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *GetMem0MemoryResponse) SetHeaders(v map[string]*string) *GetMem0MemoryResponse {
	s.Headers = v
	return s
}

func (s *GetMem0MemoryResponse) SetStatusCode(v int32) *GetMem0MemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMem0MemoryResponse) SetBody(v map[string]interface{}) *GetMem0MemoryResponse {
	s.Body = v
	return s
}

func (s *GetMem0MemoryResponse) Validate() error {
	return dara.Validate(s)
}
