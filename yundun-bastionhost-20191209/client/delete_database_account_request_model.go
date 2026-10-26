// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountId(v string) *DeleteDatabaseAccountRequest
	GetDatabaseAccountId() *string
	SetInstanceId(v string) *DeleteDatabaseAccountRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DeleteDatabaseAccountRequest
	GetProjectId() *int64
	SetRegionId(v string) *DeleteDatabaseAccountRequest
	GetRegionId() *string
}

type DeleteDatabaseAccountRequest struct {
	// The ID of the database account to delete.
	//
	// > You can call the [ListDatabaseAccounts](https://help.aliyun.com/document_detail/2758839.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The ID of the bastion host instance to which the database account to be deleted belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host to which the database account to be deleted belongs.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDatabaseAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseAccountRequest) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *DeleteDatabaseAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDatabaseAccountRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDatabaseAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDatabaseAccountRequest) SetDatabaseAccountId(v string) *DeleteDatabaseAccountRequest {
	s.DatabaseAccountId = &v
	return s
}

func (s *DeleteDatabaseAccountRequest) SetInstanceId(v string) *DeleteDatabaseAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDatabaseAccountRequest) SetProjectId(v int64) *DeleteDatabaseAccountRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDatabaseAccountRequest) SetRegionId(v string) *DeleteDatabaseAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDatabaseAccountRequest) Validate() error {
	return dara.Validate(s)
}
