// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *RevokePermissionRequest
	GetGroupId() *string
	SetNamespace(v string) *RevokePermissionRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *RevokePermissionRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *RevokePermissionRequest
	GetRegionId() *string
	SetUserId(v string) *RevokePermissionRequest
	GetUserId() *string
}

type RevokePermissionRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.defalutGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The unique identifier (UID) of the namespace. You can obtain the namespace UID on the Namespace page in the SchedulerX console.
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
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UID of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 277641081920123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokePermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokePermissionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RevokePermissionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *RevokePermissionRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *RevokePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokePermissionRequest) GetUserId() *string {
	return s.UserId
}

func (s *RevokePermissionRequest) SetGroupId(v string) *RevokePermissionRequest {
	s.GroupId = &v
	return s
}

func (s *RevokePermissionRequest) SetNamespace(v string) *RevokePermissionRequest {
	s.Namespace = &v
	return s
}

func (s *RevokePermissionRequest) SetNamespaceSource(v string) *RevokePermissionRequest {
	s.NamespaceSource = &v
	return s
}

func (s *RevokePermissionRequest) SetRegionId(v string) *RevokePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *RevokePermissionRequest) SetUserId(v string) *RevokePermissionRequest {
	s.UserId = &v
	return s
}

func (s *RevokePermissionRequest) Validate() error {
	return dara.Validate(s)
}
