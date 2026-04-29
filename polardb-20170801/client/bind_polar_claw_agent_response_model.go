// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPolarClawAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindPolarClawAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindPolarClawAgentResponse
	GetStatusCode() *int32
	SetBody(v *BindPolarClawAgentResponseBody) *BindPolarClawAgentResponse
	GetBody() *BindPolarClawAgentResponseBody
}

type BindPolarClawAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindPolarClawAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindPolarClawAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s BindPolarClawAgentResponse) GoString() string {
	return s.String()
}

func (s *BindPolarClawAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindPolarClawAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindPolarClawAgentResponse) GetBody() *BindPolarClawAgentResponseBody {
	return s.Body
}

func (s *BindPolarClawAgentResponse) SetHeaders(v map[string]*string) *BindPolarClawAgentResponse {
	s.Headers = v
	return s
}

func (s *BindPolarClawAgentResponse) SetStatusCode(v int32) *BindPolarClawAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPolarClawAgentResponse) SetBody(v *BindPolarClawAgentResponseBody) *BindPolarClawAgentResponse {
	s.Body = v
	return s
}

func (s *BindPolarClawAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
