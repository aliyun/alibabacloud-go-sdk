// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinPublicIpsToEpnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *JoinPublicIpsToEpnInstanceResponseBody
	GetRequestId() *string
}

type JoinPublicIpsToEpnInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinPublicIpsToEpnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinPublicIpsToEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *JoinPublicIpsToEpnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinPublicIpsToEpnInstanceResponseBody) SetRequestId(v string) *JoinPublicIpsToEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinPublicIpsToEpnInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
