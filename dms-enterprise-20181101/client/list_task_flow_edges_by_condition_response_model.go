// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowEdgesByConditionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskFlowEdgesByConditionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskFlowEdgesByConditionResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskFlowEdgesByConditionResponseBody) *ListTaskFlowEdgesByConditionResponse
	GetBody() *ListTaskFlowEdgesByConditionResponseBody
}

type ListTaskFlowEdgesByConditionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskFlowEdgesByConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskFlowEdgesByConditionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowEdgesByConditionResponse) GoString() string {
	return s.String()
}

func (s *ListTaskFlowEdgesByConditionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskFlowEdgesByConditionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskFlowEdgesByConditionResponse) GetBody() *ListTaskFlowEdgesByConditionResponseBody {
	return s.Body
}

func (s *ListTaskFlowEdgesByConditionResponse) SetHeaders(v map[string]*string) *ListTaskFlowEdgesByConditionResponse {
	s.Headers = v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponse) SetStatusCode(v int32) *ListTaskFlowEdgesByConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponse) SetBody(v *ListTaskFlowEdgesByConditionResponseBody) *ListTaskFlowEdgesByConditionResponse {
	s.Body = v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
