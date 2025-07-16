// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *ListPermissionsShrinkRequest
	GetDentryUuid() *string
	SetOptionShrink(v string) *ListPermissionsShrinkRequest
	GetOptionShrink() *string
	SetTenantContextShrink(v string) *ListPermissionsShrinkRequest
	GetTenantContextShrink() *string
}

type ListPermissionsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// KGZLxjv9VGkoG9YwHE5wx7k2V6EDybno
	DentryUuid          *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	OptionShrink        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s ListPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *ListPermissionsShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *ListPermissionsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListPermissionsShrinkRequest) SetDentryUuid(v string) *ListPermissionsShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *ListPermissionsShrinkRequest) SetOptionShrink(v string) *ListPermissionsShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *ListPermissionsShrinkRequest) SetTenantContextShrink(v string) *ListPermissionsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
