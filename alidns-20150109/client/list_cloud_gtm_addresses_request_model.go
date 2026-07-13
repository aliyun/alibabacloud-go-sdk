// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListCloudGtmAddressesRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *ListCloudGtmAddressesRequest
	GetAddress() *string
	SetAddressId(v string) *ListCloudGtmAddressesRequest
	GetAddressId() *string
	SetClientToken(v string) *ListCloudGtmAddressesRequest
	GetClientToken() *string
	SetEnableStatus(v string) *ListCloudGtmAddressesRequest
	GetEnableStatus() *string
	SetHealthStatus(v string) *ListCloudGtmAddressesRequest
	GetHealthStatus() *string
	SetMonitorTemplateId(v string) *ListCloudGtmAddressesRequest
	GetMonitorTemplateId() *string
	SetName(v string) *ListCloudGtmAddressesRequest
	GetName() *string
	SetPageNumber(v int32) *ListCloudGtmAddressesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmAddressesRequest
	GetPageSize() *int32
	SetType(v string) *ListCloudGtmAddressesRequest
	GetType() *string
}

type ListCloudGtmAddressesRequest struct {
	// The language of the return value. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The IP address or domain name.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The unique ID of the address.
	//
	// example:
	//
	// addr-89518218114368****
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// A client token that is used to ensure the idempotence of the request. The client must generate a unique token for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The state of the address. Valid values:
	//
	// - enable: The address is enabled.
	//
	// - disable: The address is disabled.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health status of the address. Valid values:
	//
	// - ok: All health check tasks that are associated with the address are normal.
	//
	// - ok_alert: Some health check tasks that are associated with the address are abnormal, but the address is still considered normal.
	//
	// - ok_no_monitor: No health check template is associated with the address.
	//
	// - exceptional: Some or all health check tasks that are associated with the address are abnormal, and the address is considered abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The unique ID of the health check template.
	//
	// example:
	//
	// mtp-89518052425100****
	MonitorTemplateId *string `json:"MonitorTemplateId,omitempty" xml:"MonitorTemplateId,omitempty"`
	// The name of the address.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. The value starts from **1**. The default value is **1**.
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
	// The type of the address. Valid values:
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

func (s ListCloudGtmAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListCloudGtmAddressesRequest) GetAddress() *string {
	return s.Address
}

func (s *ListCloudGtmAddressesRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *ListCloudGtmAddressesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListCloudGtmAddressesRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListCloudGtmAddressesRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListCloudGtmAddressesRequest) GetMonitorTemplateId() *string {
	return s.MonitorTemplateId
}

func (s *ListCloudGtmAddressesRequest) GetName() *string {
	return s.Name
}

func (s *ListCloudGtmAddressesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmAddressesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmAddressesRequest) GetType() *string {
	return s.Type
}

func (s *ListCloudGtmAddressesRequest) SetAcceptLanguage(v string) *ListCloudGtmAddressesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetAddress(v string) *ListCloudGtmAddressesRequest {
	s.Address = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetAddressId(v string) *ListCloudGtmAddressesRequest {
	s.AddressId = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetClientToken(v string) *ListCloudGtmAddressesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetEnableStatus(v string) *ListCloudGtmAddressesRequest {
	s.EnableStatus = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetHealthStatus(v string) *ListCloudGtmAddressesRequest {
	s.HealthStatus = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetMonitorTemplateId(v string) *ListCloudGtmAddressesRequest {
	s.MonitorTemplateId = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetName(v string) *ListCloudGtmAddressesRequest {
	s.Name = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetPageNumber(v int32) *ListCloudGtmAddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetPageSize(v int32) *ListCloudGtmAddressesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) SetType(v string) *ListCloudGtmAddressesRequest {
	s.Type = &v
	return s
}

func (s *ListCloudGtmAddressesRequest) Validate() error {
	return dara.Validate(s)
}
