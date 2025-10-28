// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListComponentsResponseBody
	GetCode() *int32
	SetComponentList(v *ListComponentsResponseBodyComponentList) *ListComponentsResponseBody
	GetComponentList() *ListComponentsResponseBodyComponentList
	SetMessage(v string) *ListComponentsResponseBody
	GetMessage() *string
}

type ListComponentsResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The components.
	ComponentList *ListComponentsResponseBodyComponentList `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListComponentsResponseBody) GetComponentList() *ListComponentsResponseBodyComponentList {
	return s.ComponentList
}

func (s *ListComponentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListComponentsResponseBody) SetCode(v int32) *ListComponentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListComponentsResponseBody) SetComponentList(v *ListComponentsResponseBodyComponentList) *ListComponentsResponseBody {
	s.ComponentList = v
	return s
}

func (s *ListComponentsResponseBody) SetMessage(v string) *ListComponentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListComponentsResponseBody) Validate() error {
	if s.ComponentList != nil {
		if err := s.ComponentList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComponentsResponseBodyComponentList struct {
	Component []*ListComponentsResponseBodyComponentListComponent `json:"Component,omitempty" xml:"Component,omitempty" type:"Repeated"`
}

func (s ListComponentsResponseBodyComponentList) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyComponentList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponentList) GetComponent() []*ListComponentsResponseBodyComponentListComponent {
	return s.Component
}

func (s *ListComponentsResponseBodyComponentList) SetComponent(v []*ListComponentsResponseBodyComponentListComponent) *ListComponentsResponseBodyComponentList {
	s.Component = v
	return s
}

func (s *ListComponentsResponseBodyComponentList) Validate() error {
	if s.Component != nil {
		for _, item := range s.Component {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComponentsResponseBodyComponentListComponent struct {
	// The ID of the component.
	//
	// example:
	//
	// 1
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The key of the component.
	//
	// example:
	//
	// JDK 7
	ComponentKey *string `json:"ComponentKey,omitempty" xml:"ComponentKey,omitempty"`
	// The description of the component.
	//
	// example:
	//
	// JDK 7
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Indicates whether the component has expired. Valid values:
	//
	// 	- false: The component has not expired.
	//
	// 	- true: The component has expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The type of the component. Valid values:
	//
	// 	- JDK
	//
	// 	- TOMCAT
	//
	// 	- TENGINE
	//
	// example:
	//
	// JDK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the component.
	//
	// example:
	//
	// oraclejdk7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListComponentsResponseBodyComponentListComponent) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyComponentListComponent) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponentListComponent) GetComponentId() *string {
	return s.ComponentId
}

func (s *ListComponentsResponseBodyComponentListComponent) GetComponentKey() *string {
	return s.ComponentKey
}

func (s *ListComponentsResponseBodyComponentListComponent) GetDesc() *string {
	return s.Desc
}

func (s *ListComponentsResponseBodyComponentListComponent) GetExpired() *bool {
	return s.Expired
}

func (s *ListComponentsResponseBodyComponentListComponent) GetType() *string {
	return s.Type
}

func (s *ListComponentsResponseBodyComponentListComponent) GetVersion() *string {
	return s.Version
}

func (s *ListComponentsResponseBodyComponentListComponent) SetComponentId(v string) *ListComponentsResponseBodyComponentListComponent {
	s.ComponentId = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetComponentKey(v string) *ListComponentsResponseBodyComponentListComponent {
	s.ComponentKey = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetDesc(v string) *ListComponentsResponseBodyComponentListComponent {
	s.Desc = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetExpired(v bool) *ListComponentsResponseBodyComponentListComponent {
	s.Expired = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetType(v string) *ListComponentsResponseBodyComponentListComponent {
	s.Type = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetVersion(v string) *ListComponentsResponseBodyComponentListComponent {
	s.Version = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) Validate() error {
	return dara.Validate(s)
}
