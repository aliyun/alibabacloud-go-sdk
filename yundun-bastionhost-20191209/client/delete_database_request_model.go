// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseId(v string) *DeleteDatabaseRequest
	GetDatabaseId() *string
	SetInstanceId(v string) *DeleteDatabaseRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DeleteDatabaseRequest
	GetProjectId() *int64
	SetRegionId(v string) *DeleteDatabaseRequest
	GetRegionId() *string
}

type DeleteDatabaseRequest struct {
	// The instance ID of the database that you want to delete.
	//
	// > You can invoke the [ListDatabases](https://help.aliyun.com/document_detail/2758822.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The instance ID of the bastion host where the database instance that you want to delete resides.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1ghxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *DeleteDatabaseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDatabaseRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDatabaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDatabaseRequest) SetDatabaseId(v string) *DeleteDatabaseRequest {
	s.DatabaseId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetInstanceId(v string) *DeleteDatabaseRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetProjectId(v int64) *DeleteDatabaseRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetRegionId(v string) *DeleteDatabaseRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
