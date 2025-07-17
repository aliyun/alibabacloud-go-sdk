// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskFlowEdgesByConditionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTaskFlowEdgesByConditionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTaskFlowEdgesByConditionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTaskFlowEdgesByConditionResponseBody) *DeleteTaskFlowEdgesByConditionResponse
	GetBody() *DeleteTaskFlowEdgesByConditionResponseBody
}

type DeleteTaskFlowEdgesByConditionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTaskFlowEdgesByConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTaskFlowEdgesByConditionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskFlowEdgesByConditionResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskFlowEdgesByConditionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTaskFlowEdgesByConditionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTaskFlowEdgesByConditionResponse) GetBody() *DeleteTaskFlowEdgesByConditionResponseBody {
	return s.Body
}

func (s *DeleteTaskFlowEdgesByConditionResponse) SetHeaders(v map[string]*string) *DeleteTaskFlowEdgesByConditionResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionResponse) SetStatusCode(v int32) *DeleteTaskFlowEdgesByConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionResponse) SetBody(v *DeleteTaskFlowEdgesByConditionResponseBody) *DeleteTaskFlowEdgesByConditionResponse {
	s.Body = v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionResponse) Validate() error {
	return dara.Validate(s)
}
