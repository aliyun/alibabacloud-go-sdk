// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpAttrMapConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMapItems(v []*IdpAttrMapConfigItem) *IdpAttrMapConfig
	GetMapItems() []*IdpAttrMapConfigItem
}

type IdpAttrMapConfig struct {
	MapItems []*IdpAttrMapConfigItem `json:"MapItems,omitempty" xml:"MapItems,omitempty" type:"Repeated"`
}

func (s IdpAttrMapConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpAttrMapConfig) GoString() string {
	return s.String()
}

func (s *IdpAttrMapConfig) GetMapItems() []*IdpAttrMapConfigItem {
	return s.MapItems
}

func (s *IdpAttrMapConfig) SetMapItems(v []*IdpAttrMapConfigItem) *IdpAttrMapConfig {
	s.MapItems = v
	return s
}

func (s *IdpAttrMapConfig) Validate() error {
	return dara.Validate(s)
}
