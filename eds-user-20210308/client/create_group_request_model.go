// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *CreateGroupRequest
	GetBizType() *string
	SetDescription(v string) *CreateGroupRequest
	GetDescription() *string
	SetGroupName(v string) *CreateGroupRequest
	GetGroupName() *string
	SetParentGroupId(v string) *CreateGroupRequest
	GetParentGroupId() *string
	SetSolutionId(v string) *CreateGroupRequest
	GetSolutionId() *string
}

type CreateGroupRequest struct {
	// example:
	//
	// ENTERPRISE
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// root
	ParentGroupId *string `json:"ParentGroupId,omitempty" xml:"ParentGroupId,omitempty"`
	// example:
	//
	// co-0esnf80jab***
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupRequest) GetParentGroupId() *string {
	return s.ParentGroupId
}

func (s *CreateGroupRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *CreateGroupRequest) SetBizType(v string) *CreateGroupRequest {
	s.BizType = &v
	return s
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetParentGroupId(v string) *CreateGroupRequest {
	s.ParentGroupId = &v
	return s
}

func (s *CreateGroupRequest) SetSolutionId(v string) *CreateGroupRequest {
	s.SolutionId = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	return dara.Validate(s)
}
