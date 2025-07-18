// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicy2ApprovalProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *AttachPolicy2ApprovalProcessRequest
	GetPolicyId() *string
	SetPolicyType(v string) *AttachPolicy2ApprovalProcessRequest
	GetPolicyType() *string
	SetProcessId(v string) *AttachPolicy2ApprovalProcessRequest
	GetProcessId() *string
}

type AttachPolicy2ApprovalProcessRequest struct {
	// example:
	//
	// ladp-27a4fedf5e73****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// approval-process-2677fcf063f5****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s AttachPolicy2ApprovalProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicy2ApprovalProcessRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicy2ApprovalProcessRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *AttachPolicy2ApprovalProcessRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AttachPolicy2ApprovalProcessRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *AttachPolicy2ApprovalProcessRequest) SetPolicyId(v string) *AttachPolicy2ApprovalProcessRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachPolicy2ApprovalProcessRequest) SetPolicyType(v string) *AttachPolicy2ApprovalProcessRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicy2ApprovalProcessRequest) SetProcessId(v string) *AttachPolicy2ApprovalProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *AttachPolicy2ApprovalProcessRequest) Validate() error {
	return dara.Validate(s)
}
