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
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enabling state of the access domain name. Valid values:
	//
	// 	- enable: The access domain name is enabled and the intelligent scheduling policy of the GTM instance takes effect.
	//
	// 	- disable: The access domain name is disabled and the intelligent scheduling policy of the GTM instance does not take effect.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The ID of the GTM 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// Remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The GTM access domain name. The value of this parameter is composed of the value of ScheduleHostname and the value of ScheduleZoneName.
	//
	// example:
	//
	// www.example.com
	ScheduleDomainName *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	// The zone (such as example.com) or subzone (such as a.example.com) of the GTM access domain name. In most cases, the zone or subzone is hosted in Authoritative DNS Resolution of the Alibaba Cloud DNS console within the account to which the GTM instance belongs.
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
