// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantDataServiceApiShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantCommandShrink(v string) *GrantDataServiceApiShrinkRequest
	GetGrantCommandShrink() *string
	SetOpTenantId(v int64) *GrantDataServiceApiShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GrantDataServiceApiShrinkRequest
	GetProjectId() *int32
}

type GrantDataServiceApiShrinkRequest struct {
	// This parameter is required.
	GrantCommandShrink *string `json:"GrantCommand,omitempty" xml:"GrantCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GrantDataServiceApiShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantDataServiceApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantDataServiceApiShrinkRequest) GetGrantCommandShrink() *string {
	return s.GrantCommandShrink
}

func (s *GrantDataServiceApiShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GrantDataServiceApiShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GrantDataServiceApiShrinkRequest) SetGrantCommandShrink(v string) *GrantDataServiceApiShrinkRequest {
	s.GrantCommandShrink = &v
	return s
}

func (s *GrantDataServiceApiShrinkRequest) SetOpTenantId(v int64) *GrantDataServiceApiShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GrantDataServiceApiShrinkRequest) SetProjectId(v int32) *GrantDataServiceApiShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GrantDataServiceApiShrinkRequest) Validate() error {
	return dara.Validate(s)
}
