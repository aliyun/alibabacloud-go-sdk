// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacPortalSmsPhoneWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNacPortalSmsPhoneWhitelistResponseBody
	GetRequestId() *string
}

type UpdateNacPortalSmsPhoneWhitelistResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNacPortalSmsPhoneWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacPortalSmsPhoneWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponseBody) SetRequestId(v string) *UpdateNacPortalSmsPhoneWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
