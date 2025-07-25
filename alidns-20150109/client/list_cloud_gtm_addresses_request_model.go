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
	// Return language value, options:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// IP address or domain name.
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Indicates the current availability of the address:
	//
	// - enable: Enabled status
	//
	// - disable: Disabled status
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health check state of the address. Valid values:
	//
	// 	- ok: The address passes all health checks of the referenced health check templates.
	//
	// 	- ok_alert: The address fails some health checks of the referenced health check templates but the address is deemed normal.
	//
	// 	- ok_no_monitor: The address does not reference a health check template.
	//
	// 	- exceptional: The address fails some or all health checks of the referenced health check templates and the address is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID of the health check template. This ID uniquely identifies the health check template.
	//
	// example:
	//
	// mtp-89518052425100**80
	MonitorTemplateId *string `json:"MonitorTemplateId,omitempty" xml:"MonitorTemplateId,omitempty"`
	// Address name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Current page number, starting from **1**, default is **1**.
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
	// Address type:
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
