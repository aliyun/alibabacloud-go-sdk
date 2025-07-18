// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalProcessShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateApprovalProcessShrinkRequest
	GetDescription() *string
	SetMatchSchemasShrink(v string) *UpdateApprovalProcessShrinkRequest
	GetMatchSchemasShrink() *string
	SetProcessId(v string) *UpdateApprovalProcessShrinkRequest
	GetProcessId() *string
	SetProcessName(v string) *UpdateApprovalProcessShrinkRequest
	GetProcessName() *string
	SetProcessNodes(v [][]*string) *UpdateApprovalProcessShrinkRequest
	GetProcessNodes() [][]*string
}

type UpdateApprovalProcessShrinkRequest struct {
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MatchSchemasShrink *string `json:"MatchSchemas,omitempty" xml:"MatchSchemas,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// approval-process-f16bf74b2b29****
	ProcessId    *string     `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ProcessName  *string     `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodes [][]*string `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
}

func (s UpdateApprovalProcessShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApprovalProcessShrinkRequest) GetMatchSchemasShrink() *string {
	return s.MatchSchemasShrink
}

func (s *UpdateApprovalProcessShrinkRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *UpdateApprovalProcessShrinkRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *UpdateApprovalProcessShrinkRequest) GetProcessNodes() [][]*string {
	return s.ProcessNodes
}

func (s *UpdateApprovalProcessShrinkRequest) SetDescription(v string) *UpdateApprovalProcessShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetMatchSchemasShrink(v string) *UpdateApprovalProcessShrinkRequest {
	s.MatchSchemasShrink = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetProcessId(v string) *UpdateApprovalProcessShrinkRequest {
	s.ProcessId = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetProcessName(v string) *UpdateApprovalProcessShrinkRequest {
	s.ProcessName = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetProcessNodes(v [][]*string) *UpdateApprovalProcessShrinkRequest {
	s.ProcessNodes = v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) Validate() error {
	return dara.Validate(s)
}
