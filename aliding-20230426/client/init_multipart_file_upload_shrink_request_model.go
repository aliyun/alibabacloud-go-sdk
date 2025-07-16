// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitMultipartFileUploadShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOptionShrink(v string) *InitMultipartFileUploadShrinkRequest
	GetOptionShrink() *string
	SetParentDentryUuid(v string) *InitMultipartFileUploadShrinkRequest
	GetParentDentryUuid() *string
	SetTenantContextShrink(v string) *InitMultipartFileUploadShrinkRequest
	GetTenantContextShrink() *string
}

type InitMultipartFileUploadShrinkRequest struct {
	OptionShrink *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// dentryUuid
	ParentDentryUuid    *string `json:"ParentDentryUuid,omitempty" xml:"ParentDentryUuid,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s InitMultipartFileUploadShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadShrinkRequest) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *InitMultipartFileUploadShrinkRequest) GetParentDentryUuid() *string {
	return s.ParentDentryUuid
}

func (s *InitMultipartFileUploadShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *InitMultipartFileUploadShrinkRequest) SetOptionShrink(v string) *InitMultipartFileUploadShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *InitMultipartFileUploadShrinkRequest) SetParentDentryUuid(v string) *InitMultipartFileUploadShrinkRequest {
	s.ParentDentryUuid = &v
	return s
}

func (s *InitMultipartFileUploadShrinkRequest) SetTenantContextShrink(v string) *InitMultipartFileUploadShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InitMultipartFileUploadShrinkRequest) Validate() error {
	return dara.Validate(s)
}
