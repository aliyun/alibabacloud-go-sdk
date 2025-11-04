// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndPublishAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAndPublishAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAndPublishAgentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAndPublishAgentResponseBody) *UpdateAndPublishAgentResponse
	GetBody() *UpdateAndPublishAgentResponseBody
}

type UpdateAndPublishAgentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAndPublishAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAndPublishAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAndPublishAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAndPublishAgentResponse) GetBody() *UpdateAndPublishAgentResponseBody {
	return s.Body
}

func (s *UpdateAndPublishAgentResponse) SetHeaders(v map[string]*string) *UpdateAndPublishAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAndPublishAgentResponse) SetStatusCode(v int32) *UpdateAndPublishAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAndPublishAgentResponse) SetBody(v *UpdateAndPublishAgentResponseBody) *UpdateAndPublishAgentResponse {
	s.Body = v
	return s
}

func (s *UpdateAndPublishAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
