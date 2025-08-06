// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMediaExportJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *CancelMediaExportJobsRequest
	GetJobIds() *string
	SetOwnerId(v int64) *CancelMediaExportJobsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelMediaExportJobsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelMediaExportJobsRequest
	GetResourceOwnerId() *int64
}

type CancelMediaExportJobsRequest struct {
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelMediaExportJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelMediaExportJobsRequest) GoString() string {
	return s.String()
}

func (s *CancelMediaExportJobsRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *CancelMediaExportJobsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelMediaExportJobsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelMediaExportJobsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelMediaExportJobsRequest) SetJobIds(v string) *CancelMediaExportJobsRequest {
	s.JobIds = &v
	return s
}

func (s *CancelMediaExportJobsRequest) SetOwnerId(v int64) *CancelMediaExportJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelMediaExportJobsRequest) SetResourceOwnerAccount(v string) *CancelMediaExportJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelMediaExportJobsRequest) SetResourceOwnerId(v int64) *CancelMediaExportJobsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelMediaExportJobsRequest) Validate() error {
	return dara.Validate(s)
}
