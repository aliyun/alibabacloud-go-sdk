// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactFlowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContactFlowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContactFlowsResponse
	GetStatusCode() *int32
	SetBody(v *ListContactFlowsResponseBody) *ListContactFlowsResponse
	GetBody() *ListContactFlowsResponseBody
}

type ListContactFlowsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContactFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContactFlowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContactFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListContactFlowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContactFlowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContactFlowsResponse) GetBody() *ListContactFlowsResponseBody {
	return s.Body
}

func (s *ListContactFlowsResponse) SetHeaders(v map[string]*string) *ListContactFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListContactFlowsResponse) SetStatusCode(v int32) *ListContactFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactFlowsResponse) SetBody(v *ListContactFlowsResponseBody) *ListContactFlowsResponse {
	s.Body = v
	return s
}

func (s *ListContactFlowsResponse) Validate() error {
	return dara.Validate(s)
}
