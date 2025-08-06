// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCDNStatisSumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndStatisTime(v string) *GetCDNStatisSumRequest
	GetEndStatisTime() *string
	SetLevel(v string) *GetCDNStatisSumRequest
	GetLevel() *string
	SetOwnerAccount(v string) *GetCDNStatisSumRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetCDNStatisSumRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetCDNStatisSumRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetCDNStatisSumRequest
	GetResourceOwnerId() *string
	SetStartStatisTime(v string) *GetCDNStatisSumRequest
	GetStartStatisTime() *string
}

type GetCDNStatisSumRequest struct {
	// This parameter is required.
	EndStatisTime        *string `json:"EndStatisTime,omitempty" xml:"EndStatisTime,omitempty"`
	Level                *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StartStatisTime *string `json:"StartStatisTime,omitempty" xml:"StartStatisTime,omitempty"`
}

func (s GetCDNStatisSumRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisSumRequest) GoString() string {
	return s.String()
}

func (s *GetCDNStatisSumRequest) GetEndStatisTime() *string {
	return s.EndStatisTime
}

func (s *GetCDNStatisSumRequest) GetLevel() *string {
	return s.Level
}

func (s *GetCDNStatisSumRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetCDNStatisSumRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetCDNStatisSumRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCDNStatisSumRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetCDNStatisSumRequest) GetStartStatisTime() *string {
	return s.StartStatisTime
}

func (s *GetCDNStatisSumRequest) SetEndStatisTime(v string) *GetCDNStatisSumRequest {
	s.EndStatisTime = &v
	return s
}

func (s *GetCDNStatisSumRequest) SetLevel(v string) *GetCDNStatisSumRequest {
	s.Level = &v
	return s
}

func (s *GetCDNStatisSumRequest) SetOwnerAccount(v string) *GetCDNStatisSumRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetCDNStatisSumRequest) SetOwnerId(v string) *GetCDNStatisSumRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCDNStatisSumRequest) SetResourceOwnerAccount(v string) *GetCDNStatisSumRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCDNStatisSumRequest) SetResourceOwnerId(v string) *GetCDNStatisSumRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCDNStatisSumRequest) SetStartStatisTime(v string) *GetCDNStatisSumRequest {
	s.StartStatisTime = &v
	return s
}

func (s *GetCDNStatisSumRequest) Validate() error {
	return dara.Validate(s)
}
