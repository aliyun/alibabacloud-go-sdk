// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteEnterpriseTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudDeleteEnterpriseTimeRequest
	GetEnterpriseId() *int64
	SetId(v int64) *CloudDeleteEnterpriseTimeRequest
	GetId() *int64
	SetName(v string) *CloudDeleteEnterpriseTimeRequest
	GetName() *string
	SetOwnerId(v int64) *CloudDeleteEnterpriseTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudDeleteEnterpriseTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteEnterpriseTimeRequest
	GetResourceOwnerId() *int64
}

type CloudDeleteEnterpriseTimeRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 时间条件id；id和name二选一，不可同时为空。
	//
	// example:
	//
	// 92
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 时间条件名称；id和name二选一，不可同时为空。
	//
	// example:
	//
	// name3
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudDeleteEnterpriseTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteEnterpriseTimeRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteEnterpriseTimeRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteEnterpriseTimeRequest) GetId() *int64 {
	return s.Id
}

func (s *CloudDeleteEnterpriseTimeRequest) GetName() *string {
	return s.Name
}

func (s *CloudDeleteEnterpriseTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteEnterpriseTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteEnterpriseTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteEnterpriseTimeRequest) SetEnterpriseId(v int64) *CloudDeleteEnterpriseTimeRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeRequest) SetId(v int64) *CloudDeleteEnterpriseTimeRequest {
	s.Id = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeRequest) SetName(v string) *CloudDeleteEnterpriseTimeRequest {
	s.Name = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeRequest) SetOwnerId(v int64) *CloudDeleteEnterpriseTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeRequest) SetResourceOwnerAccount(v string) *CloudDeleteEnterpriseTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeRequest) SetResourceOwnerId(v int64) *CloudDeleteEnterpriseTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeRequest) Validate() error {
	return dara.Validate(s)
}
