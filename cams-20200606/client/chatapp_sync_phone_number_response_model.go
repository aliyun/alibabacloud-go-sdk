// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappSyncPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatappSyncPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatappSyncPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *ChatappSyncPhoneNumberResponseBody) *ChatappSyncPhoneNumberResponse
	GetBody() *ChatappSyncPhoneNumberResponseBody
}

type ChatappSyncPhoneNumberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappSyncPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappSyncPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatappSyncPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatappSyncPhoneNumberResponse) GetBody() *ChatappSyncPhoneNumberResponseBody {
	return s.Body
}

func (s *ChatappSyncPhoneNumberResponse) SetHeaders(v map[string]*string) *ChatappSyncPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ChatappSyncPhoneNumberResponse) SetStatusCode(v int32) *ChatappSyncPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponse) SetBody(v *ChatappSyncPhoneNumberResponseBody) *ChatappSyncPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *ChatappSyncPhoneNumberResponse) Validate() error {
	return dara.Validate(s)
}
