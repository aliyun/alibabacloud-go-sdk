// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateConversationResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateConversationResponseBody) *ModelRouterCreateConversationResponse
	GetBody() *ModelRouterCreateConversationResponseBody
}

type ModelRouterCreateConversationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateConversationResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateConversationResponse) GetBody() *ModelRouterCreateConversationResponseBody {
	return s.Body
}

func (s *ModelRouterCreateConversationResponse) SetHeaders(v map[string]*string) *ModelRouterCreateConversationResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateConversationResponse) SetStatusCode(v int32) *ModelRouterCreateConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateConversationResponse) SetBody(v *ModelRouterCreateConversationResponseBody) *ModelRouterCreateConversationResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
