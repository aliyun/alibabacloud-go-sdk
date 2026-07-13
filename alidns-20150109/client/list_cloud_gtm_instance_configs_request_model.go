// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmInstanceConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListCloudGtmInstanceConfigsRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *ListCloudGtmInstanceConfigsRequest
	GetClientToken() *string
	SetEnableStatus(v string) *ListCloudGtmInstanceConfigsRequest
	GetEnableStatus() *string
	SetInstanceId(v string) *ListCloudGtmInstanceConfigsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCloudGtmInstanceConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmInstanceConfigsRequest
	GetPageSize() *int32
	SetRemark(v string) *ListCloudGtmInstanceConfigsRequest
	GetRemark() *string
	SetScheduleDomainName(v string) *ListCloudGtmInstanceConfigsRequest
	GetScheduleDomainName() *string
	SetScheduleZoneName(v string) *ListCloudGtmInstanceConfigsRequest
	GetScheduleZoneName() *string
}

type ListCloudGtmInstanceConfigsRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a unique token for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The status of the domain name instance:
	//
	// - enable: The GTM instance uses intelligent scheduling policies.
	//
	// - disable: The intelligent scheduling policies of the GTM instance are unavailable.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value starts from **1**. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. Maximum value: **100**. Default value: **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The GTM access domain name. The domain name is a combination of the host record (ScheduleHostname) and the root or subdomain (ScheduleZoneName).
	//
	// example:
	//
	// www.example.com
	ScheduleDomainName *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	// The root domain, such as example.com, or subdomain, such as a.example.com, of the GTM access domain name. This is typically a domain name that is hosted in an authoritative zone in the Alibaba Cloud DNS console and belongs to the same account as the GTM instance.
	//
	// example:
	//
	// example.com
	ScheduleZoneName *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
}

func (s ListCloudGtmInstanceConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstanceConfigsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListCloudGtmInstanceConfigsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListCloudGtmInstanceConfigsRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListCloudGtmInstanceConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudGtmInstanceConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmInstanceConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmInstanceConfigsRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListCloudGtmInstanceConfigsRequest) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *ListCloudGtmInstanceConfigsRequest) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *ListCloudGtmInstanceConfigsRequest) SetAcceptLanguage(v string) *ListCloudGtmInstanceConfigsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) SetClientToken(v string) *ListCloudGtmInstanceConfigsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) SetEnableStatus(v string) *ListCloudGtmInstanceConfigsRequest {
	s.EnableStatus = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) SetInstanceId(v string) *ListCloudGtmInstanceConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) SetPageNumber(v int32) *ListCloudGtmInstanceConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) SetPageSize(v int32) *ListCloudGtmInstanceConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) SetRemark(v string) *ListCloudGtmInstanceConfigsRequest {
	s.Remark = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) SetScheduleDomainName(v string) *ListCloudGtmInstanceConfigsRequest {
	s.ScheduleDomainName = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) SetScheduleZoneName(v string) *ListCloudGtmInstanceConfigsRequest {
	s.ScheduleZoneName = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsRequest) Validate() error {
	return dara.Validate(s)
}
