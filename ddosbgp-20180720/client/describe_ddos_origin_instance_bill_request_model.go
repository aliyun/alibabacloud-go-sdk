// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosOriginInstanceBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDdosOriginInstanceBillRequest
	GetEndTime() *int64
	SetIsShowList(v bool) *DescribeDdosOriginInstanceBillRequest
	GetIsShowList() *bool
	SetStartTime(v int64) *DescribeDdosOriginInstanceBillRequest
	GetStartTime() *int64
	SetType(v string) *DescribeDdosOriginInstanceBillRequest
	GetType() *string
}

type DescribeDdosOriginInstanceBillRequest struct {
	// The end of the time range to query. The value is a timestamp. Unit: milliseconds. The time span between StartTime and EndTime cannot exceed 30 days.
	//
	// example:
	//
	// 1711382399410
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to display the bill details. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsShowList *bool `json:"IsShowList,omitempty" xml:"IsShowList,omitempty"`
	// The beginning of the time range to query. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1711209600410
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The bill type. Valid values:
	//
	// 	- **flow_cn**: the bill for the clean bandwidth of elastic IP addresses (EIPs) with Anti-DDoS (Enhanced) enabled in the Chinese mainland.
	//
	// 	- **flow_ov**: the bill for the clean bandwidth of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland.
	//
	// 	- **standard_assets_flow_cn**: the bill for the clean bandwidth of regular Alibaba Cloud services in the Chinese mainland.
	//
	// 	- **standard_assets_flow_ov**: the bill for the clean bandwidth of regular Alibaba Cloud services outside the Chinese mainland.
	//
	// 	- **function**: the bill for the basic fee.
	//
	// 	- **ip_count**: the bill for protected IP addresses.
	//
	// 	- **monthly_summary**: the monthly summary bill.
	//
	// example:
	//
	// function
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDdosOriginInstanceBillRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDdosOriginInstanceBillRequest) GetIsShowList() *bool {
	return s.IsShowList
}

func (s *DescribeDdosOriginInstanceBillRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDdosOriginInstanceBillRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDdosOriginInstanceBillRequest) SetEndTime(v int64) *DescribeDdosOriginInstanceBillRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillRequest) SetIsShowList(v bool) *DescribeDdosOriginInstanceBillRequest {
	s.IsShowList = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillRequest) SetStartTime(v int64) *DescribeDdosOriginInstanceBillRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillRequest) SetType(v string) *DescribeDdosOriginInstanceBillRequest {
	s.Type = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillRequest) Validate() error {
	return dara.Validate(s)
}
