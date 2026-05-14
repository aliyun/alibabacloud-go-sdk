// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListQueuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkListQueuesRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *ClinkListQueuesRequest
	GetLimit() *int64
	SetOffset(v int64) *ClinkListQueuesRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkListQueuesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkListQueuesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListQueuesRequest
	GetResourceOwnerId() *int64
}

type ClinkListQueuesRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 查询记录条数，默认只为 10，最大范围100
	//
	// example:
	//
	// 33
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 偏移量，默认值为 0，最大范围10000
	//
	// example:
	//
	// 0
	Offset               *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkListQueuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListQueuesRequest) GoString() string {
	return s.String()
}

func (s *ClinkListQueuesRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListQueuesRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkListQueuesRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkListQueuesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListQueuesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListQueuesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListQueuesRequest) SetEnterpriseId(v int64) *ClinkListQueuesRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListQueuesRequest) SetLimit(v int64) *ClinkListQueuesRequest {
	s.Limit = &v
	return s
}

func (s *ClinkListQueuesRequest) SetOffset(v int64) *ClinkListQueuesRequest {
	s.Offset = &v
	return s
}

func (s *ClinkListQueuesRequest) SetOwnerId(v int64) *ClinkListQueuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListQueuesRequest) SetResourceOwnerAccount(v string) *ClinkListQueuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListQueuesRequest) SetResourceOwnerId(v int64) *ClinkListQueuesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListQueuesRequest) Validate() error {
	return dara.Validate(s)
}
