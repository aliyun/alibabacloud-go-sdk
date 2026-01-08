// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappPhoneNumberSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatappPhoneNumberSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatappPhoneNumberSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetChatappPhoneNumberSettingResponseBody) *GetChatappPhoneNumberSettingResponse
	GetBody() *GetChatappPhoneNumberSettingResponseBody
}

type GetChatappPhoneNumberSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappPhoneNumberSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappPhoneNumberSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatappPhoneNumberSettingResponse) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatappPhoneNumberSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatappPhoneNumberSettingResponse) GetBody() *GetChatappPhoneNumberSettingResponseBody {
	return s.Body
}

func (s *GetChatappPhoneNumberSettingResponse) SetHeaders(v map[string]*string) *GetChatappPhoneNumberSettingResponse {
	s.Headers = v
	return s
}

func (s *GetChatappPhoneNumberSettingResponse) SetStatusCode(v int32) *GetChatappPhoneNumberSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappPhoneNumberSettingResponse) SetBody(v *GetChatappPhoneNumberSettingResponseBody) *GetChatappPhoneNumberSettingResponse {
	s.Body = v
	return s
}

func (s *GetChatappPhoneNumberSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
