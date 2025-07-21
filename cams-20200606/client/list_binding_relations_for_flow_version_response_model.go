// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindingRelationsForFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBindingRelationsForFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBindingRelationsForFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *ListBindingRelationsForFlowVersionResponseBody) *ListBindingRelationsForFlowVersionResponse
	GetBody() *ListBindingRelationsForFlowVersionResponseBody
}

type ListBindingRelationsForFlowVersionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindingRelationsForFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindingRelationsForFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBindingRelationsForFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *ListBindingRelationsForFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBindingRelationsForFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBindingRelationsForFlowVersionResponse) GetBody() *ListBindingRelationsForFlowVersionResponseBody {
	return s.Body
}

func (s *ListBindingRelationsForFlowVersionResponse) SetHeaders(v map[string]*string) *ListBindingRelationsForFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponse) SetStatusCode(v int32) *ListBindingRelationsForFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponse) SetBody(v *ListBindingRelationsForFlowVersionResponseBody) *ListBindingRelationsForFlowVersionResponse {
	s.Body = v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponse) Validate() error {
	return dara.Validate(s)
}
