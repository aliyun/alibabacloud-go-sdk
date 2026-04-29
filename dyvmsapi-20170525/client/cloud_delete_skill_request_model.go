// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudDeleteSkillRequest
	GetEnterpriseId() *int64
	SetId(v int64) *CloudDeleteSkillRequest
	GetId() *int64
	SetName(v string) *CloudDeleteSkillRequest
	GetName() *string
	SetOwnerId(v int64) *CloudDeleteSkillRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudDeleteSkillRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteSkillRequest
	GetResourceOwnerId() *int64
}

type CloudDeleteSkillRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 技能id；正整数，id和name可选，但不能同时为空。同时都传时，以id为准
	//
	// example:
	//
	// 54
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 技能名称；id和name可选，但不能同时为空。同时都传时，以id为准
	//
	// example:
	//
	// name3
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudDeleteSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteSkillRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteSkillRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteSkillRequest) GetId() *int64 {
	return s.Id
}

func (s *CloudDeleteSkillRequest) GetName() *string {
	return s.Name
}

func (s *CloudDeleteSkillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteSkillRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteSkillRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteSkillRequest) SetEnterpriseId(v int64) *CloudDeleteSkillRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteSkillRequest) SetId(v int64) *CloudDeleteSkillRequest {
	s.Id = &v
	return s
}

func (s *CloudDeleteSkillRequest) SetName(v string) *CloudDeleteSkillRequest {
	s.Name = &v
	return s
}

func (s *CloudDeleteSkillRequest) SetOwnerId(v int64) *CloudDeleteSkillRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteSkillRequest) SetResourceOwnerAccount(v string) *CloudDeleteSkillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteSkillRequest) SetResourceOwnerId(v int64) *CloudDeleteSkillRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteSkillRequest) Validate() error {
	return dara.Validate(s)
}
