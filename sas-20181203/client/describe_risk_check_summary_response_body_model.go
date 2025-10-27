// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRiskCheckSummaryResponseBody
	GetRequestId() *string
	SetRiskCheckSummary(v *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) *DescribeRiskCheckSummaryResponseBody
	GetRiskCheckSummary() *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary
}

type DescribeRiskCheckSummaryResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 291B49F9-1685-4005-9D34-606B6F78740F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The summary information about the check results of cloud service configurations.
	RiskCheckSummary *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary `json:"RiskCheckSummary,omitempty" xml:"RiskCheckSummary,omitempty" type:"Struct"`
}

func (s DescribeRiskCheckSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskCheckSummaryResponseBody) GetRiskCheckSummary() *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	return s.RiskCheckSummary
}

func (s *DescribeRiskCheckSummaryResponseBody) SetRequestId(v string) *DescribeRiskCheckSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBody) SetRiskCheckSummary(v *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) *DescribeRiskCheckSummaryResponseBody {
	s.RiskCheckSummary = v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBody) Validate() error {
	if s.RiskCheckSummary != nil {
		if err := s.RiskCheckSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRiskCheckSummaryResponseBodyRiskCheckSummary struct {
	// The number of affected assets.
	//
	// example:
	//
	// 0
	AffectedAssetCount *int32 `json:"AffectedAssetCount,omitempty" xml:"AffectedAssetCount,omitempty"`
	// The number of the check items that failed the check.
	//
	// example:
	//
	// 0
	DisabledRiskCount *int32 `json:"DisabledRiskCount,omitempty" xml:"DisabledRiskCount,omitempty"`
	// The number of the check items that passed the check.
	//
	// example:
	//
	// 3
	EnabledRiskCount *int32 `json:"EnabledRiskCount,omitempty" xml:"EnabledRiskCount,omitempty"`
	// An array that consists of the statistics for each type of check item.
	Groups []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The number of check items.
	//
	// example:
	//
	// 4
	ItemCount *int32 `json:"ItemCount,omitempty" xml:"ItemCount,omitempty"`
	// The number of risk items detected in the last check.
	//
	// example:
	//
	// 0
	PreviousCount *int32 `json:"PreviousCount,omitempty" xml:"PreviousCount,omitempty"`
	// The timestamp of the last check. Unit: milliseconds.
	//
	// example:
	//
	// 1545012926000
	PreviousTime *int64 `json:"PreviousTime,omitempty" xml:"PreviousTime,omitempty"`
	// The number of detected risk items.
	//
	// example:
	//
	// 1
	RiskCount *int32 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// An array that consists of the number of check items at each risk level.
	RiskLevelCount []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount `json:"RiskLevelCount,omitempty" xml:"RiskLevelCount,omitempty" type:"Repeated"`
	// The proportion of risk items to all check items.
	//
	// example:
	//
	// 0.25
	RiskRate *float32 `json:"RiskRate,omitempty" xml:"RiskRate,omitempty"`
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetAffectedAssetCount() *int32 {
	return s.AffectedAssetCount
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetDisabledRiskCount() *int32 {
	return s.DisabledRiskCount
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetEnabledRiskCount() *int32 {
	return s.EnabledRiskCount
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetGroups() []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	return s.Groups
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetItemCount() *int32 {
	return s.ItemCount
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetPreviousCount() *int32 {
	return s.PreviousCount
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetPreviousTime() *int64 {
	return s.PreviousTime
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetRiskLevelCount() []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount {
	return s.RiskLevelCount
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GetRiskRate() *float32 {
	return s.RiskRate
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetAffectedAssetCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.AffectedAssetCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetDisabledRiskCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.DisabledRiskCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetEnabledRiskCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.EnabledRiskCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetGroups(v []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.Groups = v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetItemCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.ItemCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetPreviousCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.PreviousCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetPreviousTime(v int64) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.PreviousTime = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetRiskCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.RiskCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetRiskLevelCount(v []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.RiskLevelCount = v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetRiskRate(v float32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.RiskRate = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RiskLevelCount != nil {
		for _, item := range s.RiskLevelCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups struct {
	// An array that consists of the statistics about check results.
	CountByStatus []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus `json:"CountByStatus,omitempty" xml:"CountByStatus,omitempty" type:"Repeated"`
	// The ID of the check item type.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The remaining time before the check is complete.
	//
	// example:
	//
	// 0
	RemainingTime *int32 `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty"`
	// The sequence number of the check item type in the **All Types*	- drop-down list in the Security Center console.
	//
	// example:
	//
	// 1
	Sort *int32 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The status of the check. Valid values:
	//
	// 	- **finish**: The check is finished.
	//
	// 	- **running**: The check is in progress.
	//
	// 	- **waiting**: The check is pending.
	//
	// 	- **notStart**: The check is not started.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the check item type.
	//
	// example:
	//
	// Identity authentication and permissions
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) GetCountByStatus() []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus {
	return s.CountByStatus
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) GetId() *int64 {
	return s.Id
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) GetRemainingTime() *int32 {
	return s.RemainingTime
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) GetSort() *int32 {
	return s.Sort
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) GetStatus() *string {
	return s.Status
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) GetTitle() *string {
	return s.Title
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetCountByStatus(v []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.CountByStatus = v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetId(v int64) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.Id = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetRemainingTime(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.RemainingTime = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetSort(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.Sort = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetStatus(v string) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.Status = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetTitle(v string) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.Title = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) Validate() error {
	if s.CountByStatus != nil {
		for _, item := range s.CountByStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus struct {
	// The number of detected risk items.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The status of the check item after the check is finished. Valid values:
	//
	// 	- **pass**: The check item passed the check, which indicates that the check item is normal.
	//
	// 	- **failed**: The check item failed the check, which indicates that risks are detected based on the check item.
	//
	// example:
	//
	// pass
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) GetCount() *int32 {
	return s.Count
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) SetCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus {
	s.Count = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) SetStatus(v string) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus {
	s.Status = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount struct {
	// The number of check items at the specified risk level.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The risk level of the check items. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) GetCount() *int32 {
	return s.Count
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) GetKey() *string {
	return s.Key
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) SetCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount {
	s.Count = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) SetKey(v string) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount {
	s.Key = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) Validate() error {
	return dara.Validate(s)
}
