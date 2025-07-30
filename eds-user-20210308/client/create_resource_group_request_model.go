// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsResourceGroupWithOfficeSite(v int64) *CreateResourceGroupRequest
	GetIsResourceGroupWithOfficeSite() *int64
	SetPlatform(v string) *CreateResourceGroupRequest
	GetPlatform() *string
	SetResourceGroupName(v string) *CreateResourceGroupRequest
	GetResourceGroupName() *string
}

type CreateResourceGroupRequest struct {
	// example:
	//
	// 0
	IsResourceGroupWithOfficeSite *int64 `json:"IsResourceGroupWithOfficeSite,omitempty" xml:"IsResourceGroupWithOfficeSite,omitempty"`
	// example:
	//
	// AliyunConsole
	Platform          *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s CreateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) GetIsResourceGroupWithOfficeSite() *int64 {
	return s.IsResourceGroupWithOfficeSite
}

func (s *CreateResourceGroupRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateResourceGroupRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *CreateResourceGroupRequest) SetIsResourceGroupWithOfficeSite(v int64) *CreateResourceGroupRequest {
	s.IsResourceGroupWithOfficeSite = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPlatform(v string) *CreateResourceGroupRequest {
	s.Platform = &v
	return s
}

func (s *CreateResourceGroupRequest) SetResourceGroupName(v string) *CreateResourceGroupRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
