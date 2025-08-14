// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDNAJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *QueryDNAJobListRequest
	GetJobIds() *string
	SetOwnerAccount(v string) *QueryDNAJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryDNAJobListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryDNAJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryDNAJobListRequest
	GetResourceOwnerId() *int64
}

type QueryDNAJobListRequest struct {
	// The IDs of the media fingerprint analysis jobs that you want to query. We recommend that you query at most 10 jobs at a time. Separate multiple job IDs with commas (,).
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryDNAJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDNAJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryDNAJobListRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *QueryDNAJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryDNAJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryDNAJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryDNAJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryDNAJobListRequest) SetJobIds(v string) *QueryDNAJobListRequest {
	s.JobIds = &v
	return s
}

func (s *QueryDNAJobListRequest) SetOwnerAccount(v string) *QueryDNAJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryDNAJobListRequest) SetOwnerId(v int64) *QueryDNAJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryDNAJobListRequest) SetResourceOwnerAccount(v string) *QueryDNAJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryDNAJobListRequest) SetResourceOwnerId(v int64) *QueryDNAJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryDNAJobListRequest) Validate() error {
	return dara.Validate(s)
}
