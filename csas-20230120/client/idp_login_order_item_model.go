// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpLoginOrderItem interface {
	dara.Model
	String() string
	GoString() string
	SetClass(v string) *IdpLoginOrderItem
	GetClass() *string
	SetConfigId(v string) *IdpLoginOrderItem
	GetConfigId() *string
	SetDesc(v string) *IdpLoginOrderItem
	GetDesc() *string
	SetEnabled(v bool) *IdpLoginOrderItem
	GetEnabled() *bool
	SetName(v string) *IdpLoginOrderItem
	GetName() *string
	SetType(v string) *IdpLoginOrderItem
	GetType() *string
}

type IdpLoginOrderItem struct {
	Class    *string `json:"Class,omitempty" xml:"Class,omitempty"`
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Desc     *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Enabled  *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s IdpLoginOrderItem) String() string {
	return dara.Prettify(s)
}

func (s IdpLoginOrderItem) GoString() string {
	return s.String()
}

func (s *IdpLoginOrderItem) GetClass() *string {
	return s.Class
}

func (s *IdpLoginOrderItem) GetConfigId() *string {
	return s.ConfigId
}

func (s *IdpLoginOrderItem) GetDesc() *string {
	return s.Desc
}

func (s *IdpLoginOrderItem) GetEnabled() *bool {
	return s.Enabled
}

func (s *IdpLoginOrderItem) GetName() *string {
	return s.Name
}

func (s *IdpLoginOrderItem) GetType() *string {
	return s.Type
}

func (s *IdpLoginOrderItem) SetClass(v string) *IdpLoginOrderItem {
	s.Class = &v
	return s
}

func (s *IdpLoginOrderItem) SetConfigId(v string) *IdpLoginOrderItem {
	s.ConfigId = &v
	return s
}

func (s *IdpLoginOrderItem) SetDesc(v string) *IdpLoginOrderItem {
	s.Desc = &v
	return s
}

func (s *IdpLoginOrderItem) SetEnabled(v bool) *IdpLoginOrderItem {
	s.Enabled = &v
	return s
}

func (s *IdpLoginOrderItem) SetName(v string) *IdpLoginOrderItem {
	s.Name = &v
	return s
}

func (s *IdpLoginOrderItem) SetType(v string) *IdpLoginOrderItem {
	s.Type = &v
	return s
}

func (s *IdpLoginOrderItem) Validate() error {
	return dara.Validate(s)
}
