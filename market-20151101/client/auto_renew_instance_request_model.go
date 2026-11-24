// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewCycle(v string) *AutoRenewInstanceRequest
	GetAutoRenewCycle() *string
	SetAutoRenewDuration(v int32) *AutoRenewInstanceRequest
	GetAutoRenewDuration() *int32
	SetOrderBizId(v int64) *AutoRenewInstanceRequest
	GetOrderBizId() *int64
	SetOwnerId(v int64) *AutoRenewInstanceRequest
	GetOwnerId() *int64
	SetType(v string) *AutoRenewInstanceRequest
	GetType() *string
}

type AutoRenewInstanceRequest struct {
	// The cycle unit. This parameter is required when enabling auto-renewal. Valid values:
	//
	// - Day: day.
	//
	// - Month: month.
	//
	// - Year: year.
	//
	// example:
	//
	// Month
	AutoRenewCycle *string `json:"AutoRenewCycle,omitempty" xml:"AutoRenewCycle,omitempty"`
	// The renewal cycle. This parameter is required when enabling auto-renewal.
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// The Alibaba Cloud Marketplace instance ID. This parameter is required.
	//
	// This parameter is required.
	OrderBizId *int64 `json:"OrderBizId,omitempty" xml:"OrderBizId,omitempty"`
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operation type. Valid values:
	//
	// - open: enables auto-renewal.
	//
	// - close: disables auto-renewal.
	//
	// This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AutoRenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AutoRenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *AutoRenewInstanceRequest) GetAutoRenewCycle() *string {
	return s.AutoRenewCycle
}

func (s *AutoRenewInstanceRequest) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *AutoRenewInstanceRequest) GetOrderBizId() *int64 {
	return s.OrderBizId
}

func (s *AutoRenewInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AutoRenewInstanceRequest) GetType() *string {
	return s.Type
}

func (s *AutoRenewInstanceRequest) SetAutoRenewCycle(v string) *AutoRenewInstanceRequest {
	s.AutoRenewCycle = &v
	return s
}

func (s *AutoRenewInstanceRequest) SetAutoRenewDuration(v int32) *AutoRenewInstanceRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *AutoRenewInstanceRequest) SetOrderBizId(v int64) *AutoRenewInstanceRequest {
	s.OrderBizId = &v
	return s
}

func (s *AutoRenewInstanceRequest) SetOwnerId(v int64) *AutoRenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *AutoRenewInstanceRequest) SetType(v string) *AutoRenewInstanceRequest {
	s.Type = &v
	return s
}

func (s *AutoRenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
