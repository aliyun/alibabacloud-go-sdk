// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *SearchCloudGtmAddressesRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *SearchCloudGtmAddressesRequest
	GetAddress() *string
	SetAddressId(v string) *SearchCloudGtmAddressesRequest
	GetAddressId() *string
	SetAvailableStatus(v string) *SearchCloudGtmAddressesRequest
	GetAvailableStatus() *string
	SetEnableStatus(v string) *SearchCloudGtmAddressesRequest
	GetEnableStatus() *string
	SetHealthStatus(v string) *SearchCloudGtmAddressesRequest
	GetHealthStatus() *string
	SetMonitorTemplateName(v string) *SearchCloudGtmAddressesRequest
	GetMonitorTemplateName() *string
	SetNameSearchCondition(v string) *SearchCloudGtmAddressesRequest
	GetNameSearchCondition() *string
	SetNames(v []*string) *SearchCloudGtmAddressesRequest
	GetNames() []*string
	SetPageNumber(v int32) *SearchCloudGtmAddressesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmAddressesRequest
	GetPageSize() *int32
	SetRemarkSearchCondition(v string) *SearchCloudGtmAddressesRequest
	GetRemarkSearchCondition() *string
	SetRemarks(v []*string) *SearchCloudGtmAddressesRequest
	GetRemarks() []*string
	SetType(v string) *SearchCloudGtmAddressesRequest
	GetType() *string
}

type SearchCloudGtmAddressesRequest struct {
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
	// Query by service address with precise conditions, supporting IP addresses or domain names.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The address ID. This ID uniquely identifies the address.
	//
	// example:
	//
	// addr-89518218114368**92
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// Search by address availability status with precise conditions:
	//
	// - available
	//
	// - unavailable
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// Query by exact address enable status:
	//
	// - enable: enabled status
	//
	// - disable: disabled status
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health state of the addresses that you want to query. Valid values:
	//
	// 	- ok: The addresses pass all health checks of the referenced health check templates.
	//
	// 	- ok_alert: The addresses fail some health checks of the referenced health check templates, but the addresses are deemed available.
	//
	// 	- ok_no_monitor: The addresses do not reference any health check template.
	//
	// 	- exceptional: The addresses fail some or all health checks of the referenced health check templates, and the addresses are deemed unavailable.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// Health check template name.
	//
	// example:
	//
	// Ping-IPv4
	MonitorTemplateName *string `json:"MonitorTemplateName,omitempty" xml:"MonitorTemplateName,omitempty"`
	// The logical condition for querying addresses by name. This parameter is required if you want to query addresses by name. Valid values:
	//
	// 	- and: displays the results that match all search conditions.
	//
	// 	- or: displays the results that match some or all search conditions.
	//
	// example:
	//
	// or
	NameSearchCondition *string `json:"NameSearchCondition,omitempty" xml:"NameSearchCondition,omitempty"`
	// Address name, usually for users to distinguish between different addresses.
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// Current page number, starting from 1, default is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of 100 and a default of 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The logical condition for querying addresses by additional description. This parameter is required if you want to query addresses by additional description. Valid values:
	//
	// and: displays the results that match all search conditions.
	//
	// or: displays the results that match some or all search conditions.
	//
	// example:
	//
	// or
	RemarkSearchCondition *string `json:"RemarkSearchCondition,omitempty" xml:"RemarkSearchCondition,omitempty"`
	// Remarks for the address.
	Remarks []*string `json:"Remarks,omitempty" xml:"Remarks,omitempty" type:"Repeated"`
	// Search precisely by address type conditions:
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
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchCloudGtmAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesRequest) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *SearchCloudGtmAddressesRequest) GetAddress() *string {
	return s.Address
}

func (s *SearchCloudGtmAddressesRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *SearchCloudGtmAddressesRequest) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmAddressesRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmAddressesRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmAddressesRequest) GetMonitorTemplateName() *string {
	return s.MonitorTemplateName
}

func (s *SearchCloudGtmAddressesRequest) GetNameSearchCondition() *string {
	return s.NameSearchCondition
}

func (s *SearchCloudGtmAddressesRequest) GetNames() []*string {
	return s.Names
}

func (s *SearchCloudGtmAddressesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmAddressesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmAddressesRequest) GetRemarkSearchCondition() *string {
	return s.RemarkSearchCondition
}

func (s *SearchCloudGtmAddressesRequest) GetRemarks() []*string {
	return s.Remarks
}

func (s *SearchCloudGtmAddressesRequest) GetType() *string {
	return s.Type
}

func (s *SearchCloudGtmAddressesRequest) SetAcceptLanguage(v string) *SearchCloudGtmAddressesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetAddress(v string) *SearchCloudGtmAddressesRequest {
	s.Address = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetAddressId(v string) *SearchCloudGtmAddressesRequest {
	s.AddressId = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetAvailableStatus(v string) *SearchCloudGtmAddressesRequest {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetEnableStatus(v string) *SearchCloudGtmAddressesRequest {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetHealthStatus(v string) *SearchCloudGtmAddressesRequest {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetMonitorTemplateName(v string) *SearchCloudGtmAddressesRequest {
	s.MonitorTemplateName = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetNameSearchCondition(v string) *SearchCloudGtmAddressesRequest {
	s.NameSearchCondition = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetNames(v []*string) *SearchCloudGtmAddressesRequest {
	s.Names = v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetPageNumber(v int32) *SearchCloudGtmAddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetPageSize(v int32) *SearchCloudGtmAddressesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetRemarkSearchCondition(v string) *SearchCloudGtmAddressesRequest {
	s.RemarkSearchCondition = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetRemarks(v []*string) *SearchCloudGtmAddressesRequest {
	s.Remarks = v
	return s
}

func (s *SearchCloudGtmAddressesRequest) SetType(v string) *SearchCloudGtmAddressesRequest {
	s.Type = &v
	return s
}

func (s *SearchCloudGtmAddressesRequest) Validate() error {
	return dara.Validate(s)
}
