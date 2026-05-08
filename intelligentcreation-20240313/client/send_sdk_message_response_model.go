// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSdkMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendSdkMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendSdkMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendSdkMessageResponseBody) *SendSdkMessageResponse
	GetBody() *SendSdkMessageResponseBody
}

type SendSdkMessageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSdkMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSdkMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendSdkMessageResponse) GoString() string {
	return s.String()
}

func (s *SendSdkMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendSdkMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendSdkMessageResponse) GetBody() *SendSdkMessageResponseBody {
	return s.Body
}

func (s *SendSdkMessageResponse) SetHeaders(v map[string]*string) *SendSdkMessageResponse {
	s.Headers = v
	return s
}

func (s *SendSdkMessageResponse) SetStatusCode(v int32) *SendSdkMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSdkMessageResponse) SetBody(v *SendSdkMessageResponseBody) *SendSdkMessageResponse {
	s.Body = v
	return s
}

func (s *SendSdkMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
