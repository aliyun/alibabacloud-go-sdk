// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaReviewTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListQuotaReviewTasksResponseBody
	GetRequestId() *string
	SetResult(v []*ListQuotaReviewTasksResponseBodyResult) *ListQuotaReviewTasksResponseBody
	GetResult() []*ListQuotaReviewTasksResponseBodyResult
	SetTotalCount(v int32) *ListQuotaReviewTasksResponseBody
	GetTotalCount() *int32
}

type ListQuotaReviewTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// "3351A21F-705B-508C-9450-DA65A681547F"
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the tickets. For more information, see [QuotaReviewTask](https://help.aliyun.com/document_detail/173609.html).
	//
	// example:
	//
	// []
	Result []*ListQuotaReviewTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 500
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListQuotaReviewTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaReviewTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaReviewTasksResponseBody) GetResult() []*ListQuotaReviewTasksResponseBodyResult {
	return s.Result
}

func (s *ListQuotaReviewTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotaReviewTasksResponseBody) SetRequestId(v string) *ListQuotaReviewTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBody) SetResult(v []*ListQuotaReviewTasksResponseBodyResult) *ListQuotaReviewTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListQuotaReviewTasksResponseBody) SetTotalCount(v int32) *ListQuotaReviewTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQuotaReviewTasksResponseBodyResult struct {
	// The application ID.
	//
	// example:
	//
	// 120123456
	AppGroupId *int32 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// The application name.
	//
	// example:
	//
	// "td_test_os"
	AppGroupName *string `json:"appGroupName,omitempty" xml:"appGroupName,omitempty"`
	// The application type.
	//
	// example:
	//
	// "standard"
	AppGroupType *string `json:"appGroupType,omitempty" xml:"appGroupType,omitempty"`
	// Indicates whether the ticket is approved.
	//
	// example:
	//
	// true
	Approved *bool `json:"approved,omitempty" xml:"approved,omitempty"`
	// Indicates whether the application is available.
	//
	// example:
	//
	// true
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The time when the ticket was created.
	//
	// example:
	//
	// "2020-04-08T08:29:45+0000"
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the ticket was last updated.
	//
	// example:
	//
	// "2020-04-08T08:36:36+0000"
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The ticket ID.
	//
	// example:
	//
	// 142
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// The remarks.
	//
	// example:
	//
	// null
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// The computing resource quota that is applied for.
	//
	// example:
	//
	// 6000
	NewComputeResource *int32 `json:"newComputeResource,omitempty" xml:"newComputeResource,omitempty"`
	// The storage capacity quota that is applied for.
	//
	// example:
	//
	// 1100
	NewSocSize *int32 `json:"newSocSize,omitempty" xml:"newSocSize,omitempty"`
	// The application specifications that are applied for.
	//
	// example:
	//
	// "opensearch.private.common"
	NewSpec *string `json:"newSpec,omitempty" xml:"newSpec,omitempty"`
	// The original quota of computing resources.
	//
	// example:
	//
	// 500
	OldComputeResource *int32 `json:"oldComputeResource,omitempty" xml:"oldComputeResource,omitempty"`
	// The original quota of storage capacity.
	//
	// example:
	//
	// 900
	OldDocSize *int32 `json:"oldDocSize,omitempty" xml:"oldDocSize,omitempty"`
	// The original specifications of the application.
	//
	// example:
	//
	// "opensearch.private.common"
	OldSpec *string `json:"oldSpec,omitempty" xml:"oldSpec,omitempty"`
	// Indicates whether the ticket is pending.
	//
	// example:
	//
	// false
	Pending *bool `json:"pending,omitempty" xml:"pending,omitempty"`
}

func (s ListQuotaReviewTasksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaReviewTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetAppGroupId() *int32 {
	return s.AppGroupId
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetAppGroupName() *string {
	return s.AppGroupName
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetAppGroupType() *string {
	return s.AppGroupType
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetApproved() *bool {
	return s.Approved
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetAvailable() *bool {
	return s.Available
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetId() *int32 {
	return s.Id
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetMemo() *string {
	return s.Memo
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetNewComputeResource() *int32 {
	return s.NewComputeResource
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetNewSocSize() *int32 {
	return s.NewSocSize
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetNewSpec() *string {
	return s.NewSpec
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetOldComputeResource() *int32 {
	return s.OldComputeResource
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetOldDocSize() *int32 {
	return s.OldDocSize
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetOldSpec() *string {
	return s.OldSpec
}

func (s *ListQuotaReviewTasksResponseBodyResult) GetPending() *bool {
	return s.Pending
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupId(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupName(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupName = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupType(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupType = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetApproved(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Approved = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAvailable(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Available = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetGmtCreate(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetGmtModified(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetId(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetMemo(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewComputeResource(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.NewComputeResource = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewSocSize(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.NewSocSize = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewSpec(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.NewSpec = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldComputeResource(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.OldComputeResource = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldDocSize(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.OldDocSize = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldSpec(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.OldSpec = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetPending(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Pending = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
