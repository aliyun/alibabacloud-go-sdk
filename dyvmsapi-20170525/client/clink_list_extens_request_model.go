// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListExtensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaCode(v string) *ClinkListExtensRequest
	GetAreaCode() *string
	SetEnterpriseId(v int64) *ClinkListExtensRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *ClinkListExtensRequest
	GetLimit() *int64
	SetOffset(v int64) *ClinkListExtensRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkListExtensRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkListExtensRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListExtensRequest
	GetResourceOwnerId() *int64
	SetType(v int64) *ClinkListExtensRequest
	GetType() *int64
}

type ClinkListExtensRequest struct {
	// 话机区号。以 0 开头 3-4 位数字
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 查询条数。范围 10-100，默认值为 10
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 偏移量。 范围 0-10000，默认值为 0
	//
	// example:
	//
	// 0
	Offset               *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 话机类型。1: IP话机， 2: 软电话
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ClinkListExtensRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListExtensRequest) GoString() string {
	return s.String()
}

func (s *ClinkListExtensRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkListExtensRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListExtensRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkListExtensRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkListExtensRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListExtensRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListExtensRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListExtensRequest) GetType() *int64 {
	return s.Type
}

func (s *ClinkListExtensRequest) SetAreaCode(v string) *ClinkListExtensRequest {
	s.AreaCode = &v
	return s
}

func (s *ClinkListExtensRequest) SetEnterpriseId(v int64) *ClinkListExtensRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListExtensRequest) SetLimit(v int64) *ClinkListExtensRequest {
	s.Limit = &v
	return s
}

func (s *ClinkListExtensRequest) SetOffset(v int64) *ClinkListExtensRequest {
	s.Offset = &v
	return s
}

func (s *ClinkListExtensRequest) SetOwnerId(v int64) *ClinkListExtensRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListExtensRequest) SetResourceOwnerAccount(v string) *ClinkListExtensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListExtensRequest) SetResourceOwnerId(v int64) *ClinkListExtensRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListExtensRequest) SetType(v int64) *ClinkListExtensRequest {
	s.Type = &v
	return s
}

func (s *ClinkListExtensRequest) Validate() error {
	return dara.Validate(s)
}
