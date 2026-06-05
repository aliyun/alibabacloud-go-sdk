// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComfyWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComfyWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *CreateComfyWorkflowResponseBody) *CreateComfyWorkflowResponse
	GetBody() *CreateComfyWorkflowResponseBody
}

type CreateComfyWorkflowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComfyWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComfyWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyWorkflowResponse) GoString() string {
	return s.String()
}

func (s *CreateComfyWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComfyWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComfyWorkflowResponse) GetBody() *CreateComfyWorkflowResponseBody {
	return s.Body
}

func (s *CreateComfyWorkflowResponse) SetHeaders(v map[string]*string) *CreateComfyWorkflowResponse {
	s.Headers = v
	return s
}

func (s *CreateComfyWorkflowResponse) SetStatusCode(v int32) *CreateComfyWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComfyWorkflowResponse) SetBody(v *CreateComfyWorkflowResponseBody) *CreateComfyWorkflowResponse {
	s.Body = v
	return s
}

func (s *CreateComfyWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
