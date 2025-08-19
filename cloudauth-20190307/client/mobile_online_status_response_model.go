// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileOnlineStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MobileOnlineStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MobileOnlineStatusResponse
	GetStatusCode() *int32
	SetBody(v *MobileOnlineStatusResponseBody) *MobileOnlineStatusResponse
	GetBody() *MobileOnlineStatusResponseBody
}

type MobileOnlineStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MobileOnlineStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MobileOnlineStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s MobileOnlineStatusResponse) GoString() string {
	return s.String()
}

func (s *MobileOnlineStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MobileOnlineStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MobileOnlineStatusResponse) GetBody() *MobileOnlineStatusResponseBody {
	return s.Body
}

func (s *MobileOnlineStatusResponse) SetHeaders(v map[string]*string) *MobileOnlineStatusResponse {
	s.Headers = v
	return s
}

func (s *MobileOnlineStatusResponse) SetStatusCode(v int32) *MobileOnlineStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *MobileOnlineStatusResponse) SetBody(v *MobileOnlineStatusResponseBody) *MobileOnlineStatusResponse {
	s.Body = v
	return s
}

func (s *MobileOnlineStatusResponse) Validate() error {
	return dara.Validate(s)
}
