// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMultiDimTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *AddMultiDimTableShrinkRequest
	GetBaseId() *string
	SetFieldsShrink(v string) *AddMultiDimTableShrinkRequest
	GetFieldsShrink() *string
	SetName(v string) *AddMultiDimTableShrinkRequest
	GetName() *string
	SetTenantContextShrink(v string) *AddMultiDimTableShrinkRequest
	GetTenantContextShrink() *string
}

type AddMultiDimTableShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId              *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	FieldsShrink        *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s AddMultiDimTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *AddMultiDimTableShrinkRequest) GetFieldsShrink() *string {
	return s.FieldsShrink
}

func (s *AddMultiDimTableShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddMultiDimTableShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddMultiDimTableShrinkRequest) SetBaseId(v string) *AddMultiDimTableShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *AddMultiDimTableShrinkRequest) SetFieldsShrink(v string) *AddMultiDimTableShrinkRequest {
	s.FieldsShrink = &v
	return s
}

func (s *AddMultiDimTableShrinkRequest) SetName(v string) *AddMultiDimTableShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddMultiDimTableShrinkRequest) SetTenantContextShrink(v string) *AddMultiDimTableShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddMultiDimTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
