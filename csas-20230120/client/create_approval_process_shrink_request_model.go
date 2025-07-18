// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApprovalProcessShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateApprovalProcessShrinkRequest
	GetDescription() *string
	SetMatchSchemasShrink(v string) *CreateApprovalProcessShrinkRequest
	GetMatchSchemasShrink() *string
	SetProcessName(v string) *CreateApprovalProcessShrinkRequest
	GetProcessName() *string
	SetProcessNodes(v [][]*string) *CreateApprovalProcessShrinkRequest
	GetProcessNodes() [][]*string
}

type CreateApprovalProcessShrinkRequest struct {
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MatchSchemasShrink *string `json:"MatchSchemas,omitempty" xml:"MatchSchemas,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_process
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// This parameter is required.
	ProcessNodes [][]*string `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
}

func (s CreateApprovalProcessShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApprovalProcessShrinkRequest) GetMatchSchemasShrink() *string {
	return s.MatchSchemasShrink
}

func (s *CreateApprovalProcessShrinkRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *CreateApprovalProcessShrinkRequest) GetProcessNodes() [][]*string {
	return s.ProcessNodes
}

func (s *CreateApprovalProcessShrinkRequest) SetDescription(v string) *CreateApprovalProcessShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApprovalProcessShrinkRequest) SetMatchSchemasShrink(v string) *CreateApprovalProcessShrinkRequest {
	s.MatchSchemasShrink = &v
	return s
}

func (s *CreateApprovalProcessShrinkRequest) SetProcessName(v string) *CreateApprovalProcessShrinkRequest {
	s.ProcessName = &v
	return s
}

func (s *CreateApprovalProcessShrinkRequest) SetProcessNodes(v [][]*string) *CreateApprovalProcessShrinkRequest {
	s.ProcessNodes = v
	return s
}

func (s *CreateApprovalProcessShrinkRequest) Validate() error {
	return dara.Validate(s)
}
