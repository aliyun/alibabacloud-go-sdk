// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForkReview interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *ForkReview
	GetApplicant() *string
	SetForkId(v int64) *ForkReview
	GetForkId() *int64
	SetGitGroup(v string) *ForkReview
	GetGitGroup() *string
	SetId(v int64) *ForkReview
	GetId() *int64
	SetPbcName(v string) *ForkReview
	GetPbcName() *string
	SetRepoUrls(v []*string) *ForkReview
	GetRepoUrls() []*string
	SetRequestId(v string) *ForkReview
	GetRequestId() *string
	SetReviewer(v string) *ForkReview
	GetReviewer() *string
	SetReviewerId(v string) *ForkReview
	GetReviewerId() *string
	SetStatus(v string) *ForkReview
	GetStatus() *string
	SetUsage(v string) *ForkReview
	GetUsage() *string
}

type ForkReview struct {
	// example:
	//
	// 旭坤
	Applicant *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	// example:
	//
	// 1
	ForkId *int64 `json:"forkId,omitempty" xml:"forkId,omitempty"`
	// example:
	//
	// global-mall
	GitGroup *string `json:"gitGroup,omitempty" xml:"gitGroup,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// product
	PbcName   *string   `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	RepoUrls  []*string `json:"repoUrls,omitempty" xml:"repoUrls,omitempty" type:"Repeated"`
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 中驿
	Reviewer *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// example:
	//
	// 194835334
	ReviewerId *string `json:"reviewerId,omitempty" xml:"reviewerId,omitempty"`
	// example:
	//
	// 待审核
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 商城国际化版本
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ForkReview) String() string {
	return dara.Prettify(s)
}

func (s ForkReview) GoString() string {
	return s.String()
}

func (s *ForkReview) GetApplicant() *string {
	return s.Applicant
}

func (s *ForkReview) GetForkId() *int64 {
	return s.ForkId
}

func (s *ForkReview) GetGitGroup() *string {
	return s.GitGroup
}

func (s *ForkReview) GetId() *int64 {
	return s.Id
}

func (s *ForkReview) GetPbcName() *string {
	return s.PbcName
}

func (s *ForkReview) GetRepoUrls() []*string {
	return s.RepoUrls
}

func (s *ForkReview) GetRequestId() *string {
	return s.RequestId
}

func (s *ForkReview) GetReviewer() *string {
	return s.Reviewer
}

func (s *ForkReview) GetReviewerId() *string {
	return s.ReviewerId
}

func (s *ForkReview) GetStatus() *string {
	return s.Status
}

func (s *ForkReview) GetUsage() *string {
	return s.Usage
}

func (s *ForkReview) SetApplicant(v string) *ForkReview {
	s.Applicant = &v
	return s
}

func (s *ForkReview) SetForkId(v int64) *ForkReview {
	s.ForkId = &v
	return s
}

func (s *ForkReview) SetGitGroup(v string) *ForkReview {
	s.GitGroup = &v
	return s
}

func (s *ForkReview) SetId(v int64) *ForkReview {
	s.Id = &v
	return s
}

func (s *ForkReview) SetPbcName(v string) *ForkReview {
	s.PbcName = &v
	return s
}

func (s *ForkReview) SetRepoUrls(v []*string) *ForkReview {
	s.RepoUrls = v
	return s
}

func (s *ForkReview) SetRequestId(v string) *ForkReview {
	s.RequestId = &v
	return s
}

func (s *ForkReview) SetReviewer(v string) *ForkReview {
	s.Reviewer = &v
	return s
}

func (s *ForkReview) SetReviewerId(v string) *ForkReview {
	s.ReviewerId = &v
	return s
}

func (s *ForkReview) SetStatus(v string) *ForkReview {
	s.Status = &v
	return s
}

func (s *ForkReview) SetUsage(v string) *ForkReview {
	s.Usage = &v
	return s
}

func (s *ForkReview) Validate() error {
	return dara.Validate(s)
}
