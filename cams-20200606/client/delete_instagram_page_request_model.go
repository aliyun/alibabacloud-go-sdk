// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstagramPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteInstagramPageRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DeleteInstagramPageRequest
	GetOwnerId() *int64
	SetPageId(v string) *DeleteInstagramPageRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *DeleteInstagramPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteInstagramPageRequest
	GetResourceOwnerId() *int64
}

type DeleteInstagramPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 154654654654654
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteInstagramPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstagramPageRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstagramPageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstagramPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteInstagramPageRequest) GetPageId() *string {
	return s.PageId
}

func (s *DeleteInstagramPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteInstagramPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteInstagramPageRequest) SetInstanceId(v string) *DeleteInstagramPageRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstagramPageRequest) SetOwnerId(v int64) *DeleteInstagramPageRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstagramPageRequest) SetPageId(v string) *DeleteInstagramPageRequest {
	s.PageId = &v
	return s
}

func (s *DeleteInstagramPageRequest) SetResourceOwnerAccount(v string) *DeleteInstagramPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstagramPageRequest) SetResourceOwnerId(v int64) *DeleteInstagramPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInstagramPageRequest) Validate() error {
	return dara.Validate(s)
}
