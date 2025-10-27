// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkflowResponseBody) *DeleteWorkflowResponse
	GetBody() *DeleteWorkflowResponseBody
}

type DeleteWorkflowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkflowResponse) GetBody() *DeleteWorkflowResponseBody {
	return s.Body
}

func (s *DeleteWorkflowResponse) SetHeaders(v map[string]*string) *DeleteWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowResponse) SetStatusCode(v int32) *DeleteWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowResponse) SetBody(v *DeleteWorkflowResponseBody) *DeleteWorkflowResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
