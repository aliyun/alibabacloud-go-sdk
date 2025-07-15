// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSummaryDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpireInstanceNum(v int32) *DescribeSummaryDataResponseBody
	GetExpireInstanceNum() *int32
	SetRegion(v string) *DescribeSummaryDataResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeSummaryDataResponseBody
	GetRequestId() *string
	SetUsageApiNum(v int32) *DescribeSummaryDataResponseBody
	GetUsageApiNum() *int32
	SetUsageGroupNum(v int32) *DescribeSummaryDataResponseBody
	GetUsageGroupNum() *int32
	SetUsageInstanceNum(v int32) *DescribeSummaryDataResponseBody
	GetUsageInstanceNum() *int32
}

type DescribeSummaryDataResponseBody struct {
	// The number of subscription dedicated instances that expire in 14 days or less.
	//
	// example:
	//
	// 1
	ExpireInstanceNum *int32 `json:"ExpireInstanceNum,omitempty" xml:"ExpireInstanceNum,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of APIs.
	//
	// example:
	//
	// 10
	UsageApiNum *int32 `json:"UsageApiNum,omitempty" xml:"UsageApiNum,omitempty"`
	// The number of API groups.
	//
	// example:
	//
	// 1
	UsageGroupNum *int32 `json:"UsageGroupNum,omitempty" xml:"UsageGroupNum,omitempty"`
	// The number of running dedicated instances.
	//
	// example:
	//
	// 1
	UsageInstanceNum *int32 `json:"UsageInstanceNum,omitempty" xml:"UsageInstanceNum,omitempty"`
}

func (s DescribeSummaryDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSummaryDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSummaryDataResponseBody) GetExpireInstanceNum() *int32 {
	return s.ExpireInstanceNum
}

func (s *DescribeSummaryDataResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeSummaryDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSummaryDataResponseBody) GetUsageApiNum() *int32 {
	return s.UsageApiNum
}

func (s *DescribeSummaryDataResponseBody) GetUsageGroupNum() *int32 {
	return s.UsageGroupNum
}

func (s *DescribeSummaryDataResponseBody) GetUsageInstanceNum() *int32 {
	return s.UsageInstanceNum
}

func (s *DescribeSummaryDataResponseBody) SetExpireInstanceNum(v int32) *DescribeSummaryDataResponseBody {
	s.ExpireInstanceNum = &v
	return s
}

func (s *DescribeSummaryDataResponseBody) SetRegion(v string) *DescribeSummaryDataResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeSummaryDataResponseBody) SetRequestId(v string) *DescribeSummaryDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSummaryDataResponseBody) SetUsageApiNum(v int32) *DescribeSummaryDataResponseBody {
	s.UsageApiNum = &v
	return s
}

func (s *DescribeSummaryDataResponseBody) SetUsageGroupNum(v int32) *DescribeSummaryDataResponseBody {
	s.UsageGroupNum = &v
	return s
}

func (s *DescribeSummaryDataResponseBody) SetUsageInstanceNum(v int32) *DescribeSummaryDataResponseBody {
	s.UsageInstanceNum = &v
	return s
}

func (s *DescribeSummaryDataResponseBody) Validate() error {
	return dara.Validate(s)
}
