// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIProductionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *QueryIProductionJobRequest
	GetJobId() *string
	SetOwnerAccount(v string) *QueryIProductionJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryIProductionJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryIProductionJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryIProductionJobRequest
	GetResourceOwnerId() *int64
}

type QueryIProductionJobRequest struct {
	// example:
	//
	// 88c6ca184c0e432bbf5b665e2a15****
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryIProductionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryIProductionJobRequest) GoString() string {
	return s.String()
}

func (s *QueryIProductionJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryIProductionJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryIProductionJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryIProductionJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryIProductionJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryIProductionJobRequest) SetJobId(v string) *QueryIProductionJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryIProductionJobRequest) SetOwnerAccount(v string) *QueryIProductionJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryIProductionJobRequest) SetOwnerId(v int64) *QueryIProductionJobRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryIProductionJobRequest) SetResourceOwnerAccount(v string) *QueryIProductionJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryIProductionJobRequest) SetResourceOwnerId(v int64) *QueryIProductionJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryIProductionJobRequest) Validate() error {
	return dara.Validate(s)
}
