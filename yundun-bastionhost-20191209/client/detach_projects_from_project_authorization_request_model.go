// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachProjectsFromProjectAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v int64) *DetachProjectsFromProjectAuthorizationRequest
	GetAuthorizationId() *int64
	SetInstanceId(v string) *DetachProjectsFromProjectAuthorizationRequest
	GetInstanceId() *string
	SetProjectIds(v string) *DetachProjectsFromProjectAuthorizationRequest
	GetProjectIds() *string
	SetRegionId(v string) *DetachProjectsFromProjectAuthorizationRequest
	GetRegionId() *string
}

type DetachProjectsFromProjectAuthorizationRequest struct {
	// The project authorization ID.
	//
	// > You can call the ListProjectAuthorizations operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The array of project IDs to remove from the project authorization.
	//
	// > You can call the GetProjectAuthorization operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachProjectsFromProjectAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachProjectsFromProjectAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DetachProjectsFromProjectAuthorizationRequest) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *DetachProjectsFromProjectAuthorizationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachProjectsFromProjectAuthorizationRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *DetachProjectsFromProjectAuthorizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachProjectsFromProjectAuthorizationRequest) SetAuthorizationId(v int64) *DetachProjectsFromProjectAuthorizationRequest {
	s.AuthorizationId = &v
	return s
}

func (s *DetachProjectsFromProjectAuthorizationRequest) SetInstanceId(v string) *DetachProjectsFromProjectAuthorizationRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachProjectsFromProjectAuthorizationRequest) SetProjectIds(v string) *DetachProjectsFromProjectAuthorizationRequest {
	s.ProjectIds = &v
	return s
}

func (s *DetachProjectsFromProjectAuthorizationRequest) SetRegionId(v string) *DetachProjectsFromProjectAuthorizationRequest {
	s.RegionId = &v
	return s
}

func (s *DetachProjectsFromProjectAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
