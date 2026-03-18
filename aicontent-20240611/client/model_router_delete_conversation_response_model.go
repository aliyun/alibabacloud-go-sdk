// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterDeleteConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterDeleteConversationResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterDeleteConversationResponseBody) *ModelRouterDeleteConversationResponse
	GetBody() *ModelRouterDeleteConversationResponseBody
}

type ModelRouterDeleteConversationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterDeleteConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterDeleteConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteConversationResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterDeleteConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterDeleteConversationResponse) GetBody() *ModelRouterDeleteConversationResponseBody {
	return s.Body
}

func (s *ModelRouterDeleteConversationResponse) SetHeaders(v map[string]*string) *ModelRouterDeleteConversationResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterDeleteConversationResponse) SetStatusCode(v int32) *ModelRouterDeleteConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterDeleteConversationResponse) SetBody(v *ModelRouterDeleteConversationResponseBody) *ModelRouterDeleteConversationResponse {
	s.Body = v
	return s
}

func (s *ModelRouterDeleteConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
