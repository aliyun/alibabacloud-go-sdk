// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateDataServiceAppShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateDataServiceAppShrinkRequest
	GetOpTenantId() *int64
}

type CreateDataServiceAppShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataServiceAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataServiceAppShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateDataServiceAppShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataServiceAppShrinkRequest) SetCreateCommandShrink(v string) *CreateDataServiceAppShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDataServiceAppShrinkRequest) SetOpTenantId(v int64) *CreateDataServiceAppShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataServiceAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
