// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *AddGroupShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *AddGroupShrinkRequest
	GetBizExtendShrink() *string
	SetContactDetails(v string) *AddGroupShrinkRequest
	GetContactDetails() *string
	SetContactName(v string) *AddGroupShrinkRequest
	GetContactName() *string
	SetCountry(v string) *AddGroupShrinkRequest
	GetCountry() *string
	SetEmail(v string) *AddGroupShrinkRequest
	GetEmail() *string
	SetFilePath(v string) *AddGroupShrinkRequest
	GetFilePath() *string
	SetGroupName(v string) *AddGroupShrinkRequest
	GetGroupName() *string
	SetOwnerId(v int64) *AddGroupShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *AddGroupShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *AddGroupShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddGroupShrinkRequest
	GetResourceOwnerId() *int64
}

type AddGroupShrinkRequest struct {
	// example:
	//
	// 示例值示例值示例值
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// example:
	//
	// 15111111111
	ContactDetails *string `json:"ContactDetails,omitempty" xml:"ContactDetails,omitempty"`
	// example:
	//
	// mary
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// China
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// 示例值示例值
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// src/main/resources/config/promql_node.yaml
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testgroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// test
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGroupShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *AddGroupShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *AddGroupShrinkRequest) GetContactDetails() *string {
	return s.ContactDetails
}

func (s *AddGroupShrinkRequest) GetContactName() *string {
	return s.ContactName
}

func (s *AddGroupShrinkRequest) GetCountry() *string {
	return s.Country
}

func (s *AddGroupShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *AddGroupShrinkRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *AddGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddGroupShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddGroupShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddGroupShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddGroupShrinkRequest) SetBizCode(v string) *AddGroupShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *AddGroupShrinkRequest) SetBizExtendShrink(v string) *AddGroupShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *AddGroupShrinkRequest) SetContactDetails(v string) *AddGroupShrinkRequest {
	s.ContactDetails = &v
	return s
}

func (s *AddGroupShrinkRequest) SetContactName(v string) *AddGroupShrinkRequest {
	s.ContactName = &v
	return s
}

func (s *AddGroupShrinkRequest) SetCountry(v string) *AddGroupShrinkRequest {
	s.Country = &v
	return s
}

func (s *AddGroupShrinkRequest) SetEmail(v string) *AddGroupShrinkRequest {
	s.Email = &v
	return s
}

func (s *AddGroupShrinkRequest) SetFilePath(v string) *AddGroupShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *AddGroupShrinkRequest) SetGroupName(v string) *AddGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *AddGroupShrinkRequest) SetOwnerId(v int64) *AddGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddGroupShrinkRequest) SetRemark(v string) *AddGroupShrinkRequest {
	s.Remark = &v
	return s
}

func (s *AddGroupShrinkRequest) SetResourceOwnerAccount(v string) *AddGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddGroupShrinkRequest) SetResourceOwnerId(v int64) *AddGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
