// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *DeleteApplicationGroupRequest
	GetApplicationName() *string
	SetName(v string) *DeleteApplicationGroupRequest
	GetName() *string
	SetRegionId(v string) *DeleteApplicationGroupRequest
	GetRegionId() *string
	SetRetainResource(v bool) *DeleteApplicationGroupRequest
	GetRetainResource() *bool
}

type DeleteApplicationGroupRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region. Set the value to cn-hangzhou.
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

func (s DeleteApplicationGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DeleteApplicationGroupRequest) GetName() *string {
	return s.Name
}

func (s *DeleteApplicationGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApplicationGroupRequest) GetRetainResource() *bool {
	return s.RetainResource
}

func (s *DeleteApplicationGroupRequest) SetApplicationName(v string) *DeleteApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *DeleteApplicationGroupRequest) SetName(v string) *DeleteApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *DeleteApplicationGroupRequest) SetRegionId(v string) *DeleteApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApplicationGroupRequest) SetRetainResource(v bool) *DeleteApplicationGroupRequest {
	s.RetainResource = &v
	return s
}

func (s *DeleteApplicationGroupRequest) Validate() error {
	return dara.Validate(s)
}
