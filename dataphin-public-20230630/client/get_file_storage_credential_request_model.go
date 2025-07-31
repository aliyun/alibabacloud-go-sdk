// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileStorageCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetFileStorageCredentialRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetFileStorageCredentialRequest
	GetProjectId() *int64
	SetPurpose(v string) *GetFileStorageCredentialRequest
	GetPurpose() *string
	SetUseVpcEndpoint(v bool) *GetFileStorageCredentialRequest
	GetUseVpcEndpoint() *bool
}

type GetFileStorageCredentialRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 1030131021
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// RESOURCE
	Purpose        *string `json:"Purpose,omitempty" xml:"Purpose,omitempty"`
	UseVpcEndpoint *bool   `json:"UseVpcEndpoint,omitempty" xml:"UseVpcEndpoint,omitempty"`
}

func (s GetFileStorageCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileStorageCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetFileStorageCredentialRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetFileStorageCredentialRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFileStorageCredentialRequest) GetPurpose() *string {
	return s.Purpose
}

func (s *GetFileStorageCredentialRequest) GetUseVpcEndpoint() *bool {
	return s.UseVpcEndpoint
}

func (s *GetFileStorageCredentialRequest) SetOpTenantId(v int64) *GetFileStorageCredentialRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetFileStorageCredentialRequest) SetProjectId(v int64) *GetFileStorageCredentialRequest {
	s.ProjectId = &v
	return s
}

func (s *GetFileStorageCredentialRequest) SetPurpose(v string) *GetFileStorageCredentialRequest {
	s.Purpose = &v
	return s
}

func (s *GetFileStorageCredentialRequest) SetUseVpcEndpoint(v bool) *GetFileStorageCredentialRequest {
	s.UseVpcEndpoint = &v
	return s
}

func (s *GetFileStorageCredentialRequest) Validate() error {
	return dara.Validate(s)
}
