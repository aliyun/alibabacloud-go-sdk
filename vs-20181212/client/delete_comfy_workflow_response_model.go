// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComfyWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComfyWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComfyWorkflowResponseBody) *DeleteComfyWorkflowResponse
	GetBody() *DeleteComfyWorkflowResponseBody
}

type DeleteComfyWorkflowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComfyWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComfyWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DeleteComfyWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComfyWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComfyWorkflowResponse) GetBody() *DeleteComfyWorkflowResponseBody {
	return s.Body
}

func (s *DeleteComfyWorkflowResponse) SetHeaders(v map[string]*string) *DeleteComfyWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DeleteComfyWorkflowResponse) SetStatusCode(v int32) *DeleteComfyWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComfyWorkflowResponse) SetBody(v *DeleteComfyWorkflowResponseBody) *DeleteComfyWorkflowResponse {
	s.Body = v
	return s
}

func (s *DeleteComfyWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
