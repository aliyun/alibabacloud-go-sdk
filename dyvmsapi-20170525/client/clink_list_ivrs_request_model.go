// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListIvrsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkListIvrsRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *ClinkListIvrsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkListIvrsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListIvrsRequest
	GetResourceOwnerId() *int64
}

type ClinkListIvrsRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkListIvrsRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrsRequest) GoString() string {
	return s.String()
}

func (s *ClinkListIvrsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListIvrsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListIvrsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListIvrsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListIvrsRequest) SetEnterpriseId(v int64) *ClinkListIvrsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListIvrsRequest) SetOwnerId(v int64) *ClinkListIvrsRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListIvrsRequest) SetResourceOwnerAccount(v string) *ClinkListIvrsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListIvrsRequest) SetResourceOwnerId(v int64) *ClinkListIvrsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListIvrsRequest) Validate() error {
	return dara.Validate(s)
}
