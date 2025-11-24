// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceMeshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteServiceMeshRequest
	GetForce() *bool
	SetRetainResources(v string) *DeleteServiceMeshRequest
	GetRetainResources() *string
	SetServiceMeshId(v string) *DeleteServiceMeshRequest
	GetServiceMeshId() *string
}

type DeleteServiceMeshRequest struct {
	// Specifies whether to forcibly delete the ASM instance. Valid values:
	//
	// 	- `true`: forcibly deletes the ASM instance.
	//
	// 	- `false`: does not forcibly delete the ASM instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// A JSON string that can be parsed into a string array. You can use this JSON string to specify the IDs of the resource instances that need to be retained when the ASM instance is deleted.
	//
	// example:
	//
	// [" lb-bp1fxvl3q8akbj6m*****", "lb-bp1hoxkolggdw0y3*****"]
	RetainResources *string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteServiceMeshRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceMeshRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteServiceMeshRequest) GetRetainResources() *string {
	return s.RetainResources
}

func (s *DeleteServiceMeshRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DeleteServiceMeshRequest) SetForce(v bool) *DeleteServiceMeshRequest {
	s.Force = &v
	return s
}

func (s *DeleteServiceMeshRequest) SetRetainResources(v string) *DeleteServiceMeshRequest {
	s.RetainResources = &v
	return s
}

func (s *DeleteServiceMeshRequest) SetServiceMeshId(v string) *DeleteServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DeleteServiceMeshRequest) Validate() error {
	return dara.Validate(s)
}
