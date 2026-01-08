// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstagramPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstagramPageRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *ListInstagramPageRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListInstagramPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListInstagramPageRequest
	GetResourceOwnerId() *int64
}

type ListInstagramPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// inst-8t47hh***5pyj0ps
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListInstagramPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPageRequest) GoString() string {
	return s.String()
}

func (s *ListInstagramPageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstagramPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListInstagramPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListInstagramPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListInstagramPageRequest) SetInstanceId(v string) *ListInstagramPageRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstagramPageRequest) SetOwnerId(v int64) *ListInstagramPageRequest {
	s.OwnerId = &v
	return s
}

func (s *ListInstagramPageRequest) SetResourceOwnerAccount(v string) *ListInstagramPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListInstagramPageRequest) SetResourceOwnerId(v int64) *ListInstagramPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListInstagramPageRequest) Validate() error {
	return dara.Validate(s)
}
