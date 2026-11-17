// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRuntimeQrCodeConfigChannelMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendRuntimeQrCodeConfigChannelMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendRuntimeQrCodeConfigChannelMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendRuntimeQrCodeConfigChannelMessageResponseBody) *SendRuntimeQrCodeConfigChannelMessageResponse
	GetBody() *SendRuntimeQrCodeConfigChannelMessageResponseBody
}

type SendRuntimeQrCodeConfigChannelMessageResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendRuntimeQrCodeConfigChannelMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendRuntimeQrCodeConfigChannelMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendRuntimeQrCodeConfigChannelMessageResponse) GoString() string {
	return s.String()
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponse) GetBody() *SendRuntimeQrCodeConfigChannelMessageResponseBody {
	return s.Body
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponse) SetHeaders(v map[string]*string) *SendRuntimeQrCodeConfigChannelMessageResponse {
	s.Headers = v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponse) SetStatusCode(v int32) *SendRuntimeQrCodeConfigChannelMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponse) SetBody(v *SendRuntimeQrCodeConfigChannelMessageResponseBody) *SendRuntimeQrCodeConfigChannelMessageResponse {
	s.Body = v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
