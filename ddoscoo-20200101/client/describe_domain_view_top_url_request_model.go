// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewTopUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainViewTopUrlRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainViewTopUrlRequest
	GetEndTime() *int64
	SetInerval(v int64) *DescribeDomainViewTopUrlRequest
	GetInerval() *int64
	SetResourceGroupId(v string) *DescribeDomainViewTopUrlRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainViewTopUrlRequest
	GetStartTime() *int64
	SetTop(v int32) *DescribeDomainViewTopUrlRequest
	GetTop() *int32
}

type DescribeDomainViewTopUrlRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Inerval *int64 `json:"Inerval,omitempty" xml:"Inerval,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of URLs to query. Valid values: **1*	- to **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DescribeDomainViewTopUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainViewTopUrlRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainViewTopUrlRequest) GetInerval() *int64 {
	return s.Inerval
}

func (s *DescribeDomainViewTopUrlRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainViewTopUrlRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainViewTopUrlRequest) GetTop() *int32 {
	return s.Top
}

func (s *DescribeDomainViewTopUrlRequest) SetDomain(v string) *DescribeDomainViewTopUrlRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetEndTime(v int64) *DescribeDomainViewTopUrlRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetInerval(v int64) *DescribeDomainViewTopUrlRequest {
	s.Inerval = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetResourceGroupId(v string) *DescribeDomainViewTopUrlRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetStartTime(v int64) *DescribeDomainViewTopUrlRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetTop(v int32) *DescribeDomainViewTopUrlRequest {
	s.Top = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) Validate() error {
	return dara.Validate(s)
}
