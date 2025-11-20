// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlidingAssistantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlidingAssistantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlidingAssistantResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlidingAssistantResponseBody) *CreateAlidingAssistantResponse
	GetBody() *CreateAlidingAssistantResponseBody
}

type CreateAlidingAssistantResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlidingAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlidingAssistantResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlidingAssistantResponse) GoString() string {
	return s.String()
}

func (s *CreateAlidingAssistantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlidingAssistantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlidingAssistantResponse) GetBody() *CreateAlidingAssistantResponseBody {
	return s.Body
}

func (s *CreateAlidingAssistantResponse) SetHeaders(v map[string]*string) *CreateAlidingAssistantResponse {
	s.Headers = v
	return s
}

func (s *CreateAlidingAssistantResponse) SetStatusCode(v int32) *CreateAlidingAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlidingAssistantResponse) SetBody(v *CreateAlidingAssistantResponseBody) *CreateAlidingAssistantResponse {
	s.Body = v
	return s
}

func (s *CreateAlidingAssistantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
