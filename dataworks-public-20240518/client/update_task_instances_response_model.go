// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskInstancesResponseBody) *UpdateTaskInstancesResponse
	GetBody() *UpdateTaskInstancesResponseBody
}

type UpdateTaskInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskInstancesResponse) GetBody() *UpdateTaskInstancesResponseBody {
	return s.Body
}

func (s *UpdateTaskInstancesResponse) SetHeaders(v map[string]*string) *UpdateTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskInstancesResponse) SetStatusCode(v int32) *UpdateTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskInstancesResponse) SetBody(v *UpdateTaskInstancesResponseBody) *UpdateTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
