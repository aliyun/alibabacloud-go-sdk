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
	// The language of the return value. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English. This is the default value.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Performs an exact search by endpoint. IP addresses and domain names are supported.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the address. The address ID is a unique identifier.
	//
	// example:
	//
	// addr-89518218114368****
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// Performs an exact search by the availability status of the address.
	//
	// - available: The address is available.
	//
	// - unavailable: The address is unavailable.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// Performs an exact search by the status of the address.
	//
	// - enable: The address is enabled.
	//
	// - disable: The address is disabled.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// Performs an exact search by the health status of the address.
	//
	// - ok: All health check tasks for the referenced health check template are normal.
	//
	// - ok_alert: Some health check tasks for the referenced health check template are abnormal, but the address is still considered normal.
	//
	// - ok_no_monitor: The address does not reference any health check templates.
	//
	// - exceptional: Some or all health check tasks for the referenced health check template are abnormal, and the address is considered abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The name of the health check template.
	//
	// example:
	//
	// Ping-IPv4
	MonitorTemplateName *string `json:"MonitorTemplateName,omitempty" xml:"MonitorTemplateName,omitempty"`
	// The search logic for querying by address name. This parameter is required when you query by address name.
	//
	// - and: The query returns results that match all the specified keywords.
	//
	// - or: The query returns results that match some or all of the specified keywords.
	//
	// example:
	//
	// or
	NameSearchCondition *string `json:"NameSearchCondition,omitempty" xml:"NameSearchCondition,omitempty"`
	// The name of the address. This name is used for easy identification.
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// The current page number. The value starts from 1. The default value is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The maximum value is 100. The default value is 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search logic for querying by address remarks. This parameter is required when you query by address remarks.
	//
	// and: The query returns results that match all the specified keywords.
	//
	// or: The query returns results that match some or all of the specified keywords.
	//
	// example:
	//
	// or
	RemarkSearchCondition *string `json:"RemarkSearchCondition,omitempty" xml:"RemarkSearchCondition,omitempty"`
	// The remarks for the address.
	Remarks []*string `json:"Remarks,omitempty" xml:"Remarks,omitempty" type:"Repeated"`
	// Performs an exact search by the address type.
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
