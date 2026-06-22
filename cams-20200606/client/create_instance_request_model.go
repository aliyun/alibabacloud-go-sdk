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
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
}

type CreateInstanceRequest struct {
	// The channel type.
	//
	// This parameter is required.
	//
	// example:
	//
	// VIBER
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The contact email address.
	//
	// example:
	//
	// example
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// The country code.
	//
	// example:
	//
	// 1
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The ID of the Facebook Business Manager (BM).
	//
	// example:
	//
	// 393992929
	FacebookBmId *string `json:"FacebookBmId,omitempty" xml:"FacebookBmId,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// ins
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// viber_ins
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Specifies whether to confirm the audit.
	//
	// example:
	//
	// Y
	IsConfirmAudit *string `json:"IsConfirmAudit,omitempty" xml:"IsConfirmAudit,omitempty"`
	// The URL of the ISV terms file.
	//
	// example:
	//
	// https://a.com/1.pdf
	IsvTerms *string `json:"IsvTerms,omitempty" xml:"IsvTerms,omitempty"`
	// The office address of the business.
	//
	// example:
	//
	// example
	OfficeAddress *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
	// The ID of the resource group that contains the instance.
	//
	// example:
	//
	// example
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
