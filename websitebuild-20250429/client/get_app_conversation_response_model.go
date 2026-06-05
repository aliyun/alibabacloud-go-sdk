// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppConversationResponse
	GetStatusCode() *int32
	SetBody(v *GetAppConversationResponseBody) *GetAppConversationResponse
	GetBody() *GetAppConversationResponseBody
}

type GetAppConversationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppConversationResponse) GoString() string {
	return s.String()
}

func (s *GetAppConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppConversationResponse) GetBody() *GetAppConversationResponseBody {
	return s.Body
}

func (s *GetAppConversationResponse) SetHeaders(v map[string]*string) *GetAppConversationResponse {
	s.Headers = v
	return s
}

func (s *GetAppConversationResponse) SetStatusCode(v int32) *GetAppConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppConversationResponse) SetBody(v *GetAppConversationResponseBody) *GetAppConversationResponse {
	s.Body = v
	return s
}

func (s *GetAppConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
