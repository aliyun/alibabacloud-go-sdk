// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateQueueSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudCreateQueueSkillRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudCreateQueueSkillRequest
	GetOwnerId() *int64
	SetQno(v string) *CloudCreateQueueSkillRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *CloudCreateQueueSkillRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudCreateQueueSkillRequest
	GetResourceOwnerId() *int64
	SetSkillId(v int64) *CloudCreateQueueSkillRequest
	GetSkillId() *int64
	SetSkillLevel(v int64) *CloudCreateQueueSkillRequest
	GetSkillLevel() *int64
}

type CloudCreateQueueSkillRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列号
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// skill的id
	//
	// This parameter is required.
	//
	// example:
	//
	// 17
	SkillId *int64 `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// 技能值
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	SkillLevel *int64 `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s CloudCreateQueueSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueSkillRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueSkillRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateQueueSkillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudCreateQueueSkillRequest) GetQno() *string {
	return s.Qno
}

func (s *CloudCreateQueueSkillRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudCreateQueueSkillRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudCreateQueueSkillRequest) GetSkillId() *int64 {
	return s.SkillId
}

func (s *CloudCreateQueueSkillRequest) GetSkillLevel() *int64 {
	return s.SkillLevel
}

func (s *CloudCreateQueueSkillRequest) SetEnterpriseId(v int64) *CloudCreateQueueSkillRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateQueueSkillRequest) SetOwnerId(v int64) *CloudCreateQueueSkillRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudCreateQueueSkillRequest) SetQno(v string) *CloudCreateQueueSkillRequest {
	s.Qno = &v
	return s
}

func (s *CloudCreateQueueSkillRequest) SetResourceOwnerAccount(v string) *CloudCreateQueueSkillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudCreateQueueSkillRequest) SetResourceOwnerId(v int64) *CloudCreateQueueSkillRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudCreateQueueSkillRequest) SetSkillId(v int64) *CloudCreateQueueSkillRequest {
	s.SkillId = &v
	return s
}

func (s *CloudCreateQueueSkillRequest) SetSkillLevel(v int64) *CloudCreateQueueSkillRequest {
	s.SkillLevel = &v
	return s
}

func (s *CloudCreateQueueSkillRequest) Validate() error {
	return dara.Validate(s)
}
