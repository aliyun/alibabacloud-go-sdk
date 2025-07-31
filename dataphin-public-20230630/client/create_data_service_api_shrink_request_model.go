// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceApiShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateDataServiceApiShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateDataServiceApiShrinkRequest
	GetOpTenantId() *int64
}

type CreateDataServiceApiShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataServiceApiShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateDataServiceApiShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataServiceApiShrinkRequest) SetCreateCommandShrink(v string) *CreateDataServiceApiShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDataServiceApiShrinkRequest) SetOpTenantId(v int64) *CreateDataServiceApiShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataServiceApiShrinkRequest) Validate() error {
	return dara.Validate(s)
}
