// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsCodeLimitConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimitDay(v int32) *UpdateSmsCodeLimitConfigRequest
	GetLimitDay() *int32
	SetLimitHour(v int32) *UpdateSmsCodeLimitConfigRequest
	GetLimitHour() *int32
	SetLimitId(v int64) *UpdateSmsCodeLimitConfigRequest
	GetLimitId() *int64
	SetLimitMinute(v int32) *UpdateSmsCodeLimitConfigRequest
	GetLimitMinute() *int32
	SetLimitOther(v string) *UpdateSmsCodeLimitConfigRequest
	GetLimitOther() *string
	SetOwnerId(v int64) *UpdateSmsCodeLimitConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateSmsCodeLimitConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmsCodeLimitConfigRequest
	GetResourceOwnerId() *int64
}

type UpdateSmsCodeLimitConfigRequest struct {
	LimitDay  *int32 `json:"LimitDay,omitempty" xml:"LimitDay,omitempty"`
	LimitHour *int32 `json:"LimitHour,omitempty" xml:"LimitHour,omitempty"`
	// This parameter is required.
	LimitId              *int64  `json:"LimitId,omitempty" xml:"LimitId,omitempty"`
	LimitMinute          *int32  `json:"LimitMinute,omitempty" xml:"LimitMinute,omitempty"`
	LimitOther           *string `json:"LimitOther,omitempty" xml:"LimitOther,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateSmsCodeLimitConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsCodeLimitConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsCodeLimitConfigRequest) GetLimitDay() *int32 {
	return s.LimitDay
}

func (s *UpdateSmsCodeLimitConfigRequest) GetLimitHour() *int32 {
	return s.LimitHour
}

func (s *UpdateSmsCodeLimitConfigRequest) GetLimitId() *int64 {
	return s.LimitId
}

func (s *UpdateSmsCodeLimitConfigRequest) GetLimitMinute() *int32 {
	return s.LimitMinute
}

func (s *UpdateSmsCodeLimitConfigRequest) GetLimitOther() *string {
	return s.LimitOther
}

func (s *UpdateSmsCodeLimitConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmsCodeLimitConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmsCodeLimitConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmsCodeLimitConfigRequest) SetLimitDay(v int32) *UpdateSmsCodeLimitConfigRequest {
	s.LimitDay = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigRequest) SetLimitHour(v int32) *UpdateSmsCodeLimitConfigRequest {
	s.LimitHour = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigRequest) SetLimitId(v int64) *UpdateSmsCodeLimitConfigRequest {
	s.LimitId = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigRequest) SetLimitMinute(v int32) *UpdateSmsCodeLimitConfigRequest {
	s.LimitMinute = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigRequest) SetLimitOther(v string) *UpdateSmsCodeLimitConfigRequest {
	s.LimitOther = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigRequest) SetOwnerId(v int64) *UpdateSmsCodeLimitConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigRequest) SetResourceOwnerAccount(v string) *UpdateSmsCodeLimitConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigRequest) SetResourceOwnerId(v int64) *UpdateSmsCodeLimitConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigRequest) Validate() error {
	return dara.Validate(s)
}
