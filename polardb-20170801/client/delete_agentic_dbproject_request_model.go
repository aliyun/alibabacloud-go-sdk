// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAgenticDBProjectRequest
	GetDBClusterId() *string
	SetProjectId(v string) *DeleteAgenticDBProjectRequest
	GetProjectId() *string
	SetRegionId(v string) *DeleteAgenticDBProjectRequest
	GetRegionId() *string
	SetTenantId(v string) *DeleteAgenticDBProjectRequest
	GetTenantId() *string
}

type DeleteAgenticDBProjectRequest struct {
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the target project.
	//
	// This parameter is required.
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the tenant to which the project belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteAgenticDBProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBProjectRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAgenticDBProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteAgenticDBProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAgenticDBProjectRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAgenticDBProjectRequest) SetDBClusterId(v string) *DeleteAgenticDBProjectRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAgenticDBProjectRequest) SetProjectId(v string) *DeleteAgenticDBProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteAgenticDBProjectRequest) SetRegionId(v string) *DeleteAgenticDBProjectRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAgenticDBProjectRequest) SetTenantId(v string) *DeleteAgenticDBProjectRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteAgenticDBProjectRequest) Validate() error {
	return dara.Validate(s)
}
