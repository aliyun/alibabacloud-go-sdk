// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacPortalSmsPhoneWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPhones(v []*string) *UpdateNacPortalSmsPhoneWhitelistRequest
	GetPhones() []*string
}

type UpdateNacPortalSmsPhoneWhitelistRequest struct {
	// The list of phone numbers.
	Phones []*string `json:"Phones,omitempty" xml:"Phones,omitempty" type:"Repeated"`
}

func (s UpdateNacPortalSmsPhoneWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacPortalSmsPhoneWhitelistRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacPortalSmsPhoneWhitelistRequest) GetPhones() []*string {
	return s.Phones
}

func (s *UpdateNacPortalSmsPhoneWhitelistRequest) SetPhones(v []*string) *UpdateNacPortalSmsPhoneWhitelistRequest {
	s.Phones = v
	return s
}

func (s *UpdateNacPortalSmsPhoneWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
