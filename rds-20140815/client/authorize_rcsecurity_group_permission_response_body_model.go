// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeRCSecurityGroupPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeRCSecurityGroupPermissionResponseBody
	GetRequestId() *string
}

type AuthorizeRCSecurityGroupPermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0688F1D2-CDA8-5617-A43C-ADAC61D80D43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeRCSecurityGroupPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRCSecurityGroupPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeRCSecurityGroupPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeRCSecurityGroupPermissionResponseBody) SetRequestId(v string) *AuthorizeRCSecurityGroupPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
