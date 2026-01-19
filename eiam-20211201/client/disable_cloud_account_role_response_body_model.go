// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCloudAccountRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCloudAccountRoleResponseBody
	GetRequestId() *string
}

type DisableCloudAccountRoleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCloudAccountRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCloudAccountRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCloudAccountRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCloudAccountRoleResponseBody) SetRequestId(v string) *DisableCloudAccountRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCloudAccountRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
