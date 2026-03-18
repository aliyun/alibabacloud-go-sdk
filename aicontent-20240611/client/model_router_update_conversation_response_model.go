// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterUpdateConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterUpdateConversationResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterUpdateConversationResponseBody) *ModelRouterUpdateConversationResponse
	GetBody() *ModelRouterUpdateConversationResponseBody
}

type ModelRouterUpdateConversationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterUpdateConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterUpdateConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateConversationResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterUpdateConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterUpdateConversationResponse) GetBody() *ModelRouterUpdateConversationResponseBody {
	return s.Body
}

func (s *ModelRouterUpdateConversationResponse) SetHeaders(v map[string]*string) *ModelRouterUpdateConversationResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterUpdateConversationResponse) SetStatusCode(v int32) *ModelRouterUpdateConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterUpdateConversationResponse) SetBody(v *ModelRouterUpdateConversationResponseBody) *ModelRouterUpdateConversationResponse {
	s.Body = v
	return s
}

func (s *ModelRouterUpdateConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
