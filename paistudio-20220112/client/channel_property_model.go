// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelProperty interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ChannelProperty
	GetName() *string
	SetValue(v string) *ChannelProperty
	GetValue() *string
}

type ChannelProperty struct {
	// This parameter is required.
	//
	// example:
	//
	// SKlearn
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Framework
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ChannelProperty) String() string {
	return dara.Prettify(s)
}

func (s ChannelProperty) GoString() string {
	return s.String()
}

func (s *ChannelProperty) GetName() *string {
	return s.Name
}

func (s *ChannelProperty) GetValue() *string {
	return s.Value
}

func (s *ChannelProperty) SetName(v string) *ChannelProperty {
	s.Name = &v
	return s
}

func (s *ChannelProperty) SetValue(v string) *ChannelProperty {
	s.Value = &v
	return s
}

func (s *ChannelProperty) Validate() error {
	return dara.Validate(s)
}
