// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeMajorVersionPrecheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody
	GetItems() []*DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems
	SetPageNumber(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody
	GetTotalRecordCount() *int32
}

type DescribeUpgradeMajorVersionPrecheckTaskResponseBody struct {
	// The information about the upgrade check reports.
	Items []*DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D1586777-41B5-5F9E-81E8-93DFDD379024
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the upgrade check report.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeUpgradeMajorVersionPrecheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeMajorVersionPrecheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) GetItems() []*DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	return s.Items
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) SetItems(v []*DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody {
	s.Items = v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) SetPageNumber(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) SetPageRecordCount(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) SetRequestId(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) SetTotalRecordCount(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems struct {
	// The time at which the upgrade check was performed.
	//
	// The value of this parameter is a timestamp that follows the UNIX time format. Unit: milliseconds.
	//
	// example:
	//
	// 1635143903000
	CheckTime *string `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// The content of the upgrade check report.
	//
	// example:
	//
	// [user_check_report]User check success\\n[pg_upgrade_internal.log]Performing...
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The expiration time of the upgrade check report.
	//
	// The value of this parameter is a timestamp that follows the UNIX time format. Unit: milliseconds.
	//
	// example:
	//
	// 1635748703000
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The minimum recommended disk capacity during the upgrade. Unit: GB.
	//
	// >  This parameter is returned only for RDS for PostgreSQL instances.
	//
	// example:
	//
	// 100
	RecommendDiskSize *int32 `json:"RecommendDiskSize,omitempty" xml:"RecommendDiskSize,omitempty"`
	// The minimum recommended memory size during the upgrade. Unit: GB.
	//
	// >  This parameter is returned only for RDS for PostgreSQL instances.
	//
	// example:
	//
	// 8
	RecommendLeastMemSize *int32 `json:"RecommendLeastMemSize,omitempty" xml:"RecommendLeastMemSize,omitempty"`
	// The recommended memory size during the upgrade. Unit: GB.
	//
	// If the memory size of an RDS instance is greater than or equal to the recommended memory size, the RDS instance is immediately upgraded to reduce the read-only time of the instance.
	//
	// >  This parameter is returned only for RDS for PostgreSQL instances.
	//
	// example:
	//
	// 32
	RecommendMemSize *int32 `json:"RecommendMemSize,omitempty" xml:"RecommendMemSize,omitempty"`
	// The result of the upgrade check.
	//
	// Valid values:
	//
	// 	- Success
	//
	// 	- Fail
	//
	// >  If the check result is **Fail**, you must check the value of the **Detail*	- parameter to obtain the information about the errors that occurred, resolve the errors, and then try again. For more information about how to resolve common errors, see [Introduction to the check report for a major engine version upgrade to an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/218391.html).
	//
	// example:
	//
	// Success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The original major engine version of the instance.
	//
	// example:
	//
	// 11.0
	SourceMajorVersion *string `json:"SourceMajorVersion,omitempty" xml:"SourceMajorVersion,omitempty"`
	// The new major engine version of the instance.
	//
	// example:
	//
	// 12.0
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	// The ID of the upgrade check task.
	//
	// example:
	//
	// 416980000
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UpgradeMode *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
}

func (s DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetCheckTime() *string {
	return s.CheckTime
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetDetail() *string {
	return s.Detail
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetRecommendDiskSize() *int32 {
	return s.RecommendDiskSize
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetRecommendLeastMemSize() *int32 {
	return s.RecommendLeastMemSize
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetRecommendMemSize() *int32 {
	return s.RecommendMemSize
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetResult() *string {
	return s.Result
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetSourceMajorVersion() *string {
	return s.SourceMajorVersion
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetTargetMajorVersion() *string {
	return s.TargetMajorVersion
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetCheckTime(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.CheckTime = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetDetail(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.Detail = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetEffectiveTime(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.EffectiveTime = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetRecommendDiskSize(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.RecommendDiskSize = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetRecommendLeastMemSize(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.RecommendLeastMemSize = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetRecommendMemSize(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.RecommendMemSize = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetResult(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.Result = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetSourceMajorVersion(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.SourceMajorVersion = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetTargetMajorVersion(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.TargetMajorVersion = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetTaskId(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) SetUpgradeMode(v string) *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems {
	s.UpgradeMode = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
