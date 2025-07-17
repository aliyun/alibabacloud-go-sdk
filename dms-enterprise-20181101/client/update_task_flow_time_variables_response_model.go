// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowTimeVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowTimeVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowTimeVariablesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowTimeVariablesResponseBody) *UpdateTaskFlowTimeVariablesResponse
	GetBody() *UpdateTaskFlowTimeVariablesResponseBody
}

type UpdateTaskFlowTimeVariablesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowTimeVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowTimeVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowTimeVariablesResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowTimeVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowTimeVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowTimeVariablesResponse) GetBody() *UpdateTaskFlowTimeVariablesResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowTimeVariablesResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowTimeVariablesResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowTimeVariablesResponse) SetStatusCode(v int32) *UpdateTaskFlowTimeVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowTimeVariablesResponse) SetBody(v *UpdateTaskFlowTimeVariablesResponseBody) *UpdateTaskFlowTimeVariablesResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowTimeVariablesResponse) Validate() error {
	return dara.Validate(s)
}
