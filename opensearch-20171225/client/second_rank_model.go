// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecondRank interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *SecondRank
	GetActive() *bool
	SetDescription(v string) *SecondRank
	GetDescription() *string
	SetMeta(v interface{}) *SecondRank
	GetMeta() interface{}
	SetName(v string) *SecondRank
	GetName() *string
}

type SecondRank struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SecondRank) String() string {
	return dara.Prettify(s)
}

func (s SecondRank) GoString() string {
	return s.String()
}

func (s *SecondRank) GetActive() *bool {
	return s.Active
}

func (s *SecondRank) GetDescription() *string {
	return s.Description
}

func (s *SecondRank) GetMeta() interface{} {
	return s.Meta
}

func (s *SecondRank) GetName() *string {
	return s.Name
}

func (s *SecondRank) SetActive(v bool) *SecondRank {
	s.Active = &v
	return s
}

func (s *SecondRank) SetDescription(v string) *SecondRank {
	s.Description = &v
	return s
}

func (s *SecondRank) SetMeta(v interface{}) *SecondRank {
	s.Meta = v
	return s
}

func (s *SecondRank) SetName(v string) *SecondRank {
	s.Name = &v
	return s
}

func (s *SecondRank) Validate() error {
	return dara.Validate(s)
}
