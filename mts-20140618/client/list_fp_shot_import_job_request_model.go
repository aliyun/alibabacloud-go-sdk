// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotImportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *ListFpShotImportJobRequest
	GetJobIds() *string
	SetOwnerAccount(v string) *ListFpShotImportJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListFpShotImportJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListFpShotImportJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFpShotImportJobRequest
	GetResourceOwnerId() *int64
}

type ListFpShotImportJobRequest struct {
	// The job IDs. You can obtain the job IDs from the response to the [ImportFpShotJob](https://help.aliyun.com/document_detail/312262.html) operation. You can specify a maximum of 10 job IDs in a request. Separate multiple job IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****,c074b118ace44395a02063a5ab94****
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListFpShotImportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotImportJobRequest) GoString() string {
	return s.String()
}

func (s *ListFpShotImportJobRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *ListFpShotImportJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListFpShotImportJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFpShotImportJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFpShotImportJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFpShotImportJobRequest) SetJobIds(v string) *ListFpShotImportJobRequest {
	s.JobIds = &v
	return s
}

func (s *ListFpShotImportJobRequest) SetOwnerAccount(v string) *ListFpShotImportJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListFpShotImportJobRequest) SetOwnerId(v int64) *ListFpShotImportJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFpShotImportJobRequest) SetResourceOwnerAccount(v string) *ListFpShotImportJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFpShotImportJobRequest) SetResourceOwnerId(v int64) *ListFpShotImportJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFpShotImportJobRequest) Validate() error {
	return dara.Validate(s)
}
