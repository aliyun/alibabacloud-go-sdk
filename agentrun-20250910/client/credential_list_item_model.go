// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialListItem interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *CredentialListItem
	GetCreatedAt() *string
	SetCredentialAuthType(v string) *CredentialListItem
	GetCredentialAuthType() *string
	SetCredentialId(v string) *CredentialListItem
	GetCredentialId() *string
	SetCredentialName(v string) *CredentialListItem
	GetCredentialName() *string
	SetCredentialSourceType(v string) *CredentialListItem
	GetCredentialSourceType() *string
	SetEnabled(v bool) *CredentialListItem
	GetEnabled() *bool
	SetRelatedResourceCount(v int32) *CredentialListItem
	GetRelatedResourceCount() *int32
	SetUpdatedAt(v string) *CredentialListItem
	GetUpdatedAt() *string
}

type CredentialListItem struct {
	CreatedAt            *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialAuthType   *string `json:"credentialAuthType,omitempty" xml:"credentialAuthType,omitempty"`
	CredentialId         *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	CredentialName       *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	CredentialSourceType *string `json:"credentialSourceType,omitempty" xml:"credentialSourceType,omitempty"`
	Enabled              *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	RelatedResourceCount *int32  `json:"relatedResourceCount,omitempty" xml:"relatedResourceCount,omitempty"`
	UpdatedAt            *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CredentialListItem) String() string {
	return dara.Prettify(s)
}

func (s CredentialListItem) GoString() string {
	return s.String()
}

func (s *CredentialListItem) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CredentialListItem) GetCredentialAuthType() *string {
	return s.CredentialAuthType
}

func (s *CredentialListItem) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CredentialListItem) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CredentialListItem) GetCredentialSourceType() *string {
	return s.CredentialSourceType
}

func (s *CredentialListItem) GetEnabled() *bool {
	return s.Enabled
}

func (s *CredentialListItem) GetRelatedResourceCount() *int32 {
	return s.RelatedResourceCount
}

func (s *CredentialListItem) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CredentialListItem) SetCreatedAt(v string) *CredentialListItem {
	s.CreatedAt = &v
	return s
}

func (s *CredentialListItem) SetCredentialAuthType(v string) *CredentialListItem {
	s.CredentialAuthType = &v
	return s
}

func (s *CredentialListItem) SetCredentialId(v string) *CredentialListItem {
	s.CredentialId = &v
	return s
}

func (s *CredentialListItem) SetCredentialName(v string) *CredentialListItem {
	s.CredentialName = &v
	return s
}

func (s *CredentialListItem) SetCredentialSourceType(v string) *CredentialListItem {
	s.CredentialSourceType = &v
	return s
}

func (s *CredentialListItem) SetEnabled(v bool) *CredentialListItem {
	s.Enabled = &v
	return s
}

func (s *CredentialListItem) SetRelatedResourceCount(v int32) *CredentialListItem {
	s.RelatedResourceCount = &v
	return s
}

func (s *CredentialListItem) SetUpdatedAt(v string) *CredentialListItem {
	s.UpdatedAt = &v
	return s
}

func (s *CredentialListItem) Validate() error {
	return dara.Validate(s)
}
