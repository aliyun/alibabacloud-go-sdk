// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactMail(v string) *UpdateInstanceRequest
	GetContactMail() *string
	SetCountryId(v string) *UpdateInstanceRequest
	GetCountryId() *string
	SetFacebookBmId(v string) *UpdateInstanceRequest
	GetFacebookBmId() *string
	SetInstanceDescription(v string) *UpdateInstanceRequest
	GetInstanceDescription() *string
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateInstanceRequest
	GetInstanceName() *string
	SetIsConfirmAudit(v string) *UpdateInstanceRequest
	GetIsConfirmAudit() *string
	SetIsvTerms(v string) *UpdateInstanceRequest
	GetIsvTerms() *string
	SetOfficeAddress(v string) *UpdateInstanceRequest
	GetOfficeAddress() *string
	SetOwnerId(v int64) *UpdateInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateInstanceRequest
	GetResourceOwnerId() *int64
}

type UpdateInstanceRequest struct {
	// example:
	//
	// 123@alibaba.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// example:
	//
	// 1
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// fb bmId
	//
	// example:
	//
	// 3939982828
	FacebookBmId *string `json:"FacebookBmId,omitempty" xml:"FacebookBmId,omitempty"`
	// example:
	//
	// ins
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 82838838****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// https://aa.com/a.pdf
	IsvTerms *string `json:"IsvTerms,omitempty" xml:"IsvTerms,omitempty"`
	// example:
	//
	// 长沙麓谷
	OfficeAddress        *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetContactMail() *string {
	return s.ContactMail
}

func (s *UpdateInstanceRequest) GetCountryId() *string {
	return s.CountryId
}

func (s *UpdateInstanceRequest) GetFacebookBmId() *string {
	return s.FacebookBmId
}

func (s *UpdateInstanceRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceRequest) GetIsConfirmAudit() *string {
	return s.IsConfirmAudit
}

func (s *UpdateInstanceRequest) GetIsvTerms() *string {
	return s.IsvTerms
}

func (s *UpdateInstanceRequest) GetOfficeAddress() *string {
	return s.OfficeAddress
}

func (s *UpdateInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateInstanceRequest) SetContactMail(v string) *UpdateInstanceRequest {
	s.ContactMail = &v
	return s
}

func (s *UpdateInstanceRequest) SetCountryId(v string) *UpdateInstanceRequest {
	s.CountryId = &v
	return s
}

func (s *UpdateInstanceRequest) SetFacebookBmId(v string) *UpdateInstanceRequest {
	s.FacebookBmId = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceDescription(v string) *UpdateInstanceRequest {
	s.InstanceDescription = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetIsConfirmAudit(v string) *UpdateInstanceRequest {
	s.IsConfirmAudit = &v
	return s
}

func (s *UpdateInstanceRequest) SetIsvTerms(v string) *UpdateInstanceRequest {
	s.IsvTerms = &v
	return s
}

func (s *UpdateInstanceRequest) SetOfficeAddress(v string) *UpdateInstanceRequest {
	s.OfficeAddress = &v
	return s
}

func (s *UpdateInstanceRequest) SetOwnerId(v int64) *UpdateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateInstanceRequest) SetResourceOwnerAccount(v string) *UpdateInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateInstanceRequest) SetResourceOwnerId(v int64) *UpdateInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
