// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudGetQueueRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudGetQueueRequest
	GetOwnerId() *int64
	SetQno(v string) *CloudGetQueueRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *CloudGetQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudGetQueueRequest
	GetResourceOwnerId() *int64
}

type CloudGetQueueRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列编号
	//
	// This parameter is required.
	//
	// example:
	//
	// 355
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudGetQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetQueueRequest) GoString() string {
	return s.String()
}

func (s *CloudGetQueueRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudGetQueueRequest) GetQno() *string {
	return s.Qno
}

func (s *CloudGetQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudGetQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudGetQueueRequest) SetEnterpriseId(v int64) *CloudGetQueueRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetQueueRequest) SetOwnerId(v int64) *CloudGetQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudGetQueueRequest) SetQno(v string) *CloudGetQueueRequest {
	s.Qno = &v
	return s
}

func (s *CloudGetQueueRequest) SetResourceOwnerAccount(v string) *CloudGetQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudGetQueueRequest) SetResourceOwnerId(v int64) *CloudGetQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudGetQueueRequest) Validate() error {
	return dara.Validate(s)
}
