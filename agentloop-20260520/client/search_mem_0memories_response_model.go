// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMem0MemoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMem0MemoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMem0MemoriesResponse
	GetStatusCode() *int32
	SetBody(v []map[string]interface{}) *SearchMem0MemoriesResponse
	GetBody() []map[string]interface{}
}

type SearchMem0MemoriesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s SearchMem0MemoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMem0MemoriesResponse) GoString() string {
	return s.String()
}

func (s *SearchMem0MemoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMem0MemoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMem0MemoriesResponse) GetBody() []map[string]interface{} {
	return s.Body
}

func (s *SearchMem0MemoriesResponse) SetHeaders(v map[string]*string) *SearchMem0MemoriesResponse {
	s.Headers = v
	return s
}

func (s *SearchMem0MemoriesResponse) SetStatusCode(v int32) *SearchMem0MemoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMem0MemoriesResponse) SetBody(v []map[string]interface{}) *SearchMem0MemoriesResponse {
	s.Body = v
	return s
}

func (s *SearchMem0MemoriesResponse) Validate() error {
	return dara.Validate(s)
}
