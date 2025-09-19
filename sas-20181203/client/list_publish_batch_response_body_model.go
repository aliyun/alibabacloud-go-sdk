// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatchList(v []*ListPublishBatchResponseBodyBatchList) *ListPublishBatchResponseBody
	GetBatchList() []*ListPublishBatchResponseBodyBatchList
	SetPageInfo(v *ListPublishBatchResponseBodyPageInfo) *ListPublishBatchResponseBody
	GetPageInfo() *ListPublishBatchResponseBodyPageInfo
	SetRequestId(v string) *ListPublishBatchResponseBody
	GetRequestId() *string
}

type ListPublishBatchResponseBody struct {
	// The information about the release batches.
	BatchList []*ListPublishBatchResponseBodyBatchList `json:"BatchList,omitempty" xml:"BatchList,omitempty" type:"Repeated"`
	// The page information.
	PageInfo *ListPublishBatchResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDCAE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublishBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishBatchResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishBatchResponseBody) GetBatchList() []*ListPublishBatchResponseBodyBatchList {
	return s.BatchList
}

func (s *ListPublishBatchResponseBody) GetPageInfo() *ListPublishBatchResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListPublishBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishBatchResponseBody) SetBatchList(v []*ListPublishBatchResponseBodyBatchList) *ListPublishBatchResponseBody {
	s.BatchList = v
	return s
}

func (s *ListPublishBatchResponseBody) SetPageInfo(v *ListPublishBatchResponseBodyPageInfo) *ListPublishBatchResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListPublishBatchResponseBody) SetRequestId(v string) *ListPublishBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishBatchResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPublishBatchResponseBodyBatchList struct {
	// The ID of the release batch.
	//
	// example:
	//
	// 1371
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The interval between two release batches. Unit: hours.
	//
	// example:
	//
	// 12
	BatchInterval *int32 `json:"BatchInterval,omitempty" xml:"BatchInterval,omitempty"`
	// The name of the release batch.
	//
	// example:
	//
	// test
	BatchName *string `json:"BatchName,omitempty" xml:"BatchName,omitempty"`
	// The current batch number during a batch release.
	//
	// example:
	//
	// 2147483647
	BatchNo *int32 `json:"BatchNo,omitempty" xml:"BatchNo,omitempty"`
	// The progress of the release batch. This parameter indicates the number of servers that are upgraded in the release batch.
	//
	// example:
	//
	// 12
	BatchProcess *int32 `json:"BatchProcess,omitempty" xml:"BatchProcess,omitempty"`
	// The total number of batches.
	//
	// example:
	//
	// 3
	BatchTotal *int32 `json:"BatchTotal,omitempty" xml:"BatchTotal,omitempty"`
	// The asset selection dimension. Valid values:
	//
	// 	- **0**: instance.
	//
	// 	- **1**: machine group.
	//
	// 	- **2**: Virtual Private Cloud (VPC) ID.
	//
	// example:
	//
	// 0
	OperationBase *int32 `json:"OperationBase,omitempty" xml:"OperationBase,omitempty"`
	// The publish status of the Security Center agent. Valid values:
	//
	// 	- **0**: not started.
	//
	// 	- **1**: publishing.
	//
	// 	- **2**: published.
	//
	// 	- **3**: publish suspended.
	//
	// 	- **4**: forcibly upgrading.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The destination version of the Security Center agent.
	//
	// example:
	//
	// 0.0.9
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPublishBatchResponseBodyBatchList) String() string {
	return dara.Prettify(s)
}

func (s ListPublishBatchResponseBodyBatchList) GoString() string {
	return s.String()
}

func (s *ListPublishBatchResponseBodyBatchList) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ListPublishBatchResponseBodyBatchList) GetBatchInterval() *int32 {
	return s.BatchInterval
}

func (s *ListPublishBatchResponseBodyBatchList) GetBatchName() *string {
	return s.BatchName
}

func (s *ListPublishBatchResponseBodyBatchList) GetBatchNo() *int32 {
	return s.BatchNo
}

func (s *ListPublishBatchResponseBodyBatchList) GetBatchProcess() *int32 {
	return s.BatchProcess
}

func (s *ListPublishBatchResponseBodyBatchList) GetBatchTotal() *int32 {
	return s.BatchTotal
}

func (s *ListPublishBatchResponseBodyBatchList) GetOperationBase() *int32 {
	return s.OperationBase
}

func (s *ListPublishBatchResponseBodyBatchList) GetStatus() *int32 {
	return s.Status
}

func (s *ListPublishBatchResponseBodyBatchList) GetVersion() *string {
	return s.Version
}

func (s *ListPublishBatchResponseBodyBatchList) SetBatchId(v int64) *ListPublishBatchResponseBodyBatchList {
	s.BatchId = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) SetBatchInterval(v int32) *ListPublishBatchResponseBodyBatchList {
	s.BatchInterval = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) SetBatchName(v string) *ListPublishBatchResponseBodyBatchList {
	s.BatchName = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) SetBatchNo(v int32) *ListPublishBatchResponseBodyBatchList {
	s.BatchNo = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) SetBatchProcess(v int32) *ListPublishBatchResponseBodyBatchList {
	s.BatchProcess = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) SetBatchTotal(v int32) *ListPublishBatchResponseBodyBatchList {
	s.BatchTotal = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) SetOperationBase(v int32) *ListPublishBatchResponseBodyBatchList {
	s.OperationBase = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) SetStatus(v int32) *ListPublishBatchResponseBodyBatchList {
	s.Status = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) SetVersion(v string) *ListPublishBatchResponseBodyBatchList {
	s.Version = &v
	return s
}

func (s *ListPublishBatchResponseBodyBatchList) Validate() error {
	return dara.Validate(s)
}

type ListPublishBatchResponseBodyPageInfo struct {
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
	// The total number of entries returned.
	//
	// example:
	//
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublishBatchResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPublishBatchResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListPublishBatchResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListPublishBatchResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishBatchResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublishBatchResponseBodyPageInfo) SetCurrentPage(v int32) *ListPublishBatchResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListPublishBatchResponseBodyPageInfo) SetPageSize(v int32) *ListPublishBatchResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListPublishBatchResponseBodyPageInfo) SetTotalCount(v int32) *ListPublishBatchResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListPublishBatchResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
