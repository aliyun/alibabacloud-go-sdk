// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlidingAssistantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlidingAssistantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlidingAssistantResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlidingAssistantResponseBody) *UpdateAlidingAssistantResponse
	GetBody() *UpdateAlidingAssistantResponseBody
}

type UpdateAlidingAssistantResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlidingAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlidingAssistantResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlidingAssistantResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlidingAssistantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlidingAssistantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlidingAssistantResponse) GetBody() *UpdateAlidingAssistantResponseBody {
	return s.Body
}

func (s *UpdateAlidingAssistantResponse) SetHeaders(v map[string]*string) *UpdateAlidingAssistantResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlidingAssistantResponse) SetStatusCode(v int32) *UpdateAlidingAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlidingAssistantResponse) SetBody(v *UpdateAlidingAssistantResponseBody) *UpdateAlidingAssistantResponse {
	s.Body = v
	return s
}

func (s *UpdateAlidingAssistantResponse) Validate() error {
	return dara.Validate(s)
}
