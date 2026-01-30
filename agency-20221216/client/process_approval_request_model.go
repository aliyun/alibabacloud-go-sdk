// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationSheetId(v string) *ProcessApprovalRequest
	GetApplicationSheetId() *string
	SetApprovalAction(v string) *ProcessApprovalRequest
	GetApprovalAction() *string
	SetApprovalComments(v string) *ProcessApprovalRequest
	GetApprovalComments() *string
}

type ProcessApprovalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d54ca949-9b88-4514-add3-c6029c4027f4
	ApplicationSheetId *string `json:"ApplicationSheetId,omitempty" xml:"ApplicationSheetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Approve
	ApprovalAction *string `json:"ApprovalAction,omitempty" xml:"ApprovalAction,omitempty"`
	// This parameter is required.
	ApprovalComments *string `json:"ApprovalComments,omitempty" xml:"ApprovalComments,omitempty"`
}

func (s ProcessApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s ProcessApprovalRequest) GoString() string {
	return s.String()
}

func (s *ProcessApprovalRequest) GetApplicationSheetId() *string {
	return s.ApplicationSheetId
}

func (s *ProcessApprovalRequest) GetApprovalAction() *string {
	return s.ApprovalAction
}

func (s *ProcessApprovalRequest) GetApprovalComments() *string {
	return s.ApprovalComments
}

func (s *ProcessApprovalRequest) SetApplicationSheetId(v string) *ProcessApprovalRequest {
	s.ApplicationSheetId = &v
	return s
}

func (s *ProcessApprovalRequest) SetApprovalAction(v string) *ProcessApprovalRequest {
	s.ApprovalAction = &v
	return s
}

func (s *ProcessApprovalRequest) SetApprovalComments(v string) *ProcessApprovalRequest {
	s.ApprovalComments = &v
	return s
}

func (s *ProcessApprovalRequest) Validate() error {
	return dara.Validate(s)
}
