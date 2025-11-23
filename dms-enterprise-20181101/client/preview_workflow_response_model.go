// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *PreviewWorkflowResponseBody) *PreviewWorkflowResponse
	GetBody() *PreviewWorkflowResponseBody
}

type PreviewWorkflowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewWorkflowResponse) GoString() string {
	return s.String()
}

func (s *PreviewWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewWorkflowResponse) GetBody() *PreviewWorkflowResponseBody {
	return s.Body
}

func (s *PreviewWorkflowResponse) SetHeaders(v map[string]*string) *PreviewWorkflowResponse {
	s.Headers = v
	return s
}

func (s *PreviewWorkflowResponse) SetStatusCode(v int32) *PreviewWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewWorkflowResponse) SetBody(v *PreviewWorkflowResponseBody) *PreviewWorkflowResponse {
	s.Body = v
	return s
}

func (s *PreviewWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
