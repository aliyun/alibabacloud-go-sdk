// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInnerNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInnerNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInnerNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListInnerNodesResponseBody) *ListInnerNodesResponse
	GetBody() *ListInnerNodesResponseBody
}

type ListInnerNodesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInnerNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInnerNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInnerNodesResponse) GoString() string {
	return s.String()
}

func (s *ListInnerNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInnerNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInnerNodesResponse) GetBody() *ListInnerNodesResponseBody {
	return s.Body
}

func (s *ListInnerNodesResponse) SetHeaders(v map[string]*string) *ListInnerNodesResponse {
	s.Headers = v
	return s
}

func (s *ListInnerNodesResponse) SetStatusCode(v int32) *ListInnerNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInnerNodesResponse) SetBody(v *ListInnerNodesResponseBody) *ListInnerNodesResponse {
	s.Body = v
	return s
}

func (s *ListInnerNodesResponse) Validate() error {
	return dara.Validate(s)
}
