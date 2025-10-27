// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOverallItemStatistic(v *GetCheckSummaryResponseBodyOverallItemStatistic) *GetCheckSummaryResponseBody
	GetOverallItemStatistic() *GetCheckSummaryResponseBodyOverallItemStatistic
	SetOverallStatistic(v *GetCheckSummaryResponseBodyOverallStatistic) *GetCheckSummaryResponseBody
	GetOverallStatistic() *GetCheckSummaryResponseBodyOverallStatistic
	SetRequestId(v string) *GetCheckSummaryResponseBody
	GetRequestId() *string
	SetSummarys(v []*GetCheckSummaryResponseBodySummarys) *GetCheckSummaryResponseBody
	GetSummarys() []*GetCheckSummaryResponseBodySummarys
}

type GetCheckSummaryResponseBody struct {
	// The statistics about the number of check items.
	OverallItemStatistic *GetCheckSummaryResponseBodyOverallItemStatistic `json:"OverallItemStatistic,omitempty" xml:"OverallItemStatistic,omitempty" type:"Struct"`
	// The overall risk statistics.
	OverallStatistic *GetCheckSummaryResponseBodyOverallStatistic `json:"OverallStatistic,omitempty" xml:"OverallStatistic,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 843E4805-****-7EE12FA8DBFD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The summary information about the configuration checks on cloud services.
	Summarys []*GetCheckSummaryResponseBodySummarys `json:"Summarys,omitempty" xml:"Summarys,omitempty" type:"Repeated"`
}

func (s GetCheckSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryResponseBody) GetOverallItemStatistic() *GetCheckSummaryResponseBodyOverallItemStatistic {
	return s.OverallItemStatistic
}

func (s *GetCheckSummaryResponseBody) GetOverallStatistic() *GetCheckSummaryResponseBodyOverallStatistic {
	return s.OverallStatistic
}

func (s *GetCheckSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckSummaryResponseBody) GetSummarys() []*GetCheckSummaryResponseBodySummarys {
	return s.Summarys
}

func (s *GetCheckSummaryResponseBody) SetOverallItemStatistic(v *GetCheckSummaryResponseBodyOverallItemStatistic) *GetCheckSummaryResponseBody {
	s.OverallItemStatistic = v
	return s
}

func (s *GetCheckSummaryResponseBody) SetOverallStatistic(v *GetCheckSummaryResponseBodyOverallStatistic) *GetCheckSummaryResponseBody {
	s.OverallStatistic = v
	return s
}

func (s *GetCheckSummaryResponseBody) SetRequestId(v string) *GetCheckSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckSummaryResponseBody) SetSummarys(v []*GetCheckSummaryResponseBodySummarys) *GetCheckSummaryResponseBody {
	s.Summarys = v
	return s
}

func (s *GetCheckSummaryResponseBody) Validate() error {
	if s.OverallItemStatistic != nil {
		if err := s.OverallItemStatistic.Validate(); err != nil {
			return err
		}
	}
	if s.OverallStatistic != nil {
		if err := s.OverallStatistic.Validate(); err != nil {
			return err
		}
	}
	if s.Summarys != nil {
		for _, item := range s.Summarys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCheckSummaryResponseBodyOverallItemStatistic struct {
	// The number of check items supported by the system.
	//
	// example:
	//
	// 620
	ReleaseCount *int32 `json:"ReleaseCount,omitempty" xml:"ReleaseCount,omitempty"`
	// The number of check items available to you.
	//
	// example:
	//
	// 25
	ResultCount *int32 `json:"ResultCount,omitempty" xml:"ResultCount,omitempty"`
}

func (s GetCheckSummaryResponseBodyOverallItemStatistic) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryResponseBodyOverallItemStatistic) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryResponseBodyOverallItemStatistic) GetReleaseCount() *int32 {
	return s.ReleaseCount
}

func (s *GetCheckSummaryResponseBodyOverallItemStatistic) GetResultCount() *int32 {
	return s.ResultCount
}

func (s *GetCheckSummaryResponseBodyOverallItemStatistic) SetReleaseCount(v int32) *GetCheckSummaryResponseBodyOverallItemStatistic {
	s.ReleaseCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallItemStatistic) SetResultCount(v int32) *GetCheckSummaryResponseBodyOverallItemStatistic {
	s.ResultCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallItemStatistic) Validate() error {
	return dara.Validate(s)
}

type GetCheckSummaryResponseBodyOverallStatistic struct {
	// The number of unchecked check items.
	//
	// example:
	//
	// 3
	NotCheckCount *int32 `json:"NotCheckCount,omitempty" xml:"NotCheckCount,omitempty"`
	// The number of unchecked high-risk check items.
	//
	// example:
	//
	// 1
	NotCheckHighCount *int32 `json:"NotCheckHighCount,omitempty" xml:"NotCheckHighCount,omitempty"`
	// The number of unchecked low-risk check items.
	//
	// example:
	//
	// 1
	NotCheckLowCount *int32 `json:"NotCheckLowCount,omitempty" xml:"NotCheckLowCount,omitempty"`
	// The number of unchecked medium-risk check items.
	//
	// example:
	//
	// 1
	NotCheckMediumCount *int32 `json:"NotCheckMediumCount,omitempty" xml:"NotCheckMediumCount,omitempty"`
	// The number of check items that failed to pass the check.
	//
	// example:
	//
	// 3
	NotPassCount *int32 `json:"NotPassCount,omitempty" xml:"NotPassCount,omitempty"`
	// The number of high-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassHighCount *int32 `json:"NotPassHighCount,omitempty" xml:"NotPassHighCount,omitempty"`
	// The number of low-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassLowCount *int32 `json:"NotPassLowCount,omitempty" xml:"NotPassLowCount,omitempty"`
	// The number of medium-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassMediumCount *int32 `json:"NotPassMediumCount,omitempty" xml:"NotPassMediumCount,omitempty"`
	// The number of check items that pass the check.
	//
	// example:
	//
	// 3
	PassCount *int32 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	// The number of high-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassHighCount *int32 `json:"PassHighCount,omitempty" xml:"PassHighCount,omitempty"`
	// The number of low-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassLowCount *int32 `json:"PassLowCount,omitempty" xml:"PassLowCount,omitempty"`
	// The number of medium-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassMediumCount *int32 `json:"PassMediumCount,omitempty" xml:"PassMediumCount,omitempty"`
}

func (s GetCheckSummaryResponseBodyOverallStatistic) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryResponseBodyOverallStatistic) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetNotCheckCount() *int32 {
	return s.NotCheckCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetNotCheckHighCount() *int32 {
	return s.NotCheckHighCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetNotCheckLowCount() *int32 {
	return s.NotCheckLowCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetNotCheckMediumCount() *int32 {
	return s.NotCheckMediumCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetNotPassCount() *int32 {
	return s.NotPassCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetNotPassHighCount() *int32 {
	return s.NotPassHighCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetNotPassLowCount() *int32 {
	return s.NotPassLowCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetNotPassMediumCount() *int32 {
	return s.NotPassMediumCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetPassCount() *int32 {
	return s.PassCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetPassHighCount() *int32 {
	return s.PassHighCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetPassLowCount() *int32 {
	return s.PassLowCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) GetPassMediumCount() *int32 {
	return s.PassMediumCount
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetNotCheckCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.NotCheckCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetNotCheckHighCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.NotCheckHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetNotCheckLowCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.NotCheckLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetNotCheckMediumCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.NotCheckMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetNotPassCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.NotPassCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetNotPassHighCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.NotPassHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetNotPassLowCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.NotPassLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetNotPassMediumCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.NotPassMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetPassCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.PassCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetPassHighCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.PassHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetPassLowCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.PassLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) SetPassMediumCount(v int32) *GetCheckSummaryResponseBodyOverallStatistic {
	s.PassMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodyOverallStatistic) Validate() error {
	return dara.Validate(s)
}

type GetCheckSummaryResponseBodySummarys struct {
	// The number of detected risk items.
	//
	// example:
	//
	// 5
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The number of check items that pass the check.
	//
	// example:
	//
	// 10
	PassCount *int64 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	// The information about the check items.
	Standards []*GetCheckSummaryResponseBodySummarysStandards `json:"Standards,omitempty" xml:"Standards,omitempty" type:"Repeated"`
	// The type of the check item. Valid values:
	//
	// 	- **COMPLIANCE**
	//
	// 	- **RISK**
	//
	// 	- **IDENTITY_PERMISSION**
	//
	// example:
	//
	// IDENTITY_PERMISSION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The risk statistics by type.
	TypeStatistic *GetCheckSummaryResponseBodySummarysTypeStatistic `json:"TypeStatistic,omitempty" xml:"TypeStatistic,omitempty" type:"Struct"`
}

func (s GetCheckSummaryResponseBodySummarys) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryResponseBodySummarys) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryResponseBodySummarys) GetFailCount() *int32 {
	return s.FailCount
}

func (s *GetCheckSummaryResponseBodySummarys) GetPassCount() *int64 {
	return s.PassCount
}

func (s *GetCheckSummaryResponseBodySummarys) GetStandards() []*GetCheckSummaryResponseBodySummarysStandards {
	return s.Standards
}

func (s *GetCheckSummaryResponseBodySummarys) GetType() *string {
	return s.Type
}

func (s *GetCheckSummaryResponseBodySummarys) GetTypeStatistic() *GetCheckSummaryResponseBodySummarysTypeStatistic {
	return s.TypeStatistic
}

func (s *GetCheckSummaryResponseBodySummarys) SetFailCount(v int32) *GetCheckSummaryResponseBodySummarys {
	s.FailCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarys) SetPassCount(v int64) *GetCheckSummaryResponseBodySummarys {
	s.PassCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarys) SetStandards(v []*GetCheckSummaryResponseBodySummarysStandards) *GetCheckSummaryResponseBodySummarys {
	s.Standards = v
	return s
}

func (s *GetCheckSummaryResponseBodySummarys) SetType(v string) *GetCheckSummaryResponseBodySummarys {
	s.Type = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarys) SetTypeStatistic(v *GetCheckSummaryResponseBodySummarysTypeStatistic) *GetCheckSummaryResponseBodySummarys {
	s.TypeStatistic = v
	return s
}

func (s *GetCheckSummaryResponseBodySummarys) Validate() error {
	if s.Standards != nil {
		for _, item := range s.Standards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TypeStatistic != nil {
		if err := s.TypeStatistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCheckSummaryResponseBodySummarysStandards struct {
	// The number of check items that failed to pass the check.
	//
	// example:
	//
	// 1
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The ID of the check item.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of check items that pass the check.
	//
	// example:
	//
	// 1
	PassCount *int32 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	// The number of **high-risk*	- items.
	//
	// example:
	//
	// 1
	RiskLevelHighCount *int32 `json:"RiskLevelHighCount,omitempty" xml:"RiskLevelHighCount,omitempty"`
	// The number of **low-risk*	- items.
	//
	// example:
	//
	// 1
	RiskLevelLowCount *int32 `json:"RiskLevelLowCount,omitempty" xml:"RiskLevelLowCount,omitempty"`
	// The number of **medium-risk*	- items.
	//
	// example:
	//
	// 1
	RiskLevelMediumCount *int32 `json:"RiskLevelMediumCount,omitempty" xml:"RiskLevelMediumCount,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// Identity and permission management
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The standard statistics of the check items.
	StandardStatistic *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic `json:"StandardStatistic,omitempty" xml:"StandardStatistic,omitempty" type:"Struct"`
}

func (s GetCheckSummaryResponseBodySummarysStandards) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryResponseBodySummarysStandards) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryResponseBodySummarysStandards) GetFailCount() *int32 {
	return s.FailCount
}

func (s *GetCheckSummaryResponseBodySummarysStandards) GetId() *int64 {
	return s.Id
}

func (s *GetCheckSummaryResponseBodySummarysStandards) GetPassCount() *int32 {
	return s.PassCount
}

func (s *GetCheckSummaryResponseBodySummarysStandards) GetRiskLevelHighCount() *int32 {
	return s.RiskLevelHighCount
}

func (s *GetCheckSummaryResponseBodySummarysStandards) GetRiskLevelLowCount() *int32 {
	return s.RiskLevelLowCount
}

func (s *GetCheckSummaryResponseBodySummarysStandards) GetRiskLevelMediumCount() *int32 {
	return s.RiskLevelMediumCount
}

func (s *GetCheckSummaryResponseBodySummarysStandards) GetShowName() *string {
	return s.ShowName
}

func (s *GetCheckSummaryResponseBodySummarysStandards) GetStandardStatistic() *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	return s.StandardStatistic
}

func (s *GetCheckSummaryResponseBodySummarysStandards) SetFailCount(v int32) *GetCheckSummaryResponseBodySummarysStandards {
	s.FailCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandards) SetId(v int64) *GetCheckSummaryResponseBodySummarysStandards {
	s.Id = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandards) SetPassCount(v int32) *GetCheckSummaryResponseBodySummarysStandards {
	s.PassCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandards) SetRiskLevelHighCount(v int32) *GetCheckSummaryResponseBodySummarysStandards {
	s.RiskLevelHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandards) SetRiskLevelLowCount(v int32) *GetCheckSummaryResponseBodySummarysStandards {
	s.RiskLevelLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandards) SetRiskLevelMediumCount(v int32) *GetCheckSummaryResponseBodySummarysStandards {
	s.RiskLevelMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandards) SetShowName(v string) *GetCheckSummaryResponseBodySummarysStandards {
	s.ShowName = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandards) SetStandardStatistic(v *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) *GetCheckSummaryResponseBodySummarysStandards {
	s.StandardStatistic = v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandards) Validate() error {
	if s.StandardStatistic != nil {
		if err := s.StandardStatistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCheckSummaryResponseBodySummarysStandardsStandardStatistic struct {
	// The number of unchecked check items.
	//
	// example:
	//
	// 3
	NotCheckCount *int32 `json:"NotCheckCount,omitempty" xml:"NotCheckCount,omitempty"`
	// The number of unchecked high-risk check items.
	//
	// example:
	//
	// 1
	NotCheckHighCount *int32 `json:"NotCheckHighCount,omitempty" xml:"NotCheckHighCount,omitempty"`
	// The number of unchecked low-risk check items.
	//
	// example:
	//
	// 1
	NotCheckLowCount *int32 `json:"NotCheckLowCount,omitempty" xml:"NotCheckLowCount,omitempty"`
	// The number of unchecked medium-risk check items.
	//
	// example:
	//
	// 1
	NotCheckMediumCount *int32 `json:"NotCheckMediumCount,omitempty" xml:"NotCheckMediumCount,omitempty"`
	// The number of check items that failed to pass the check.
	//
	// example:
	//
	// 3
	NotPassCount *int32 `json:"NotPassCount,omitempty" xml:"NotPassCount,omitempty"`
	// The number of high-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassHighCount *int32 `json:"NotPassHighCount,omitempty" xml:"NotPassHighCount,omitempty"`
	// The number of low-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassLowCount *int32 `json:"NotPassLowCount,omitempty" xml:"NotPassLowCount,omitempty"`
	// The number of medium-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassMediumCount *int32 `json:"NotPassMediumCount,omitempty" xml:"NotPassMediumCount,omitempty"`
	// The number of check items that pass the check.
	//
	// example:
	//
	// 3
	PassCount *int32 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	// The number of high-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassHighCount *int32 `json:"PassHighCount,omitempty" xml:"PassHighCount,omitempty"`
	// The number of low-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassLowCount *int32 `json:"PassLowCount,omitempty" xml:"PassLowCount,omitempty"`
	// The number of medium-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassMediumCount *int32 `json:"PassMediumCount,omitempty" xml:"PassMediumCount,omitempty"`
}

func (s GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetNotCheckCount() *int32 {
	return s.NotCheckCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetNotCheckHighCount() *int32 {
	return s.NotCheckHighCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetNotCheckLowCount() *int32 {
	return s.NotCheckLowCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetNotCheckMediumCount() *int32 {
	return s.NotCheckMediumCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetNotPassCount() *int32 {
	return s.NotPassCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetNotPassHighCount() *int32 {
	return s.NotPassHighCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetNotPassLowCount() *int32 {
	return s.NotPassLowCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetNotPassMediumCount() *int32 {
	return s.NotPassMediumCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetPassCount() *int32 {
	return s.PassCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetPassHighCount() *int32 {
	return s.PassHighCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetPassLowCount() *int32 {
	return s.PassLowCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) GetPassMediumCount() *int32 {
	return s.PassMediumCount
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetNotCheckCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.NotCheckCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetNotCheckHighCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.NotCheckHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetNotCheckLowCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.NotCheckLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetNotCheckMediumCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.NotCheckMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetNotPassCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.NotPassCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetNotPassHighCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.NotPassHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetNotPassLowCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.NotPassLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetNotPassMediumCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.NotPassMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetPassCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.PassCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetPassHighCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.PassHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetPassLowCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.PassLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) SetPassMediumCount(v int32) *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic {
	s.PassMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysStandardsStandardStatistic) Validate() error {
	return dara.Validate(s)
}

type GetCheckSummaryResponseBodySummarysTypeStatistic struct {
	// The number of unchecked check items.
	//
	// example:
	//
	// 3
	NotCheckCount *int32 `json:"NotCheckCount,omitempty" xml:"NotCheckCount,omitempty"`
	// The number of unchecked high-risk check items.
	//
	// example:
	//
	// 1
	NotCheckHighCount *int32 `json:"NotCheckHighCount,omitempty" xml:"NotCheckHighCount,omitempty"`
	// The number of unchecked low-risk check items.
	//
	// example:
	//
	// 1
	NotCheckLowCount *int32 `json:"NotCheckLowCount,omitempty" xml:"NotCheckLowCount,omitempty"`
	// The number of unchecked medium-risk check items.
	//
	// example:
	//
	// 1
	NotCheckMediumCount *int32 `json:"NotCheckMediumCount,omitempty" xml:"NotCheckMediumCount,omitempty"`
	// The number of check items that failed to pass the check.
	//
	// example:
	//
	// 3
	NotPassCount *int32 `json:"NotPassCount,omitempty" xml:"NotPassCount,omitempty"`
	// The number of high-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassHighCount *int32 `json:"NotPassHighCount,omitempty" xml:"NotPassHighCount,omitempty"`
	// The number of low-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassLowCount *int32 `json:"NotPassLowCount,omitempty" xml:"NotPassLowCount,omitempty"`
	// The number of medium-risk check items that failed to pass the check.
	//
	// example:
	//
	// 1
	NotPassMediumCount *int32 `json:"NotPassMediumCount,omitempty" xml:"NotPassMediumCount,omitempty"`
	// The number of check items that pass the check.
	//
	// example:
	//
	// 3
	PassCount *int32 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	// The number of high-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassHighCount *int32 `json:"PassHighCount,omitempty" xml:"PassHighCount,omitempty"`
	// The number of low-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassLowCount *int32 `json:"PassLowCount,omitempty" xml:"PassLowCount,omitempty"`
	// The number of medium-risk check items that pass the check.
	//
	// example:
	//
	// 1
	PassMediumCount *int32 `json:"PassMediumCount,omitempty" xml:"PassMediumCount,omitempty"`
}

func (s GetCheckSummaryResponseBodySummarysTypeStatistic) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryResponseBodySummarysTypeStatistic) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetNotCheckCount() *int32 {
	return s.NotCheckCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetNotCheckHighCount() *int32 {
	return s.NotCheckHighCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetNotCheckLowCount() *int32 {
	return s.NotCheckLowCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetNotCheckMediumCount() *int32 {
	return s.NotCheckMediumCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetNotPassCount() *int32 {
	return s.NotPassCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetNotPassHighCount() *int32 {
	return s.NotPassHighCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetNotPassLowCount() *int32 {
	return s.NotPassLowCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetNotPassMediumCount() *int32 {
	return s.NotPassMediumCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetPassCount() *int32 {
	return s.PassCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetPassHighCount() *int32 {
	return s.PassHighCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetPassLowCount() *int32 {
	return s.PassLowCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) GetPassMediumCount() *int32 {
	return s.PassMediumCount
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetNotCheckCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.NotCheckCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetNotCheckHighCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.NotCheckHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetNotCheckLowCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.NotCheckLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetNotCheckMediumCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.NotCheckMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetNotPassCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.NotPassCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetNotPassHighCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.NotPassHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetNotPassLowCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.NotPassLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetNotPassMediumCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.NotPassMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetPassCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.PassCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetPassHighCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.PassHighCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetPassLowCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.PassLowCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) SetPassMediumCount(v int32) *GetCheckSummaryResponseBodySummarysTypeStatistic {
	s.PassMediumCount = &v
	return s
}

func (s *GetCheckSummaryResponseBodySummarysTypeStatistic) Validate() error {
	return dara.Validate(s)
}
