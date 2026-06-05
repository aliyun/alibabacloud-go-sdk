// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchAppConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchAppConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchAppConversationResponse
	GetStatusCode() *int32
	SetBody(v *SwitchAppConversationResponseBody) *SwitchAppConversationResponse
	GetBody() *SwitchAppConversationResponseBody
}

type SwitchAppConversationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchAppConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchAppConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchAppConversationResponse) GoString() string {
	return s.String()
}

func (s *SwitchAppConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchAppConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchAppConversationResponse) GetBody() *SwitchAppConversationResponseBody {
	return s.Body
}

func (s *SwitchAppConversationResponse) SetHeaders(v map[string]*string) *SwitchAppConversationResponse {
	s.Headers = v
	return s
}

func (s *SwitchAppConversationResponse) SetStatusCode(v int32) *SwitchAppConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchAppConversationResponse) SetBody(v *SwitchAppConversationResponseBody) *SwitchAppConversationResponse {
	s.Body = v
	return s
}

func (s *SwitchAppConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
