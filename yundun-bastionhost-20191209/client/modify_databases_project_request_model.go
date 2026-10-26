// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabasesProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseIds(v string) *ModifyDatabasesProjectRequest
	GetDatabaseIds() *string
	SetInstanceId(v string) *ModifyDatabasesProjectRequest
	GetInstanceId() *string
	SetProjectId(v int64) *ModifyDatabasesProjectRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyDatabasesProjectRequest
	GetRegionId() *string
	SetTargetProjectId(v int64) *ModifyDatabasesProjectRequest
	GetTargetProjectId() *int64
}

type ModifyDatabasesProjectRequest struct {
	// The IDs of the database instances that you want to move. The value is a JSON string. You can specify up to 100 database instance IDs. Separate multiple IDs with commas (,).
	//
	// > You can call the [ListDatabases](https://help.aliyun.com/document_detail/2758822.html) operation to query the database instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2,3
	DatabaseIds *string `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zmb2y9ydw08
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the project to which the database instances currently belong.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the destination project to which you want to move the database instances.
	//
	// > You can call the ListProjectsForProjectAuthorizationPrincipal operation to query the project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s ModifyDatabasesProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabasesProjectRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabasesProjectRequest) GetDatabaseIds() *string {
	return s.DatabaseIds
}

func (s *ModifyDatabasesProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDatabasesProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ModifyDatabasesProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDatabasesProjectRequest) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *ModifyDatabasesProjectRequest) SetDatabaseIds(v string) *ModifyDatabasesProjectRequest {
	s.DatabaseIds = &v
	return s
}

func (s *ModifyDatabasesProjectRequest) SetInstanceId(v string) *ModifyDatabasesProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabasesProjectRequest) SetProjectId(v int64) *ModifyDatabasesProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyDatabasesProjectRequest) SetRegionId(v string) *ModifyDatabasesProjectRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDatabasesProjectRequest) SetTargetProjectId(v int64) *ModifyDatabasesProjectRequest {
	s.TargetProjectId = &v
	return s
}

func (s *ModifyDatabasesProjectRequest) Validate() error {
	return dara.Validate(s)
}
