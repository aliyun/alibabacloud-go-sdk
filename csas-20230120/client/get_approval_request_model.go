// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalId(v string) *GetApprovalRequest
	GetApprovalId() *string
}

type GetApprovalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// approval-872b5e911b35****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
}

func (s GetApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalRequest) GoString() string {
	return s.String()
}

func (s *GetApprovalRequest) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *GetApprovalRequest) SetApprovalId(v string) *GetApprovalRequest {
	s.ApprovalId = &v
	return s
}

func (s *GetApprovalRequest) Validate() error {
	return dara.Validate(s)
}
