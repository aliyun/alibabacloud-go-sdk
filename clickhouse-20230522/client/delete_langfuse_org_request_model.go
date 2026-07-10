// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteLangfuseOrgRequest
	GetDBInstanceId() *string
	SetOrganizationId(v string) *DeleteLangfuseOrgRequest
	GetOrganizationId() *string
	SetRegionId(v string) *DeleteLangfuseOrgRequest
	GetRegionId() *string
}

type DeleteLangfuseOrgRequest struct {
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLangfuseOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseOrgRequest) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseOrgRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteLangfuseOrgRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteLangfuseOrgRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLangfuseOrgRequest) SetDBInstanceId(v string) *DeleteLangfuseOrgRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteLangfuseOrgRequest) SetOrganizationId(v string) *DeleteLangfuseOrgRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteLangfuseOrgRequest) SetRegionId(v string) *DeleteLangfuseOrgRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLangfuseOrgRequest) Validate() error {
	return dara.Validate(s)
}
