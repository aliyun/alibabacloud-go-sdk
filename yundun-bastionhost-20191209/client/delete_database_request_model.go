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
	SetRegionId(v string) *DeleteDatabaseRequest
	GetRegionId() *string
}

type DeleteDatabaseRequest struct {
	// The ID of the database that you want to delete.
	//
	// > You can call the [ListDatabases](https://help.aliyun.com/document_detail/2758822.html) operation to query the database ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The ID of the bastion host from which you want to delete the database.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1ghxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *DeleteDatabaseRequest) SetRegionId(v string) *DeleteDatabaseRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
