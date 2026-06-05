// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConversationLockStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppConversationLockStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppConversationLockStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAppConversationLockStatusResponseBody) *GetAppConversationLockStatusResponse
	GetBody() *GetAppConversationLockStatusResponseBody
}

type GetAppConversationLockStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppConversationLockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppConversationLockStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppConversationLockStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAppConversationLockStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppConversationLockStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppConversationLockStatusResponse) GetBody() *GetAppConversationLockStatusResponseBody {
	return s.Body
}

func (s *GetAppConversationLockStatusResponse) SetHeaders(v map[string]*string) *GetAppConversationLockStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAppConversationLockStatusResponse) SetStatusCode(v int32) *GetAppConversationLockStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppConversationLockStatusResponse) SetBody(v *GetAppConversationLockStatusResponseBody) *GetAppConversationLockStatusResponse {
	s.Body = v
	return s
}

func (s *GetAppConversationLockStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
