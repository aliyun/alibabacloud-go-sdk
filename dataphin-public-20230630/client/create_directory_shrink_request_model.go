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
}

type CreateDirectoryShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *CreateDirectoryShrinkRequest) SetCreateCommandShrink(v string) *CreateDirectoryShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDirectoryShrinkRequest) SetOpTenantId(v int64) *CreateDirectoryShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDirectoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
