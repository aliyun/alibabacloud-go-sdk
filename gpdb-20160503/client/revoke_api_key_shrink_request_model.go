// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApiKeyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *RevokeApiKeyShrinkRequest
	GetKeyId() *string
	SetRegionId(v string) *RevokeApiKeyShrinkRequest
	GetRegionId() *string
	SetServiceIdsShrink(v string) *RevokeApiKeyShrinkRequest
	GetServiceIdsShrink() *string
	SetWorkspaceId(v string) *RevokeApiKeyShrinkRequest
	GetWorkspaceId() *string
}

type RevokeApiKeyShrinkRequest struct {
	// API KEY ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// api-xxxxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of service IDs to be authorized.
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

func (s RevokeApiKeyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeApiKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeApiKeyShrinkRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *RevokeApiKeyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeApiKeyShrinkRequest) GetServiceIdsShrink() *string {
	return s.ServiceIdsShrink
}

func (s *RevokeApiKeyShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RevokeApiKeyShrinkRequest) SetKeyId(v string) *RevokeApiKeyShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *RevokeApiKeyShrinkRequest) SetRegionId(v string) *RevokeApiKeyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeApiKeyShrinkRequest) SetServiceIdsShrink(v string) *RevokeApiKeyShrinkRequest {
	s.ServiceIdsShrink = &v
	return s
}

func (s *RevokeApiKeyShrinkRequest) SetWorkspaceId(v string) *RevokeApiKeyShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RevokeApiKeyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
