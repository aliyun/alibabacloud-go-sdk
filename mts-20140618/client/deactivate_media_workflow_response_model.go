// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateMediaWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactivateMediaWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactivateMediaWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *DeactivateMediaWorkflowResponseBody) *DeactivateMediaWorkflowResponse
	GetBody() *DeactivateMediaWorkflowResponseBody
}

type DeactivateMediaWorkflowResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactivateMediaWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactivateMediaWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactivateMediaWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DeactivateMediaWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactivateMediaWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactivateMediaWorkflowResponse) GetBody() *DeactivateMediaWorkflowResponseBody {
	return s.Body
}

func (s *DeactivateMediaWorkflowResponse) SetHeaders(v map[string]*string) *DeactivateMediaWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DeactivateMediaWorkflowResponse) SetStatusCode(v int32) *DeactivateMediaWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateMediaWorkflowResponse) SetBody(v *DeactivateMediaWorkflowResponseBody) *DeactivateMediaWorkflowResponse {
	s.Body = v
	return s
}

func (s *DeactivateMediaWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
