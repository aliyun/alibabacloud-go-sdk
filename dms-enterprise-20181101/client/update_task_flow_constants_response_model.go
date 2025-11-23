// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowConstantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowConstantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowConstantsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowConstantsResponseBody) *UpdateTaskFlowConstantsResponse
	GetBody() *UpdateTaskFlowConstantsResponseBody
}

type UpdateTaskFlowConstantsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowConstantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowConstantsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowConstantsResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowConstantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowConstantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowConstantsResponse) GetBody() *UpdateTaskFlowConstantsResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowConstantsResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowConstantsResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowConstantsResponse) SetStatusCode(v int32) *UpdateTaskFlowConstantsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowConstantsResponse) SetBody(v *UpdateTaskFlowConstantsResponseBody) *UpdateTaskFlowConstantsResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowConstantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
