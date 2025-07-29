// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantOption(v bool) *GrantPermissionRequest
	GetGrantOption() *bool
	SetGroupId(v string) *GrantPermissionRequest
	GetGroupId() *string
	SetNamespace(v string) *GrantPermissionRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GrantPermissionRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *GrantPermissionRequest
	GetRegionId() *string
	SetUserId(v string) *GrantPermissionRequest
	GetUserId() *string
	SetUserName(v string) *GrantPermissionRequest
	GetUserName() *string
}

type GrantPermissionRequest struct {
	// Specifies whether to grant the permissions with the GRANT option. Valid values: -**true*	- -**false**
	//
	// example:
	//
	// false
	GrantOption *bool `json:"GrantOption,omitempty" xml:"GrantOption,omitempty"`
	// The application group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffcdf01
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 277641081920123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// lilei
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GrantPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantPermissionRequest) GetGrantOption() *bool {
	return s.GrantOption
}

func (s *GrantPermissionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GrantPermissionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GrantPermissionRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GrantPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantPermissionRequest) GetUserId() *string {
	return s.UserId
}

func (s *GrantPermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *GrantPermissionRequest) SetGrantOption(v bool) *GrantPermissionRequest {
	s.GrantOption = &v
	return s
}

func (s *GrantPermissionRequest) SetGroupId(v string) *GrantPermissionRequest {
	s.GroupId = &v
	return s
}

func (s *GrantPermissionRequest) SetNamespace(v string) *GrantPermissionRequest {
	s.Namespace = &v
	return s
}

func (s *GrantPermissionRequest) SetNamespaceSource(v string) *GrantPermissionRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GrantPermissionRequest) SetRegionId(v string) *GrantPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *GrantPermissionRequest) SetUserId(v string) *GrantPermissionRequest {
	s.UserId = &v
	return s
}

func (s *GrantPermissionRequest) SetUserName(v string) *GrantPermissionRequest {
	s.UserName = &v
	return s
}

func (s *GrantPermissionRequest) Validate() error {
	return dara.Validate(s)
}
