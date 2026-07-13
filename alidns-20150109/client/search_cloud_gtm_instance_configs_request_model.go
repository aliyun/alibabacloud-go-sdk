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
	// - `zh-CN`: Chinese
	//
	// - `en-US` (default): English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The service availability status of the instance configuration. Valid values:
	//
	// - `available`: The service for the GTM access domain is available when the instance configuration is **enabled*	- and its health status is **Normal*	- or Alert.
	//
	// - `unavailable`: The service for the GTM access domain is unavailable when the instance configuration is **disabled*	- or its health status is **Exceptional**.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// A client-generated token to ensure the idempotence of the request. The token must be unique across requests and can contain up to 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The status of the instance configuration. Valid values:
	//
	// - `enable`: Enabled. The intelligent scheduling policy of the GTM instance is in effect.
	//
	// - `disable`: Disabled. The intelligent scheduling policy of the GTM instance is unavailable.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health status of the instance configuration. Valid values:
	//
	// - `ok`: Normal. All address pools referenced by the GTM access domain are available.
	//
	// - `ok_alert`: Alert. Some address pools referenced by the GTM access domain are unavailable. In this state, DNS resolution continues for available address pools but stops for unavailable ones.
	//
	// - `exceptional`: Exceptional. All address pools referenced by the GTM access domain are unavailable. In this case, failover resolution uses the addresses from the non-empty address pool with the smallest sequence number to ensure clients receive a resolution result.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Pages start from 1. The default value is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. The maximum value is **100*	- and the default value is **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A note for the instance configuration.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The GTM access domain. It is formed by combining the host record (`ScheduleHostname`) with the primary or subdomain name (`ScheduleZoneName`).
	//
	// example:
	//
	// www.example.com
	ScheduleDomainName *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	// The primary domain name (for example, `example.com`) or subdomain name (for example, `a.example.com`) of the GTM access domain. This is typically a domain name managed by Alibaba Cloud DNS under the same account that owns the GTM instance.
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
