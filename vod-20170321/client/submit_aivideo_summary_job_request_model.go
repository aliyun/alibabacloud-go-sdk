// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoSummaryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoSummaryConfig(v string) *SubmitAIVideoSummaryJobRequest
	GetAIVideoSummaryConfig() *string
	SetMediaId(v string) *SubmitAIVideoSummaryJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIVideoSummaryJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIVideoSummaryJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIVideoSummaryJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIVideoSummaryJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIVideoSummaryJobRequest
	GetUserData() *string
}

type SubmitAIVideoSummaryJobRequest struct {
	AIVideoSummaryConfig *string `json:"AIVideoSummaryConfig,omitempty" xml:"AIVideoSummaryConfig,omitempty"`
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIVideoSummaryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoSummaryJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoSummaryJobRequest) GetAIVideoSummaryConfig() *string {
	return s.AIVideoSummaryConfig
}

func (s *SubmitAIVideoSummaryJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIVideoSummaryJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIVideoSummaryJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIVideoSummaryJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIVideoSummaryJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIVideoSummaryJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIVideoSummaryJobRequest) SetAIVideoSummaryConfig(v string) *SubmitAIVideoSummaryJobRequest {
	s.AIVideoSummaryConfig = &v
	return s
}

func (s *SubmitAIVideoSummaryJobRequest) SetMediaId(v string) *SubmitAIVideoSummaryJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIVideoSummaryJobRequest) SetOwnerAccount(v string) *SubmitAIVideoSummaryJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIVideoSummaryJobRequest) SetOwnerId(v string) *SubmitAIVideoSummaryJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIVideoSummaryJobRequest) SetResourceOwnerAccount(v string) *SubmitAIVideoSummaryJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIVideoSummaryJobRequest) SetResourceOwnerId(v string) *SubmitAIVideoSummaryJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIVideoSummaryJobRequest) SetUserData(v string) *SubmitAIVideoSummaryJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIVideoSummaryJobRequest) Validate() error {
	return dara.Validate(s)
}
