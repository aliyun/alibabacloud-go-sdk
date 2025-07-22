// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAgentResponse
	GetStatusCode() *int32
	SetBody(v *StartAgentResponseBody) *StartAgentResponse
	GetBody() *StartAgentResponseBody
}

type StartAgentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAgentResponse) GoString() string {
	return s.String()
}

func (s *StartAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAgentResponse) GetBody() *StartAgentResponseBody {
	return s.Body
}

func (s *StartAgentResponse) SetHeaders(v map[string]*string) *StartAgentResponse {
	s.Headers = v
	return s
}

func (s *StartAgentResponse) SetStatusCode(v int32) *StartAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAgentResponse) SetBody(v *StartAgentResponseBody) *StartAgentResponse {
	s.Body = v
	return s
}

func (s *StartAgentResponse) Validate() error {
	return dara.Validate(s)
}
