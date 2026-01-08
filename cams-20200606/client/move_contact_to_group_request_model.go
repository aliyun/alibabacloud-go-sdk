// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveContactToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *MoveContactToGroupRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *MoveContactToGroupRequest
	GetBizExtend() map[string]interface{}
	SetContacts(v string) *MoveContactToGroupRequest
	GetContacts() *string
	SetLinkExistGroups(v string) *MoveContactToGroupRequest
	GetLinkExistGroups() *string
	SetLinkNewGroups(v string) *MoveContactToGroupRequest
	GetLinkNewGroups() *string
	SetOwnerId(v int64) *MoveContactToGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MoveContactToGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MoveContactToGroupRequest
	GetResourceOwnerId() *int64
}

type MoveContactToGroupRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"id":1}]
	Contacts *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	// example:
	//
	// [{"id":1}]
	LinkExistGroups *string `json:"LinkExistGroups,omitempty" xml:"LinkExistGroups,omitempty"`
	// example:
	//
	// [{"groupName":"aaa"}]
	LinkNewGroups        *string `json:"LinkNewGroups,omitempty" xml:"LinkNewGroups,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s MoveContactToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveContactToGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveContactToGroupRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *MoveContactToGroupRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *MoveContactToGroupRequest) GetContacts() *string {
	return s.Contacts
}

func (s *MoveContactToGroupRequest) GetLinkExistGroups() *string {
	return s.LinkExistGroups
}

func (s *MoveContactToGroupRequest) GetLinkNewGroups() *string {
	return s.LinkNewGroups
}

func (s *MoveContactToGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MoveContactToGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MoveContactToGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MoveContactToGroupRequest) SetBizCode(v string) *MoveContactToGroupRequest {
	s.BizCode = &v
	return s
}

func (s *MoveContactToGroupRequest) SetBizExtend(v map[string]interface{}) *MoveContactToGroupRequest {
	s.BizExtend = v
	return s
}

func (s *MoveContactToGroupRequest) SetContacts(v string) *MoveContactToGroupRequest {
	s.Contacts = &v
	return s
}

func (s *MoveContactToGroupRequest) SetLinkExistGroups(v string) *MoveContactToGroupRequest {
	s.LinkExistGroups = &v
	return s
}

func (s *MoveContactToGroupRequest) SetLinkNewGroups(v string) *MoveContactToGroupRequest {
	s.LinkNewGroups = &v
	return s
}

func (s *MoveContactToGroupRequest) SetOwnerId(v int64) *MoveContactToGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *MoveContactToGroupRequest) SetResourceOwnerAccount(v string) *MoveContactToGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MoveContactToGroupRequest) SetResourceOwnerId(v int64) *MoveContactToGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MoveContactToGroupRequest) Validate() error {
	return dara.Validate(s)
}
