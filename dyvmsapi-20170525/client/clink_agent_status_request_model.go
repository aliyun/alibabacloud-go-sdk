// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkAgentStatusRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *ClinkAgentStatusRequest
	GetLimit() *int64
	SetOffset(v int64) *ClinkAgentStatusRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkAgentStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkAgentStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkAgentStatusRequest
	GetResourceOwnerId() *int64
}

type ClinkAgentStatusRequest struct {
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
	// 10
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

func (s ClinkAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkAgentStatusRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkAgentStatusRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkAgentStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkAgentStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkAgentStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkAgentStatusRequest) SetEnterpriseId(v int64) *ClinkAgentStatusRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkAgentStatusRequest) SetLimit(v int64) *ClinkAgentStatusRequest {
	s.Limit = &v
	return s
}

func (s *ClinkAgentStatusRequest) SetOffset(v int64) *ClinkAgentStatusRequest {
	s.Offset = &v
	return s
}

func (s *ClinkAgentStatusRequest) SetOwnerId(v int64) *ClinkAgentStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkAgentStatusRequest) SetResourceOwnerAccount(v string) *ClinkAgentStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkAgentStatusRequest) SetResourceOwnerId(v int64) *ClinkAgentStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
