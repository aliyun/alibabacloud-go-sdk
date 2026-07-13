// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmAddressPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *SearchCloudGtmAddressPoolsRequest
	GetAcceptLanguage() *string
	SetAddressPoolName(v string) *SearchCloudGtmAddressPoolsRequest
	GetAddressPoolName() *string
	SetAddressPoolType(v string) *SearchCloudGtmAddressPoolsRequest
	GetAddressPoolType() *string
	SetAvailableStatus(v string) *SearchCloudGtmAddressPoolsRequest
	GetAvailableStatus() *string
	SetClientToken(v string) *SearchCloudGtmAddressPoolsRequest
	GetClientToken() *string
	SetEnableStatus(v string) *SearchCloudGtmAddressPoolsRequest
	GetEnableStatus() *string
	SetHealthStatus(v string) *SearchCloudGtmAddressPoolsRequest
	GetHealthStatus() *string
	SetPageNumber(v int32) *SearchCloudGtmAddressPoolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmAddressPoolsRequest
	GetPageSize() *int32
	SetRemark(v string) *SearchCloudGtmAddressPoolsRequest
	GetRemark() *string
}

type SearchCloudGtmAddressPoolsRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US (default): English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The name of the address pool. Fuzzy match is supported.
	//
	// example:
	//
	// AddressPool-1
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// The type of the address pool. Exact match is supported. Valid values:
	//
	// - IPv4
	//
	// - IPv6
	//
	// - domain
	//
	// example:
	//
	// IPv4
	AddressPoolType *string `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	// The availability of the address pool. Exact match is supported. Valid values:
	//
	// - available: The address pool is available.
	//
	// - unavailable: The address pool is unavailable.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a unique value from your client for this parameter. The client token can contain only ASCII characters and must be a maximum of 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The status of the address pool. Exact match is supported. Valid values:
	//
	// - enable: The address pool is enabled.
	//
	// - disable: The address pool is disabled.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health status of the address pool. Exact match is supported. Valid values:
	//
	// ok: Normal. All addresses in the address pool are available.
	//
	// ok_alert: Warning. Some addresses in the address pool are unavailable, but the address pool is still considered normal. In the warning state, available addresses are resolved as expected, while unavailable addresses are not.
	//
	// exceptional: Abnormal. Some or all addresses in the address pool are unavailable, and the address pool is considered abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The page number. The value starts from 1. The default value is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is 100. The default value is 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The remarks for the address pool. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s SearchCloudGtmAddressPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsRequest) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *SearchCloudGtmAddressPoolsRequest) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *SearchCloudGtmAddressPoolsRequest) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *SearchCloudGtmAddressPoolsRequest) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmAddressPoolsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SearchCloudGtmAddressPoolsRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmAddressPoolsRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmAddressPoolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmAddressPoolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmAddressPoolsRequest) GetRemark() *string {
	return s.Remark
}

func (s *SearchCloudGtmAddressPoolsRequest) SetAcceptLanguage(v string) *SearchCloudGtmAddressPoolsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetAddressPoolName(v string) *SearchCloudGtmAddressPoolsRequest {
	s.AddressPoolName = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetAddressPoolType(v string) *SearchCloudGtmAddressPoolsRequest {
	s.AddressPoolType = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetAvailableStatus(v string) *SearchCloudGtmAddressPoolsRequest {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetClientToken(v string) *SearchCloudGtmAddressPoolsRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetEnableStatus(v string) *SearchCloudGtmAddressPoolsRequest {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetHealthStatus(v string) *SearchCloudGtmAddressPoolsRequest {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetPageNumber(v int32) *SearchCloudGtmAddressPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetPageSize(v int32) *SearchCloudGtmAddressPoolsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) SetRemark(v string) *SearchCloudGtmAddressPoolsRequest {
	s.Remark = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsRequest) Validate() error {
	return dara.Validate(s)
}
