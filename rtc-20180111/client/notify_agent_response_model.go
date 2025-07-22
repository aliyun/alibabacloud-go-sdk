// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *NotifyAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *NotifyAgentResponse
	GetStatusCode() *int32
	SetBody(v *NotifyAgentResponseBody) *NotifyAgentResponse
	GetBody() *NotifyAgentResponseBody
}

type NotifyAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s NotifyAgentResponse) GoString() string {
	return s.String()
}

func (s *NotifyAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *NotifyAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *NotifyAgentResponse) GetBody() *NotifyAgentResponseBody {
	return s.Body
}

func (s *NotifyAgentResponse) SetHeaders(v map[string]*string) *NotifyAgentResponse {
	s.Headers = v
	return s
}

func (s *NotifyAgentResponse) SetStatusCode(v int32) *NotifyAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyAgentResponse) SetBody(v *NotifyAgentResponseBody) *NotifyAgentResponse {
	s.Body = v
	return s
}

func (s *NotifyAgentResponse) Validate() error {
	return dara.Validate(s)
}
