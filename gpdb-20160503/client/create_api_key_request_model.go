// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateApiKeyRequest
	GetDescription() *string
	SetKeyName(v string) *CreateApiKeyRequest
	GetKeyName() *string
	SetRegionId(v string) *CreateApiKeyRequest
	GetRegionId() *string
	SetServiceIds(v []*string) *CreateApiKeyRequest
	GetServiceIds() []*string
	SetWorkspaceId(v string) *CreateApiKeyRequest
	GetWorkspaceId() *string
}

type CreateApiKeyRequest struct {
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

func (s CreateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiKeyRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *CreateApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApiKeyRequest) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *CreateApiKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateApiKeyRequest) SetDescription(v string) *CreateApiKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateApiKeyRequest) SetKeyName(v string) *CreateApiKeyRequest {
	s.KeyName = &v
	return s
}

func (s *CreateApiKeyRequest) SetRegionId(v string) *CreateApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApiKeyRequest) SetServiceIds(v []*string) *CreateApiKeyRequest {
	s.ServiceIds = v
	return s
}

func (s *CreateApiKeyRequest) SetWorkspaceId(v string) *CreateApiKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
