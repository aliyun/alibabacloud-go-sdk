// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryConversationListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryConversationListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryConversationListResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryConversationListResponseBody) *ModelRouterQueryConversationListResponse
	GetBody() *ModelRouterQueryConversationListResponseBody
}

type ModelRouterQueryConversationListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryConversationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryConversationListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryConversationListResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryConversationListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryConversationListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryConversationListResponse) GetBody() *ModelRouterQueryConversationListResponseBody {
	return s.Body
}

func (s *ModelRouterQueryConversationListResponse) SetHeaders(v map[string]*string) *ModelRouterQueryConversationListResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryConversationListResponse) SetStatusCode(v int32) *ModelRouterQueryConversationListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryConversationListResponse) SetBody(v *ModelRouterQueryConversationListResponseBody) *ModelRouterQueryConversationListResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryConversationListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
