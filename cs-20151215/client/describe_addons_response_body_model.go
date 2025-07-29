// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentGroups(v []*DescribeAddonsResponseBodyComponentGroups) *DescribeAddonsResponseBody
	GetComponentGroups() []*DescribeAddonsResponseBodyComponentGroups
	SetStandardComponents(v map[string]*StandardComponentsValue) *DescribeAddonsResponseBody
	GetStandardComponents() map[string]*StandardComponentsValue
}

type DescribeAddonsResponseBody struct {
	// The list of the returned components.
	ComponentGroups []*DescribeAddonsResponseBodyComponentGroups `json:"ComponentGroups,omitempty" xml:"ComponentGroups,omitempty" type:"Repeated"`
	// Standard components.
	StandardComponents map[string]*StandardComponentsValue `json:"StandardComponents,omitempty" xml:"StandardComponents,omitempty"`
}

func (s DescribeAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBody) GetComponentGroups() []*DescribeAddonsResponseBodyComponentGroups {
	return s.ComponentGroups
}

func (s *DescribeAddonsResponseBody) GetStandardComponents() map[string]*StandardComponentsValue {
	return s.StandardComponents
}

func (s *DescribeAddonsResponseBody) SetComponentGroups(v []*DescribeAddonsResponseBodyComponentGroups) *DescribeAddonsResponseBody {
	s.ComponentGroups = v
	return s
}

func (s *DescribeAddonsResponseBody) SetStandardComponents(v map[string]*StandardComponentsValue) *DescribeAddonsResponseBody {
	s.StandardComponents = v
	return s
}

func (s *DescribeAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAddonsResponseBodyComponentGroups struct {
	// The name of the component group.
	//
	// example:
	//
	// storage
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// The names of the components in the component group.
	Items []*DescribeAddonsResponseBodyComponentGroupsItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s DescribeAddonsResponseBodyComponentGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonsResponseBodyComponentGroups) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyComponentGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeAddonsResponseBodyComponentGroups) GetItems() []*DescribeAddonsResponseBodyComponentGroupsItems {
	return s.Items
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetGroupName(v string) *DescribeAddonsResponseBodyComponentGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetItems(v []*DescribeAddonsResponseBodyComponentGroupsItems) *DescribeAddonsResponseBodyComponentGroups {
	s.Items = v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeAddonsResponseBodyComponentGroupsItems struct {
	// The name of the component.
	//
	// example:
	//
	// flexvolume
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAddonsResponseBodyComponentGroupsItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonsResponseBodyComponentGroupsItems) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) GetName() *string {
	return s.Name
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetName(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Name = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) Validate() error {
	return dara.Validate(s)
}
