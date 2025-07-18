// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpAttrMapConfigItem interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *IdpAttrMapConfigItem
	GetSource() *string
	SetTarget(v string) *IdpAttrMapConfigItem
	GetTarget() *string
	SetTargetType(v string) *IdpAttrMapConfigItem
	GetTargetType() *string
}

type IdpAttrMapConfigItem struct {
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Target     *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s IdpAttrMapConfigItem) String() string {
	return dara.Prettify(s)
}

func (s IdpAttrMapConfigItem) GoString() string {
	return s.String()
}

func (s *IdpAttrMapConfigItem) GetSource() *string {
	return s.Source
}

func (s *IdpAttrMapConfigItem) GetTarget() *string {
	return s.Target
}

func (s *IdpAttrMapConfigItem) GetTargetType() *string {
	return s.TargetType
}

func (s *IdpAttrMapConfigItem) SetSource(v string) *IdpAttrMapConfigItem {
	s.Source = &v
	return s
}

func (s *IdpAttrMapConfigItem) SetTarget(v string) *IdpAttrMapConfigItem {
	s.Target = &v
	return s
}

func (s *IdpAttrMapConfigItem) SetTargetType(v string) *IdpAttrMapConfigItem {
	s.TargetType = &v
	return s
}

func (s *IdpAttrMapConfigItem) Validate() error {
	return dara.Validate(s)
}
