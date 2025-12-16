// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFirstRank interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *FirstRank
	GetActive() *bool
	SetDescription(v string) *FirstRank
	GetDescription() *string
	SetMeta(v interface{}) *FirstRank
	GetMeta() interface{}
	SetName(v string) *FirstRank
	GetName() *string
	SetType(v string) *FirstRank
	GetType() *string
}

type FirstRank struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
	Type        *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FirstRank) String() string {
	return dara.Prettify(s)
}

func (s FirstRank) GoString() string {
	return s.String()
}

func (s *FirstRank) GetActive() *bool {
	return s.Active
}

func (s *FirstRank) GetDescription() *string {
	return s.Description
}

func (s *FirstRank) GetMeta() interface{} {
	return s.Meta
}

func (s *FirstRank) GetName() *string {
	return s.Name
}

func (s *FirstRank) GetType() *string {
	return s.Type
}

func (s *FirstRank) SetActive(v bool) *FirstRank {
	s.Active = &v
	return s
}

func (s *FirstRank) SetDescription(v string) *FirstRank {
	s.Description = &v
	return s
}

func (s *FirstRank) SetMeta(v interface{}) *FirstRank {
	s.Meta = v
	return s
}

func (s *FirstRank) SetName(v string) *FirstRank {
	s.Name = &v
	return s
}

func (s *FirstRank) SetType(v string) *FirstRank {
	s.Type = &v
	return s
}

func (s *FirstRank) Validate() error {
	return dara.Validate(s)
}
