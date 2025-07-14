// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmenu interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*SubmenuItems) *Submenu
	GetItems() []*SubmenuItems
	SetSubmenuDesc(v string) *Submenu
	GetSubmenuDesc() *string
	SetSubmenuType(v string) *Submenu
	GetSubmenuType() *string
	SetSubmenus(v []*Submenu) *Submenu
	GetSubmenus() []*Submenu
}

type Submenu struct {
	Items       []*SubmenuItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	SubmenuDesc *string         `json:"SubmenuDesc,omitempty" xml:"SubmenuDesc,omitempty"`
	SubmenuType *string         `json:"SubmenuType,omitempty" xml:"SubmenuType,omitempty"`
	Submenus    []*Submenu      `json:"Submenus,omitempty" xml:"Submenus,omitempty" type:"Repeated"`
}

func (s Submenu) String() string {
	return dara.Prettify(s)
}

func (s Submenu) GoString() string {
	return s.String()
}

func (s *Submenu) GetItems() []*SubmenuItems {
	return s.Items
}

func (s *Submenu) GetSubmenuDesc() *string {
	return s.SubmenuDesc
}

func (s *Submenu) GetSubmenuType() *string {
	return s.SubmenuType
}

func (s *Submenu) GetSubmenus() []*Submenu {
	return s.Submenus
}

func (s *Submenu) SetItems(v []*SubmenuItems) *Submenu {
	s.Items = v
	return s
}

func (s *Submenu) SetSubmenuDesc(v string) *Submenu {
	s.SubmenuDesc = &v
	return s
}

func (s *Submenu) SetSubmenuType(v string) *Submenu {
	s.SubmenuType = &v
	return s
}

func (s *Submenu) SetSubmenus(v []*Submenu) *Submenu {
	s.Submenus = v
	return s
}

func (s *Submenu) Validate() error {
	return dara.Validate(s)
}

type SubmenuItems struct {
	DefaultSelected *bool     `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	ItemDesc        *string   `json:"ItemDesc,omitempty" xml:"ItemDesc,omitempty"`
	ItemType        *string   `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	RelatingItems   []*string `json:"RelatingItems,omitempty" xml:"RelatingItems,omitempty" type:"Repeated"`
}

func (s SubmenuItems) String() string {
	return dara.Prettify(s)
}

func (s SubmenuItems) GoString() string {
	return s.String()
}

func (s *SubmenuItems) GetDefaultSelected() *bool {
	return s.DefaultSelected
}

func (s *SubmenuItems) GetItemDesc() *string {
	return s.ItemDesc
}

func (s *SubmenuItems) GetItemType() *string {
	return s.ItemType
}

func (s *SubmenuItems) GetRelatingItems() []*string {
	return s.RelatingItems
}

func (s *SubmenuItems) SetDefaultSelected(v bool) *SubmenuItems {
	s.DefaultSelected = &v
	return s
}

func (s *SubmenuItems) SetItemDesc(v string) *SubmenuItems {
	s.ItemDesc = &v
	return s
}

func (s *SubmenuItems) SetItemType(v string) *SubmenuItems {
	s.ItemType = &v
	return s
}

func (s *SubmenuItems) SetRelatingItems(v []*string) *SubmenuItems {
	s.RelatingItems = v
	return s
}

func (s *SubmenuItems) Validate() error {
	return dara.Validate(s)
}
