// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoPornRecogJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoPornRecogConfig(v string) *SubmitAIVideoPornRecogJobRequest
	GetAIVideoPornRecogConfig() *string
	SetMediaId(v string) *SubmitAIVideoPornRecogJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIVideoPornRecogJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIVideoPornRecogJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIVideoPornRecogJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIVideoPornRecogJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIVideoPornRecogJobRequest
	GetUserData() *string
}

type SubmitAIVideoPornRecogJobRequest struct {
	AIVideoPornRecogConfig *string `json:"AIVideoPornRecogConfig,omitempty" xml:"AIVideoPornRecogConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIVideoPornRecogJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoPornRecogJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoPornRecogJobRequest) GetAIVideoPornRecogConfig() *string {
	return s.AIVideoPornRecogConfig
}

func (s *SubmitAIVideoPornRecogJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoPornRecogJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIVideoPornRecogJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIVideoPornRecogJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIVideoPornRecogJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIVideoPornRecogJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIVideoPornRecogJobRequest) SetAIVideoPornRecogConfig(v string) *SubmitAIVideoPornRecogJobRequest {
	s.AIVideoPornRecogConfig = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobRequest) SetMediaId(v string) *SubmitAIVideoPornRecogJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobRequest) SetOwnerAccount(v string) *SubmitAIVideoPornRecogJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobRequest) SetOwnerId(v string) *SubmitAIVideoPornRecogJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobRequest) SetResourceOwnerAccount(v string) *SubmitAIVideoPornRecogJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobRequest) SetResourceOwnerId(v string) *SubmitAIVideoPornRecogJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobRequest) SetUserData(v string) *SubmitAIVideoPornRecogJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIVideoPornRecogJobRequest) Validate() error {
	return dara.Validate(s)
}
