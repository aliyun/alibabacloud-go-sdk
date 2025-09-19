// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanTaskStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDealedRiskNum(v int32) *DescribeScanTaskStatisticsResponseBody
	GetDealedRiskNum() *int32
	SetPersonalTaskNum(v int64) *DescribeScanTaskStatisticsResponseBody
	GetPersonalTaskNum() *int64
	SetRequestId(v string) *DescribeScanTaskStatisticsResponseBody
	GetRequestId() *string
	SetTotalTaskNum(v int64) *DescribeScanTaskStatisticsResponseBody
	GetTotalTaskNum() *int64
	SetUserNum(v int64) *DescribeScanTaskStatisticsResponseBody
	GetUserNum() *int64
}

type DescribeScanTaskStatisticsResponseBody struct {
	// The number of risks that are handled for the user.
	//
	// example:
	//
	// 11
	DealedRiskNum *int32 `json:"DealedRiskNum,omitempty" xml:"DealedRiskNum,omitempty"`
	// The total number of tasks that are created for the user.
	//
	// example:
	//
	// 11
	PersonalTaskNum *int64 `json:"PersonalTaskNum,omitempty" xml:"PersonalTaskNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 765EDBDE-1686-5DBA-B76F-2E0XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of virus detection tasks.
	//
	// example:
	//
	// 11
	TotalTaskNum *int64 `json:"TotalTaskNum,omitempty" xml:"TotalTaskNum,omitempty"`
	// The number of risks that are detected for the user.
	//
	// example:
	//
	// 11
	UserNum *int64 `json:"UserNum,omitempty" xml:"UserNum,omitempty"`
}

func (s DescribeScanTaskStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanTaskStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskStatisticsResponseBody) GetDealedRiskNum() *int32 {
	return s.DealedRiskNum
}

func (s *DescribeScanTaskStatisticsResponseBody) GetPersonalTaskNum() *int64 {
	return s.PersonalTaskNum
}

func (s *DescribeScanTaskStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScanTaskStatisticsResponseBody) GetTotalTaskNum() *int64 {
	return s.TotalTaskNum
}

func (s *DescribeScanTaskStatisticsResponseBody) GetUserNum() *int64 {
	return s.UserNum
}

func (s *DescribeScanTaskStatisticsResponseBody) SetDealedRiskNum(v int32) *DescribeScanTaskStatisticsResponseBody {
	s.DealedRiskNum = &v
	return s
}

func (s *DescribeScanTaskStatisticsResponseBody) SetPersonalTaskNum(v int64) *DescribeScanTaskStatisticsResponseBody {
	s.PersonalTaskNum = &v
	return s
}

func (s *DescribeScanTaskStatisticsResponseBody) SetRequestId(v string) *DescribeScanTaskStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScanTaskStatisticsResponseBody) SetTotalTaskNum(v int64) *DescribeScanTaskStatisticsResponseBody {
	s.TotalTaskNum = &v
	return s
}

func (s *DescribeScanTaskStatisticsResponseBody) SetUserNum(v int64) *DescribeScanTaskStatisticsResponseBody {
	s.UserNum = &v
	return s
}

func (s *DescribeScanTaskStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
