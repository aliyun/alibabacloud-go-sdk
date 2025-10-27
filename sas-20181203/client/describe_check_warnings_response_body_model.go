// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckWarnings(v []*DescribeCheckWarningsResponseBodyCheckWarnings) *DescribeCheckWarningsResponseBody
	GetCheckWarnings() []*DescribeCheckWarningsResponseBodyCheckWarnings
	SetCount(v int32) *DescribeCheckWarningsResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeCheckWarningsResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeCheckWarningsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCheckWarningsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCheckWarningsResponseBody
	GetTotalCount() *int32
}

type DescribeCheckWarningsResponseBody struct {
	// The information about the check item.
	CheckWarnings []*DescribeCheckWarningsResponseBodyCheckWarnings `json:"CheckWarnings,omitempty" xml:"CheckWarnings,omitempty" type:"Repeated"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0DFCADBA-7065-42DA-AF17-6868B9C2A8CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCheckWarningsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningsResponseBody) GetCheckWarnings() []*DescribeCheckWarningsResponseBodyCheckWarnings {
	return s.CheckWarnings
}

func (s *DescribeCheckWarningsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCheckWarningsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCheckWarningsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCheckWarningsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckWarningsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCheckWarningsResponseBody) SetCheckWarnings(v []*DescribeCheckWarningsResponseBodyCheckWarnings) *DescribeCheckWarningsResponseBody {
	s.CheckWarnings = v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetCount(v int32) *DescribeCheckWarningsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetCurrentPage(v int32) *DescribeCheckWarningsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetPageSize(v int32) *DescribeCheckWarningsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetRequestId(v string) *DescribeCheckWarningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetTotalCount(v int32) *DescribeCheckWarningsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) Validate() error {
	if s.CheckWarnings != nil {
		for _, item := range s.CheckWarnings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCheckWarningsResponseBodyCheckWarnings struct {
	// The ID of the check item.
	//
	// example:
	//
	// 2546
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The ID of the alert that is generated for the baseline check result.
	//
	// example:
	//
	// 212251441
	CheckWarningId *int64 `json:"CheckWarningId,omitempty" xml:"CheckWarningId,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// 8de456b00ff0a2009ee8ef7fc59fd0457fa44f20b8282af3e79c2a0e2492****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// /svn-host
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The error message in the check result.
	//
	// example:
	//
	// ScriptKilledOfCpuHigh
	ExecErrorMessage *string `json:"ExecErrorMessage,omitempty" xml:"ExecErrorMessage,omitempty"`
	// Indicates whether fixing is supported. Valid values:
	//
	// 	- **0**: Fixing is not supported.
	//
	// 	- **1**: Fixing is supported.
	//
	// example:
	//
	// 0
	FixStatus *int32 `json:"FixStatus,omitempty" xml:"FixStatus,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// passwordExpire
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The timestamp of the latest processing of the check item risk of the machine. Unit: milliseconds.
	//
	// example:
	//
	// 1694692471000
	LastHandleTime *int64 `json:"LastHandleTime,omitempty" xml:"LastHandleTime,omitempty"`
	// The risk level of the risk item. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The description.
	//
	// example:
	//
	// ignore
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- **1**: failed.
	//
	// 	- **2**: verifying.
	//
	// 	- **3**: passed.
	//
	// 	- **5**: expired.
	//
	// 	- **6**: ignored.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the check item.
	//
	// example:
	//
	// hc.check.type.identity_auth
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the server on which the baseline check is performed.
	//
	// example:
	//
	// d42f938c-d962-48a0-90f9-****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeCheckWarningsResponseBodyCheckWarnings) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningsResponseBodyCheckWarnings) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetCheckWarningId() *int64 {
	return s.CheckWarningId
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetExecErrorMessage() *string {
	return s.ExecErrorMessage
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetFixStatus() *int32 {
	return s.FixStatus
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetItem() *string {
	return s.Item
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetLastHandleTime() *int64 {
	return s.LastHandleTime
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetLevel() *string {
	return s.Level
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetReason() *string {
	return s.Reason
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetType() *string {
	return s.Type
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetCheckId(v int64) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetCheckWarningId(v int64) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.CheckWarningId = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetContainerId(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.ContainerId = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetContainerName(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.ContainerName = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetExecErrorMessage(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.ExecErrorMessage = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetFixStatus(v int32) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.FixStatus = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetItem(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Item = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetLastHandleTime(v int64) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.LastHandleTime = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetLevel(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Level = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetReason(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Reason = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetStatus(v int32) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Status = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetType(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Type = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetUuid(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Uuid = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) Validate() error {
	return dara.Validate(s)
}
