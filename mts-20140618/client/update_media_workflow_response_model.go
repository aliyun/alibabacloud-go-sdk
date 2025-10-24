// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaWorkflowResponseBody) *UpdateMediaWorkflowResponse
	GetBody() *UpdateMediaWorkflowResponseBody
}

type UpdateMediaWorkflowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaWorkflowResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaWorkflowResponse) GetBody() *UpdateMediaWorkflowResponseBody {
	return s.Body
}

func (s *UpdateMediaWorkflowResponse) SetHeaders(v map[string]*string) *UpdateMediaWorkflowResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaWorkflowResponse) SetStatusCode(v int32) *UpdateMediaWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaWorkflowResponse) SetBody(v *UpdateMediaWorkflowResponseBody) *UpdateMediaWorkflowResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
