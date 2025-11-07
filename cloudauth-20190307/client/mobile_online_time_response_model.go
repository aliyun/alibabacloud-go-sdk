// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileOnlineTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MobileOnlineTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MobileOnlineTimeResponse
	GetStatusCode() *int32
	SetBody(v *MobileOnlineTimeResponseBody) *MobileOnlineTimeResponse
	GetBody() *MobileOnlineTimeResponseBody
}

type MobileOnlineTimeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MobileOnlineTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MobileOnlineTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s MobileOnlineTimeResponse) GoString() string {
	return s.String()
}

func (s *MobileOnlineTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MobileOnlineTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MobileOnlineTimeResponse) GetBody() *MobileOnlineTimeResponseBody {
	return s.Body
}

func (s *MobileOnlineTimeResponse) SetHeaders(v map[string]*string) *MobileOnlineTimeResponse {
	s.Headers = v
	return s
}

func (s *MobileOnlineTimeResponse) SetStatusCode(v int32) *MobileOnlineTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *MobileOnlineTimeResponse) SetBody(v *MobileOnlineTimeResponseBody) *MobileOnlineTimeResponse {
	s.Body = v
	return s
}

func (s *MobileOnlineTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
