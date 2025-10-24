// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaWorkflowTriggerModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaWorkflowTriggerModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaWorkflowTriggerModeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaWorkflowTriggerModeResponseBody) *UpdateMediaWorkflowTriggerModeResponse
	GetBody() *UpdateMediaWorkflowTriggerModeResponseBody
}

type UpdateMediaWorkflowTriggerModeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaWorkflowTriggerModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaWorkflowTriggerModeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaWorkflowTriggerModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaWorkflowTriggerModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaWorkflowTriggerModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaWorkflowTriggerModeResponse) GetBody() *UpdateMediaWorkflowTriggerModeResponseBody {
	return s.Body
}

func (s *UpdateMediaWorkflowTriggerModeResponse) SetHeaders(v map[string]*string) *UpdateMediaWorkflowTriggerModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponse) SetStatusCode(v int32) *UpdateMediaWorkflowTriggerModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponse) SetBody(v *UpdateMediaWorkflowTriggerModeResponseBody) *UpdateMediaWorkflowTriggerModeResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
