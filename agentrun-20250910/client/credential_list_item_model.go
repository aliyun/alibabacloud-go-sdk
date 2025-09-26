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
	SetId(v string) *CredentialListItem
	GetId() *string
	SetName(v string) *CredentialListItem
	GetName() *string
	SetRelatedWorloads(v []*RelatedWorkload) *CredentialListItem
	GetRelatedWorloads() []*RelatedWorkload
	SetType(v string) *CredentialListItem
	GetType() *string
	SetUpdatedAt(v string) *CredentialListItem
	GetUpdatedAt() *string
}

type CredentialListItem struct {
	CreatedAt       *string            `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id              *string            `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string            `json:"name,omitempty" xml:"name,omitempty"`
	RelatedWorloads []*RelatedWorkload `json:"relatedWorloads,omitempty" xml:"relatedWorloads,omitempty" type:"Repeated"`
	Type            *string            `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt       *string            `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
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

func (s *CredentialListItem) GetId() *string {
	return s.Id
}

func (s *CredentialListItem) GetName() *string {
	return s.Name
}

func (s *CredentialListItem) GetRelatedWorloads() []*RelatedWorkload {
	return s.RelatedWorloads
}

func (s *CredentialListItem) GetType() *string {
	return s.Type
}

func (s *CredentialListItem) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CredentialListItem) SetCreatedAt(v string) *CredentialListItem {
	s.CreatedAt = &v
	return s
}

func (s *CredentialListItem) SetId(v string) *CredentialListItem {
	s.Id = &v
	return s
}

func (s *CredentialListItem) SetName(v string) *CredentialListItem {
	s.Name = &v
	return s
}

func (s *CredentialListItem) SetRelatedWorloads(v []*RelatedWorkload) *CredentialListItem {
	s.RelatedWorloads = v
	return s
}

func (s *CredentialListItem) SetType(v string) *CredentialListItem {
	s.Type = &v
	return s
}

func (s *CredentialListItem) SetUpdatedAt(v string) *CredentialListItem {
	s.UpdatedAt = &v
	return s
}

func (s *CredentialListItem) Validate() error {
	return dara.Validate(s)
}
