// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeReview(v *GetCodeReviewResponseBodyCodeReview) *GetCodeReviewResponseBody
	GetCodeReview() *GetCodeReviewResponseBodyCodeReview
	SetRequestId(v string) *GetCodeReviewResponseBody
	GetRequestId() *string
}

type GetCodeReviewResponseBody struct {
	// The code review metadata.
	CodeReview *GetCodeReviewResponseBodyCodeReview `json:"CodeReview,omitempty" xml:"CodeReview,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// request-id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCodeReviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCodeReviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeReviewResponseBody) GetCodeReview() *GetCodeReviewResponseBodyCodeReview {
	return s.CodeReview
}

func (s *GetCodeReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCodeReviewResponseBody) SetCodeReview(v *GetCodeReviewResponseBodyCodeReview) *GetCodeReviewResponseBody {
	s.CodeReview = v
	return s
}

func (s *GetCodeReviewResponseBody) SetRequestId(v string) *GetCodeReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeReviewResponseBody) Validate() error {
	if s.CodeReview != nil {
		if err := s.CodeReview.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCodeReviewResponseBodyCodeReview struct {
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
	// The review description provided at creation.
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
	// The unique identifier of the review.
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
	// The object type of the root object.
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
	// The account IDs of the designated reviewers.
	Reviewers []*string `json:"Reviewers,omitempty" xml:"Reviewers,omitempty" type:"Repeated"`
	// The review status.
	//
	// example:
	//
	// Reviewing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCodeReviewResponseBodyCodeReview) String() string {
	return dara.Prettify(s)
}

func (s GetCodeReviewResponseBodyCodeReview) GoString() string {
	return s.String()
}

func (s *GetCodeReviewResponseBodyCodeReview) GetCommand() *string {
	return s.Command
}

func (s *GetCodeReviewResponseBodyCodeReview) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCodeReviewResponseBodyCodeReview) GetCreator() *string {
	return s.Creator
}

func (s *GetCodeReviewResponseBodyCodeReview) GetDescription() *string {
	return s.Description
}

func (s *GetCodeReviewResponseBodyCodeReview) GetFinalReviewer() *string {
	return s.FinalReviewer
}

func (s *GetCodeReviewResponseBodyCodeReview) GetId() *string {
	return s.Id
}

func (s *GetCodeReviewResponseBodyCodeReview) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetCodeReviewResponseBodyCodeReview) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetCodeReviewResponseBodyCodeReview) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetCodeReviewResponseBodyCodeReview) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetCodeReviewResponseBodyCodeReview) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCodeReviewResponseBodyCodeReview) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCodeReviewResponseBodyCodeReview) GetReviewTime() *int64 {
	return s.ReviewTime
}

func (s *GetCodeReviewResponseBodyCodeReview) GetReviewers() []*string {
	return s.Reviewers
}

func (s *GetCodeReviewResponseBodyCodeReview) GetStatus() *string {
	return s.Status
}

func (s *GetCodeReviewResponseBodyCodeReview) SetCommand(v string) *GetCodeReviewResponseBodyCodeReview {
	s.Command = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetCreateTime(v int64) *GetCodeReviewResponseBodyCodeReview {
	s.CreateTime = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetCreator(v string) *GetCodeReviewResponseBodyCodeReview {
	s.Creator = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetDescription(v string) *GetCodeReviewResponseBodyCodeReview {
	s.Description = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetFinalReviewer(v string) *GetCodeReviewResponseBodyCodeReview {
	s.FinalReviewer = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetId(v string) *GetCodeReviewResponseBodyCodeReview {
	s.Id = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetModifyTime(v int64) *GetCodeReviewResponseBodyCodeReview {
	s.ModifyTime = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetObjectId(v string) *GetCodeReviewResponseBodyCodeReview {
	s.ObjectId = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetObjectName(v string) *GetCodeReviewResponseBodyCodeReview {
	s.ObjectName = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetObjectType(v string) *GetCodeReviewResponseBodyCodeReview {
	s.ObjectType = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetProjectId(v int64) *GetCodeReviewResponseBodyCodeReview {
	s.ProjectId = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetRegionId(v string) *GetCodeReviewResponseBodyCodeReview {
	s.RegionId = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetReviewTime(v int64) *GetCodeReviewResponseBodyCodeReview {
	s.ReviewTime = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetReviewers(v []*string) *GetCodeReviewResponseBodyCodeReview {
	s.Reviewers = v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) SetStatus(v string) *GetCodeReviewResponseBodyCodeReview {
	s.Status = &v
	return s
}

func (s *GetCodeReviewResponseBodyCodeReview) Validate() error {
	return dara.Validate(s)
}
