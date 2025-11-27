// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRCSecurityGroupPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeRCSecurityGroupPermissionResponseBody
	GetRequestId() *string
}

type RevokeRCSecurityGroupPermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 847BA085-B377-4BFA-8267-F82345ECE1D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeRCSecurityGroupPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeRCSecurityGroupPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeRCSecurityGroupPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeRCSecurityGroupPermissionResponseBody) SetRequestId(v string) *RevokeRCSecurityGroupPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
