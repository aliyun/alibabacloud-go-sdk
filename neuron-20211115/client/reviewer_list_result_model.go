// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewerListResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReviewerListResult
	GetRequestId() *string
	SetReviewers(v []*Reviewer) *ReviewerListResult
	GetReviewers() []*Reviewer
}

type ReviewerListResult struct {
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Reviewers []*Reviewer `json:"reviewers,omitempty" xml:"reviewers,omitempty" type:"Repeated"`
}

func (s ReviewerListResult) String() string {
	return dara.Prettify(s)
}

func (s ReviewerListResult) GoString() string {
	return s.String()
}

func (s *ReviewerListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ReviewerListResult) GetReviewers() []*Reviewer {
	return s.Reviewers
}

func (s *ReviewerListResult) SetRequestId(v string) *ReviewerListResult {
	s.RequestId = &v
	return s
}

func (s *ReviewerListResult) SetReviewers(v []*Reviewer) *ReviewerListResult {
	s.Reviewers = v
	return s
}

func (s *ReviewerListResult) Validate() error {
	if s.Reviewers != nil {
		for _, item := range s.Reviewers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
