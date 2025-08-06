// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDNAInitializationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *SubmitDNAInitializationJobRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *SubmitDNAInitializationJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitDNAInitializationJobRequest
	GetOwnerId() *string
	SetRecentNumber(v int32) *SubmitDNAInitializationJobRequest
	GetRecentNumber() *int32
	SetResourceOwnerAccount(v string) *SubmitDNAInitializationJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitDNAInitializationJobRequest
	GetResourceOwnerId() *string
	SetStartTime(v string) *SubmitDNAInitializationJobRequest
	GetStartTime() *string
	SetType(v string) *SubmitDNAInitializationJobRequest
	GetType() *string
}

type SubmitDNAInitializationJobRequest struct {
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RecentNumber         *int32  `json:"RecentNumber,omitempty" xml:"RecentNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitDNAInitializationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDNAInitializationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDNAInitializationJobRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SubmitDNAInitializationJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitDNAInitializationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitDNAInitializationJobRequest) GetRecentNumber() *int32 {
	return s.RecentNumber
}

func (s *SubmitDNAInitializationJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitDNAInitializationJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitDNAInitializationJobRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitDNAInitializationJobRequest) GetType() *string {
	return s.Type
}

func (s *SubmitDNAInitializationJobRequest) SetEndTime(v string) *SubmitDNAInitializationJobRequest {
	s.EndTime = &v
	return s
}

func (s *SubmitDNAInitializationJobRequest) SetOwnerAccount(v string) *SubmitDNAInitializationJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitDNAInitializationJobRequest) SetOwnerId(v string) *SubmitDNAInitializationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitDNAInitializationJobRequest) SetRecentNumber(v int32) *SubmitDNAInitializationJobRequest {
	s.RecentNumber = &v
	return s
}

func (s *SubmitDNAInitializationJobRequest) SetResourceOwnerAccount(v string) *SubmitDNAInitializationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitDNAInitializationJobRequest) SetResourceOwnerId(v string) *SubmitDNAInitializationJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitDNAInitializationJobRequest) SetStartTime(v string) *SubmitDNAInitializationJobRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitDNAInitializationJobRequest) SetType(v string) *SubmitDNAInitializationJobRequest {
	s.Type = &v
	return s
}

func (s *SubmitDNAInitializationJobRequest) Validate() error {
	return dara.Validate(s)
}
