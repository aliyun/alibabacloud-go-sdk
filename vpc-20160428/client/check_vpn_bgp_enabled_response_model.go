// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVpnBgpEnabledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckVpnBgpEnabledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckVpnBgpEnabledResponse
	GetStatusCode() *int32
	SetBody(v *CheckVpnBgpEnabledResponseBody) *CheckVpnBgpEnabledResponse
	GetBody() *CheckVpnBgpEnabledResponseBody
}

type CheckVpnBgpEnabledResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckVpnBgpEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckVpnBgpEnabledResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckVpnBgpEnabledResponse) GoString() string {
	return s.String()
}

func (s *CheckVpnBgpEnabledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckVpnBgpEnabledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckVpnBgpEnabledResponse) GetBody() *CheckVpnBgpEnabledResponseBody {
	return s.Body
}

func (s *CheckVpnBgpEnabledResponse) SetHeaders(v map[string]*string) *CheckVpnBgpEnabledResponse {
	s.Headers = v
	return s
}

func (s *CheckVpnBgpEnabledResponse) SetStatusCode(v int32) *CheckVpnBgpEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckVpnBgpEnabledResponse) SetBody(v *CheckVpnBgpEnabledResponseBody) *CheckVpnBgpEnabledResponse {
	s.Body = v
	return s
}

func (s *CheckVpnBgpEnabledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
