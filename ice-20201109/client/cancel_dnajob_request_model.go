// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDNAJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CancelDNAJobRequest
	GetJobId() *string
	SetOwnerAccount(v string) *CancelDNAJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelDNAJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelDNAJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelDNAJobRequest
	GetResourceOwnerId() *int64
}

type CancelDNAJobRequest struct {
	// The ID of the media fingerprint analysis job that you want to cancel.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelDNAJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDNAJobRequest) GoString() string {
	return s.String()
}

func (s *CancelDNAJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelDNAJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelDNAJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelDNAJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelDNAJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelDNAJobRequest) SetJobId(v string) *CancelDNAJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelDNAJobRequest) SetOwnerAccount(v string) *CancelDNAJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelDNAJobRequest) SetOwnerId(v int64) *CancelDNAJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelDNAJobRequest) SetResourceOwnerAccount(v string) *CancelDNAJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelDNAJobRequest) SetResourceOwnerId(v int64) *CancelDNAJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelDNAJobRequest) Validate() error {
	return dara.Validate(s)
}
