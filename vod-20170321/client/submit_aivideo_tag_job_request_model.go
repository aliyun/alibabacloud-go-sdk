// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoTagJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoTagConfig(v string) *SubmitAIVideoTagJobRequest
	GetAIVideoTagConfig() *string
	SetMediaId(v string) *SubmitAIVideoTagJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIVideoTagJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIVideoTagJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIVideoTagJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIVideoTagJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIVideoTagJobRequest
	GetUserData() *string
}

type SubmitAIVideoTagJobRequest struct {
	AIVideoTagConfig *string `json:"AIVideoTagConfig,omitempty" xml:"AIVideoTagConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIVideoTagJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoTagJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoTagJobRequest) GetAIVideoTagConfig() *string {
	return s.AIVideoTagConfig
}

func (s *SubmitAIVideoTagJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoTagJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIVideoTagJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIVideoTagJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIVideoTagJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIVideoTagJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIVideoTagJobRequest) SetAIVideoTagConfig(v string) *SubmitAIVideoTagJobRequest {
	s.AIVideoTagConfig = &v
	return s
}

func (s *SubmitAIVideoTagJobRequest) SetMediaId(v string) *SubmitAIVideoTagJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoTagJobRequest) SetOwnerAccount(v string) *SubmitAIVideoTagJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIVideoTagJobRequest) SetOwnerId(v string) *SubmitAIVideoTagJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIVideoTagJobRequest) SetResourceOwnerAccount(v string) *SubmitAIVideoTagJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIVideoTagJobRequest) SetResourceOwnerId(v string) *SubmitAIVideoTagJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIVideoTagJobRequest) SetUserData(v string) *SubmitAIVideoTagJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIVideoTagJobRequest) Validate() error {
	return dara.Validate(s)
}
