// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *RevokeApiKeyRequest
	GetKeyId() *string
	SetRegionId(v string) *RevokeApiKeyRequest
	GetRegionId() *string
	SetServiceIds(v []*string) *RevokeApiKeyRequest
	GetServiceIds() []*string
	SetWorkspaceId(v string) *RevokeApiKeyRequest
	GetWorkspaceId() *string
}

type RevokeApiKeyRequest struct {
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
	ServiceIds []*string `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RevokeApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeApiKeyRequest) GoString() string {
	return s.String()
}

func (s *RevokeApiKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *RevokeApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeApiKeyRequest) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *RevokeApiKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RevokeApiKeyRequest) SetKeyId(v string) *RevokeApiKeyRequest {
	s.KeyId = &v
	return s
}

func (s *RevokeApiKeyRequest) SetRegionId(v string) *RevokeApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeApiKeyRequest) SetServiceIds(v []*string) *RevokeApiKeyRequest {
	s.ServiceIds = v
	return s
}

func (s *RevokeApiKeyRequest) SetWorkspaceId(v string) *RevokeApiKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RevokeApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
