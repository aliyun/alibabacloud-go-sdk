// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpFileDeleteJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *QueryFpFileDeleteJobListRequest
	GetJobIds() *string
	SetOwnerAccount(v string) *QueryFpFileDeleteJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryFpFileDeleteJobListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryFpFileDeleteJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryFpFileDeleteJobListRequest
	GetResourceOwnerId() *int64
}

type QueryFpFileDeleteJobListRequest struct {
	// The IDs of the jobs of deleting media files from a media fingerprint library. You can obtain the job IDs from the response parameters of the [SubmitFpFileDeleteJob](https://help.aliyun.com/document_detail/209274.html) operation. Separate multiple job IDs with commas (,). If you leave this parameter empty, the system returns the latest 20 jobs that are submitted.
	//
	// example:
	//
	// d98459323c024947a104f6a50cbf****,c2dc694696f1441591c5012a73c1****
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryFpFileDeleteJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFpFileDeleteJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryFpFileDeleteJobListRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *QueryFpFileDeleteJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryFpFileDeleteJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryFpFileDeleteJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryFpFileDeleteJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryFpFileDeleteJobListRequest) SetJobIds(v string) *QueryFpFileDeleteJobListRequest {
	s.JobIds = &v
	return s
}

func (s *QueryFpFileDeleteJobListRequest) SetOwnerAccount(v string) *QueryFpFileDeleteJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryFpFileDeleteJobListRequest) SetOwnerId(v int64) *QueryFpFileDeleteJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryFpFileDeleteJobListRequest) SetResourceOwnerAccount(v string) *QueryFpFileDeleteJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryFpFileDeleteJobListRequest) SetResourceOwnerId(v int64) *QueryFpFileDeleteJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryFpFileDeleteJobListRequest) Validate() error {
	return dara.Validate(s)
}
