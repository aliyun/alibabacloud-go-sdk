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
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Address pool name, supports fuzzy search for the entered address pool name.
	//
	// example:
	//
	// AddressPool-1
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// Address pool type, supports precise query for address pool types:
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
	// Address pool availability status, supporting precise queries for address pool availability:
	//
	// - available: Available
	//
	// - unavailable: Unavailable
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Address pool enable status, supports precise query of address pool enable status:
	//
	// - enable: Enabled status
	//
	// - disable: Disabled status
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health state of the address pool. You can enter a health state for exact search. Valid values:
	//
	// ok: The health state of the address pool is normal and all addresses that are referenced by the address pool are available.
	//
	// ok_alert: The health state of the address pool is warning and some of the addresses that are referenced by the address pool are unavailable. However, the address pool is deemed normal. In this case, only the available addresses are returned for Domain Name System (DNS) requests.
	//
	// exceptional: The health state of the address pool is abnormal and some or all of the addresses that are referenced by the address pool are unavailable. In this case, the address pool is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// Current page number, starting from 1, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of 100 and a default of 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Address pool remarks, supporting fuzzy search for the input remarks.
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
