// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *StartWorkflowResponseBody) *StartWorkflowResponse
	GetBody() *StartWorkflowResponseBody
}

type StartWorkflowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s StartWorkflowResponse) GoString() string {
	return s.String()
}

func (s *StartWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartWorkflowResponse) GetBody() *StartWorkflowResponseBody {
	return s.Body
}

func (s *StartWorkflowResponse) SetHeaders(v map[string]*string) *StartWorkflowResponse {
	s.Headers = v
	return s
}

func (s *StartWorkflowResponse) SetStatusCode(v int32) *StartWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *StartWorkflowResponse) SetBody(v *StartWorkflowResponseBody) *StartWorkflowResponse {
	s.Body = v
	return s
}

func (s *StartWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
