// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowOwnerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowOwnerResponseBody) *UpdateTaskFlowOwnerResponse
	GetBody() *UpdateTaskFlowOwnerResponseBody
}

type UpdateTaskFlowOwnerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowOwnerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowOwnerResponse) GetBody() *UpdateTaskFlowOwnerResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowOwnerResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowOwnerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowOwnerResponse) SetStatusCode(v int32) *UpdateTaskFlowOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowOwnerResponse) SetBody(v *UpdateTaskFlowOwnerResponseBody) *UpdateTaskFlowOwnerResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowOwnerResponse) Validate() error {
	return dara.Validate(s)
}
