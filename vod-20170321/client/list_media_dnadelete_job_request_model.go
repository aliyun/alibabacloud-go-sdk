// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaDNADeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *ListMediaDNADeleteJobRequest
	GetJobIds() *string
	SetOwnerAccount(v string) *ListMediaDNADeleteJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListMediaDNADeleteJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListMediaDNADeleteJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListMediaDNADeleteJobRequest
	GetResourceOwnerId() *string
}

type ListMediaDNADeleteJobRequest struct {
	// This parameter is required.
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListMediaDNADeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNADeleteJobRequest) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *ListMediaDNADeleteJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListMediaDNADeleteJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListMediaDNADeleteJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMediaDNADeleteJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListMediaDNADeleteJobRequest) SetJobIds(v string) *ListMediaDNADeleteJobRequest {
	s.JobIds = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) SetOwnerAccount(v string) *ListMediaDNADeleteJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) SetOwnerId(v string) *ListMediaDNADeleteJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) SetResourceOwnerAccount(v string) *ListMediaDNADeleteJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) SetResourceOwnerId(v string) *ListMediaDNADeleteJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
