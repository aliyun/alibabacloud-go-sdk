// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *ListAIJobRequest
	GetJobIds() *string
	SetOwnerAccount(v string) *ListAIJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIJobRequest
	GetResourceOwnerId() *string
}

type ListAIJobRequest struct {
	// The list of job IDs. You can obtain the job ID from the PlayInfo parameter in the response to the [GetPlayInfo](https://help.aliyun.com/document_detail/56124.html) operation.
	//
	// >  You can specify a maximum of 10 job IDs in a request. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// a718a3a1e8bb42ee3bc88921e94****,aasdcsfg782740asd3****,k2l3ibaskod98wrns9d****
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIJobRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *ListAIJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIJobRequest) SetJobIds(v string) *ListAIJobRequest {
	s.JobIds = &v
	return s
}

func (s *ListAIJobRequest) SetOwnerAccount(v string) *ListAIJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIJobRequest) SetOwnerId(v string) *ListAIJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIJobRequest) SetResourceOwnerAccount(v string) *ListAIJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIJobRequest) SetResourceOwnerId(v string) *ListAIJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIJobRequest) Validate() error {
	return dara.Validate(s)
}
