// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinVSwitchesToEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinVSwitchesToEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinVSwitchesToEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *JoinVSwitchesToEpnInstanceResponseBody) *JoinVSwitchesToEpnInstanceResponse
	GetBody() *JoinVSwitchesToEpnInstanceResponseBody
}

type JoinVSwitchesToEpnInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinVSwitchesToEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinVSwitchesToEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinVSwitchesToEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *JoinVSwitchesToEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinVSwitchesToEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinVSwitchesToEpnInstanceResponse) GetBody() *JoinVSwitchesToEpnInstanceResponseBody {
	return s.Body
}

func (s *JoinVSwitchesToEpnInstanceResponse) SetHeaders(v map[string]*string) *JoinVSwitchesToEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *JoinVSwitchesToEpnInstanceResponse) SetStatusCode(v int32) *JoinVSwitchesToEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinVSwitchesToEpnInstanceResponse) SetBody(v *JoinVSwitchesToEpnInstanceResponseBody) *JoinVSwitchesToEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *JoinVSwitchesToEpnInstanceResponse) Validate() error {
	return dara.Validate(s)
}
