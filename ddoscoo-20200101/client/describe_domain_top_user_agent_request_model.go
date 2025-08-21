// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopUserAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainTopUserAgentRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainTopUserAgentRequest
	GetEndTime() *int64
	SetLimit(v int64) *DescribeDomainTopUserAgentRequest
	GetLimit() *int64
	SetRegion(v string) *DescribeDomainTopUserAgentRequest
	GetRegion() *string
	SetStartTime(v int64) *DescribeDomainTopUserAgentRequest
	GetStartTime() *int64
}

type DescribeDomainTopUserAgentRequest struct {
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1708352700
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The region in which your service is deployed. Valid values:
	//
	// 	- **cn**: a region in the Chinese mainland.
	//
	// 	- **cn-hongkong**: a region outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1609430400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTopUserAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUserAgentRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUserAgentRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopUserAgentRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainTopUserAgentRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeDomainTopUserAgentRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDomainTopUserAgentRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainTopUserAgentRequest) SetDomain(v string) *DescribeDomainTopUserAgentRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopUserAgentRequest) SetEndTime(v int64) *DescribeDomainTopUserAgentRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopUserAgentRequest) SetLimit(v int64) *DescribeDomainTopUserAgentRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDomainTopUserAgentRequest) SetRegion(v string) *DescribeDomainTopUserAgentRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainTopUserAgentRequest) SetStartTime(v int64) *DescribeDomainTopUserAgentRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopUserAgentRequest) Validate() error {
	return dara.Validate(s)
}
