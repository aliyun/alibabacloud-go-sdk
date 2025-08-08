// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceivedShare interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ReceivedShare
	GetCatalogName() *string
	SetComment(v string) *ReceivedShare
	GetComment() *string
	SetCreatedAt(v int64) *ReceivedShare
	GetCreatedAt() *int64
	SetCreatedBy(v string) *ReceivedShare
	GetCreatedBy() *string
	SetOwner(v string) *ReceivedShare
	GetOwner() *string
	SetProviderTenantId(v int64) *ReceivedShare
	GetProviderTenantId() *int64
	SetShareId(v string) *ReceivedShare
	GetShareId() *string
	SetShareName(v string) *ReceivedShare
	GetShareName() *string
	SetUpdatedAt(v int64) *ReceivedShare
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *ReceivedShare
	GetUpdatedBy() *string
}

type ReceivedShare struct {
	CatalogName      *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	Comment          *string `json:"comment,omitempty" xml:"comment,omitempty"`
	CreatedAt        *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy        *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Owner            *string `json:"owner,omitempty" xml:"owner,omitempty"`
	ProviderTenantId *int64  `json:"providerTenantId,omitempty" xml:"providerTenantId,omitempty"`
	ShareId          *string `json:"shareId,omitempty" xml:"shareId,omitempty"`
	ShareName        *string `json:"shareName,omitempty" xml:"shareName,omitempty"`
	UpdatedAt        *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy        *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s ReceivedShare) String() string {
	return dara.Prettify(s)
}

func (s ReceivedShare) GoString() string {
	return s.String()
}

func (s *ReceivedShare) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ReceivedShare) GetComment() *string {
	return s.Comment
}

func (s *ReceivedShare) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ReceivedShare) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ReceivedShare) GetOwner() *string {
	return s.Owner
}

func (s *ReceivedShare) GetProviderTenantId() *int64 {
	return s.ProviderTenantId
}

func (s *ReceivedShare) GetShareId() *string {
	return s.ShareId
}

func (s *ReceivedShare) GetShareName() *string {
	return s.ShareName
}

func (s *ReceivedShare) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ReceivedShare) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ReceivedShare) SetCatalogName(v string) *ReceivedShare {
	s.CatalogName = &v
	return s
}

func (s *ReceivedShare) SetComment(v string) *ReceivedShare {
	s.Comment = &v
	return s
}

func (s *ReceivedShare) SetCreatedAt(v int64) *ReceivedShare {
	s.CreatedAt = &v
	return s
}

func (s *ReceivedShare) SetCreatedBy(v string) *ReceivedShare {
	s.CreatedBy = &v
	return s
}

func (s *ReceivedShare) SetOwner(v string) *ReceivedShare {
	s.Owner = &v
	return s
}

func (s *ReceivedShare) SetProviderTenantId(v int64) *ReceivedShare {
	s.ProviderTenantId = &v
	return s
}

func (s *ReceivedShare) SetShareId(v string) *ReceivedShare {
	s.ShareId = &v
	return s
}

func (s *ReceivedShare) SetShareName(v string) *ReceivedShare {
	s.ShareName = &v
	return s
}

func (s *ReceivedShare) SetUpdatedAt(v int64) *ReceivedShare {
	s.UpdatedAt = &v
	return s
}

func (s *ReceivedShare) SetUpdatedBy(v string) *ReceivedShare {
	s.UpdatedBy = &v
	return s
}

func (s *ReceivedShare) Validate() error {
	return dara.Validate(s)
}
