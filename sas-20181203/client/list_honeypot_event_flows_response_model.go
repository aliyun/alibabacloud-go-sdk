// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotEventFlowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotEventFlowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotEventFlowsResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotEventFlowsResponseBody) *ListHoneypotEventFlowsResponse
	GetBody() *ListHoneypotEventFlowsResponseBody
}

type ListHoneypotEventFlowsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotEventFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotEventFlowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventFlowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotEventFlowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotEventFlowsResponse) GetBody() *ListHoneypotEventFlowsResponseBody {
	return s.Body
}

func (s *ListHoneypotEventFlowsResponse) SetHeaders(v map[string]*string) *ListHoneypotEventFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotEventFlowsResponse) SetStatusCode(v int32) *ListHoneypotEventFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotEventFlowsResponse) SetBody(v *ListHoneypotEventFlowsResponseBody) *ListHoneypotEventFlowsResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotEventFlowsResponse) Validate() error {
	return dara.Validate(s)
}
