// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateEdgeContainerAppVersionShrinkRequest
	GetAppId() *string
	SetContainersShrink(v string) *CreateEdgeContainerAppVersionShrinkRequest
	GetContainersShrink() *string
	SetName(v string) *CreateEdgeContainerAppVersionShrinkRequest
	GetName() *string
	SetRemarks(v string) *CreateEdgeContainerAppVersionShrinkRequest
	GetRemarks() *string
}

type CreateEdgeContainerAppVersionShrinkRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	ContainersShrink *string `json:"Containers,omitempty" xml:"Containers,omitempty"`
	// This parameter is required.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
}

func (s CreateEdgeContainerAppVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) GetContainersShrink() *string {
	return s.ContainersShrink
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) SetAppId(v string) *CreateEdgeContainerAppVersionShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) SetContainersShrink(v string) *CreateEdgeContainerAppVersionShrinkRequest {
	s.ContainersShrink = &v
	return s
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) SetName(v string) *CreateEdgeContainerAppVersionShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) SetRemarks(v string) *CreateEdgeContainerAppVersionShrinkRequest {
	s.Remarks = &v
	return s
}

func (s *CreateEdgeContainerAppVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
