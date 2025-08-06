// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaExportJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *DeleteMediaExportJobsRequest
	GetJobIds() *string
	SetOwnerId(v int64) *DeleteMediaExportJobsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteMediaExportJobsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMediaExportJobsRequest
	GetResourceOwnerId() *int64
}

type DeleteMediaExportJobsRequest struct {
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMediaExportJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaExportJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaExportJobsRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *DeleteMediaExportJobsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMediaExportJobsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMediaExportJobsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMediaExportJobsRequest) SetJobIds(v string) *DeleteMediaExportJobsRequest {
	s.JobIds = &v
	return s
}

func (s *DeleteMediaExportJobsRequest) SetOwnerId(v int64) *DeleteMediaExportJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMediaExportJobsRequest) SetResourceOwnerAccount(v string) *DeleteMediaExportJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMediaExportJobsRequest) SetResourceOwnerId(v int64) *DeleteMediaExportJobsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMediaExportJobsRequest) Validate() error {
	return dara.Validate(s)
}
