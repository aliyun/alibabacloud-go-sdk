// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinPublicIpsToEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinPublicIpsToEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinPublicIpsToEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *JoinPublicIpsToEpnInstanceResponseBody) *JoinPublicIpsToEpnInstanceResponse
	GetBody() *JoinPublicIpsToEpnInstanceResponseBody
}

type JoinPublicIpsToEpnInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinPublicIpsToEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinPublicIpsToEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinPublicIpsToEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *JoinPublicIpsToEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinPublicIpsToEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinPublicIpsToEpnInstanceResponse) GetBody() *JoinPublicIpsToEpnInstanceResponseBody {
	return s.Body
}

func (s *JoinPublicIpsToEpnInstanceResponse) SetHeaders(v map[string]*string) *JoinPublicIpsToEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *JoinPublicIpsToEpnInstanceResponse) SetStatusCode(v int32) *JoinPublicIpsToEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinPublicIpsToEpnInstanceResponse) SetBody(v *JoinPublicIpsToEpnInstanceResponseBody) *JoinPublicIpsToEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *JoinPublicIpsToEpnInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
