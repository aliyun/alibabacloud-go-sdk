// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlidingAssistantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlidingAssistantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlidingAssistantResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlidingAssistantResponseBody) *DeleteAlidingAssistantResponse
	GetBody() *DeleteAlidingAssistantResponseBody
}

type DeleteAlidingAssistantResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlidingAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlidingAssistantResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlidingAssistantResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlidingAssistantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlidingAssistantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlidingAssistantResponse) GetBody() *DeleteAlidingAssistantResponseBody {
	return s.Body
}

func (s *DeleteAlidingAssistantResponse) SetHeaders(v map[string]*string) *DeleteAlidingAssistantResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlidingAssistantResponse) SetStatusCode(v int32) *DeleteAlidingAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlidingAssistantResponse) SetBody(v *DeleteAlidingAssistantResponseBody) *DeleteAlidingAssistantResponse {
	s.Body = v
	return s
}

func (s *DeleteAlidingAssistantResponse) Validate() error {
	return dara.Validate(s)
}
