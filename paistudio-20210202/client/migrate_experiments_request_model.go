// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateExperimentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *MigrateExperimentsRequest
	GetAccessibility() *string
	SetDestFolderId(v string) *MigrateExperimentsRequest
	GetDestFolderId() *string
	SetIsOwner(v bool) *MigrateExperimentsRequest
	GetIsOwner() *bool
	SetSourceExpId(v int64) *MigrateExperimentsRequest
	GetSourceExpId() *int64
	SetUpdateIfExists(v bool) *MigrateExperimentsRequest
	GetUpdateIfExists() *bool
	SetWorkspaceId(v string) *MigrateExperimentsRequest
	GetWorkspaceId() *string
}

type MigrateExperimentsRequest struct {
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// folder-12321313
	DestFolderId *string `json:"DestFolderId,omitempty" xml:"DestFolderId,omitempty"`
	// example:
	//
	// true
	IsOwner *bool `json:"IsOwner,omitempty" xml:"IsOwner,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	SourceExpId *int64 `json:"SourceExpId,omitempty" xml:"SourceExpId,omitempty"`
	// example:
	//
	// true
	UpdateIfExists *bool `json:"UpdateIfExists,omitempty" xml:"UpdateIfExists,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MigrateExperimentsRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateExperimentsRequest) GoString() string {
	return s.String()
}

func (s *MigrateExperimentsRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *MigrateExperimentsRequest) GetDestFolderId() *string {
	return s.DestFolderId
}

func (s *MigrateExperimentsRequest) GetIsOwner() *bool {
	return s.IsOwner
}

func (s *MigrateExperimentsRequest) GetSourceExpId() *int64 {
	return s.SourceExpId
}

func (s *MigrateExperimentsRequest) GetUpdateIfExists() *bool {
	return s.UpdateIfExists
}

func (s *MigrateExperimentsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *MigrateExperimentsRequest) SetAccessibility(v string) *MigrateExperimentsRequest {
	s.Accessibility = &v
	return s
}

func (s *MigrateExperimentsRequest) SetDestFolderId(v string) *MigrateExperimentsRequest {
	s.DestFolderId = &v
	return s
}

func (s *MigrateExperimentsRequest) SetIsOwner(v bool) *MigrateExperimentsRequest {
	s.IsOwner = &v
	return s
}

func (s *MigrateExperimentsRequest) SetSourceExpId(v int64) *MigrateExperimentsRequest {
	s.SourceExpId = &v
	return s
}

func (s *MigrateExperimentsRequest) SetUpdateIfExists(v bool) *MigrateExperimentsRequest {
	s.UpdateIfExists = &v
	return s
}

func (s *MigrateExperimentsRequest) SetWorkspaceId(v string) *MigrateExperimentsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *MigrateExperimentsRequest) Validate() error {
	return dara.Validate(s)
}
