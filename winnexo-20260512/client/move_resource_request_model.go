// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceDirectoryId(v string) *MoveResourceRequest
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveResourceRequest
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveResourceRequest
	GetTargetDirectoryId() *string
	SetTenantId(v string) *MoveResourceRequest
	GetTenantId() *string
}

type MoveResourceRequest struct {
	// The source directory ID, which is the personal directory where the resource currently resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceDirectoryId
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// The ID of the resource to be moved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The target directory ID, which is the personal directory to which the resource will be moved.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTargetDirectoryId
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
	// The tenant ID. You can view the tenant ID by logging on to the MaxCompute console and choosing **Tenant Management*	- > **Tenant Properties*	- in the left-side navigation pane.
	//
	// example:
	//
	// PiPklI1iSRTm6VFFqlY9VzbgiEiE
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MoveResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceRequest) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveResourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveResourceRequest) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MoveResourceRequest) SetSourceDirectoryId(v string) *MoveResourceRequest {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveResourceRequest) SetSourceId(v string) *MoveResourceRequest {
	s.SourceId = &v
	return s
}

func (s *MoveResourceRequest) SetTargetDirectoryId(v string) *MoveResourceRequest {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveResourceRequest) SetTenantId(v string) *MoveResourceRequest {
	s.TenantId = &v
	return s
}

func (s *MoveResourceRequest) Validate() error {
	return dara.Validate(s)
}
