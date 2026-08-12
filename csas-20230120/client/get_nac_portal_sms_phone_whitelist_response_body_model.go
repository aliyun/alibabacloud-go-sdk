// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacPortalSmsPhoneWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhones(v []*string) *GetNacPortalSmsPhoneWhitelistResponseBody
	GetPhones() []*string
	SetRequestId(v string) *GetNacPortalSmsPhoneWhitelistResponseBody
	GetRequestId() *string
}

type GetNacPortalSmsPhoneWhitelistResponseBody struct {
	// The list of phone numbers.
	Phones []*string `json:"Phones,omitempty" xml:"Phones,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B608C6AE-623D-55C4-9454-601B88AE937E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNacPortalSmsPhoneWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNacPortalSmsPhoneWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *GetNacPortalSmsPhoneWhitelistResponseBody) GetPhones() []*string {
	return s.Phones
}

func (s *GetNacPortalSmsPhoneWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNacPortalSmsPhoneWhitelistResponseBody) SetPhones(v []*string) *GetNacPortalSmsPhoneWhitelistResponseBody {
	s.Phones = v
	return s
}

func (s *GetNacPortalSmsPhoneWhitelistResponseBody) SetRequestId(v string) *GetNacPortalSmsPhoneWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNacPortalSmsPhoneWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
