// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryConversationResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryConversationResponseBody) *ModelRouterQueryConversationResponse
	GetBody() *ModelRouterQueryConversationResponseBody
}

type ModelRouterQueryConversationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryConversationResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryConversationResponse) GetBody() *ModelRouterQueryConversationResponseBody {
	return s.Body
}

func (s *ModelRouterQueryConversationResponse) SetHeaders(v map[string]*string) *ModelRouterQueryConversationResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryConversationResponse) SetStatusCode(v int32) *ModelRouterQueryConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryConversationResponse) SetBody(v *ModelRouterQueryConversationResponseBody) *ModelRouterQueryConversationResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
