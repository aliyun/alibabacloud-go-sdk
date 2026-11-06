// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeReviewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCodeReviewsResponseBodyPagingInfo) *ListCodeReviewsResponseBody
	GetPagingInfo() *ListCodeReviewsResponseBodyPagingInfo
	SetRequestId(v string) *ListCodeReviewsResponseBody
	GetRequestId() *string
}

type ListCodeReviewsResponseBody struct {
	// The pagination information.
	PagingInfo *ListCodeReviewsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// request-id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCodeReviewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCodeReviewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCodeReviewsResponseBody) GetPagingInfo() *ListCodeReviewsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCodeReviewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCodeReviewsResponseBody) SetPagingInfo(v *ListCodeReviewsResponseBodyPagingInfo) *ListCodeReviewsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCodeReviewsResponseBody) SetRequestId(v string) *ListCodeReviewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCodeReviewsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCodeReviewsResponseBodyPagingInfo struct {
	// The list of code reviews.
	CodeReviews []*ListCodeReviewsResponseBodyPagingInfoCodeReviews `json:"CodeReviews,omitempty" xml:"CodeReviews,omitempty" type:"Repeated"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCodeReviewsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCodeReviewsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCodeReviewsResponseBodyPagingInfo) GetCodeReviews() []*ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	return s.CodeReviews
}

func (s *ListCodeReviewsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCodeReviewsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCodeReviewsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCodeReviewsResponseBodyPagingInfo) SetCodeReviews(v []*ListCodeReviewsResponseBodyPagingInfoCodeReviews) *ListCodeReviewsResponseBodyPagingInfo {
	s.CodeReviews = v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfo) SetPageNumber(v int32) *ListCodeReviewsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfo) SetPageSize(v int32) *ListCodeReviewsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfo) SetTotalCount(v int32) *ListCodeReviewsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfo) Validate() error {
	if s.CodeReviews != nil {
		for _, item := range s.CodeReviews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCodeReviewsResponseBodyPagingInfoCodeReviews struct {
	// The node command of the root object.
	//
	// example:
	//
	// WORKFLOW
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1793145600000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the actual submitter.
	//
	// example:
	//
	// 209000000000000002
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The review description provided upon creation.
	//
	// example:
	//
	// Code review
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The account ID of the final approver. This field is empty if no approval decision has been made.
	//
	// example:
	//
	// 209000000000000001
	FinalReviewer *string `json:"FinalReviewer,omitempty" xml:"FinalReviewer,omitempty"`
	// The unique ID of the review.
	//
	// example:
	//
	// review-id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1793145600000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the root review object.
	//
	// example:
	//
	// 8600000000000000001
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the root review object.
	//
	// example:
	//
	// sales_workflow
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The root object type.
	//
	// example:
	//
	// WorkflowDefinition
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 23620
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The approval time. This field is empty if no approval decision has been made.
	//
	// example:
	//
	// 1793145660000
	ReviewTime *int64 `json:"ReviewTime,omitempty" xml:"ReviewTime,omitempty"`
	// The list of account IDs of the specified reviewers.
	Reviewers []*string `json:"Reviewers,omitempty" xml:"Reviewers,omitempty" type:"Repeated"`
	// The review status.
	//
	// example:
	//
	// Reviewing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCodeReviewsResponseBodyPagingInfoCodeReviews) String() string {
	return dara.Prettify(s)
}

func (s ListCodeReviewsResponseBodyPagingInfoCodeReviews) GoString() string {
	return s.String()
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetCommand() *string {
	return s.Command
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetCreator() *string {
	return s.Creator
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetDescription() *string {
	return s.Description
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetFinalReviewer() *string {
	return s.FinalReviewer
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetId() *string {
	return s.Id
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetReviewTime() *int64 {
	return s.ReviewTime
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetReviewers() []*string {
	return s.Reviewers
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) GetStatus() *string {
	return s.Status
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetCommand(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.Command = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetCreateTime(v int64) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.CreateTime = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetCreator(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.Creator = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetDescription(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.Description = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetFinalReviewer(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.FinalReviewer = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetId(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.Id = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetModifyTime(v int64) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.ModifyTime = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetObjectId(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.ObjectId = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetObjectName(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.ObjectName = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetObjectType(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.ObjectType = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetProjectId(v int64) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.ProjectId = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetRegionId(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.RegionId = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetReviewTime(v int64) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.ReviewTime = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetReviewers(v []*string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.Reviewers = v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) SetStatus(v string) *ListCodeReviewsResponseBodyPagingInfoCodeReviews {
	s.Status = &v
	return s
}

func (s *ListCodeReviewsResponseBodyPagingInfoCodeReviews) Validate() error {
	return dara.Validate(s)
}
