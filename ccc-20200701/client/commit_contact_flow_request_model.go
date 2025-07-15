// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *CommitContactFlowRequest
	GetContactFlowId() *string
	SetDefinition(v string) *CommitContactFlowRequest
	GetDefinition() *string
	SetDescription(v string) *CommitContactFlowRequest
	GetDescription() *string
	SetDraftId(v string) *CommitContactFlowRequest
	GetDraftId() *string
	SetInstanceId(v string) *CommitContactFlowRequest
	GetInstanceId() *string
}

type CommitContactFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7d706489-d06d-4a92-8666-8c9dba2c5cb1
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// This parameter is required.
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b28f74ca-5846-4496-8bbd-34fb1750798c
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CommitContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CommitContactFlowRequest) GoString() string {
	return s.String()
}

func (s *CommitContactFlowRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *CommitContactFlowRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CommitContactFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *CommitContactFlowRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *CommitContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CommitContactFlowRequest) SetContactFlowId(v string) *CommitContactFlowRequest {
	s.ContactFlowId = &v
	return s
}

func (s *CommitContactFlowRequest) SetDefinition(v string) *CommitContactFlowRequest {
	s.Definition = &v
	return s
}

func (s *CommitContactFlowRequest) SetDescription(v string) *CommitContactFlowRequest {
	s.Description = &v
	return s
}

func (s *CommitContactFlowRequest) SetDraftId(v string) *CommitContactFlowRequest {
	s.DraftId = &v
	return s
}

func (s *CommitContactFlowRequest) SetInstanceId(v string) *CommitContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *CommitContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
