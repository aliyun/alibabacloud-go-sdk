// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkDescribeQueueRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *ClinkDescribeQueueRequest
	GetOwnerId() *int64
	SetQno(v string) *ClinkDescribeQueueRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *ClinkDescribeQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDescribeQueueRequest
	GetResourceOwnerId() *int64
}

type ClinkDescribeQueueRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列号
	//
	// This parameter is required.
	//
	// example:
	//
	// 2233
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDescribeQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeQueueRequest) GoString() string {
	return s.String()
}

func (s *ClinkDescribeQueueRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDescribeQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDescribeQueueRequest) GetQno() *string {
	return s.Qno
}

func (s *ClinkDescribeQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDescribeQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDescribeQueueRequest) SetEnterpriseId(v int64) *ClinkDescribeQueueRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDescribeQueueRequest) SetOwnerId(v int64) *ClinkDescribeQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDescribeQueueRequest) SetQno(v string) *ClinkDescribeQueueRequest {
	s.Qno = &v
	return s
}

func (s *ClinkDescribeQueueRequest) SetResourceOwnerAccount(v string) *ClinkDescribeQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDescribeQueueRequest) SetResourceOwnerId(v int64) *ClinkDescribeQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDescribeQueueRequest) Validate() error {
	return dara.Validate(s)
}
