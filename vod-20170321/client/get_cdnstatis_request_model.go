// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCDNStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetCDNStatisRequest
	GetEndTime() *int64
	SetLevel(v string) *GetCDNStatisRequest
	GetLevel() *string
	SetOwnerAccount(v string) *GetCDNStatisRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetCDNStatisRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetCDNStatisRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetCDNStatisRequest
	GetResourceOwnerId() *string
	SetStartTime(v int64) *GetCDNStatisRequest
	GetStartTime() *int64
}

type GetCDNStatisRequest struct {
	// This parameter is required.
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Level                *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetCDNStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisRequest) GoString() string {
	return s.String()
}

func (s *GetCDNStatisRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetCDNStatisRequest) GetLevel() *string {
	return s.Level
}

func (s *GetCDNStatisRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetCDNStatisRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetCDNStatisRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCDNStatisRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetCDNStatisRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetCDNStatisRequest) SetEndTime(v int64) *GetCDNStatisRequest {
	s.EndTime = &v
	return s
}

func (s *GetCDNStatisRequest) SetLevel(v string) *GetCDNStatisRequest {
	s.Level = &v
	return s
}

func (s *GetCDNStatisRequest) SetOwnerAccount(v string) *GetCDNStatisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetCDNStatisRequest) SetOwnerId(v string) *GetCDNStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCDNStatisRequest) SetResourceOwnerAccount(v string) *GetCDNStatisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCDNStatisRequest) SetResourceOwnerId(v string) *GetCDNStatisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCDNStatisRequest) SetStartTime(v int64) *GetCDNStatisRequest {
	s.StartTime = &v
	return s
}

func (s *GetCDNStatisRequest) Validate() error {
	return dara.Validate(s)
}
