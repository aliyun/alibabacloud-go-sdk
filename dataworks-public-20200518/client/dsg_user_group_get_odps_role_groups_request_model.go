// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupGetOdpsRoleGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *DsgUserGroupGetOdpsRoleGroupsRequest
	GetProjectName() *string
}

type DsgUserGroupGetOdpsRoleGroupsRequest struct {
	// The name of the MaxCompute project. You can call the [DsgPlatformQueryProjectsAndSchemaFromMeta](https://help.aliyun.com/document_detail/2786303.html) operation to query a list of MaxCompute projects.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DsgUserGroupGetOdpsRoleGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupGetOdpsRoleGroupsRequest) GoString() string {
	return s.String()
}

func (s *DsgUserGroupGetOdpsRoleGroupsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DsgUserGroupGetOdpsRoleGroupsRequest) SetProjectName(v string) *DsgUserGroupGetOdpsRoleGroupsRequest {
	s.ProjectName = &v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsRequest) Validate() error {
	return dara.Validate(s)
}
