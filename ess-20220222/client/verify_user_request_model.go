// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *VerifyUserRequest
	GetOwnerId() *int64
	SetRegionId(v string) *VerifyUserRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *VerifyUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *VerifyUserRequest
	GetResourceOwnerId() *int64
}

type VerifyUserRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where Auto Scaling is required.
	//
	// Examples: `cn-hangzhou` and `cn-shanghai`. For more information, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s VerifyUserRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyUserRequest) GoString() string {
	return s.String()
}

func (s *VerifyUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *VerifyUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *VerifyUserRequest) SetOwnerId(v int64) *VerifyUserRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyUserRequest) SetRegionId(v string) *VerifyUserRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyUserRequest) SetResourceOwnerAccount(v string) *VerifyUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyUserRequest) SetResourceOwnerId(v int64) *VerifyUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyUserRequest) Validate() error {
	return dara.Validate(s)
}
