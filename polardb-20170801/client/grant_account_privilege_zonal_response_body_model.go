// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAccountPrivilegeZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantAccountPrivilegeZonalResponseBody
	GetRequestId() *string
}

type GrantAccountPrivilegeZonalResponseBody struct {
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantAccountPrivilegeZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantAccountPrivilegeZonalResponseBody) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantAccountPrivilegeZonalResponseBody) SetRequestId(v string) *GrantAccountPrivilegeZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantAccountPrivilegeZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
