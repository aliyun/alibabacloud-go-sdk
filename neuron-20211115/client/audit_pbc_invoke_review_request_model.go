// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditPbcInvokeReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprove(v bool) *AuditPbcInvokeReviewRequest
	GetApprove() *bool
}

type AuditPbcInvokeReviewRequest struct {
	Approve *bool `json:"approve,omitempty" xml:"approve,omitempty"`
}

func (s AuditPbcInvokeReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s AuditPbcInvokeReviewRequest) GoString() string {
	return s.String()
}

func (s *AuditPbcInvokeReviewRequest) GetApprove() *bool {
	return s.Approve
}

func (s *AuditPbcInvokeReviewRequest) SetApprove(v bool) *AuditPbcInvokeReviewRequest {
	s.Approve = &v
	return s
}

func (s *AuditPbcInvokeReviewRequest) Validate() error {
	return dara.Validate(s)
}
