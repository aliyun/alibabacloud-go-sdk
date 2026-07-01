// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversionDataIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversionRate(v string) *ConversionDataIntlRequest
	GetConversionRate() *string
	SetOwnerId(v int64) *ConversionDataIntlRequest
	GetOwnerId() *int64
	SetReportTime(v int64) *ConversionDataIntlRequest
	GetReportTime() *int64
	SetResourceOwnerAccount(v string) *ConversionDataIntlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConversionDataIntlRequest
	GetResourceOwnerId() *int64
}

type ConversionDataIntlRequest struct {
	// The conversion rate monitoring value.
	//
	// >This parameter is of the double type. Valid values: [0,1].
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.53
	ConversionRate *string `json:"ConversionRate,omitempty" xml:"ConversionRate,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The point in time when the conversion rate is observed. The value must be a Unix timestamp in milliseconds, represented as a long integer.
	//
	// >If you do not specify this parameter, the current timestamp is used by default.
	//
	// example:
	//
	// 1349055900000
	ReportTime           *int64  `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ConversionDataIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s ConversionDataIntlRequest) GoString() string {
	return s.String()
}

func (s *ConversionDataIntlRequest) GetConversionRate() *string {
	return s.ConversionRate
}

func (s *ConversionDataIntlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConversionDataIntlRequest) GetReportTime() *int64 {
	return s.ReportTime
}

func (s *ConversionDataIntlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConversionDataIntlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConversionDataIntlRequest) SetConversionRate(v string) *ConversionDataIntlRequest {
	s.ConversionRate = &v
	return s
}

func (s *ConversionDataIntlRequest) SetOwnerId(v int64) *ConversionDataIntlRequest {
	s.OwnerId = &v
	return s
}

func (s *ConversionDataIntlRequest) SetReportTime(v int64) *ConversionDataIntlRequest {
	s.ReportTime = &v
	return s
}

func (s *ConversionDataIntlRequest) SetResourceOwnerAccount(v string) *ConversionDataIntlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConversionDataIntlRequest) SetResourceOwnerId(v int64) *ConversionDataIntlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConversionDataIntlRequest) Validate() error {
	return dara.Validate(s)
}
