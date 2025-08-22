// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDcdnRealTimeDeliveryProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *ListDcdnRealTimeDeliveryProjectRequest
	GetBusinessType() *string
	SetDomainName(v string) *ListDcdnRealTimeDeliveryProjectRequest
	GetDomainName() *string
	SetPageNumber(v int32) *ListDcdnRealTimeDeliveryProjectRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDcdnRealTimeDeliveryProjectRequest
	GetPageSize() *int32
}

type ListDcdnRealTimeDeliveryProjectRequest struct {
	// The type of the collected logs. Default value: cdn_log_access_l1. Valid values:
	//
	// 	- **cdn_log_access_l1**: access logs of Dynamic Content Delivery Network (DCDN) points of presence (POPs)
	//
	// 	- **cdn_log_origin**: back-to-origin logs
	//
	// 	- **cdn_log_er**: EdgeRoutine logs
	//
	// 	- By default, this parameter is left empty, and all logs are returned.
	//
	// example:
	//
	// cdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The domain name. You can specify only one domain name in each request. If this parameter is not specified, all domain names are queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The number of the page to return. Valid values: **1*	- to **100000**. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDcdnRealTimeDeliveryProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectRequest) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetBusinessType(v string) *ListDcdnRealTimeDeliveryProjectRequest {
	s.BusinessType = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetDomainName(v string) *ListDcdnRealTimeDeliveryProjectRequest {
	s.DomainName = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetPageNumber(v int32) *ListDcdnRealTimeDeliveryProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetPageSize(v int32) *ListDcdnRealTimeDeliveryProjectRequest {
	s.PageSize = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) Validate() error {
	return dara.Validate(s)
}
