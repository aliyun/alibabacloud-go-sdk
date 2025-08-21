// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinVSwitchesToEpnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *JoinVSwitchesToEpnInstanceResponseBody
	GetRequestId() *string
}

type JoinVSwitchesToEpnInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinVSwitchesToEpnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinVSwitchesToEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *JoinVSwitchesToEpnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinVSwitchesToEpnInstanceResponseBody) SetRequestId(v string) *JoinVSwitchesToEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinVSwitchesToEpnInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
