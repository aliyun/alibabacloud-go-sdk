// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *GrantApiKeyRequest
	GetKeyId() *string
	SetRegionId(v string) *GrantApiKeyRequest
	GetRegionId() *string
	SetServiceIds(v []*string) *GrantApiKeyRequest
	GetServiceIds() []*string
	SetWorkspaceId(v string) *GrantApiKeyRequest
	GetWorkspaceId() *string
}

type GrantApiKeyRequest struct {
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

func (s GrantApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantApiKeyRequest) GoString() string {
	return s.String()
}

func (s *GrantApiKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GrantApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantApiKeyRequest) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *GrantApiKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GrantApiKeyRequest) SetKeyId(v string) *GrantApiKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GrantApiKeyRequest) SetRegionId(v string) *GrantApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *GrantApiKeyRequest) SetServiceIds(v []*string) *GrantApiKeyRequest {
	s.ServiceIds = v
	return s
}

func (s *GrantApiKeyRequest) SetWorkspaceId(v string) *GrantApiKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GrantApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
