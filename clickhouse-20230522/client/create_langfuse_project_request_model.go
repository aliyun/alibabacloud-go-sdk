// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateLangfuseProjectRequest
	GetDBInstanceId() *string
	SetName(v string) *CreateLangfuseProjectRequest
	GetName() *string
	SetOrganizationId(v string) *CreateLangfuseProjectRequest
	GetOrganizationId() *string
	SetRegionId(v string) *CreateLangfuseProjectRequest
	GetRegionId() *string
}

type CreateLangfuseProjectRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The Langfuse project name.
	//
	// This parameter is required.
	//
	// example:
	//
	// project_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The organization ID to which the Langfuse project belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateLangfuseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateLangfuseProjectRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateLangfuseProjectRequest) GetName() *string {
	return s.Name
}

func (s *CreateLangfuseProjectRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateLangfuseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLangfuseProjectRequest) SetDBInstanceId(v string) *CreateLangfuseProjectRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateLangfuseProjectRequest) SetName(v string) *CreateLangfuseProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateLangfuseProjectRequest) SetOrganizationId(v string) *CreateLangfuseProjectRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateLangfuseProjectRequest) SetRegionId(v string) *CreateLangfuseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLangfuseProjectRequest) Validate() error {
	return dara.Validate(s)
}
