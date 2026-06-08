// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveProcessInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalAction(v string) *ApproveProcessInstanceRequest
	GetApprovalAction() *string
	SetApprovalComment(v string) *ApproveProcessInstanceRequest
	GetApprovalComment() *string
	SetClientToken(v string) *ApproveProcessInstanceRequest
	GetClientToken() *string
	SetProcessInstanceId(v string) *ApproveProcessInstanceRequest
	GetProcessInstanceId() *string
}

type ApproveProcessInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Agree
	ApprovalAction *string `json:"ApprovalAction,omitempty" xml:"ApprovalAction,omitempty"`
	// This parameter is required.
	ApprovalComment *string `json:"ApprovalComment,omitempty" xml:"ApprovalComment,omitempty"`
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 332066440109224007
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
}

func (s ApproveProcessInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *ApproveProcessInstanceRequest) GetApprovalAction() *string {
	return s.ApprovalAction
}

func (s *ApproveProcessInstanceRequest) GetApprovalComment() *string {
	return s.ApprovalComment
}

func (s *ApproveProcessInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ApproveProcessInstanceRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *ApproveProcessInstanceRequest) SetApprovalAction(v string) *ApproveProcessInstanceRequest {
	s.ApprovalAction = &v
	return s
}

func (s *ApproveProcessInstanceRequest) SetApprovalComment(v string) *ApproveProcessInstanceRequest {
	s.ApprovalComment = &v
	return s
}

func (s *ApproveProcessInstanceRequest) SetClientToken(v string) *ApproveProcessInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ApproveProcessInstanceRequest) SetProcessInstanceId(v string) *ApproveProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ApproveProcessInstanceRequest) Validate() error {
	return dara.Validate(s)
}
