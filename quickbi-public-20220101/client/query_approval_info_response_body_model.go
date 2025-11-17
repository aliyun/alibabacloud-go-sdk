// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApprovalInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryApprovalInfoResponseBody
	GetRequestId() *string
	SetResult(v *QueryApprovalInfoResponseBodyResult) *QueryApprovalInfoResponseBody
	GetResult() *QueryApprovalInfoResponseBodyResult
	SetSuccess(v bool) *QueryApprovalInfoResponseBody
	GetSuccess() *bool
}

type QueryApprovalInfoResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return the result of the interface execution.
	Result *QueryApprovalInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the API call was successful. Possible values are:
	//
	// - true: success
	//
	// - false: failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryApprovalInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryApprovalInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryApprovalInfoResponseBody) GetResult() *QueryApprovalInfoResponseBodyResult {
	return s.Result
}

func (s *QueryApprovalInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryApprovalInfoResponseBody) SetRequestId(v string) *QueryApprovalInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryApprovalInfoResponseBody) SetResult(v *QueryApprovalInfoResponseBodyResult) *QueryApprovalInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryApprovalInfoResponseBody) SetSuccess(v bool) *QueryApprovalInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryApprovalInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryApprovalInfoResponseBodyResult struct {
	// Array of approval flow information.
	Data []*QueryApprovalInfoResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of records requested per page.
	//
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The starting position of the current page.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryApprovalInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryApprovalInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoResponseBodyResult) GetData() []*QueryApprovalInfoResponseBodyResultData {
	return s.Data
}

func (s *QueryApprovalInfoResponseBodyResult) GetPage() *int32 {
	return s.Page
}

func (s *QueryApprovalInfoResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryApprovalInfoResponseBodyResult) GetStart() *int32 {
	return s.Start
}

func (s *QueryApprovalInfoResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *QueryApprovalInfoResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryApprovalInfoResponseBodyResult) SetData(v []*QueryApprovalInfoResponseBodyResultData) *QueryApprovalInfoResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetPage(v int32) *QueryApprovalInfoResponseBodyResult {
	s.Page = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetPageSize(v int32) *QueryApprovalInfoResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetStart(v int32) *QueryApprovalInfoResponseBodyResult {
	s.Start = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetTotal(v int32) *QueryApprovalInfoResponseBodyResult {
	s.Total = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetTotalPages(v int32) *QueryApprovalInfoResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryApprovalInfoResponseBodyResultData struct {
	// Applicant\\"s user ID, qbi user ID.
	//
	// example:
	//
	// 1359508
	ApplicantId *string `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	// Applicant\\"s nickname.
	//
	// example:
	//
	// Li Fei
	ApplicantName *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	// Application ID.
	//
	// example:
	//
	// 64813ef6da58e80eef8ed2f9
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Application reason.
	//
	// example:
	//
	// Development needs
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// Approver\\"s user ID, qbi user ID.
	//
	// example:
	//
	// sdasascasxasd
	ApproverId *string `json:"ApproverId,omitempty" xml:"ApproverId,omitempty"`
	// Approver\\"s nickname.
	//
	// example:
	//
	// data_fusion_002
	ApproverName *string `json:"ApproverName,omitempty" xml:"ApproverName,omitempty"`
	// Whether the resource has been deleted:
	//
	// - true: Deleted
	//
	// - false: Not deleted
	//
	// example:
	//
	// true
	DeleteFlag *bool `json:"DeleteFlag,omitempty" xml:"DeleteFlag,omitempty"`
	// Permission expiration date, timestamp.
	//
	// example:
	//
	// 1708568097135
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// Permission approval status:
	//
	// - 0: Under review, corresponding to 0 in the request parameters
	//
	// - 1: Approved, corresponding to 1 in the request parameters
	//
	// - 2: Rejected, corresponding to 1 in the request parameters
	//
	// example:
	//
	// 0
	FlagStatus *int32 `json:"FlagStatus,omitempty" xml:"FlagStatus,omitempty"`
	// Application creation time, timestamp.
	//
	// example:
	//
	// 1687315758
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Application modification time, timestamp.
	//
	// example:
	//
	// 1640595729000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Handling reason.
	//
	// example:
	//
	// Development needs
	HandleReason *string `json:"HandleReason,omitempty" xml:"HandleReason,omitempty"`
	// The ID of the resource for which permission is requested.
	//
	// example:
	//
	// acl-ct4t2e4u2x4ej1bzur
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource for which permission is requested (e.g., report name, space name...).
	//
	// example:
	//
	// Test Resources
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// DASHBOARD
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// Test Workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryApprovalInfoResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryApprovalInfoResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoResponseBodyResultData) GetApplicantId() *string {
	return s.ApplicantId
}

func (s *QueryApprovalInfoResponseBodyResultData) GetApplicantName() *string {
	return s.ApplicantName
}

func (s *QueryApprovalInfoResponseBodyResultData) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *QueryApprovalInfoResponseBodyResultData) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *QueryApprovalInfoResponseBodyResultData) GetApproverId() *string {
	return s.ApproverId
}

func (s *QueryApprovalInfoResponseBodyResultData) GetApproverName() *string {
	return s.ApproverName
}

func (s *QueryApprovalInfoResponseBodyResultData) GetDeleteFlag() *bool {
	return s.DeleteFlag
}

func (s *QueryApprovalInfoResponseBodyResultData) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *QueryApprovalInfoResponseBodyResultData) GetFlagStatus() *int32 {
	return s.FlagStatus
}

func (s *QueryApprovalInfoResponseBodyResultData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryApprovalInfoResponseBodyResultData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryApprovalInfoResponseBodyResultData) GetHandleReason() *string {
	return s.HandleReason
}

func (s *QueryApprovalInfoResponseBodyResultData) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryApprovalInfoResponseBodyResultData) GetResourceName() *string {
	return s.ResourceName
}

func (s *QueryApprovalInfoResponseBodyResultData) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryApprovalInfoResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApplicantId(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApplicantId = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApplicantName(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApplicantName = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApplicationId(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApplicationId = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApplyReason(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApplyReason = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApproverId(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApproverId = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApproverName(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApproverName = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetDeleteFlag(v bool) *QueryApprovalInfoResponseBodyResultData {
	s.DeleteFlag = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetExpireDate(v int64) *QueryApprovalInfoResponseBodyResultData {
	s.ExpireDate = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetFlagStatus(v int32) *QueryApprovalInfoResponseBodyResultData {
	s.FlagStatus = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetGmtCreate(v int64) *QueryApprovalInfoResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetGmtModified(v int64) *QueryApprovalInfoResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetHandleReason(v string) *QueryApprovalInfoResponseBodyResultData {
	s.HandleReason = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetResourceId(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ResourceId = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetResourceName(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ResourceName = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetResourceType(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ResourceType = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetWorkspaceName(v string) *QueryApprovalInfoResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
