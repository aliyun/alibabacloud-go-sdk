// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForkReviewAuditCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApprove(v bool) *ForkReviewAuditCmd
	GetApprove() *bool
}

type ForkReviewAuditCmd struct {
	// example:
	//
	// true
	Approve *bool `json:"approve,omitempty" xml:"approve,omitempty"`
}

func (s ForkReviewAuditCmd) String() string {
	return dara.Prettify(s)
}

func (s ForkReviewAuditCmd) GoString() string {
	return s.String()
}

func (s *ForkReviewAuditCmd) GetApprove() *bool {
	return s.Approve
}

func (s *ForkReviewAuditCmd) SetApprove(v bool) *ForkReviewAuditCmd {
	s.Approve = &v
	return s
}

func (s *ForkReviewAuditCmd) Validate() error {
	return dara.Validate(s)
}
