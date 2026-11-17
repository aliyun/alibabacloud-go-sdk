// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopKnowledgeBaseBuildTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopKnowledgeBaseBuildTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopKnowledgeBaseBuildTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopKnowledgeBaseBuildTaskInstanceResponseBody) *StopKnowledgeBaseBuildTaskInstanceResponse
	GetBody() *StopKnowledgeBaseBuildTaskInstanceResponseBody
}

type StopKnowledgeBaseBuildTaskInstanceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopKnowledgeBaseBuildTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponse) GetBody() *StopKnowledgeBaseBuildTaskInstanceResponseBody {
	return s.Body
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponse) SetHeaders(v map[string]*string) *StopKnowledgeBaseBuildTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponse) SetStatusCode(v int32) *StopKnowledgeBaseBuildTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponse) SetBody(v *StopKnowledgeBaseBuildTaskInstanceResponseBody) *StopKnowledgeBaseBuildTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
