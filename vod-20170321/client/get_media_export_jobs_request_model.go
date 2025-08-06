// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaExportJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *GetMediaExportJobsRequest
	GetJobIds() *string
	SetOwnerId(v int64) *GetMediaExportJobsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetMediaExportJobsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetMediaExportJobsRequest
	GetResourceOwnerId() *int64
}

type GetMediaExportJobsRequest struct {
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMediaExportJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaExportJobsRequest) GoString() string {
	return s.String()
}

func (s *GetMediaExportJobsRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *GetMediaExportJobsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetMediaExportJobsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMediaExportJobsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetMediaExportJobsRequest) SetJobIds(v string) *GetMediaExportJobsRequest {
	s.JobIds = &v
	return s
}

func (s *GetMediaExportJobsRequest) SetOwnerId(v int64) *GetMediaExportJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMediaExportJobsRequest) SetResourceOwnerAccount(v string) *GetMediaExportJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMediaExportJobsRequest) SetResourceOwnerId(v int64) *GetMediaExportJobsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMediaExportJobsRequest) Validate() error {
	return dara.Validate(s)
}
