// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCSecurityGroupPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCSecurityGroupPermissionResponseBody
	GetRequestId() *string
}

type ModifyRCSecurityGroupPermissionResponseBody struct {
	// example:
	//
	// AB44DC0A-7E77-442A-97A9-C6418694CB22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCSecurityGroupPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCSecurityGroupPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCSecurityGroupPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCSecurityGroupPermissionResponseBody) SetRequestId(v string) *ModifyRCSecurityGroupPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
