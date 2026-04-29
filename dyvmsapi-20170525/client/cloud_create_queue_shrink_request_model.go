// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateQueueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudCreateQueueShrinkRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudCreateQueueShrinkRequest
	GetOwnerId() *int64
	SetQueueShrink(v string) *CloudCreateQueueShrinkRequest
	GetQueueShrink() *string
	SetQueueSkillsShrink(v string) *CloudCreateQueueShrinkRequest
	GetQueueSkillsShrink() *string
	SetResourceOwnerAccount(v string) *CloudCreateQueueShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudCreateQueueShrinkRequest
	GetResourceOwnerId() *int64
}

type CloudCreateQueueShrinkRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列信息
	//
	// This parameter is required.
	QueueShrink *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// 队列技能
	//
	// This parameter is required.
	QueueSkillsShrink    *string `json:"QueueSkills,omitempty" xml:"QueueSkills,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudCreateQueueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateQueueShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudCreateQueueShrinkRequest) GetQueueShrink() *string {
	return s.QueueShrink
}

func (s *CloudCreateQueueShrinkRequest) GetQueueSkillsShrink() *string {
	return s.QueueSkillsShrink
}

func (s *CloudCreateQueueShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudCreateQueueShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudCreateQueueShrinkRequest) SetEnterpriseId(v int64) *CloudCreateQueueShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateQueueShrinkRequest) SetOwnerId(v int64) *CloudCreateQueueShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudCreateQueueShrinkRequest) SetQueueShrink(v string) *CloudCreateQueueShrinkRequest {
	s.QueueShrink = &v
	return s
}

func (s *CloudCreateQueueShrinkRequest) SetQueueSkillsShrink(v string) *CloudCreateQueueShrinkRequest {
	s.QueueSkillsShrink = &v
	return s
}

func (s *CloudCreateQueueShrinkRequest) SetResourceOwnerAccount(v string) *CloudCreateQueueShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudCreateQueueShrinkRequest) SetResourceOwnerId(v int64) *CloudCreateQueueShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudCreateQueueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
