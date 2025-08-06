// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoSummaryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoSummaryJobIds(v string) *ListAIVideoSummaryJobRequest
	GetAIVideoSummaryJobIds() *string
	SetOwnerAccount(v string) *ListAIVideoSummaryJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIVideoSummaryJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIVideoSummaryJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIVideoSummaryJobRequest
	GetResourceOwnerId() *string
}

type ListAIVideoSummaryJobRequest struct {
	// This parameter is required.
	AIVideoSummaryJobIds *string `json:"AIVideoSummaryJobIds,omitempty" xml:"AIVideoSummaryJobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIVideoSummaryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoSummaryJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIVideoSummaryJobRequest) GetAIVideoSummaryJobIds() *string {
	return s.AIVideoSummaryJobIds
}

func (s *ListAIVideoSummaryJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIVideoSummaryJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIVideoSummaryJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIVideoSummaryJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIVideoSummaryJobRequest) SetAIVideoSummaryJobIds(v string) *ListAIVideoSummaryJobRequest {
	s.AIVideoSummaryJobIds = &v
	return s
}

func (s *ListAIVideoSummaryJobRequest) SetOwnerAccount(v string) *ListAIVideoSummaryJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIVideoSummaryJobRequest) SetOwnerId(v string) *ListAIVideoSummaryJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIVideoSummaryJobRequest) SetResourceOwnerAccount(v string) *ListAIVideoSummaryJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIVideoSummaryJobRequest) SetResourceOwnerId(v string) *ListAIVideoSummaryJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIVideoSummaryJobRequest) Validate() error {
	return dara.Validate(s)
}
