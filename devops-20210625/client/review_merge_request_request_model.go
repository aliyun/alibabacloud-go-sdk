// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewMergeRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ReviewMergeRequestRequest
	GetAccessToken() *string
	SetDraftCommentIds(v []*string) *ReviewMergeRequestRequest
	GetDraftCommentIds() []*string
	SetReviewComment(v string) *ReviewMergeRequestRequest
	GetReviewComment() *string
	SetReviewOpinion(v string) *ReviewMergeRequestRequest
	GetReviewOpinion() *string
	SetOrganizationId(v string) *ReviewMergeRequestRequest
	GetOrganizationId() *string
}

type ReviewMergeRequestRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken     *string   `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	DraftCommentIds []*string `json:"draftCommentIds,omitempty" xml:"draftCommentIds,omitempty" type:"Repeated"`
	ReviewComment   *string   `json:"reviewComment,omitempty" xml:"reviewComment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PASS
	ReviewOpinion *string `json:"reviewOpinion,omitempty" xml:"reviewOpinion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ReviewMergeRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s ReviewMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *ReviewMergeRequestRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ReviewMergeRequestRequest) GetDraftCommentIds() []*string {
	return s.DraftCommentIds
}

func (s *ReviewMergeRequestRequest) GetReviewComment() *string {
	return s.ReviewComment
}

func (s *ReviewMergeRequestRequest) GetReviewOpinion() *string {
	return s.ReviewOpinion
}

func (s *ReviewMergeRequestRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ReviewMergeRequestRequest) SetAccessToken(v string) *ReviewMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *ReviewMergeRequestRequest) SetDraftCommentIds(v []*string) *ReviewMergeRequestRequest {
	s.DraftCommentIds = v
	return s
}

func (s *ReviewMergeRequestRequest) SetReviewComment(v string) *ReviewMergeRequestRequest {
	s.ReviewComment = &v
	return s
}

func (s *ReviewMergeRequestRequest) SetReviewOpinion(v string) *ReviewMergeRequestRequest {
	s.ReviewOpinion = &v
	return s
}

func (s *ReviewMergeRequestRequest) SetOrganizationId(v string) *ReviewMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *ReviewMergeRequestRequest) Validate() error {
	return dara.Validate(s)
}
