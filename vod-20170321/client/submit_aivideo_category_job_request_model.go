// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCategoryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCategoryConfig(v string) *SubmitAIVideoCategoryJobRequest
	GetAIVideoCategoryConfig() *string
	SetMediaId(v string) *SubmitAIVideoCategoryJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIVideoCategoryJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIVideoCategoryJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIVideoCategoryJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIVideoCategoryJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIVideoCategoryJobRequest
	GetUserData() *string
}

type SubmitAIVideoCategoryJobRequest struct {
	AIVideoCategoryConfig *string `json:"AIVideoCategoryConfig,omitempty" xml:"AIVideoCategoryConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIVideoCategoryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCategoryJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCategoryJobRequest) GetAIVideoCategoryConfig() *string {
	return s.AIVideoCategoryConfig
}

func (s *SubmitAIVideoCategoryJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoCategoryJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIVideoCategoryJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIVideoCategoryJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIVideoCategoryJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIVideoCategoryJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIVideoCategoryJobRequest) SetAIVideoCategoryConfig(v string) *SubmitAIVideoCategoryJobRequest {
	s.AIVideoCategoryConfig = &v
	return s
}

func (s *SubmitAIVideoCategoryJobRequest) SetMediaId(v string) *SubmitAIVideoCategoryJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoCategoryJobRequest) SetOwnerAccount(v string) *SubmitAIVideoCategoryJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIVideoCategoryJobRequest) SetOwnerId(v string) *SubmitAIVideoCategoryJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIVideoCategoryJobRequest) SetResourceOwnerAccount(v string) *SubmitAIVideoCategoryJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIVideoCategoryJobRequest) SetResourceOwnerId(v string) *SubmitAIVideoCategoryJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIVideoCategoryJobRequest) SetUserData(v string) *SubmitAIVideoCategoryJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIVideoCategoryJobRequest) Validate() error {
	return dara.Validate(s)
}
