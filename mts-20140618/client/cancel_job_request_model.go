// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CancelJobRequest
	GetJobId() *string
	SetOwnerAccount(v string) *CancelJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelJobRequest
	GetResourceOwnerId() *int64
}

type CancelJobRequest struct {
	// The ID of the transcoding job to be canceled. You can log on to the **MPS console*	- and click **Tasks*	- in the left-side navigation pane to obtain job IDs. Alternatively, you can obtain job IDs from the response of the [SubmitJobs](https://help.aliyun.com/document_detail/29226.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// d1ce4d3efcb549419193f50f1fcd****
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelJobRequest) GoString() string {
	return s.String()
}

func (s *CancelJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelJobRequest) SetJobId(v string) *CancelJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelJobRequest) SetOwnerAccount(v string) *CancelJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelJobRequest) SetOwnerId(v int64) *CancelJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelJobRequest) SetResourceOwnerAccount(v string) *CancelJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelJobRequest) SetResourceOwnerId(v int64) *CancelJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelJobRequest) Validate() error {
	return dara.Validate(s)
}
