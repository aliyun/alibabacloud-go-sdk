// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBannerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendBannerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendBannerResponse
	GetStatusCode() *int32
	SetBody(v *SendBannerResponseBody) *SendBannerResponse
	GetBody() *SendBannerResponseBody
}

type SendBannerResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendBannerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendBannerResponse) String() string {
	return dara.Prettify(s)
}

func (s SendBannerResponse) GoString() string {
	return s.String()
}

func (s *SendBannerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendBannerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendBannerResponse) GetBody() *SendBannerResponseBody {
	return s.Body
}

func (s *SendBannerResponse) SetHeaders(v map[string]*string) *SendBannerResponse {
	s.Headers = v
	return s
}

func (s *SendBannerResponse) SetStatusCode(v int32) *SendBannerResponse {
	s.StatusCode = &v
	return s
}

func (s *SendBannerResponse) SetBody(v *SendBannerResponseBody) *SendBannerResponse {
	s.Body = v
	return s
}

func (s *SendBannerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
