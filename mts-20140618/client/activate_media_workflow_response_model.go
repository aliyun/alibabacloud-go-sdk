// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateMediaWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateMediaWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateMediaWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *ActivateMediaWorkflowResponseBody) *ActivateMediaWorkflowResponse
	GetBody() *ActivateMediaWorkflowResponseBody
}

type ActivateMediaWorkflowResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateMediaWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateMediaWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateMediaWorkflowResponse) GoString() string {
	return s.String()
}

func (s *ActivateMediaWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateMediaWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateMediaWorkflowResponse) GetBody() *ActivateMediaWorkflowResponseBody {
	return s.Body
}

func (s *ActivateMediaWorkflowResponse) SetHeaders(v map[string]*string) *ActivateMediaWorkflowResponse {
	s.Headers = v
	return s
}

func (s *ActivateMediaWorkflowResponse) SetStatusCode(v int32) *ActivateMediaWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateMediaWorkflowResponse) SetBody(v *ActivateMediaWorkflowResponseBody) *ActivateMediaWorkflowResponse {
	s.Body = v
	return s
}

func (s *ActivateMediaWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
