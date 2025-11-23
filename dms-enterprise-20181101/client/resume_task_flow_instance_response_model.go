// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTaskFlowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeTaskFlowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeTaskFlowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ResumeTaskFlowInstanceResponseBody) *ResumeTaskFlowInstanceResponse
	GetBody() *ResumeTaskFlowInstanceResponseBody
}

type ResumeTaskFlowInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeTaskFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeTaskFlowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeTaskFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResumeTaskFlowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeTaskFlowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeTaskFlowInstanceResponse) GetBody() *ResumeTaskFlowInstanceResponseBody {
	return s.Body
}

func (s *ResumeTaskFlowInstanceResponse) SetHeaders(v map[string]*string) *ResumeTaskFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResumeTaskFlowInstanceResponse) SetStatusCode(v int32) *ResumeTaskFlowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeTaskFlowInstanceResponse) SetBody(v *ResumeTaskFlowInstanceResponseBody) *ResumeTaskFlowInstanceResponse {
	s.Body = v
	return s
}

func (s *ResumeTaskFlowInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
