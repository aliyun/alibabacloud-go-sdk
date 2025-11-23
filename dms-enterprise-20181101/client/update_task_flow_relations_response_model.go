// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowRelationsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowRelationsResponseBody) *UpdateTaskFlowRelationsResponse
	GetBody() *UpdateTaskFlowRelationsResponseBody
}

type UpdateTaskFlowRelationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowRelationsResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowRelationsResponse) GetBody() *UpdateTaskFlowRelationsResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowRelationsResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowRelationsResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowRelationsResponse) SetStatusCode(v int32) *UpdateTaskFlowRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowRelationsResponse) SetBody(v *UpdateTaskFlowRelationsResponseBody) *UpdateTaskFlowRelationsResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
