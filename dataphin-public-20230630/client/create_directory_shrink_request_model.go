// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirectoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateDirectoryShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateDirectoryShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateDirectoryShrinkRequest
	GetOpUserId() *string
}

type CreateDirectoryShrinkRequest struct {
	// The create request.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateDirectoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDirectoryShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateDirectoryShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDirectoryShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateDirectoryShrinkRequest) SetCreateCommandShrink(v string) *CreateDirectoryShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDirectoryShrinkRequest) SetOpTenantId(v int64) *CreateDirectoryShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDirectoryShrinkRequest) SetOpUserId(v string) *CreateDirectoryShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateDirectoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
