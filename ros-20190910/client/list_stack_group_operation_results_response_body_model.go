// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupOperationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListStackGroupOperationResultsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStackGroupOperationResultsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListStackGroupOperationResultsResponseBody
	GetRequestId() *string
	SetStackGroupOperationResults(v []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) *ListStackGroupOperationResultsResponseBody
	GetStackGroupOperationResults() []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults
	SetTotalCount(v int32) *ListStackGroupOperationResultsResponseBody
	GetTotalCount() *int32
}

type ListStackGroupOperationResultsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the results of the operation.
	StackGroupOperationResults []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults `json:"StackGroupOperationResults,omitempty" xml:"StackGroupOperationResults,omitempty" type:"Repeated"`
	// The total number of results.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackGroupOperationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupOperationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStackGroupOperationResultsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStackGroupOperationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackGroupOperationResultsResponseBody) GetStackGroupOperationResults() []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	return s.StackGroupOperationResults
}

func (s *ListStackGroupOperationResultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListStackGroupOperationResultsResponseBody) SetPageNumber(v int32) *ListStackGroupOperationResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetPageSize(v int32) *ListStackGroupOperationResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetRequestId(v string) *ListStackGroupOperationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetStackGroupOperationResults(v []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) *ListStackGroupOperationResultsResponseBody {
	s.StackGroupOperationResults = v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetTotalCount(v int32) *ListStackGroupOperationResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListStackGroupOperationResultsResponseBodyStackGroupOperationResults struct {
	// The ID of the account to which the stack instance belongs.
	//
	// 	- If the stack group has self-managed permissions, the stack instance belongs to an Alibaba Cloud account.
	//
	// 	- If the stack group has service-managed permissions, the stack instance belongs to a member account in the resource directory.
	//
	// >  For more information about the account, see [Overview](https://help.aliyun.com/document_detail/154578.html).
	//
	// example:
	//
	// 175458090349****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The folder ID of the resource directory.
	//
	// >  This parameter is returned only when the stack group is granted service-managed permissions.
	//
	// example:
	//
	// "fd-4PvlVLOL8v"
	RdFolderId *string `json:"RdFolderId,omitempty" xml:"RdFolderId,omitempty"`
	// The region ID of the stack instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the operation.
	//
	// Valid values:
	//
	// 	- RUNNING: The operation is being performed.
	//
	// 	- SUCCEEDED: The operation succeeded.
	//
	// 	- FAILED: The operation failed.
	//
	// 	- STOPPING: The operation is being stopped.
	//
	// 	- STOPPED: The operation is stopped.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the operation is in a specific state.
	//
	// >  This parameter is returned only when stack instances are in the OUTDATED state.
	//
	// example:
	//
	// User initiated operation
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) GetAccountId() *string {
	return s.AccountId
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) GetRdFolderId() *string {
	return s.RdFolderId
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) GetStatus() *string {
	return s.Status
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetAccountId(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.AccountId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetRdFolderId(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.RdFolderId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetRegionId(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetStatus(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.Status = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetStatusReason(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.StatusReason = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) Validate() error {
	return dara.Validate(s)
}
