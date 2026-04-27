// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnos(v string) *CloudQueryAgentSkillRequest
	GetCnos() *string
	SetEnterpriseId(v int64) *CloudQueryAgentSkillRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudQueryAgentSkillRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudQueryAgentSkillRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudQueryAgentSkillRequest
	GetResourceOwnerId() *int64
}

type CloudQueryAgentSkillRequest struct {
	// 座席工号，可传多个，中间以英文逗号分割
	//
	// This parameter is required.
	//
	// example:
	//
	// 111,222
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudQueryAgentSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentSkillRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentSkillRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudQueryAgentSkillRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryAgentSkillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudQueryAgentSkillRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudQueryAgentSkillRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudQueryAgentSkillRequest) SetCnos(v string) *CloudQueryAgentSkillRequest {
	s.Cnos = &v
	return s
}

func (s *CloudQueryAgentSkillRequest) SetEnterpriseId(v int64) *CloudQueryAgentSkillRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryAgentSkillRequest) SetOwnerId(v int64) *CloudQueryAgentSkillRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudQueryAgentSkillRequest) SetResourceOwnerAccount(v string) *CloudQueryAgentSkillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudQueryAgentSkillRequest) SetResourceOwnerId(v int64) *CloudQueryAgentSkillRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudQueryAgentSkillRequest) Validate() error {
	return dara.Validate(s)
}
