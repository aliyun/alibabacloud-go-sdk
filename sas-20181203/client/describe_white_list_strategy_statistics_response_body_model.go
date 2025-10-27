// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeWhiteListStrategyStatisticsResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeWhiteListStrategyStatisticsResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWhiteListStrategyStatisticsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWhiteListStrategyStatisticsResponseBody
	GetRequestId() *string
	SetStrategies(v []*DescribeWhiteListStrategyStatisticsResponseBodyStrategies) *DescribeWhiteListStrategyStatisticsResponseBody
	GetStrategies() []*DescribeWhiteListStrategyStatisticsResponseBodyStrategies
	SetTotalCount(v int32) *DescribeWhiteListStrategyStatisticsResponseBody
	GetTotalCount() *int32
}

type DescribeWhiteListStrategyStatisticsResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 571B2642-BF51-5BDD-906B-D2340DB9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of the policies.
	Strategies []*DescribeWhiteListStrategyStatisticsResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWhiteListStrategyStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) GetStrategies() []*DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	return s.Strategies
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) SetCount(v int32) *DescribeWhiteListStrategyStatisticsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) SetCurrentPage(v int32) *DescribeWhiteListStrategyStatisticsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) SetPageSize(v int32) *DescribeWhiteListStrategyStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) SetRequestId(v string) *DescribeWhiteListStrategyStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) SetStrategies(v []*DescribeWhiteListStrategyStatisticsResponseBodyStrategies) *DescribeWhiteListStrategyStatisticsResponseBody {
	s.Strategies = v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) SetTotalCount(v int32) *DescribeWhiteListStrategyStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBody) Validate() error {
	if s.Strategies != nil {
		for _, item := range s.Strategies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWhiteListStrategyStatisticsResponseBodyStrategies struct {
	// The number of the servers on which the policy takes effect.
	//
	// example:
	//
	// 9
	AssetCount *int32 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The learning progress of the policy. Unit: percent (%).
	//
	// example:
	//
	// 80
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: learning
	//
	// 	- **2**: paused
	//
	// 	- **3**: learning completed
	//
	// 	- **4**: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 11906
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// win
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The intelligent learning duration. Unit: hour.
	//
	// example:
	//
	// 9
	StudyTime *int32 `json:"StudyTime,omitempty" xml:"StudyTime,omitempty"`
	// The number of suspicious processes.
	//
	// example:
	//
	// 1
	SuspiciousProcCount *int32 `json:"SuspiciousProcCount,omitempty" xml:"SuspiciousProcCount,omitempty"`
	// The number of trusted processes.
	//
	// example:
	//
	// 2
	TrustProcCount *int32 `json:"TrustProcCount,omitempty" xml:"TrustProcCount,omitempty"`
	// The number of malicious processes.
	//
	// example:
	//
	// 2
	VirusProcCount *int32 `json:"VirusProcCount,omitempty" xml:"VirusProcCount,omitempty"`
}

func (s DescribeWhiteListStrategyStatisticsResponseBodyStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetAssetCount() *int32 {
	return s.AssetCount
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetStudyTime() *int32 {
	return s.StudyTime
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetSuspiciousProcCount() *int32 {
	return s.SuspiciousProcCount
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetTrustProcCount() *int32 {
	return s.TrustProcCount
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) GetVirusProcCount() *int32 {
	return s.VirusProcCount
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetAssetCount(v int32) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.AssetCount = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetProgress(v int32) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.Progress = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetStatus(v int32) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.Status = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetStrategyId(v int64) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.StrategyId = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetStrategyName(v string) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.StrategyName = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetStudyTime(v int32) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.StudyTime = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetSuspiciousProcCount(v int32) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.SuspiciousProcCount = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetTrustProcCount(v int32) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.TrustProcCount = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) SetVirusProcCount(v int32) *DescribeWhiteListStrategyStatisticsResponseBodyStrategies {
	s.VirusProcCount = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponseBodyStrategies) Validate() error {
	return dara.Validate(s)
}
