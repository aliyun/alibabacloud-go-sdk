// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CommitFileShrinkRequest
	GetName() *string
	SetOptionShrink(v string) *CommitFileShrinkRequest
	GetOptionShrink() *string
	SetParentDentryUuid(v string) *CommitFileShrinkRequest
	GetParentDentryUuid() *string
	SetTenantContextShrink(v string) *CommitFileShrinkRequest
	GetTenantContextShrink() *string
	SetUploadKey(v string) *CommitFileShrinkRequest
	GetUploadKey() *string
}

type CommitFileShrinkRequest struct {
	// example:
	//
	// None
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OptionShrink *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// dentryUuid
	ParentDentryUuid    *string `json:"ParentDentryUuid,omitempty" xml:"ParentDentryUuid,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// upload_key
	UploadKey *string `json:"UploadKey,omitempty" xml:"UploadKey,omitempty"`
}

func (s CommitFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CommitFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *CommitFileShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CommitFileShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *CommitFileShrinkRequest) GetParentDentryUuid() *string {
	return s.ParentDentryUuid
}

func (s *CommitFileShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CommitFileShrinkRequest) GetUploadKey() *string {
	return s.UploadKey
}

func (s *CommitFileShrinkRequest) SetName(v string) *CommitFileShrinkRequest {
	s.Name = &v
	return s
}

func (s *CommitFileShrinkRequest) SetOptionShrink(v string) *CommitFileShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *CommitFileShrinkRequest) SetParentDentryUuid(v string) *CommitFileShrinkRequest {
	s.ParentDentryUuid = &v
	return s
}

func (s *CommitFileShrinkRequest) SetTenantContextShrink(v string) *CommitFileShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CommitFileShrinkRequest) SetUploadKey(v string) *CommitFileShrinkRequest {
	s.UploadKey = &v
	return s
}

func (s *CommitFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
