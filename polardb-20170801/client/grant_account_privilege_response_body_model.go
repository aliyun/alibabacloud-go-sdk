// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAccountPrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantAccountPrivilegeResponseBody
	GetRequestId() *string
}

type GrantAccountPrivilegeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantAccountPrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantAccountPrivilegeResponseBody) SetRequestId(v string) *GrantAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantAccountPrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}
