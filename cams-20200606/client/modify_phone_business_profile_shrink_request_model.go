// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhoneBusinessProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbout(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetAbout() *string
	SetAddress(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetAddress() *string
	SetCustSpaceId(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetCustSpaceId() *string
	SetDescription(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetDescription() *string
	SetEmail(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetEmail() *string
	SetOwnerId(v int64) *ModifyPhoneBusinessProfileShrinkRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetPhoneNumber() *string
	SetProfilePictureUrl(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetProfilePictureUrl() *string
	SetResourceOwnerAccount(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPhoneBusinessProfileShrinkRequest
	GetResourceOwnerId() *int64
	SetVertical(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetVertical() *string
	SetWebsitesShrink(v string) *ModifyPhoneBusinessProfileShrinkRequest
	GetWebsitesShrink() *string
}

type ModifyPhoneBusinessProfileShrinkRequest struct {
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
	WebsitesShrink *string `json:"Websites,omitempty" xml:"Websites,omitempty"`
}

func (s ModifyPhoneBusinessProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhoneBusinessProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetAbout() *string {
	return s.About
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetProfilePictureUrl() *string {
	return s.ProfilePictureUrl
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetVertical() *string {
	return s.Vertical
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) GetWebsitesShrink() *string {
	return s.WebsitesShrink
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetAbout(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.About = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetAddress(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.Address = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetCustSpaceId(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetDescription(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetEmail(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.Email = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetOwnerId(v int64) *ModifyPhoneBusinessProfileShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetPhoneNumber(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetProfilePictureUrl(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.ProfilePictureUrl = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetResourceOwnerAccount(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetResourceOwnerId(v int64) *ModifyPhoneBusinessProfileShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetVertical(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.Vertical = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetWebsitesShrink(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.WebsitesShrink = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
