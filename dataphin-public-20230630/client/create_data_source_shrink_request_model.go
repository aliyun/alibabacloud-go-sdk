// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateDataSourceShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateDataSourceShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateDataSourceShrinkRequest
	GetOpUserId() *string
}

type CreateDataSourceShrinkRequest struct {
	// Request object
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// Tenant ID
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

func (s CreateDataSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateDataSourceShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataSourceShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateDataSourceShrinkRequest) SetCreateCommandShrink(v string) *CreateDataSourceShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetOpTenantId(v int64) *CreateDataSourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetOpUserId(v string) *CreateDataSourceShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
