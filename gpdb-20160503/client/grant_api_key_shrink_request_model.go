// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantApiKeyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *GrantApiKeyShrinkRequest
	GetKeyId() *string
	SetRegionId(v string) *GrantApiKeyShrinkRequest
	GetRegionId() *string
	SetServiceIdsShrink(v string) *GrantApiKeyShrinkRequest
	GetServiceIdsShrink() *string
	SetWorkspaceId(v string) *GrantApiKeyShrinkRequest
	GetWorkspaceId() *string
}

type GrantApiKeyShrinkRequest struct {
	// The ID of the API key.
	//
	// This parameter is required.
	//
	// example:
	//
	// api-xxxxxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of service IDs to authorize.
	//
	// This parameter is required.
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

func (s GrantApiKeyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantApiKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantApiKeyShrinkRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GrantApiKeyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantApiKeyShrinkRequest) GetServiceIdsShrink() *string {
	return s.ServiceIdsShrink
}

func (s *GrantApiKeyShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GrantApiKeyShrinkRequest) SetKeyId(v string) *GrantApiKeyShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *GrantApiKeyShrinkRequest) SetRegionId(v string) *GrantApiKeyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GrantApiKeyShrinkRequest) SetServiceIdsShrink(v string) *GrantApiKeyShrinkRequest {
	s.ServiceIdsShrink = &v
	return s
}

func (s *GrantApiKeyShrinkRequest) SetWorkspaceId(v string) *GrantApiKeyShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GrantApiKeyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
