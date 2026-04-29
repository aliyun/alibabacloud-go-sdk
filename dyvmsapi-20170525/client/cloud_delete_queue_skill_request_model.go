// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteQueueSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudDeleteQueueSkillRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudDeleteQueueSkillRequest
	GetOwnerId() *int64
	SetQno(v string) *CloudDeleteQueueSkillRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *CloudDeleteQueueSkillRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteQueueSkillRequest
	GetResourceOwnerId() *int64
	SetSkillId(v int64) *CloudDeleteQueueSkillRequest
	GetSkillId() *int64
}

type CloudDeleteQueueSkillRequest struct {
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
	// 10000
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// skill的id
	//
	// This parameter is required.
	//
	// example:
	//
	// 76
	SkillId *int64 `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s CloudDeleteQueueSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteQueueSkillRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteQueueSkillRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteQueueSkillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteQueueSkillRequest) GetQno() *string {
	return s.Qno
}

func (s *CloudDeleteQueueSkillRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteQueueSkillRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteQueueSkillRequest) GetSkillId() *int64 {
	return s.SkillId
}

func (s *CloudDeleteQueueSkillRequest) SetEnterpriseId(v int64) *CloudDeleteQueueSkillRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteQueueSkillRequest) SetOwnerId(v int64) *CloudDeleteQueueSkillRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteQueueSkillRequest) SetQno(v string) *CloudDeleteQueueSkillRequest {
	s.Qno = &v
	return s
}

func (s *CloudDeleteQueueSkillRequest) SetResourceOwnerAccount(v string) *CloudDeleteQueueSkillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteQueueSkillRequest) SetResourceOwnerId(v int64) *CloudDeleteQueueSkillRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteQueueSkillRequest) SetSkillId(v int64) *CloudDeleteQueueSkillRequest {
	s.SkillId = &v
	return s
}

func (s *CloudDeleteQueueSkillRequest) Validate() error {
	return dara.Validate(s)
}
