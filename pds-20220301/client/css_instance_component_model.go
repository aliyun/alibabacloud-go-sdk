// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCssInstanceComponent interface {
	dara.Model
	String() string
	GoString() string
	SetComponentCode(v string) *CssInstanceComponent
	GetComponentCode() *string
	SetComponentName(v string) *CssInstanceComponent
	GetComponentName() *string
	SetGlobalKey(v string) *CssInstanceComponent
	GetGlobalKey() *string
	SetInstanceProperty(v []*CssInstanceProperty) *CssInstanceComponent
	GetInstanceProperty() []*CssInstanceProperty
	SetModuleAttrStatus(v int64) *CssInstanceComponent
	GetModuleAttrStatus() *int64
	SetTag(v string) *CssInstanceComponent
	GetTag() *string
}

type CssInstanceComponent struct {
	ComponentCode    *string                `json:"componentCode,omitempty" xml:"componentCode,omitempty"`
	ComponentName    *string                `json:"componentName,omitempty" xml:"componentName,omitempty"`
	GlobalKey        *string                `json:"globalKey,omitempty" xml:"globalKey,omitempty"`
	InstanceProperty []*CssInstanceProperty `json:"instanceProperty,omitempty" xml:"instanceProperty,omitempty" type:"Repeated"`
	ModuleAttrStatus *int64                 `json:"moduleAttrStatus,omitempty" xml:"moduleAttrStatus,omitempty"`
	Tag              *string                `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s CssInstanceComponent) String() string {
	return dara.Prettify(s)
}

func (s CssInstanceComponent) GoString() string {
	return s.String()
}

func (s *CssInstanceComponent) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *CssInstanceComponent) GetComponentName() *string {
	return s.ComponentName
}

func (s *CssInstanceComponent) GetGlobalKey() *string {
	return s.GlobalKey
}

func (s *CssInstanceComponent) GetInstanceProperty() []*CssInstanceProperty {
	return s.InstanceProperty
}

func (s *CssInstanceComponent) GetModuleAttrStatus() *int64 {
	return s.ModuleAttrStatus
}

func (s *CssInstanceComponent) GetTag() *string {
	return s.Tag
}

func (s *CssInstanceComponent) SetComponentCode(v string) *CssInstanceComponent {
	s.ComponentCode = &v
	return s
}

func (s *CssInstanceComponent) SetComponentName(v string) *CssInstanceComponent {
	s.ComponentName = &v
	return s
}

func (s *CssInstanceComponent) SetGlobalKey(v string) *CssInstanceComponent {
	s.GlobalKey = &v
	return s
}

func (s *CssInstanceComponent) SetInstanceProperty(v []*CssInstanceProperty) *CssInstanceComponent {
	s.InstanceProperty = v
	return s
}

func (s *CssInstanceComponent) SetModuleAttrStatus(v int64) *CssInstanceComponent {
	s.ModuleAttrStatus = &v
	return s
}

func (s *CssInstanceComponent) SetTag(v string) *CssInstanceComponent {
	s.Tag = &v
	return s
}

func (s *CssInstanceComponent) Validate() error {
	if s.InstanceProperty != nil {
		for _, item := range s.InstanceProperty {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
