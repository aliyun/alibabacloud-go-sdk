// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListQueueSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudListQueueSkillRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudListQueueSkillRequest
	GetOwnerId() *int64
	SetQno(v string) *CloudListQueueSkillRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *CloudListQueueSkillRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListQueueSkillRequest
	GetResourceOwnerId() *int64
}

type CloudListQueueSkillRequest struct {
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
}

func (s CloudListQueueSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueSkillRequest) GoString() string {
	return s.String()
}

func (s *CloudListQueueSkillRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListQueueSkillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListQueueSkillRequest) GetQno() *string {
	return s.Qno
}

func (s *CloudListQueueSkillRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListQueueSkillRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListQueueSkillRequest) SetEnterpriseId(v int64) *CloudListQueueSkillRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListQueueSkillRequest) SetOwnerId(v int64) *CloudListQueueSkillRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListQueueSkillRequest) SetQno(v string) *CloudListQueueSkillRequest {
	s.Qno = &v
	return s
}

func (s *CloudListQueueSkillRequest) SetResourceOwnerAccount(v string) *CloudListQueueSkillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListQueueSkillRequest) SetResourceOwnerId(v int64) *CloudListQueueSkillRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListQueueSkillRequest) Validate() error {
	return dara.Validate(s)
}
