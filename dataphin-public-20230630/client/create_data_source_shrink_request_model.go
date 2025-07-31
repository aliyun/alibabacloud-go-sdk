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
}

type CreateDataSourceShrinkRequest struct {
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *CreateDataSourceShrinkRequest) SetCreateCommandShrink(v string) *CreateDataSourceShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetOpTenantId(v int64) *CreateDataSourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
