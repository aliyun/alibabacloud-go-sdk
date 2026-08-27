// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceObjectBindingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectBindingsShrink(v string) *ReplaceObjectBindingsShrinkRequest
	GetObjectBindingsShrink() *string
	SetSourceId(v string) *ReplaceObjectBindingsShrinkRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceObjectBindingsShrinkRequest
	GetTenantId() *string
}

type ReplaceObjectBindingsShrinkRequest struct {
	// The new list of object bindings (full replacement. Pass an empty list to clear all bindings).
	//
	// This parameter is required.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
	// The ID of the personal FILE data source to be replaced (unique within the tenant).
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly through winnexo-cli using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceObjectBindingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *ReplaceObjectBindingsShrinkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceObjectBindingsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceObjectBindingsShrinkRequest) SetObjectBindingsShrink(v string) *ReplaceObjectBindingsShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetSourceId(v string) *ReplaceObjectBindingsShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetTenantId(v string) *ReplaceObjectBindingsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
