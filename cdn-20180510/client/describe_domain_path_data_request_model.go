// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainPathDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainPathDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainPathDataRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeDomainPathDataRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDomainPathDataRequest
	GetPageSize() *int32
	SetPath(v string) *DescribeDomainPathDataRequest
	GetPath() *string
	SetStartTime(v string) *DescribeDomainPathDataRequest
	GetStartTime() *string
}

type DescribeDomainPathDataRequest struct {
	// The accelerated domain name.
	//
	// > You can specify only one domain name in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The interval between the start time and end time must be less than 30 days. Example: 2016-10-21T04:00:00Z.
	//
	// example:
	//
	// 2016-10-21T04:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. Pages start from page **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: integers from **1*	- to **1000**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The paths that you want to query. Separate paths with forward slashes (/). If you do not set this parameter, all paths are queried. If you set the value to a directory, it must end with a forward slash (/).
	//
	// > Fuzzy match is not supported. If you want data to be collected based on a directory, you can specify a specific directory, for example, directory/path/. In this case, bandwidth data is collected based on directory/path/.
	//
	// example:
	//
	// /path/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2016-10-20T04:00:00Z.
	//
	// example:
	//
	// 2016-10-20T04:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainPathDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPathDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainPathDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainPathDataRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDomainPathDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainPathDataRequest) GetPath() *string {
	return s.Path
}

func (s *DescribeDomainPathDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainPathDataRequest) SetDomainName(v string) *DescribeDomainPathDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetEndTime(v string) *DescribeDomainPathDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetPageNumber(v int32) *DescribeDomainPathDataRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetPageSize(v int32) *DescribeDomainPathDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetPath(v string) *DescribeDomainPathDataRequest {
	s.Path = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetStartTime(v string) *DescribeDomainPathDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainPathDataRequest) Validate() error {
	return dara.Validate(s)
}
