// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemoryNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemoryNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListMemoryNodesResponseBody) *ListMemoryNodesResponse
	GetBody() *ListMemoryNodesResponseBody
}

type ListMemoryNodesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoryNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoryNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryNodesResponse) GoString() string {
	return s.String()
}

func (s *ListMemoryNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemoryNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemoryNodesResponse) GetBody() *ListMemoryNodesResponseBody {
	return s.Body
}

func (s *ListMemoryNodesResponse) SetHeaders(v map[string]*string) *ListMemoryNodesResponse {
	s.Headers = v
	return s
}

func (s *ListMemoryNodesResponse) SetStatusCode(v int32) *ListMemoryNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoryNodesResponse) SetBody(v *ListMemoryNodesResponseBody) *ListMemoryNodesResponse {
	s.Body = v
	return s
}

func (s *ListMemoryNodesResponse) Validate() error {
	return dara.Validate(s)
}
