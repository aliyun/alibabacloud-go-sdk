// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowNameAndDescResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowNameAndDescResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowNameAndDescResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowNameAndDescResponseBody) *UpdateTaskFlowNameAndDescResponse
	GetBody() *UpdateTaskFlowNameAndDescResponseBody
}

type UpdateTaskFlowNameAndDescResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowNameAndDescResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowNameAndDescResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowNameAndDescResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowNameAndDescResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowNameAndDescResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowNameAndDescResponse) GetBody() *UpdateTaskFlowNameAndDescResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowNameAndDescResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowNameAndDescResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowNameAndDescResponse) SetStatusCode(v int32) *UpdateTaskFlowNameAndDescResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescResponse) SetBody(v *UpdateTaskFlowNameAndDescResponseBody) *UpdateTaskFlowNameAndDescResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowNameAndDescResponse) Validate() error {
	return dara.Validate(s)
}
