// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseId(v string) *GetDatabaseRequest
	GetDatabaseId() *string
	SetInstanceId(v string) *GetDatabaseRequest
	GetInstanceId() *string
	SetProjectId(v int64) *GetDatabaseRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetDatabaseRequest
	GetRegionId() *string
}

type GetDatabaseRequest struct {
	// The ID of the database to query.
	//
	// > You can call the [ListDatabases](https://help.aliyun.com/document_detail/2758822.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The instance ID of the bastion host where the database resides.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-wwo36qbv601
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host where the database resides.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *GetDatabaseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDatabaseRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDatabaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDatabaseRequest) SetDatabaseId(v string) *GetDatabaseRequest {
	s.DatabaseId = &v
	return s
}

func (s *GetDatabaseRequest) SetInstanceId(v string) *GetDatabaseRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDatabaseRequest) SetProjectId(v int64) *GetDatabaseRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDatabaseRequest) SetRegionId(v string) *GetDatabaseRequest {
	s.RegionId = &v
	return s
}

func (s *GetDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
