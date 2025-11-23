// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTaskFlowEdgesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTaskFlowEdgesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTaskFlowEdgesResponse
	GetStatusCode() *int32
	SetBody(v *AddTaskFlowEdgesResponseBody) *AddTaskFlowEdgesResponse
	GetBody() *AddTaskFlowEdgesResponseBody
}

type AddTaskFlowEdgesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTaskFlowEdgesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTaskFlowEdgesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTaskFlowEdgesResponse) GoString() string {
	return s.String()
}

func (s *AddTaskFlowEdgesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTaskFlowEdgesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTaskFlowEdgesResponse) GetBody() *AddTaskFlowEdgesResponseBody {
	return s.Body
}

func (s *AddTaskFlowEdgesResponse) SetHeaders(v map[string]*string) *AddTaskFlowEdgesResponse {
	s.Headers = v
	return s
}

func (s *AddTaskFlowEdgesResponse) SetStatusCode(v int32) *AddTaskFlowEdgesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTaskFlowEdgesResponse) SetBody(v *AddTaskFlowEdgesResponseBody) *AddTaskFlowEdgesResponse {
	s.Body = v
	return s
}

func (s *AddTaskFlowEdgesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
