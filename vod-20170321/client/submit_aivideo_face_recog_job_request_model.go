// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoFaceRecogJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoFaceRecogConfig(v string) *SubmitAIVideoFaceRecogJobRequest
	GetAIVideoFaceRecogConfig() *string
	SetMediaId(v string) *SubmitAIVideoFaceRecogJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIVideoFaceRecogJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIVideoFaceRecogJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIVideoFaceRecogJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIVideoFaceRecogJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIVideoFaceRecogJobRequest
	GetUserData() *string
}

type SubmitAIVideoFaceRecogJobRequest struct {
	AIVideoFaceRecogConfig *string `json:"AIVideoFaceRecogConfig,omitempty" xml:"AIVideoFaceRecogConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIVideoFaceRecogJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoFaceRecogJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoFaceRecogJobRequest) GetAIVideoFaceRecogConfig() *string {
	return s.AIVideoFaceRecogConfig
}

func (s *SubmitAIVideoFaceRecogJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoFaceRecogJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIVideoFaceRecogJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIVideoFaceRecogJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIVideoFaceRecogJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIVideoFaceRecogJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIVideoFaceRecogJobRequest) SetAIVideoFaceRecogConfig(v string) *SubmitAIVideoFaceRecogJobRequest {
	s.AIVideoFaceRecogConfig = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobRequest) SetMediaId(v string) *SubmitAIVideoFaceRecogJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobRequest) SetOwnerAccount(v string) *SubmitAIVideoFaceRecogJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobRequest) SetOwnerId(v string) *SubmitAIVideoFaceRecogJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobRequest) SetResourceOwnerAccount(v string) *SubmitAIVideoFaceRecogJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobRequest) SetResourceOwnerId(v string) *SubmitAIVideoFaceRecogJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobRequest) SetUserData(v string) *SubmitAIVideoFaceRecogJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIVideoFaceRecogJobRequest) Validate() error {
	return dara.Validate(s)
}
