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
	SetNewExpiration(v int64) *ApproveProcessInstanceRequest
	GetNewExpiration() *int64
	SetProcessInstanceId(v string) *ApproveProcessInstanceRequest
	GetProcessInstanceId() *string
}

type ApproveProcessInstanceRequest struct {
	// The approval action. Valid values:
	//
	// - Agree: approves the request.
	//
	// - Deny: rejects the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// Agree
	ApprovalAction *string `json:"ApprovalAction,omitempty" xml:"ApprovalAction,omitempty"`
	// The approval comment.
	//
	// This parameter is required.
	//
	// example:
	//
	// Approve authorization
	ApprovalComment *string `json:"ApprovalComment,omitempty" xml:"ApprovalComment,omitempty"`
	// The idempotency token. We recommend that you use a UUID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The modified authorization expiration time. The value is a millisecond-level timestamp.
	//
	// example:
	//
	// 1782541464000
	NewExpiration *int64 `json:"NewExpiration,omitempty" xml:"NewExpiration,omitempty"`
	// The flow instance ID. Both new and legacy Security Center approval orders are supported.
	//
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

func (s *ApproveProcessInstanceRequest) GetNewExpiration() *int64 {
	return s.NewExpiration
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

func (s *ApproveProcessInstanceRequest) SetNewExpiration(v int64) *ApproveProcessInstanceRequest {
	s.NewExpiration = &v
	return s
}

func (s *ApproveProcessInstanceRequest) SetProcessInstanceId(v string) *ApproveProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ApproveProcessInstanceRequest) Validate() error {
	return dara.Validate(s)
}
