// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteInstanceRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DeleteInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteInstanceRequest
	GetResourceOwnerId() *int64
}

type DeleteInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2399393****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerId(v int64) *DeleteInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerAccount(v string) *DeleteInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerId(v int64) *DeleteInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
