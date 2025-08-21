// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlaEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeSlaEventListRequest
	GetEndTime() *int64
	SetIp(v string) *DescribeSlaEventListRequest
	GetIp() *string
	SetPage(v int64) *DescribeSlaEventListRequest
	GetPage() *int64
	SetPageSize(v int64) *DescribeSlaEventListRequest
	GetPageSize() *int64
	SetRegion(v string) *DescribeSlaEventListRequest
	GetRegion() *string
	SetStartTime(v int64) *DescribeSlaEventListRequest
	GetStartTime() *int64
}

type DescribeSlaEventListRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3289457398
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the service Valid values:
	//
	// 	- **cn**: Anti-DDoS Pro
	//
	// 	- **cn-hongkong**: Anti-DDoS Premium
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
	// 3289457398
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlaEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlaEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlaEventListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSlaEventListRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeSlaEventListRequest) GetPage() *int64 {
	return s.Page
}

func (s *DescribeSlaEventListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSlaEventListRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeSlaEventListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSlaEventListRequest) SetEndTime(v int64) *DescribeSlaEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlaEventListRequest) SetIp(v string) *DescribeSlaEventListRequest {
	s.Ip = &v
	return s
}

func (s *DescribeSlaEventListRequest) SetPage(v int64) *DescribeSlaEventListRequest {
	s.Page = &v
	return s
}

func (s *DescribeSlaEventListRequest) SetPageSize(v int64) *DescribeSlaEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlaEventListRequest) SetRegion(v string) *DescribeSlaEventListRequest {
	s.Region = &v
	return s
}

func (s *DescribeSlaEventListRequest) SetStartTime(v int64) *DescribeSlaEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlaEventListRequest) Validate() error {
	return dara.Validate(s)
}
