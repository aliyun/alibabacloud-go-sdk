// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowEdgesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowEdgesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowEdgesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowEdgesResponseBody) *UpdateTaskFlowEdgesResponse
	GetBody() *UpdateTaskFlowEdgesResponseBody
}

type UpdateTaskFlowEdgesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowEdgesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowEdgesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowEdgesResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowEdgesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowEdgesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowEdgesResponse) GetBody() *UpdateTaskFlowEdgesResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowEdgesResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowEdgesResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowEdgesResponse) SetStatusCode(v int32) *UpdateTaskFlowEdgesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowEdgesResponse) SetBody(v *UpdateTaskFlowEdgesResponseBody) *UpdateTaskFlowEdgesResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowEdgesResponse) Validate() error {
	return dara.Validate(s)
}
