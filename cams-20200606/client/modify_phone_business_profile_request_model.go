// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhoneBusinessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbout(v string) *ModifyPhoneBusinessProfileRequest
	GetAbout() *string
	SetAddress(v string) *ModifyPhoneBusinessProfileRequest
	GetAddress() *string
	SetCustSpaceId(v string) *ModifyPhoneBusinessProfileRequest
	GetCustSpaceId() *string
	SetDescription(v string) *ModifyPhoneBusinessProfileRequest
	GetDescription() *string
	SetEmail(v string) *ModifyPhoneBusinessProfileRequest
	GetEmail() *string
	SetOwnerId(v int64) *ModifyPhoneBusinessProfileRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ModifyPhoneBusinessProfileRequest
	GetPhoneNumber() *string
	SetProfilePictureUrl(v string) *ModifyPhoneBusinessProfileRequest
	GetProfilePictureUrl() *string
	SetResourceOwnerAccount(v string) *ModifyPhoneBusinessProfileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPhoneBusinessProfileRequest
	GetResourceOwnerId() *int64
	SetVertical(v string) *ModifyPhoneBusinessProfileRequest
	GetVertical() *string
	SetWebsites(v []*string) *ModifyPhoneBusinessProfileRequest
	GetWebsites() []*string
}

type ModifyPhoneBusinessProfileRequest struct {
	// The business information.
	//
	// example:
	//
	// business profile
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// The address.
	//
	// example:
	//
	// The phone number.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The space ID of the RAM user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The description of the phone number.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The email address.
	//
	// example:
	//
	// aa@aliyun.com
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The URL of the profile picture.
	//
	// example:
	//
	// http://a.img
	ProfilePictureUrl    *string `json:"ProfilePictureUrl,omitempty" xml:"ProfilePictureUrl,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The industry.
	//
	// >  Valid values: OTHER, AUTO, BEAUTY, APPAREL, EDU, ENTERTAIN, EVENT_PLAN, FINANCE, GROCERY, GOVT, HOTEL, HEALTH, NONPROFIT, PROF_SERVICES, RETAIL, TRAVEL, and RESTAURANT.
	//
	// example:
	//
	// OTHER
	Vertical *string `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
	// The URLs of the websites.
	Websites []*string `json:"Websites,omitempty" xml:"Websites,omitempty" type:"Repeated"`
}

func (s ModifyPhoneBusinessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhoneBusinessProfileRequest) GoString() string {
	return s.String()
}

func (s *ModifyPhoneBusinessProfileRequest) GetAbout() *string {
	return s.About
}

func (s *ModifyPhoneBusinessProfileRequest) GetAddress() *string {
	return s.Address
}

func (s *ModifyPhoneBusinessProfileRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ModifyPhoneBusinessProfileRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyPhoneBusinessProfileRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyPhoneBusinessProfileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPhoneBusinessProfileRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ModifyPhoneBusinessProfileRequest) GetProfilePictureUrl() *string {
	return s.ProfilePictureUrl
}

func (s *ModifyPhoneBusinessProfileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPhoneBusinessProfileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPhoneBusinessProfileRequest) GetVertical() *string {
	return s.Vertical
}

func (s *ModifyPhoneBusinessProfileRequest) GetWebsites() []*string {
	return s.Websites
}

func (s *ModifyPhoneBusinessProfileRequest) SetAbout(v string) *ModifyPhoneBusinessProfileRequest {
	s.About = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetAddress(v string) *ModifyPhoneBusinessProfileRequest {
	s.Address = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetCustSpaceId(v string) *ModifyPhoneBusinessProfileRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetDescription(v string) *ModifyPhoneBusinessProfileRequest {
	s.Description = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetEmail(v string) *ModifyPhoneBusinessProfileRequest {
	s.Email = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetOwnerId(v int64) *ModifyPhoneBusinessProfileRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetPhoneNumber(v string) *ModifyPhoneBusinessProfileRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetProfilePictureUrl(v string) *ModifyPhoneBusinessProfileRequest {
	s.ProfilePictureUrl = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetResourceOwnerAccount(v string) *ModifyPhoneBusinessProfileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetResourceOwnerId(v int64) *ModifyPhoneBusinessProfileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetVertical(v string) *ModifyPhoneBusinessProfileRequest {
	s.Vertical = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetWebsites(v []*string) *ModifyPhoneBusinessProfileRequest {
	s.Websites = v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) Validate() error {
	return dara.Validate(s)
}
