// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDailyTokenQuota(v int64) *CreateApiKeyRequest
	GetDailyTokenQuota() *int64
	SetInstanceId(v string) *CreateApiKeyRequest
	GetInstanceId() *string
	SetKeyName(v string) *CreateApiKeyRequest
	GetKeyName() *string
	SetLimitRate(v float64) *CreateApiKeyRequest
	GetLimitRate() *float64
	SetLimitType(v string) *CreateApiKeyRequest
	GetLimitType() *string
	SetQuantity(v int32) *CreateApiKeyRequest
	GetQuantity() *int32
	SetTokenQuota(v int64) *CreateApiKeyRequest
	GetTokenQuota() *int64
}

type CreateApiKeyRequest struct {
	// example:
	//
	// 1000000000
	DailyTokenQuota *int64 `json:"DailyTokenQuota,omitempty" xml:"DailyTokenQuota,omitempty"`
	// The instance name.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the API key.
	//
	// example:
	//
	// api-*****
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The quota percentage.
	//
	// example:
	//
	// 0.2
	LimitRate *float64 `json:"LimitRate,omitempty" xml:"LimitRate,omitempty"`
	// The quota type. Valid values:
	//
	// - ratio: by percentage.
	//
	// - fixed: by fixed value.
	//
	// - auto: automatic allocation.
	//
	// example:
	//
	// fixed
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// The number of API keys to create. Default value: **1**.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The quota for the current key.
	//
	// example:
	//
	// 100000
	TokenQuota *int64 `json:"TokenQuota,omitempty" xml:"TokenQuota,omitempty"`
}

func (s CreateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateApiKeyRequest) GetDailyTokenQuota() *int64 {
	return s.DailyTokenQuota
}

func (s *CreateApiKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApiKeyRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *CreateApiKeyRequest) GetLimitRate() *float64 {
	return s.LimitRate
}

func (s *CreateApiKeyRequest) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateApiKeyRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *CreateApiKeyRequest) GetTokenQuota() *int64 {
	return s.TokenQuota
}

func (s *CreateApiKeyRequest) SetDailyTokenQuota(v int64) *CreateApiKeyRequest {
	s.DailyTokenQuota = &v
	return s
}

func (s *CreateApiKeyRequest) SetInstanceId(v string) *CreateApiKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApiKeyRequest) SetKeyName(v string) *CreateApiKeyRequest {
	s.KeyName = &v
	return s
}

func (s *CreateApiKeyRequest) SetLimitRate(v float64) *CreateApiKeyRequest {
	s.LimitRate = &v
	return s
}

func (s *CreateApiKeyRequest) SetLimitType(v string) *CreateApiKeyRequest {
	s.LimitType = &v
	return s
}

func (s *CreateApiKeyRequest) SetQuantity(v int32) *CreateApiKeyRequest {
	s.Quantity = &v
	return s
}

func (s *CreateApiKeyRequest) SetTokenQuota(v int64) *CreateApiKeyRequest {
	s.TokenQuota = &v
	return s
}

func (s *CreateApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
