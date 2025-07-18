// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalId(v string) *UpdateApprovalStatusRequest
	GetApprovalId() *string
	SetStatus(v string) *UpdateApprovalStatusRequest
	GetStatus() *string
}

type UpdateApprovalStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// approval-872b5e911b35****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateApprovalStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateApprovalStatusRequest) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *UpdateApprovalStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateApprovalStatusRequest) SetApprovalId(v string) *UpdateApprovalStatusRequest {
	s.ApprovalId = &v
	return s
}

func (s *UpdateApprovalStatusRequest) SetStatus(v string) *UpdateApprovalStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateApprovalStatusRequest) Validate() error {
	return dara.Validate(s)
}
