// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIAgentPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIAgentPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *ListAIAgentPhoneNumberResponseBody) *ListAIAgentPhoneNumberResponse
	GetBody() *ListAIAgentPhoneNumberResponseBody
}

type ListAIAgentPhoneNumberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIAgentPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIAgentPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ListAIAgentPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIAgentPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIAgentPhoneNumberResponse) GetBody() *ListAIAgentPhoneNumberResponseBody {
	return s.Body
}

func (s *ListAIAgentPhoneNumberResponse) SetHeaders(v map[string]*string) *ListAIAgentPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ListAIAgentPhoneNumberResponse) SetStatusCode(v int32) *ListAIAgentPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIAgentPhoneNumberResponse) SetBody(v *ListAIAgentPhoneNumberResponseBody) *ListAIAgentPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *ListAIAgentPhoneNumberResponse) Validate() error {
	return dara.Validate(s)
}
