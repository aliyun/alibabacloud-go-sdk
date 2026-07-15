// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserBusinessFormRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompany(v string) *AddUserBusinessFormRequest
	GetCompany() *string
	SetEmail(v string) *AddUserBusinessFormRequest
	GetEmail() *string
	SetPhoneNumber(v string) *AddUserBusinessFormRequest
	GetPhoneNumber() *string
	SetPosition(v string) *AddUserBusinessFormRequest
	GetPosition() *string
	SetRemark(v string) *AddUserBusinessFormRequest
	GetRemark() *string
	SetUserName(v string) *AddUserBusinessFormRequest
	GetUserName() *string
	SetWebsite(v string) *AddUserBusinessFormRequest
	GetWebsite() *string
}

type AddUserBusinessFormRequest struct {
	// The company.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx有限公司
	Company *string `json:"Company,omitempty" xml:"Company,omitempty"`
	// The email address.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx@alibaba.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 158********
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The job title.
	//
	// This parameter is required.
	//
	// example:
	//
	// 经理
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// Additional remarks.
	//
	// example:
	//
	// 请尽快联系我们
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// 张三
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The company website.
	//
	// example:
	//
	// xxx.com
	Website *string `json:"Website,omitempty" xml:"Website,omitempty"`
}

func (s AddUserBusinessFormRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserBusinessFormRequest) GoString() string {
	return s.String()
}

func (s *AddUserBusinessFormRequest) GetCompany() *string {
	return s.Company
}

func (s *AddUserBusinessFormRequest) GetEmail() *string {
	return s.Email
}

func (s *AddUserBusinessFormRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *AddUserBusinessFormRequest) GetPosition() *string {
	return s.Position
}

func (s *AddUserBusinessFormRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddUserBusinessFormRequest) GetUserName() *string {
	return s.UserName
}

func (s *AddUserBusinessFormRequest) GetWebsite() *string {
	return s.Website
}

func (s *AddUserBusinessFormRequest) SetCompany(v string) *AddUserBusinessFormRequest {
	s.Company = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetEmail(v string) *AddUserBusinessFormRequest {
	s.Email = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetPhoneNumber(v string) *AddUserBusinessFormRequest {
	s.PhoneNumber = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetPosition(v string) *AddUserBusinessFormRequest {
	s.Position = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetRemark(v string) *AddUserBusinessFormRequest {
	s.Remark = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetUserName(v string) *AddUserBusinessFormRequest {
	s.UserName = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetWebsite(v string) *AddUserBusinessFormRequest {
	s.Website = &v
	return s
}

func (s *AddUserBusinessFormRequest) Validate() error {
	return dara.Validate(s)
}
