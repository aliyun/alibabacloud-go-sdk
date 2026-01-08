// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *CreateInstanceRequest
	GetChannelType() *string
	SetContactMail(v string) *CreateInstanceRequest
	GetContactMail() *string
	SetCountryId(v string) *CreateInstanceRequest
	GetCountryId() *string
	SetFacebookBmId(v string) *CreateInstanceRequest
	GetFacebookBmId() *string
	SetInstanceDescription(v string) *CreateInstanceRequest
	GetInstanceDescription() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetIsConfirmAudit(v string) *CreateInstanceRequest
	GetIsConfirmAudit() *string
	SetIsvTerms(v string) *CreateInstanceRequest
	GetIsvTerms() *string
	SetOfficeAddress(v string) *CreateInstanceRequest
	GetOfficeAddress() *string
	SetOwnerId(v int64) *CreateInstanceRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateInstanceRequest
	GetResourceOwnerId() *int64
}

type CreateInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// VIBER
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 示例值示例值
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// example:
	//
	// 1
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// FacebookBmId
	//
	// example:
	//
	// 393992929
	FacebookBmId *string `json:"FacebookBmId,omitempty" xml:"FacebookBmId,omitempty"`
	// example:
	//
	// ins
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// viber_ins
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Y
	IsConfirmAudit *string `json:"IsConfirmAudit,omitempty" xml:"IsConfirmAudit,omitempty"`
	// example:
	//
	// https://a.com/1.pdf
	IsvTerms *string `json:"IsvTerms,omitempty" xml:"IsvTerms,omitempty"`
	// example:
	//
	// 长沙麓谷
	OfficeAddress *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *CreateInstanceRequest) GetContactMail() *string {
	return s.ContactMail
}

func (s *CreateInstanceRequest) GetCountryId() *string {
	return s.CountryId
}

func (s *CreateInstanceRequest) GetFacebookBmId() *string {
	return s.FacebookBmId
}

func (s *CreateInstanceRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetIsConfirmAudit() *string {
	return s.IsConfirmAudit
}

func (s *CreateInstanceRequest) GetIsvTerms() *string {
	return s.IsvTerms
}

func (s *CreateInstanceRequest) GetOfficeAddress() *string {
	return s.OfficeAddress
}

func (s *CreateInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateInstanceRequest) SetChannelType(v string) *CreateInstanceRequest {
	s.ChannelType = &v
	return s
}

func (s *CreateInstanceRequest) SetContactMail(v string) *CreateInstanceRequest {
	s.ContactMail = &v
	return s
}

func (s *CreateInstanceRequest) SetCountryId(v string) *CreateInstanceRequest {
	s.CountryId = &v
	return s
}

func (s *CreateInstanceRequest) SetFacebookBmId(v string) *CreateInstanceRequest {
	s.FacebookBmId = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceDescription(v string) *CreateInstanceRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetIsConfirmAudit(v string) *CreateInstanceRequest {
	s.IsConfirmAudit = &v
	return s
}

func (s *CreateInstanceRequest) SetIsvTerms(v string) *CreateInstanceRequest {
	s.IsvTerms = &v
	return s
}

func (s *CreateInstanceRequest) SetOfficeAddress(v string) *CreateInstanceRequest {
	s.OfficeAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerAccount(v string) *CreateInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerId(v int64) *CreateInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
