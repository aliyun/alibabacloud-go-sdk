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
}

type UpdateInstanceRequest struct {
	// The contact email address.
	//
	// example:
	//
	// 123@alibaba.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// The country code.
	//
	// > For a list of country codes, see [Country Codes](https://help.aliyun.com/document_detail/608210.html).
	//
	// example:
	//
	// 1
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The Facebook Business Manager ID.
	//
	// example:
	//
	// 393998****
	FacebookBmId *string `json:"FacebookBmId,omitempty" xml:"FacebookBmId,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// ins
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 82838838****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// viber_ins
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Specifies whether to enable automatic audit confirmation.
	//
	// example:
	//
	// Y
	IsConfirmAudit *string `json:"IsConfirmAudit,omitempty" xml:"IsConfirmAudit,omitempty"`
	// The URL of the Independent Software Vendor (ISV) agreement file.
	//
	// example:
	//
	// https://aa.com/a.pdf
	IsvTerms *string `json:"IsvTerms,omitempty" xml:"IsvTerms,omitempty"`
	// The company address.
	//
	// example:
	//
	// example
	OfficeAddress *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
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

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
