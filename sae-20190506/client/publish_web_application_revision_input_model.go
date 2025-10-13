// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishWebApplicationRevisionInput interface {
	dara.Model
	String() string
	GoString() string
	SetContainers(v []*Container) *PublishWebApplicationRevisionInput
	GetContainers() []*Container
	SetDescription(v string) *PublishWebApplicationRevisionInput
	GetDescription() *string
	SetEnableArmsMetrics(v bool) *PublishWebApplicationRevisionInput
	GetEnableArmsMetrics() *bool
	SetTakeEffect(v bool) *PublishWebApplicationRevisionInput
	GetTakeEffect() *bool
}

type PublishWebApplicationRevisionInput struct {
	// This parameter is required.
	Containers []*Container `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// example:
	//
	// test version
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableArmsMetrics *bool   `json:"EnableArmsMetrics,omitempty" xml:"EnableArmsMetrics,omitempty"`
	TakeEffect        *bool   `json:"TakeEffect,omitempty" xml:"TakeEffect,omitempty"`
}

func (s PublishWebApplicationRevisionInput) String() string {
	return dara.Prettify(s)
}

func (s PublishWebApplicationRevisionInput) GoString() string {
	return s.String()
}

func (s *PublishWebApplicationRevisionInput) GetContainers() []*Container {
	return s.Containers
}

func (s *PublishWebApplicationRevisionInput) GetDescription() *string {
	return s.Description
}

func (s *PublishWebApplicationRevisionInput) GetEnableArmsMetrics() *bool {
	return s.EnableArmsMetrics
}

func (s *PublishWebApplicationRevisionInput) GetTakeEffect() *bool {
	return s.TakeEffect
}

func (s *PublishWebApplicationRevisionInput) SetContainers(v []*Container) *PublishWebApplicationRevisionInput {
	s.Containers = v
	return s
}

func (s *PublishWebApplicationRevisionInput) SetDescription(v string) *PublishWebApplicationRevisionInput {
	s.Description = &v
	return s
}

func (s *PublishWebApplicationRevisionInput) SetEnableArmsMetrics(v bool) *PublishWebApplicationRevisionInput {
	s.EnableArmsMetrics = &v
	return s
}

func (s *PublishWebApplicationRevisionInput) SetTakeEffect(v bool) *PublishWebApplicationRevisionInput {
	s.TakeEffect = &v
	return s
}

func (s *PublishWebApplicationRevisionInput) Validate() error {
	if s.Containers != nil {
		for _, item := range s.Containers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
