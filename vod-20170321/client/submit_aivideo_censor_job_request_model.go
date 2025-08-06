// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCensorJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCensorConfig(v string) *SubmitAIVideoCensorJobRequest
	GetAIVideoCensorConfig() *string
	SetMediaId(v string) *SubmitAIVideoCensorJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIVideoCensorJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIVideoCensorJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIVideoCensorJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIVideoCensorJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIVideoCensorJobRequest
	GetUserData() *string
}

type SubmitAIVideoCensorJobRequest struct {
	AIVideoCensorConfig *string `json:"AIVideoCensorConfig,omitempty" xml:"AIVideoCensorConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIVideoCensorJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCensorJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCensorJobRequest) GetAIVideoCensorConfig() *string {
	return s.AIVideoCensorConfig
}

func (s *SubmitAIVideoCensorJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoCensorJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIVideoCensorJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIVideoCensorJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIVideoCensorJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIVideoCensorJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIVideoCensorJobRequest) SetAIVideoCensorConfig(v string) *SubmitAIVideoCensorJobRequest {
	s.AIVideoCensorConfig = &v
	return s
}

func (s *SubmitAIVideoCensorJobRequest) SetMediaId(v string) *SubmitAIVideoCensorJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoCensorJobRequest) SetOwnerAccount(v string) *SubmitAIVideoCensorJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIVideoCensorJobRequest) SetOwnerId(v string) *SubmitAIVideoCensorJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIVideoCensorJobRequest) SetResourceOwnerAccount(v string) *SubmitAIVideoCensorJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIVideoCensorJobRequest) SetResourceOwnerId(v string) *SubmitAIVideoCensorJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIVideoCensorJobRequest) SetUserData(v string) *SubmitAIVideoCensorJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIVideoCensorJobRequest) Validate() error {
	return dara.Validate(s)
}
