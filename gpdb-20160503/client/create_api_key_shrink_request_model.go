// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateApiKeyShrinkRequest
	GetDescription() *string
	SetKeyName(v string) *CreateApiKeyShrinkRequest
	GetKeyName() *string
	SetRegionId(v string) *CreateApiKeyShrinkRequest
	GetRegionId() *string
	SetServiceIdsShrink(v string) *CreateApiKeyShrinkRequest
	GetServiceIdsShrink() *string
	SetWorkspaceId(v string) *CreateApiKeyShrinkRequest
	GetWorkspaceId() *string
}

type CreateApiKeyShrinkRequest struct {
	// The description.
	//
	// example:
	//
	// test secret
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the API key.
	//
	// This parameter is required.
	//
	// example:
	//
	// my api key
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of service IDs to authorize.
	ServiceIdsShrink *string `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateApiKeyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApiKeyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiKeyShrinkRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *CreateApiKeyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApiKeyShrinkRequest) GetServiceIdsShrink() *string {
	return s.ServiceIdsShrink
}

func (s *CreateApiKeyShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateApiKeyShrinkRequest) SetDescription(v string) *CreateApiKeyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApiKeyShrinkRequest) SetKeyName(v string) *CreateApiKeyShrinkRequest {
	s.KeyName = &v
	return s
}

func (s *CreateApiKeyShrinkRequest) SetRegionId(v string) *CreateApiKeyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApiKeyShrinkRequest) SetServiceIdsShrink(v string) *CreateApiKeyShrinkRequest {
	s.ServiceIdsShrink = &v
	return s
}

func (s *CreateApiKeyShrinkRequest) SetWorkspaceId(v string) *CreateApiKeyShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateApiKeyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
