// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetStackGroupRequest
	GetRegionId() *string
	SetStackGroupId(v string) *GetStackGroupRequest
	GetStackGroupId() *string
	SetStackGroupName(v string) *GetStackGroupRequest
	GetStackGroupName() *string
}

type GetStackGroupRequest struct {
	// The name of the stack group. The name must be unique within a region.
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	//
	// >  You must specify one of the StackGroupName and StackGroupId parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// fd0ddef9-9540-4b42-a464-94f77835****
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The ID of the stack group.
	//
	// >  You must specify one of the StackGroupName and StackGroupId parameters.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s GetStackGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupRequest) GoString() string {
	return s.String()
}

func (s *GetStackGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackGroupRequest) GetStackGroupId() *string {
	return s.StackGroupId
}

func (s *GetStackGroupRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *GetStackGroupRequest) SetRegionId(v string) *GetStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackGroupRequest) SetStackGroupId(v string) *GetStackGroupRequest {
	s.StackGroupId = &v
	return s
}

func (s *GetStackGroupRequest) SetStackGroupName(v string) *GetStackGroupRequest {
	s.StackGroupName = &v
	return s
}

func (s *GetStackGroupRequest) Validate() error {
	return dara.Validate(s)
}
