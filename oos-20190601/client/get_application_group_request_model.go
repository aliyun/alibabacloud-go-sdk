// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *GetApplicationGroupRequest
	GetApplicationName() *string
	SetName(v string) *GetApplicationGroupRequest
	GetName() *string
	SetRegionId(v string) *GetApplicationGroupRequest
	GetRegionId() *string
}

type GetApplicationGroupRequest struct {
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
}

func (s GetApplicationGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetApplicationGroupRequest) GetName() *string {
	return s.Name
}

func (s *GetApplicationGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApplicationGroupRequest) SetApplicationName(v string) *GetApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationGroupRequest) SetName(v string) *GetApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *GetApplicationGroupRequest) SetRegionId(v string) *GetApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetApplicationGroupRequest) Validate() error {
	return dara.Validate(s)
}
