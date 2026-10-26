// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectAuthorizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationIds(v string) *DeleteProjectAuthorizationsRequest
	GetAuthorizationIds() *string
	SetInstanceId(v string) *DeleteProjectAuthorizationsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteProjectAuthorizationsRequest
	GetRegionId() *string
}

type DeleteProjectAuthorizationsRequest struct {
	// The IDs of the project authorizations to delete.
	//
	// > You can call the ListProjectAuthorizations operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	AuthorizationIds *string `json:"AuthorizationIds,omitempty" xml:"AuthorizationIds,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteProjectAuthorizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectAuthorizationsRequest) GetAuthorizationIds() *string {
	return s.AuthorizationIds
}

func (s *DeleteProjectAuthorizationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteProjectAuthorizationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteProjectAuthorizationsRequest) SetAuthorizationIds(v string) *DeleteProjectAuthorizationsRequest {
	s.AuthorizationIds = &v
	return s
}

func (s *DeleteProjectAuthorizationsRequest) SetInstanceId(v string) *DeleteProjectAuthorizationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteProjectAuthorizationsRequest) SetRegionId(v string) *DeleteProjectAuthorizationsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteProjectAuthorizationsRequest) Validate() error {
	return dara.Validate(s)
}
