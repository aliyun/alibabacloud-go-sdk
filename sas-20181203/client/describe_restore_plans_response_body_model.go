// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestorePlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeRestorePlansResponseBodyPageInfo) *DescribeRestorePlansResponseBody
	GetPageInfo() *DescribeRestorePlansResponseBodyPageInfo
	SetRequestId(v string) *DescribeRestorePlansResponseBody
	GetRequestId() *string
	SetRestorePlans(v []*DescribeRestorePlansResponseBodyRestorePlans) *DescribeRestorePlansResponseBody
	GetRestorePlans() []*DescribeRestorePlansResponseBodyRestorePlans
}

type DescribeRestorePlansResponseBody struct {
	// The pagination information.
	PageInfo *DescribeRestorePlansResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the restoration tasks.
	RestorePlans []*DescribeRestorePlansResponseBodyRestorePlans `json:"RestorePlans,omitempty" xml:"RestorePlans,omitempty" type:"Repeated"`
}

func (s DescribeRestorePlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestorePlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestorePlansResponseBody) GetPageInfo() *DescribeRestorePlansResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeRestorePlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestorePlansResponseBody) GetRestorePlans() []*DescribeRestorePlansResponseBodyRestorePlans {
	return s.RestorePlans
}

func (s *DescribeRestorePlansResponseBody) SetPageInfo(v *DescribeRestorePlansResponseBodyPageInfo) *DescribeRestorePlansResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeRestorePlansResponseBody) SetRequestId(v string) *DescribeRestorePlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestorePlansResponseBody) SetRestorePlans(v []*DescribeRestorePlansResponseBodyRestorePlans) *DescribeRestorePlansResponseBody {
	s.RestorePlans = v
	return s
}

func (s *DescribeRestorePlansResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RestorePlans != nil {
		for _, item := range s.RestorePlans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRestorePlansResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 33
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRestorePlansResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestorePlansResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeRestorePlansResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeRestorePlansResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRestorePlansResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestorePlansResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRestorePlansResponseBodyPageInfo) SetCount(v int32) *DescribeRestorePlansResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeRestorePlansResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyPageInfo) SetPageSize(v int32) *DescribeRestorePlansResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyPageInfo) SetTotalCount(v int32) *DescribeRestorePlansResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeRestorePlansResponseBodyRestorePlans struct {
	// The timestamp when the restoration task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1655174753****
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// Bankup****
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The name of the server on which the database resides.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the anti-ransomware policy.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the anti-ransomware policy.
	//
	// example:
	//
	// KtDataBase
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The point in time to which data is restored.
	//
	// example:
	//
	// 165875100****
	RestorePoint *int64 `json:"RestorePoint,omitempty" xml:"RestorePoint,omitempty"`
	// The status of the restoration task. Valid values:
	//
	// 	- **init**: initializing
	//
	// 	- **created**: creating
	//
	// 	- **running**: running
	//
	// 	- **completed**: complete
	//
	// 	- **error**: failed
	//
	// 	- **restoring**: restoring
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the destination database.
	//
	// example:
	//
	// OABak
	TargetDatabaseName *string `json:"TargetDatabaseName,omitempty" xml:"TargetDatabaseName,omitempty"`
	// The ID of the destination server.
	//
	// example:
	//
	// i-2zehqflgbl9ep2he****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The name of the destination server.
	//
	// example:
	//
	// hbr-detection-hh
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	// The timestamp when the restoration task was last updated. Unit: milliseconds.
	//
	// example:
	//
	// 166849080****
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeRestorePlansResponseBodyRestorePlans) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestorePlansResponseBodyRestorePlans) GoString() string {
	return s.String()
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetRestorePoint() *int64 {
	return s.RestorePoint
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetStatus() *string {
	return s.Status
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetTargetDatabaseName() *string {
	return s.TargetDatabaseName
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetCreatedTime(v int64) *DescribeRestorePlansResponseBodyRestorePlans {
	s.CreatedTime = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetDatabaseName(v string) *DescribeRestorePlansResponseBodyRestorePlans {
	s.DatabaseName = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetInstanceName(v string) *DescribeRestorePlansResponseBodyRestorePlans {
	s.InstanceName = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetPolicyId(v int64) *DescribeRestorePlansResponseBodyRestorePlans {
	s.PolicyId = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetPolicyName(v string) *DescribeRestorePlansResponseBodyRestorePlans {
	s.PolicyName = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetRestorePoint(v int64) *DescribeRestorePlansResponseBodyRestorePlans {
	s.RestorePoint = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetStatus(v string) *DescribeRestorePlansResponseBodyRestorePlans {
	s.Status = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetTargetDatabaseName(v string) *DescribeRestorePlansResponseBodyRestorePlans {
	s.TargetDatabaseName = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetTargetInstanceId(v string) *DescribeRestorePlansResponseBodyRestorePlans {
	s.TargetInstanceId = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetTargetInstanceName(v string) *DescribeRestorePlansResponseBodyRestorePlans {
	s.TargetInstanceName = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) SetUpdatedTime(v int64) *DescribeRestorePlansResponseBodyRestorePlans {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeRestorePlansResponseBodyRestorePlans) Validate() error {
	return dara.Validate(s)
}
