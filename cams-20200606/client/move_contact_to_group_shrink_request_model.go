// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveContactToGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *MoveContactToGroupShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *MoveContactToGroupShrinkRequest
	GetBizExtendShrink() *string
	SetContacts(v string) *MoveContactToGroupShrinkRequest
	GetContacts() *string
	SetLinkExistGroups(v string) *MoveContactToGroupShrinkRequest
	GetLinkExistGroups() *string
	SetLinkNewGroups(v string) *MoveContactToGroupShrinkRequest
	GetLinkNewGroups() *string
	SetOwnerId(v int64) *MoveContactToGroupShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MoveContactToGroupShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MoveContactToGroupShrinkRequest
	GetResourceOwnerId() *int64
}

type MoveContactToGroupShrinkRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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

func (s MoveContactToGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveContactToGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *MoveContactToGroupShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *MoveContactToGroupShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *MoveContactToGroupShrinkRequest) GetContacts() *string {
	return s.Contacts
}

func (s *MoveContactToGroupShrinkRequest) GetLinkExistGroups() *string {
	return s.LinkExistGroups
}

func (s *MoveContactToGroupShrinkRequest) GetLinkNewGroups() *string {
	return s.LinkNewGroups
}

func (s *MoveContactToGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MoveContactToGroupShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MoveContactToGroupShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MoveContactToGroupShrinkRequest) SetBizCode(v string) *MoveContactToGroupShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *MoveContactToGroupShrinkRequest) SetBizExtendShrink(v string) *MoveContactToGroupShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *MoveContactToGroupShrinkRequest) SetContacts(v string) *MoveContactToGroupShrinkRequest {
	s.Contacts = &v
	return s
}

func (s *MoveContactToGroupShrinkRequest) SetLinkExistGroups(v string) *MoveContactToGroupShrinkRequest {
	s.LinkExistGroups = &v
	return s
}

func (s *MoveContactToGroupShrinkRequest) SetLinkNewGroups(v string) *MoveContactToGroupShrinkRequest {
	s.LinkNewGroups = &v
	return s
}

func (s *MoveContactToGroupShrinkRequest) SetOwnerId(v int64) *MoveContactToGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *MoveContactToGroupShrinkRequest) SetResourceOwnerAccount(v string) *MoveContactToGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MoveContactToGroupShrinkRequest) SetResourceOwnerId(v int64) *MoveContactToGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MoveContactToGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
