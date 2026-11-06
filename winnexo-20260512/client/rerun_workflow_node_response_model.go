// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunWorkflowNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RerunWorkflowNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RerunWorkflowNodeResponse
	GetStatusCode() *int32
	SetBody(v *RerunWorkflowNodeResponseBody) *RerunWorkflowNodeResponse
	GetBody() *RerunWorkflowNodeResponseBody
}

type RerunWorkflowNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RerunWorkflowNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RerunWorkflowNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RerunWorkflowNodeResponse) GoString() string {
	return s.String()
}

func (s *RerunWorkflowNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RerunWorkflowNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RerunWorkflowNodeResponse) GetBody() *RerunWorkflowNodeResponseBody {
	return s.Body
}

func (s *RerunWorkflowNodeResponse) SetHeaders(v map[string]*string) *RerunWorkflowNodeResponse {
	s.Headers = v
	return s
}

func (s *RerunWorkflowNodeResponse) SetStatusCode(v int32) *RerunWorkflowNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunWorkflowNodeResponse) SetBody(v *RerunWorkflowNodeResponseBody) *RerunWorkflowNodeResponse {
	s.Body = v
	return s
}

func (s *RerunWorkflowNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
