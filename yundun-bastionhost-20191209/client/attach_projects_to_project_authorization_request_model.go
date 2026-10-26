// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachProjectsToProjectAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v int64) *AttachProjectsToProjectAuthorizationRequest
	GetAuthorizationId() *int64
	SetInstanceId(v string) *AttachProjectsToProjectAuthorizationRequest
	GetInstanceId() *string
	SetProjectIds(v string) *AttachProjectsToProjectAuthorizationRequest
	GetProjectIds() *string
	SetRegionId(v string) *AttachProjectsToProjectAuthorizationRequest
	GetRegionId() *string
}

type AttachProjectsToProjectAuthorizationRequest struct {
	// The ID of the project authorization.
	//
	// > You can call the ListProjectAuthorizations operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The instance ID of Bastionhost.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the projects that can be managed.
	//
	// > You can call the ListProjects operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachProjectsToProjectAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachProjectsToProjectAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *AttachProjectsToProjectAuthorizationRequest) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *AttachProjectsToProjectAuthorizationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachProjectsToProjectAuthorizationRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *AttachProjectsToProjectAuthorizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachProjectsToProjectAuthorizationRequest) SetAuthorizationId(v int64) *AttachProjectsToProjectAuthorizationRequest {
	s.AuthorizationId = &v
	return s
}

func (s *AttachProjectsToProjectAuthorizationRequest) SetInstanceId(v string) *AttachProjectsToProjectAuthorizationRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachProjectsToProjectAuthorizationRequest) SetProjectIds(v string) *AttachProjectsToProjectAuthorizationRequest {
	s.ProjectIds = &v
	return s
}

func (s *AttachProjectsToProjectAuthorizationRequest) SetRegionId(v string) *AttachProjectsToProjectAuthorizationRequest {
	s.RegionId = &v
	return s
}

func (s *AttachProjectsToProjectAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
