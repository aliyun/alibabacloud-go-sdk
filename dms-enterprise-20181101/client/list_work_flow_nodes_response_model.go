// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkFlowNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkFlowNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkFlowNodesResponseBody) *ListWorkFlowNodesResponse
	GetBody() *ListWorkFlowNodesResponseBody
}

type ListWorkFlowNodesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkFlowNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkFlowNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowNodesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkFlowNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkFlowNodesResponse) GetBody() *ListWorkFlowNodesResponseBody {
	return s.Body
}

func (s *ListWorkFlowNodesResponse) SetHeaders(v map[string]*string) *ListWorkFlowNodesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkFlowNodesResponse) SetStatusCode(v int32) *ListWorkFlowNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkFlowNodesResponse) SetBody(v *ListWorkFlowNodesResponseBody) *ListWorkFlowNodesResponse {
	s.Body = v
	return s
}

func (s *ListWorkFlowNodesResponse) Validate() error {
	return dara.Validate(s)
}
