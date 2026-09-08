// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAgentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAgentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAgentJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAgentJobResponseBody) *SubmitAgentJobResponse
	GetBody() *SubmitAgentJobResponseBody
}

type SubmitAgentJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAgentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAgentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAgentJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAgentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAgentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAgentJobResponse) GetBody() *SubmitAgentJobResponseBody {
	return s.Body
}

func (s *SubmitAgentJobResponse) SetHeaders(v map[string]*string) *SubmitAgentJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAgentJobResponse) SetStatusCode(v int32) *SubmitAgentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAgentJobResponse) SetBody(v *SubmitAgentJobResponseBody) *SubmitAgentJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAgentJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
