// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacPortalSmsPhoneWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetNacPortalSmsPhoneWhitelistRequest struct {
}

func (s GetNacPortalSmsPhoneWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNacPortalSmsPhoneWhitelistRequest) GoString() string {
	return s.String()
}

func (s *GetNacPortalSmsPhoneWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
