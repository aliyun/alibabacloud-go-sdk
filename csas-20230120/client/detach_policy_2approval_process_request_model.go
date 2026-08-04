// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicy2ApprovalProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DetachPolicy2ApprovalProcessRequest
	GetPolicyId() *string
	SetPolicyType(v string) *DetachPolicy2ApprovalProcessRequest
	GetPolicyType() *string
	SetProcessId(v string) *DetachPolicy2ApprovalProcessRequest
	GetProcessId() *string
}

type DetachPolicy2ApprovalProcessRequest struct {
	// Business policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ladp-27a4fedf5e73****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Policy type. Valid values:
	//
	// - **DomainBlacklist**: Domain blacklist.
	//
	// - **DomainWhitelist**: Domain whitelist.
	//
	// - **SoftwareBlock**: Software disable.
	//
	// - **AppUninstall**: Terminal uninstall.
	//
	// - **DlpSend**: File outbound transfer.
	//
	// - **PeripheralBlock**: Peripheral control.
	//
	// This parameter is required.
	//
	// example:
	//
	// PeripheralBlock
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// Approval process ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// approval-process-2677fcf063f5****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s DetachPolicy2ApprovalProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicy2ApprovalProcessRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicy2ApprovalProcessRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DetachPolicy2ApprovalProcessRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetachPolicy2ApprovalProcessRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *DetachPolicy2ApprovalProcessRequest) SetPolicyId(v string) *DetachPolicy2ApprovalProcessRequest {
	s.PolicyId = &v
	return s
}

func (s *DetachPolicy2ApprovalProcessRequest) SetPolicyType(v string) *DetachPolicy2ApprovalProcessRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicy2ApprovalProcessRequest) SetProcessId(v string) *DetachPolicy2ApprovalProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *DetachPolicy2ApprovalProcessRequest) Validate() error {
	return dara.Validate(s)
}
