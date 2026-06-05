// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComfyWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyComfyWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyComfyWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *ModifyComfyWorkflowResponseBody) *ModifyComfyWorkflowResponse
	GetBody() *ModifyComfyWorkflowResponseBody
}

type ModifyComfyWorkflowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyComfyWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyComfyWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyComfyWorkflowResponse) GoString() string {
	return s.String()
}

func (s *ModifyComfyWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyComfyWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyComfyWorkflowResponse) GetBody() *ModifyComfyWorkflowResponseBody {
	return s.Body
}

func (s *ModifyComfyWorkflowResponse) SetHeaders(v map[string]*string) *ModifyComfyWorkflowResponse {
	s.Headers = v
	return s
}

func (s *ModifyComfyWorkflowResponse) SetStatusCode(v int32) *ModifyComfyWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyComfyWorkflowResponse) SetBody(v *ModifyComfyWorkflowResponseBody) *ModifyComfyWorkflowResponse {
	s.Body = v
	return s
}

func (s *ModifyComfyWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
