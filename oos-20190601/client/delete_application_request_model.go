// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteApplicationRequest
	GetForce() *bool
	SetName(v string) *DeleteApplicationRequest
	GetName() *string
	SetRegionId(v string) *DeleteApplicationRequest
	GetRegionId() *string
	SetRetainResource(v bool) *DeleteApplicationRequest
	GetRetainResource() *bool
}

type DeleteApplicationRequest struct {
	// Specifies whether to forcibly delete the application. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// False
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to retain resources created by application manager when deleting the application. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	RetainResource *bool `json:"RetainResource,omitempty" xml:"RetainResource,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteApplicationRequest) GetName() *string {
	return s.Name
}

func (s *DeleteApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApplicationRequest) GetRetainResource() *bool {
	return s.RetainResource
}

func (s *DeleteApplicationRequest) SetForce(v bool) *DeleteApplicationRequest {
	s.Force = &v
	return s
}

func (s *DeleteApplicationRequest) SetName(v string) *DeleteApplicationRequest {
	s.Name = &v
	return s
}

func (s *DeleteApplicationRequest) SetRegionId(v string) *DeleteApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApplicationRequest) SetRetainResource(v bool) *DeleteApplicationRequest {
	s.RetainResource = &v
	return s
}

func (s *DeleteApplicationRequest) Validate() error {
	return dara.Validate(s)
}
