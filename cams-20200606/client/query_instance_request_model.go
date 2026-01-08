// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryInstanceRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *QueryInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryInstanceRequest
	GetResourceOwnerId() *int64
}

type QueryInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9293938****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryInstanceRequest) SetInstanceId(v string) *QueryInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstanceRequest) SetOwnerId(v int64) *QueryInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryInstanceRequest) SetResourceOwnerAccount(v string) *QueryInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryInstanceRequest) SetResourceOwnerId(v int64) *QueryInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryInstanceRequest) Validate() error {
	return dara.Validate(s)
}
