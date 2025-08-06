// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCoverJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCoverConfig(v string) *SubmitAIVideoCoverJobRequest
	GetAIVideoCoverConfig() *string
	SetMediaId(v string) *SubmitAIVideoCoverJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIVideoCoverJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIVideoCoverJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIVideoCoverJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIVideoCoverJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIVideoCoverJobRequest
	GetUserData() *string
}

type SubmitAIVideoCoverJobRequest struct {
	AIVideoCoverConfig *string `json:"AIVideoCoverConfig,omitempty" xml:"AIVideoCoverConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIVideoCoverJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCoverJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCoverJobRequest) GetAIVideoCoverConfig() *string {
	return s.AIVideoCoverConfig
}

func (s *SubmitAIVideoCoverJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoCoverJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIVideoCoverJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIVideoCoverJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIVideoCoverJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIVideoCoverJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIVideoCoverJobRequest) SetAIVideoCoverConfig(v string) *SubmitAIVideoCoverJobRequest {
	s.AIVideoCoverConfig = &v
	return s
}

func (s *SubmitAIVideoCoverJobRequest) SetMediaId(v string) *SubmitAIVideoCoverJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoCoverJobRequest) SetOwnerAccount(v string) *SubmitAIVideoCoverJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIVideoCoverJobRequest) SetOwnerId(v string) *SubmitAIVideoCoverJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIVideoCoverJobRequest) SetResourceOwnerAccount(v string) *SubmitAIVideoCoverJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIVideoCoverJobRequest) SetResourceOwnerId(v string) *SubmitAIVideoCoverJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIVideoCoverJobRequest) SetUserData(v string) *SubmitAIVideoCoverJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIVideoCoverJobRequest) Validate() error {
	return dara.Validate(s)
}
