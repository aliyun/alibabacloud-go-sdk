// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowCooperatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowCooperatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowCooperatorsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowCooperatorsResponseBody) *UpdateTaskFlowCooperatorsResponse
	GetBody() *UpdateTaskFlowCooperatorsResponseBody
}

type UpdateTaskFlowCooperatorsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowCooperatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowCooperatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowCooperatorsResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowCooperatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowCooperatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowCooperatorsResponse) GetBody() *UpdateTaskFlowCooperatorsResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowCooperatorsResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowCooperatorsResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowCooperatorsResponse) SetStatusCode(v int32) *UpdateTaskFlowCooperatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsResponse) SetBody(v *UpdateTaskFlowCooperatorsResponseBody) *UpdateTaskFlowCooperatorsResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowCooperatorsResponse) Validate() error {
	return dara.Validate(s)
}
