// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginPolarClawChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LoginPolarClawChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LoginPolarClawChannelResponse
	GetStatusCode() *int32
	SetBody(v *LoginPolarClawChannelResponseBody) *LoginPolarClawChannelResponse
	GetBody() *LoginPolarClawChannelResponseBody
}

type LoginPolarClawChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoginPolarClawChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoginPolarClawChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s LoginPolarClawChannelResponse) GoString() string {
	return s.String()
}

func (s *LoginPolarClawChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LoginPolarClawChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LoginPolarClawChannelResponse) GetBody() *LoginPolarClawChannelResponseBody {
	return s.Body
}

func (s *LoginPolarClawChannelResponse) SetHeaders(v map[string]*string) *LoginPolarClawChannelResponse {
	s.Headers = v
	return s
}

func (s *LoginPolarClawChannelResponse) SetStatusCode(v int32) *LoginPolarClawChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginPolarClawChannelResponse) SetBody(v *LoginPolarClawChannelResponseBody) *LoginPolarClawChannelResponse {
	s.Body = v
	return s
}

func (s *LoginPolarClawChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
