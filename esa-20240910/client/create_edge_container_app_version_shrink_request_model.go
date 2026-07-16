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
	// The application ID. You can call the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation to obtain the application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The container group to deploy for this version, including specific image information. The image information consists of the image address, startup commands, parameters, environment variables, and probe rules. Multiple images are supported. This parameter is a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "Name": "container1",
	//
	//             "Image": "image1",
	//
	//             "Spec": "1C2G",
	//
	//             "Command": "/bin/sh",
	//
	//             "Args": "-c hello",
	//
	//             "ProbeType": "tcpSocket",
	//
	//             "ProbeContent": "{\\"Port\\":8080}"
	//
	//       },
	//
	//       {
	//
	//             "Name": "container2",
	//
	//             "Image": "image2",
	//
	//             "Spec": "2C4G",
	//
	//             "ProbeType": "httpGet",
	//
	//             "ProbeContent": "{\\"Path\\":\\"/\\",\\"Port\\":80,\\"InitialDelaySeconds\\":10}"
	//
	//       }
	//
	// ]
	ContainersShrink *string `json:"Containers,omitempty" xml:"Containers,omitempty"`
	// The version name. The name must be **6 to 128*	- characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// verson1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test app
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
