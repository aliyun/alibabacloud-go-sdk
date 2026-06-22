// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentInstanceSelector interface {
	dara.Model
	String() string
	GoString() string
	SetActionScope(v string) *ComponentInstanceSelector
	GetActionScope() *string
	SetApplicationName(v string) *ComponentInstanceSelector
	GetApplicationName() *string
	SetComponentInstances(v []*ComponentInstanceSelectorComponentInstances) *ComponentInstanceSelector
	GetComponentInstances() []*ComponentInstanceSelectorComponentInstances
	SetComponents(v []*ComponentInstanceSelectorComponents) *ComponentInstanceSelector
	GetComponents() []*ComponentInstanceSelectorComponents
	SetRunActionScope(v string) *ComponentInstanceSelector
	GetRunActionScope() *string
}

type ComponentInstanceSelector struct {
	// Deprecated
	//
	// The action scope. Valid values:
	//
	// - APPLICATION: The application level.
	//
	// - COMPONENT: The component level.
	//
	// - COMPONENT_INSTANCE: The component instance level.
	//
	// example:
	//
	// APPLICATION
	ActionScope *string `json:"ActionScope,omitempty" xml:"ActionScope,omitempty"`
	// The application name.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// A list of component instances. This parameter is used when `RunActionScope` is set to `COMPONENT_INSTANCE`.
	ComponentInstances []*ComponentInstanceSelectorComponentInstances `json:"ComponentInstances,omitempty" xml:"ComponentInstances,omitempty" type:"Repeated"`
	// A list of components. This parameter is used when `RunActionScope` is set to `COMPONENT`.
	Components []*ComponentInstanceSelectorComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The action scope. Valid values:
	//
	// - APPLICATION: The application level.
	//
	// - COMPONENT: The component level.
	//
	// - COMPONENT_INSTANCE: The component instance level.
	//
	// This parameter is required.
	//
	// example:
	//
	// APPLICATION
	RunActionScope *string `json:"RunActionScope,omitempty" xml:"RunActionScope,omitempty"`
}

func (s ComponentInstanceSelector) String() string {
	return dara.Prettify(s)
}

func (s ComponentInstanceSelector) GoString() string {
	return s.String()
}

func (s *ComponentInstanceSelector) GetActionScope() *string {
	return s.ActionScope
}

func (s *ComponentInstanceSelector) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ComponentInstanceSelector) GetComponentInstances() []*ComponentInstanceSelectorComponentInstances {
	return s.ComponentInstances
}

func (s *ComponentInstanceSelector) GetComponents() []*ComponentInstanceSelectorComponents {
	return s.Components
}

func (s *ComponentInstanceSelector) GetRunActionScope() *string {
	return s.RunActionScope
}

func (s *ComponentInstanceSelector) SetActionScope(v string) *ComponentInstanceSelector {
	s.ActionScope = &v
	return s
}

func (s *ComponentInstanceSelector) SetApplicationName(v string) *ComponentInstanceSelector {
	s.ApplicationName = &v
	return s
}

func (s *ComponentInstanceSelector) SetComponentInstances(v []*ComponentInstanceSelectorComponentInstances) *ComponentInstanceSelector {
	s.ComponentInstances = v
	return s
}

func (s *ComponentInstanceSelector) SetComponents(v []*ComponentInstanceSelectorComponents) *ComponentInstanceSelector {
	s.Components = v
	return s
}

func (s *ComponentInstanceSelector) SetRunActionScope(v string) *ComponentInstanceSelector {
	s.RunActionScope = &v
	return s
}

func (s *ComponentInstanceSelector) Validate() error {
	if s.ComponentInstances != nil {
		for _, item := range s.ComponentInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ComponentInstanceSelectorComponentInstances struct {
	// The application name.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The component name.
	//
	// example:
	//
	// DataNode
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The node ID.
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ComponentInstanceSelectorComponentInstances) String() string {
	return dara.Prettify(s)
}

func (s ComponentInstanceSelectorComponentInstances) GoString() string {
	return s.String()
}

func (s *ComponentInstanceSelectorComponentInstances) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ComponentInstanceSelectorComponentInstances) GetComponentName() *string {
	return s.ComponentName
}

func (s *ComponentInstanceSelectorComponentInstances) GetNodeId() *string {
	return s.NodeId
}

func (s *ComponentInstanceSelectorComponentInstances) SetApplicationName(v string) *ComponentInstanceSelectorComponentInstances {
	s.ApplicationName = &v
	return s
}

func (s *ComponentInstanceSelectorComponentInstances) SetComponentName(v string) *ComponentInstanceSelectorComponentInstances {
	s.ComponentName = &v
	return s
}

func (s *ComponentInstanceSelectorComponentInstances) SetNodeId(v string) *ComponentInstanceSelectorComponentInstances {
	s.NodeId = &v
	return s
}

func (s *ComponentInstanceSelectorComponentInstances) Validate() error {
	return dara.Validate(s)
}

type ComponentInstanceSelectorComponents struct {
	// The application name.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The component name.
	//
	// example:
	//
	// DataNode
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
}

func (s ComponentInstanceSelectorComponents) String() string {
	return dara.Prettify(s)
}

func (s ComponentInstanceSelectorComponents) GoString() string {
	return s.String()
}

func (s *ComponentInstanceSelectorComponents) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ComponentInstanceSelectorComponents) GetComponentName() *string {
	return s.ComponentName
}

func (s *ComponentInstanceSelectorComponents) SetApplicationName(v string) *ComponentInstanceSelectorComponents {
	s.ApplicationName = &v
	return s
}

func (s *ComponentInstanceSelectorComponents) SetComponentName(v string) *ComponentInstanceSelectorComponents {
	s.ComponentName = &v
	return s
}

func (s *ComponentInstanceSelectorComponents) Validate() error {
	return dara.Validate(s)
}
