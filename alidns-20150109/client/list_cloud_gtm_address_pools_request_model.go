// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAddressPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListCloudGtmAddressPoolsRequest
	GetAcceptLanguage() *string
	SetAddressPoolName(v string) *ListCloudGtmAddressPoolsRequest
	GetAddressPoolName() *string
	SetAddressPoolType(v string) *ListCloudGtmAddressPoolsRequest
	GetAddressPoolType() *string
	SetClientToken(v string) *ListCloudGtmAddressPoolsRequest
	GetClientToken() *string
	SetEnableStatus(v string) *ListCloudGtmAddressPoolsRequest
	GetEnableStatus() *string
	SetPageNumber(v int32) *ListCloudGtmAddressPoolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmAddressPoolsRequest
	GetPageSize() *int32
	SetRemark(v string) *ListCloudGtmAddressPoolsRequest
	GetRemark() *string
}

type ListCloudGtmAddressPoolsRequest struct {
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
	// Address pool name.
	//
	// example:
	//
	// AddressPool-1
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// The type of the address pool. Valid values:
	//
	// 	- IPv4: indicates that the service address to be resolved is an IPv4 address.
	//
	// 	- IPv6: indicates that the service address to be resolved is an IPv6 address.
	//
	// 	- domain: indicates that the service address to be resolved is a domain name.
	//
	// example:
	//
	// IPv4
	AddressPoolType *string `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enabling state of the address pool. Valid values:
	//
	// 	- enable: The address pool is enabled.
	//
	// 	- disable: The address pool is disabled.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// Current page number, starting at **1**, default is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of **100**, and a default of **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The additional description of the address pool.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ListCloudGtmAddressPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListCloudGtmAddressPoolsRequest) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *ListCloudGtmAddressPoolsRequest) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *ListCloudGtmAddressPoolsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListCloudGtmAddressPoolsRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListCloudGtmAddressPoolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmAddressPoolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmAddressPoolsRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListCloudGtmAddressPoolsRequest) SetAcceptLanguage(v string) *ListCloudGtmAddressPoolsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCloudGtmAddressPoolsRequest) SetAddressPoolName(v string) *ListCloudGtmAddressPoolsRequest {
	s.AddressPoolName = &v
	return s
}

func (s *ListCloudGtmAddressPoolsRequest) SetAddressPoolType(v string) *ListCloudGtmAddressPoolsRequest {
	s.AddressPoolType = &v
	return s
}

func (s *ListCloudGtmAddressPoolsRequest) SetClientToken(v string) *ListCloudGtmAddressPoolsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListCloudGtmAddressPoolsRequest) SetEnableStatus(v string) *ListCloudGtmAddressPoolsRequest {
	s.EnableStatus = &v
	return s
}

func (s *ListCloudGtmAddressPoolsRequest) SetPageNumber(v int32) *ListCloudGtmAddressPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmAddressPoolsRequest) SetPageSize(v int32) *ListCloudGtmAddressPoolsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmAddressPoolsRequest) SetRemark(v string) *ListCloudGtmAddressPoolsRequest {
	s.Remark = &v
	return s
}

func (s *ListCloudGtmAddressPoolsRequest) Validate() error {
	return dara.Validate(s)
}
