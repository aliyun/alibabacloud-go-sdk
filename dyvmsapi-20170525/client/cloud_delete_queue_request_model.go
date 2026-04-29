// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudDeleteQueueRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudDeleteQueueRequest
	GetOwnerId() *int64
	SetQno(v string) *CloudDeleteQueueRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *CloudDeleteQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteQueueRequest
	GetResourceOwnerId() *int64
}

type CloudDeleteQueueRequest struct {
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
	// 1112
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudDeleteQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteQueueRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteQueueRequest) GetQno() *string {
	return s.Qno
}

func (s *CloudDeleteQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteQueueRequest) SetEnterpriseId(v int64) *CloudDeleteQueueRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteQueueRequest) SetOwnerId(v int64) *CloudDeleteQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteQueueRequest) SetQno(v string) *CloudDeleteQueueRequest {
	s.Qno = &v
	return s
}

func (s *CloudDeleteQueueRequest) SetResourceOwnerAccount(v string) *CloudDeleteQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteQueueRequest) SetResourceOwnerId(v int64) *CloudDeleteQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteQueueRequest) Validate() error {
	return dara.Validate(s)
}
