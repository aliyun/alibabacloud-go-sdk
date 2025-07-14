// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceCodeAccount interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarUrl(v string) *SourceCodeAccount
	GetAvatarUrl() *string
	SetId(v string) *SourceCodeAccount
	GetId() *string
	SetName(v string) *SourceCodeAccount
	GetName() *string
	SetOrganizations(v []*SourceCodeAccountOrganizations) *SourceCodeAccount
	GetOrganizations() []*SourceCodeAccountOrganizations
}

type SourceCodeAccount struct {
	AvatarUrl     *string                           `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Id            *string                           `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Organizations []*SourceCodeAccountOrganizations `json:"Organizations,omitempty" xml:"Organizations,omitempty" type:"Repeated"`
}

func (s SourceCodeAccount) String() string {
	return dara.Prettify(s)
}

func (s SourceCodeAccount) GoString() string {
	return s.String()
}

func (s *SourceCodeAccount) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *SourceCodeAccount) GetId() *string {
	return s.Id
}

func (s *SourceCodeAccount) GetName() *string {
	return s.Name
}

func (s *SourceCodeAccount) GetOrganizations() []*SourceCodeAccountOrganizations {
	return s.Organizations
}

func (s *SourceCodeAccount) SetAvatarUrl(v string) *SourceCodeAccount {
	s.AvatarUrl = &v
	return s
}

func (s *SourceCodeAccount) SetId(v string) *SourceCodeAccount {
	s.Id = &v
	return s
}

func (s *SourceCodeAccount) SetName(v string) *SourceCodeAccount {
	s.Name = &v
	return s
}

func (s *SourceCodeAccount) SetOrganizations(v []*SourceCodeAccountOrganizations) *SourceCodeAccount {
	s.Organizations = v
	return s
}

func (s *SourceCodeAccount) Validate() error {
	return dara.Validate(s)
}

type SourceCodeAccountOrganizations struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SourceCodeAccountOrganizations) String() string {
	return dara.Prettify(s)
}

func (s SourceCodeAccountOrganizations) GoString() string {
	return s.String()
}

func (s *SourceCodeAccountOrganizations) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *SourceCodeAccountOrganizations) GetId() *string {
	return s.Id
}

func (s *SourceCodeAccountOrganizations) GetName() *string {
	return s.Name
}

func (s *SourceCodeAccountOrganizations) SetAvatarUrl(v string) *SourceCodeAccountOrganizations {
	s.AvatarUrl = &v
	return s
}

func (s *SourceCodeAccountOrganizations) SetId(v string) *SourceCodeAccountOrganizations {
	s.Id = &v
	return s
}

func (s *SourceCodeAccountOrganizations) SetName(v string) *SourceCodeAccountOrganizations {
	s.Name = &v
	return s
}

func (s *SourceCodeAccountOrganizations) Validate() error {
	return dara.Validate(s)
}
