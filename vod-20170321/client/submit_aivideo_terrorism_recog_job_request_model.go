// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoTerrorismRecogJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoTerrorismRecogConfig(v string) *SubmitAIVideoTerrorismRecogJobRequest
	GetAIVideoTerrorismRecogConfig() *string
	SetMediaId(v string) *SubmitAIVideoTerrorismRecogJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIVideoTerrorismRecogJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIVideoTerrorismRecogJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIVideoTerrorismRecogJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIVideoTerrorismRecogJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIVideoTerrorismRecogJobRequest
	GetUserData() *string
}

type SubmitAIVideoTerrorismRecogJobRequest struct {
	AIVideoTerrorismRecogConfig *string `json:"AIVideoTerrorismRecogConfig,omitempty" xml:"AIVideoTerrorismRecogConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIVideoTerrorismRecogJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoTerrorismRecogJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) GetAIVideoTerrorismRecogConfig() *string {
	return s.AIVideoTerrorismRecogConfig
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) SetAIVideoTerrorismRecogConfig(v string) *SubmitAIVideoTerrorismRecogJobRequest {
	s.AIVideoTerrorismRecogConfig = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) SetMediaId(v string) *SubmitAIVideoTerrorismRecogJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) SetOwnerAccount(v string) *SubmitAIVideoTerrorismRecogJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) SetOwnerId(v string) *SubmitAIVideoTerrorismRecogJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) SetResourceOwnerAccount(v string) *SubmitAIVideoTerrorismRecogJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) SetResourceOwnerId(v string) *SubmitAIVideoTerrorismRecogJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) SetUserData(v string) *SubmitAIVideoTerrorismRecogJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIVideoTerrorismRecogJobRequest) Validate() error {
	return dara.Validate(s)
}
