// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidateApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalId(v string) *InvalidateApprovalRequest
	GetApprovalId() *string
}

type InvalidateApprovalRequest struct {
	// The ID of the approval instance to be immediately invalidated. You can obtain the ID by calling the ListApprovals operation. You can specify only one approval instance ID under the current Alibaba Cloud account at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// approval-6b5188a28634****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
}

func (s InvalidateApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s InvalidateApprovalRequest) GoString() string {
	return s.String()
}

func (s *InvalidateApprovalRequest) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *InvalidateApprovalRequest) SetApprovalId(v string) *InvalidateApprovalRequest {
	s.ApprovalId = &v
	return s
}

func (s *InvalidateApprovalRequest) Validate() error {
	return dara.Validate(s)
}
