// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0MemoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMem0MemoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMem0MemoriesResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *DeleteMem0MemoriesResponse
	GetBody() map[string]interface{}
}

type DeleteMem0MemoriesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMem0MemoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0MemoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMem0MemoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMem0MemoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMem0MemoriesResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *DeleteMem0MemoriesResponse) SetHeaders(v map[string]*string) *DeleteMem0MemoriesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMem0MemoriesResponse) SetStatusCode(v int32) *DeleteMem0MemoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMem0MemoriesResponse) SetBody(v map[string]interface{}) *DeleteMem0MemoriesResponse {
	s.Body = v
	return s
}

func (s *DeleteMem0MemoriesResponse) Validate() error {
	return dara.Validate(s)
}
