// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIASRJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIASRConfig(v string) *SubmitAIASRJobRequest
	GetAIASRConfig() *string
	SetMediaId(v string) *SubmitAIASRJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIASRJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIASRJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIASRJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIASRJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIASRJobRequest
	GetUserData() *string
}

type SubmitAIASRJobRequest struct {
	AIASRConfig *string `json:"AIASRConfig,omitempty" xml:"AIASRConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIASRJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIASRJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIASRJobRequest) GetAIASRConfig() *string {
	return s.AIASRConfig
}

func (s *SubmitAIASRJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIASRJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIASRJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIASRJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIASRJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIASRJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIASRJobRequest) SetAIASRConfig(v string) *SubmitAIASRJobRequest {
	s.AIASRConfig = &v
	return s
}

func (s *SubmitAIASRJobRequest) SetMediaId(v string) *SubmitAIASRJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIASRJobRequest) SetOwnerAccount(v string) *SubmitAIASRJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIASRJobRequest) SetOwnerId(v string) *SubmitAIASRJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIASRJobRequest) SetResourceOwnerAccount(v string) *SubmitAIASRJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIASRJobRequest) SetResourceOwnerId(v string) *SubmitAIASRJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIASRJobRequest) SetUserData(v string) *SubmitAIASRJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIASRJobRequest) Validate() error {
	return dara.Validate(s)
}
