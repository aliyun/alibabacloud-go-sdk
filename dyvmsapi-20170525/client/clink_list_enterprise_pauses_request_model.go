// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListEnterprisePausesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkListEnterprisePausesRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *ClinkListEnterprisePausesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkListEnterprisePausesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListEnterprisePausesRequest
	GetResourceOwnerId() *int64
}

type ClinkListEnterprisePausesRequest struct {
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

func (s ClinkListEnterprisePausesRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListEnterprisePausesRequest) GoString() string {
	return s.String()
}

func (s *ClinkListEnterprisePausesRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListEnterprisePausesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListEnterprisePausesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListEnterprisePausesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListEnterprisePausesRequest) SetEnterpriseId(v int64) *ClinkListEnterprisePausesRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListEnterprisePausesRequest) SetOwnerId(v int64) *ClinkListEnterprisePausesRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListEnterprisePausesRequest) SetResourceOwnerAccount(v string) *ClinkListEnterprisePausesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListEnterprisePausesRequest) SetResourceOwnerId(v int64) *ClinkListEnterprisePausesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListEnterprisePausesRequest) Validate() error {
	return dara.Validate(s)
}
