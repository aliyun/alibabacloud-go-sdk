// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *AddGroupRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *AddGroupRequest
	GetBizExtend() map[string]interface{}
	SetContactDetails(v string) *AddGroupRequest
	GetContactDetails() *string
	SetContactName(v string) *AddGroupRequest
	GetContactName() *string
	SetCountry(v string) *AddGroupRequest
	GetCountry() *string
	SetEmail(v string) *AddGroupRequest
	GetEmail() *string
	SetFilePath(v string) *AddGroupRequest
	GetFilePath() *string
	SetGroupName(v string) *AddGroupRequest
	GetGroupName() *string
	SetOwnerId(v int64) *AddGroupRequest
	GetOwnerId() *int64
	SetRemark(v string) *AddGroupRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *AddGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddGroupRequest
	GetResourceOwnerId() *int64
}

type AddGroupRequest struct {
	// example:
	//
	// 示例值示例值示例值
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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

func (s AddGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGroupRequest) GoString() string {
	return s.String()
}

func (s *AddGroupRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *AddGroupRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *AddGroupRequest) GetContactDetails() *string {
	return s.ContactDetails
}

func (s *AddGroupRequest) GetContactName() *string {
	return s.ContactName
}

func (s *AddGroupRequest) GetCountry() *string {
	return s.Country
}

func (s *AddGroupRequest) GetEmail() *string {
	return s.Email
}

func (s *AddGroupRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *AddGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddGroupRequest) SetBizCode(v string) *AddGroupRequest {
	s.BizCode = &v
	return s
}

func (s *AddGroupRequest) SetBizExtend(v map[string]interface{}) *AddGroupRequest {
	s.BizExtend = v
	return s
}

func (s *AddGroupRequest) SetContactDetails(v string) *AddGroupRequest {
	s.ContactDetails = &v
	return s
}

func (s *AddGroupRequest) SetContactName(v string) *AddGroupRequest {
	s.ContactName = &v
	return s
}

func (s *AddGroupRequest) SetCountry(v string) *AddGroupRequest {
	s.Country = &v
	return s
}

func (s *AddGroupRequest) SetEmail(v string) *AddGroupRequest {
	s.Email = &v
	return s
}

func (s *AddGroupRequest) SetFilePath(v string) *AddGroupRequest {
	s.FilePath = &v
	return s
}

func (s *AddGroupRequest) SetGroupName(v string) *AddGroupRequest {
	s.GroupName = &v
	return s
}

func (s *AddGroupRequest) SetOwnerId(v int64) *AddGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *AddGroupRequest) SetRemark(v string) *AddGroupRequest {
	s.Remark = &v
	return s
}

func (s *AddGroupRequest) SetResourceOwnerAccount(v string) *AddGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddGroupRequest) SetResourceOwnerId(v int64) *AddGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddGroupRequest) Validate() error {
	return dara.Validate(s)
}
