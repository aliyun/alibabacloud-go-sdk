// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIImageJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *GetAIImageJobsRequest
	GetJobIds() *string
	SetOwnerAccount(v string) *GetAIImageJobsRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetAIImageJobsRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetAIImageJobsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetAIImageJobsRequest
	GetResourceOwnerId() *string
}

type GetAIImageJobsRequest struct {
	// The ID of the image AI processing job. You can obtain the value of JobId from the response to the [SubmitAIImageJob](~~SubmitAIImageJob~~) operation.
	//
	// 	- You can specify a maximum of 10 IDs.
	//
	// 	- Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// cf08a2c6e11e*****de1711b738b9067
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAIImageJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIImageJobsRequest) GoString() string {
	return s.String()
}

func (s *GetAIImageJobsRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *GetAIImageJobsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetAIImageJobsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAIImageJobsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAIImageJobsRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetAIImageJobsRequest) SetJobIds(v string) *GetAIImageJobsRequest {
	s.JobIds = &v
	return s
}

func (s *GetAIImageJobsRequest) SetOwnerAccount(v string) *GetAIImageJobsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAIImageJobsRequest) SetOwnerId(v string) *GetAIImageJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAIImageJobsRequest) SetResourceOwnerAccount(v string) *GetAIImageJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAIImageJobsRequest) SetResourceOwnerId(v string) *GetAIImageJobsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAIImageJobsRequest) Validate() error {
	return dara.Validate(s)
}
