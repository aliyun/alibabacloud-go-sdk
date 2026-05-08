// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseAgentResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseAgentResponseBody) *ReleaseAgentResponse
	GetBody() *ReleaseAgentResponseBody
}

type ReleaseAgentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAgentResponse) GoString() string {
	return s.String()
}

func (s *ReleaseAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseAgentResponse) GetBody() *ReleaseAgentResponseBody {
	return s.Body
}

func (s *ReleaseAgentResponse) SetHeaders(v map[string]*string) *ReleaseAgentResponse {
	s.Headers = v
	return s
}

func (s *ReleaseAgentResponse) SetStatusCode(v int32) *ReleaseAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseAgentResponse) SetBody(v *ReleaseAgentResponseBody) *ReleaseAgentResponse {
	s.Body = v
	return s
}

func (s *ReleaseAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
