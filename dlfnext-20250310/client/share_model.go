// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShare interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Share
	GetComment() *string
	SetCreatedAt(v int64) *Share
	GetCreatedAt() *int64
	SetCreatedBy(v string) *Share
	GetCreatedBy() *string
	SetEnableWrite(v bool) *Share
	GetEnableWrite() *bool
	SetOwner(v string) *Share
	GetOwner() *string
	SetProviderTenantId(v int64) *Share
	GetProviderTenantId() *int64
	SetShareId(v string) *Share
	GetShareId() *string
	SetShareName(v string) *Share
	GetShareName() *string
	SetUpdatedAt(v int64) *Share
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Share
	GetUpdatedBy() *string
}

type Share struct {
	Comment          *string `json:"comment,omitempty" xml:"comment,omitempty"`
	CreatedAt        *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy        *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	EnableWrite      *bool   `json:"enableWrite,omitempty" xml:"enableWrite,omitempty"`
	Owner            *string `json:"owner,omitempty" xml:"owner,omitempty"`
	ProviderTenantId *int64  `json:"providerTenantId,omitempty" xml:"providerTenantId,omitempty"`
	ShareId          *string `json:"shareId,omitempty" xml:"shareId,omitempty"`
	ShareName        *string `json:"shareName,omitempty" xml:"shareName,omitempty"`
	UpdatedAt        *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy        *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Share) String() string {
	return dara.Prettify(s)
}

func (s Share) GoString() string {
	return s.String()
}

func (s *Share) GetComment() *string {
	return s.Comment
}

func (s *Share) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Share) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Share) GetEnableWrite() *bool {
	return s.EnableWrite
}

func (s *Share) GetOwner() *string {
	return s.Owner
}

func (s *Share) GetProviderTenantId() *int64 {
	return s.ProviderTenantId
}

func (s *Share) GetShareId() *string {
	return s.ShareId
}

func (s *Share) GetShareName() *string {
	return s.ShareName
}

func (s *Share) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Share) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Share) SetComment(v string) *Share {
	s.Comment = &v
	return s
}

func (s *Share) SetCreatedAt(v int64) *Share {
	s.CreatedAt = &v
	return s
}

func (s *Share) SetCreatedBy(v string) *Share {
	s.CreatedBy = &v
	return s
}

func (s *Share) SetEnableWrite(v bool) *Share {
	s.EnableWrite = &v
	return s
}

func (s *Share) SetOwner(v string) *Share {
	s.Owner = &v
	return s
}

func (s *Share) SetProviderTenantId(v int64) *Share {
	s.ProviderTenantId = &v
	return s
}

func (s *Share) SetShareId(v string) *Share {
	s.ShareId = &v
	return s
}

func (s *Share) SetShareName(v string) *Share {
	s.ShareName = &v
	return s
}

func (s *Share) SetUpdatedAt(v int64) *Share {
	s.UpdatedAt = &v
	return s
}

func (s *Share) SetUpdatedBy(v string) *Share {
	s.UpdatedBy = &v
	return s
}

func (s *Share) Validate() error {
	return dara.Validate(s)
}
