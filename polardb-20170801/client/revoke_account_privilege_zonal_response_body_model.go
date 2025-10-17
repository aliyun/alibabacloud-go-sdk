// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAccountPrivilegeZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeAccountPrivilegeZonalResponseBody
	GetRequestId() *string
}

type RevokeAccountPrivilegeZonalResponseBody struct {
	// example:
	//
	// F9F1CB1A-B1D5-4EF5-A53A-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeAccountPrivilegeZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeAccountPrivilegeZonalResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeAccountPrivilegeZonalResponseBody) SetRequestId(v string) *RevokeAccountPrivilegeZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
