// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpDBDeleteJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *QueryFpDBDeleteJobListRequest
	GetJobIds() *string
	SetOwnerAccount(v string) *QueryFpDBDeleteJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryFpDBDeleteJobListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryFpDBDeleteJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryFpDBDeleteJobListRequest
	GetResourceOwnerId() *int64
}

type QueryFpDBDeleteJobListRequest struct {
	// The IDs of the jobs of clearing or deleting a media fingerprint library. You can obtain the job IDs from the response parameters of the [SubmitFpDBDeleteJob](https://help.aliyun.com/document_detail/209341.html) operation. Separate multiple job IDs with commas (,). If you leave this parameter empty, the system returns the latest 20 jobs that are submitted.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****,78dc866518b843259669df58ed30****
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryFpDBDeleteJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFpDBDeleteJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryFpDBDeleteJobListRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *QueryFpDBDeleteJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryFpDBDeleteJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryFpDBDeleteJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryFpDBDeleteJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryFpDBDeleteJobListRequest) SetJobIds(v string) *QueryFpDBDeleteJobListRequest {
	s.JobIds = &v
	return s
}

func (s *QueryFpDBDeleteJobListRequest) SetOwnerAccount(v string) *QueryFpDBDeleteJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryFpDBDeleteJobListRequest) SetOwnerId(v int64) *QueryFpDBDeleteJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryFpDBDeleteJobListRequest) SetResourceOwnerAccount(v string) *QueryFpDBDeleteJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryFpDBDeleteJobListRequest) SetResourceOwnerId(v int64) *QueryFpDBDeleteJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryFpDBDeleteJobListRequest) Validate() error {
	return dara.Validate(s)
}
