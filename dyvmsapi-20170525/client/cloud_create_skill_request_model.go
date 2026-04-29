// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CloudCreateSkillRequest
	GetComment() *string
	SetEnterpriseId(v int64) *CloudCreateSkillRequest
	GetEnterpriseId() *int64
	SetName(v string) *CloudCreateSkillRequest
	GetName() *string
	SetOwnerId(v int64) *CloudCreateSkillRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudCreateSkillRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudCreateSkillRequest
	GetResourceOwnerId() *int64
}

type CloudCreateSkillRequest struct {
	// 描述
	//
	// example:
	//
	// Comment3
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 技能名称
	//
	// This parameter is required.
	//
	// example:
	//
	// skillName
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudCreateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateSkillRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateSkillRequest) GetComment() *string {
	return s.Comment
}

func (s *CloudCreateSkillRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateSkillRequest) GetName() *string {
	return s.Name
}

func (s *CloudCreateSkillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudCreateSkillRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudCreateSkillRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudCreateSkillRequest) SetComment(v string) *CloudCreateSkillRequest {
	s.Comment = &v
	return s
}

func (s *CloudCreateSkillRequest) SetEnterpriseId(v int64) *CloudCreateSkillRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateSkillRequest) SetName(v string) *CloudCreateSkillRequest {
	s.Name = &v
	return s
}

func (s *CloudCreateSkillRequest) SetOwnerId(v int64) *CloudCreateSkillRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudCreateSkillRequest) SetResourceOwnerAccount(v string) *CloudCreateSkillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudCreateSkillRequest) SetResourceOwnerId(v int64) *CloudCreateSkillRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudCreateSkillRequest) Validate() error {
	return dara.Validate(s)
}
