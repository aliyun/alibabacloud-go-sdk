// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerInfo interface {
	dara.Model
	String() string
	GoString() string
	SetMainContainer(v string) *ContainerInfo
	GetMainContainer() *string
	SetSidecarContainers(v []*string) *ContainerInfo
	GetSidecarContainers() []*string
}

type ContainerInfo struct {
	// The name of the main container.
	//
	// example:
	//
	// pytorch
	MainContainer *string `json:"MainContainer,omitempty" xml:"MainContainer,omitempty"`
	// The list of sidecar container names.
	SidecarContainers []*string `json:"SidecarContainers,omitempty" xml:"SidecarContainers,omitempty" type:"Repeated"`
}

func (s ContainerInfo) String() string {
	return dara.Prettify(s)
}

func (s ContainerInfo) GoString() string {
	return s.String()
}

func (s *ContainerInfo) GetMainContainer() *string {
	return s.MainContainer
}

func (s *ContainerInfo) GetSidecarContainers() []*string {
	return s.SidecarContainers
}

func (s *ContainerInfo) SetMainContainer(v string) *ContainerInfo {
	s.MainContainer = &v
	return s
}

func (s *ContainerInfo) SetSidecarContainers(v []*string) *ContainerInfo {
	s.SidecarContainers = v
	return s
}

func (s *ContainerInfo) Validate() error {
	return dara.Validate(s)
}
