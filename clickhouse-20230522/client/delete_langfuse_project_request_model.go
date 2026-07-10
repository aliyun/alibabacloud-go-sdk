// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteLangfuseProjectRequest
	GetDBInstanceId() *string
	SetOrganizationId(v string) *DeleteLangfuseProjectRequest
	GetOrganizationId() *string
	SetProjectId(v string) *DeleteLangfuseProjectRequest
	GetProjectId() *string
	SetRegionId(v string) *DeleteLangfuseProjectRequest
	GetRegionId() *string
}

type DeleteLangfuseProjectRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The Langfuse organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The Langfuse project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLangfuseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseProjectRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteLangfuseProjectRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteLangfuseProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteLangfuseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLangfuseProjectRequest) SetDBInstanceId(v string) *DeleteLangfuseProjectRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteLangfuseProjectRequest) SetOrganizationId(v string) *DeleteLangfuseProjectRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteLangfuseProjectRequest) SetProjectId(v string) *DeleteLangfuseProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteLangfuseProjectRequest) SetRegionId(v string) *DeleteLangfuseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLangfuseProjectRequest) Validate() error {
	return dara.Validate(s)
}
