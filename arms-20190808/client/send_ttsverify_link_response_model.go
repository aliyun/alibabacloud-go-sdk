// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTTSVerifyLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendTTSVerifyLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendTTSVerifyLinkResponse
	GetStatusCode() *int32
	SetBody(v *SendTTSVerifyLinkResponseBody) *SendTTSVerifyLinkResponse
	GetBody() *SendTTSVerifyLinkResponseBody
}

type SendTTSVerifyLinkResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTTSVerifyLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTTSVerifyLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s SendTTSVerifyLinkResponse) GoString() string {
	return s.String()
}

func (s *SendTTSVerifyLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendTTSVerifyLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendTTSVerifyLinkResponse) GetBody() *SendTTSVerifyLinkResponseBody {
	return s.Body
}

func (s *SendTTSVerifyLinkResponse) SetHeaders(v map[string]*string) *SendTTSVerifyLinkResponse {
	s.Headers = v
	return s
}

func (s *SendTTSVerifyLinkResponse) SetStatusCode(v int32) *SendTTSVerifyLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTTSVerifyLinkResponse) SetBody(v *SendTTSVerifyLinkResponseBody) *SendTTSVerifyLinkResponse {
	s.Body = v
	return s
}

func (s *SendTTSVerifyLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
