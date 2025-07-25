// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmInstanceConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *SearchCloudGtmInstanceConfigsRequest
	GetAcceptLanguage() *string
	SetAvailableStatus(v string) *SearchCloudGtmInstanceConfigsRequest
	GetAvailableStatus() *string
	SetClientToken(v string) *SearchCloudGtmInstanceConfigsRequest
	GetClientToken() *string
	SetEnableStatus(v string) *SearchCloudGtmInstanceConfigsRequest
	GetEnableStatus() *string
	SetHealthStatus(v string) *SearchCloudGtmInstanceConfigsRequest
	GetHealthStatus() *string
	SetInstanceId(v string) *SearchCloudGtmInstanceConfigsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *SearchCloudGtmInstanceConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmInstanceConfigsRequest
	GetPageSize() *int32
	SetRemark(v string) *SearchCloudGtmInstanceConfigsRequest
	GetRemark() *string
	SetScheduleDomainName(v string) *SearchCloudGtmInstanceConfigsRequest
	GetScheduleDomainName() *string
	SetScheduleZoneName(v string) *SearchCloudGtmInstanceConfigsRequest
	GetScheduleZoneName() *string
}

type SearchCloudGtmInstanceConfigsRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The availability state of the access domain name. Valid values:
	//
	// 	- available: If the access domain name is **enabled*	- and the health state is **normal**, the access domain name is deemed **available**.
	//
	// 	- unavailable: If the access domain name is **disabled*	- or the health state is **abnormal**, the access domain name is deemed **unavailable**.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enabling state of the access domain name. Valid values:
	//
	// 	- enable: The access domain name is enabled and the intelligent scheduling policy of the corresponding GTM instance takes effect.
	//
	// 	- disable: The access domain name is disabled and the intelligent scheduling policy of the corresponding GTM instance does not take effect.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health state of the access domain name. Valid values:
	//
	// 	- ok: The health state of the access domain name is normal and all address pools that are referenced by the access domain name are available.
	//
	// 	- ok_alert: The health state of the access domain name is warning and some of the address pools that are referenced by the access domain name are unavailable. In this case, only the available address pools are returned for Domain Name System (DNS) requests.
	//
	// 	- exceptional: The health state of the access domain name is abnormal and all address pools that are referenced by the access domain name are unavailable. In this case, addresses in the non-empty address pool with the smallest sequence number are preferentially used for fallback resolution. This returns DNS results for clients as much as possible.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Current page number, starting from 1, default is 1.
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
	// Remarks for the domain instance.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The access domain name. The value of this parameter is composed of the value of ScheduleHostname and the value of ScheduleZoneName.
	//
	// example:
	//
	// www.example.com
	ScheduleDomainName *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	// The zone such as example.com or subzone such as a.example.com of the access domain name. In most cases, the zone or subzone is hosted by the Public Authoritative DNS module of Alibaba Cloud DNS. This zone belongs to the account to which the GTM instance belongs.
	//
	// example:
	//
	// example.com
	ScheduleZoneName *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
}

func (s SearchCloudGtmInstanceConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetRemark() *string {
	return s.Remark
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *SearchCloudGtmInstanceConfigsRequest) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetAcceptLanguage(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetAvailableStatus(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetClientToken(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetEnableStatus(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetHealthStatus(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetInstanceId(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetPageNumber(v int32) *SearchCloudGtmInstanceConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetPageSize(v int32) *SearchCloudGtmInstanceConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetRemark(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.Remark = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetScheduleDomainName(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.ScheduleDomainName = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) SetScheduleZoneName(v string) *SearchCloudGtmInstanceConfigsRequest {
	s.ScheduleZoneName = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsRequest) Validate() error {
	return dara.Validate(s)
}
