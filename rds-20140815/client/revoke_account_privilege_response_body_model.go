// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAccountPrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeAccountPrivilegeResponseBody
	GetRequestId() *string
}

type RevokeAccountPrivilegeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E22099CA-A61E-4992-A0B7-CE82DC175626
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeAccountPrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeAccountPrivilegeResponseBody) SetRequestId(v string) *RevokeAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeAccountPrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}
