// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterChatCompletionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterChatCompletionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterChatCompletionsResponse
	GetStatusCode() *int32
	SetId(v string) *ModelRouterChatCompletionsResponse
	GetId() *string
	SetEvent(v string) *ModelRouterChatCompletionsResponse
	GetEvent() *string
	SetBody(v *ModelRouterChatCompletionsResponseBody) *ModelRouterChatCompletionsResponse
	GetBody() *ModelRouterChatCompletionsResponseBody
}

type ModelRouterChatCompletionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                                 `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                                 `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ModelRouterChatCompletionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterChatCompletionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterChatCompletionsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterChatCompletionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterChatCompletionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterChatCompletionsResponse) GetId() *string {
	return s.Id
}

func (s *ModelRouterChatCompletionsResponse) GetEvent() *string {
	return s.Event
}

func (s *ModelRouterChatCompletionsResponse) GetBody() *ModelRouterChatCompletionsResponseBody {
	return s.Body
}

func (s *ModelRouterChatCompletionsResponse) SetHeaders(v map[string]*string) *ModelRouterChatCompletionsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterChatCompletionsResponse) SetStatusCode(v int32) *ModelRouterChatCompletionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterChatCompletionsResponse) SetId(v string) *ModelRouterChatCompletionsResponse {
	s.Id = &v
	return s
}

func (s *ModelRouterChatCompletionsResponse) SetEvent(v string) *ModelRouterChatCompletionsResponse {
	s.Event = &v
	return s
}

func (s *ModelRouterChatCompletionsResponse) SetBody(v *ModelRouterChatCompletionsResponseBody) *ModelRouterChatCompletionsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterChatCompletionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
