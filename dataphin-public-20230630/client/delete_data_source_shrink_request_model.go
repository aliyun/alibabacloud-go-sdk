// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteDataSourceShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteDataSourceShrinkRequest
	GetOpTenantId() *int64
}

type DeleteDataSourceShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteDataSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteDataSourceShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteDataSourceShrinkRequest) SetDeleteCommandShrink(v string) *DeleteDataSourceShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteDataSourceShrinkRequest) SetOpTenantId(v int64) *DeleteDataSourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteDataSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
