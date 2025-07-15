// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCloudPhoneNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewCloudPhoneNodesRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewCloudPhoneNodesRequest
	GetAutoRenew() *bool
	SetNodeIds(v []*string) *RenewCloudPhoneNodesRequest
	GetNodeIds() []*string
	SetPeriod(v int32) *RenewCloudPhoneNodesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewCloudPhoneNodesRequest
	GetPeriodUnit() *string
}

type RenewCloudPhoneNodesRequest struct {
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-renewal feature. In this case, the system automatically renews the instance upon expiration.
	//
	// 	- false (default): disables the auto-renewal feature. In this case, you need to manually renew the instance upon expiration.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The cloud phone matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The subscription duration. The unit is specified by `PeriodUnit`. Valid values:
	//
	// 	- When `PeriodUnit` is set to **year**: 1.
	//
	// 	- When `PeriodUnit` is set to **month**: 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s RenewCloudPhoneNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewCloudPhoneNodesRequest) GoString() string {
	return s.String()
}

func (s *RenewCloudPhoneNodesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewCloudPhoneNodesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewCloudPhoneNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *RenewCloudPhoneNodesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewCloudPhoneNodesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewCloudPhoneNodesRequest) SetAutoPay(v bool) *RenewCloudPhoneNodesRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetAutoRenew(v bool) *RenewCloudPhoneNodesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetNodeIds(v []*string) *RenewCloudPhoneNodesRequest {
	s.NodeIds = v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetPeriod(v int32) *RenewCloudPhoneNodesRequest {
	s.Period = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetPeriodUnit(v string) *RenewCloudPhoneNodesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) Validate() error {
	return dara.Validate(s)
}
