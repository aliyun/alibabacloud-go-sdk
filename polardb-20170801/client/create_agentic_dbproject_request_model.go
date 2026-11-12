// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAgenticDBProjectRequest
	GetDBClusterId() *string
	SetDefaultBranchName(v string) *CreateAgenticDBProjectRequest
	GetDefaultBranchName() *string
	SetDescription(v string) *CreateAgenticDBProjectRequest
	GetDescription() *string
	SetHistoryRetentionSeconds(v int64) *CreateAgenticDBProjectRequest
	GetHistoryRetentionSeconds() *int64
	SetProjectName(v string) *CreateAgenticDBProjectRequest
	GetProjectName() *string
	SetRegionId(v string) *CreateAgenticDBProjectRequest
	GetRegionId() *string
	SetTenantId(v string) *CreateAgenticDBProjectRequest
	GetTenantId() *string
	SetZoneId(v string) *CreateAgenticDBProjectRequest
	GetZoneId() *string
}

type CreateAgenticDBProjectRequest struct {
	// The ID of the AgenticDB cluster. The ID must have the pagc- prefix. You can obtain the ID after calling CreateDBCluster and specifying AgenticDbType=dedicated and ServerlessType=AgileServerless. A regular cluster ID with the pc- prefix is not supported. If an invalid ID is provided, the InvalidDBClusterId.NotFound error may be returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the default branch. Default value: main.
	//
	// example:
	//
	// main
	DefaultBranchName *string `json:"DefaultBranchName,omitempty" xml:"DefaultBranchName,omitempty"`
	// The description of the project.
	//
	// example:
	//
	// Production analytics database
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The retention period of historical data, in seconds.
	//
	// example:
	//
	// 604800
	HistoryRetentionSeconds *int64 `json:"HistoryRetentionSeconds,omitempty" xml:"HistoryRetentionSeconds,omitempty"`
	// The name of the project. The name must be unique within the same tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// analytics-prod
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the tenant. You can obtain the ID from the response after calling CreateAgenticDBTenant to create a tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the zone where the project storage resides. The zone must be the same as the zone where the cluster compute resides. If this parameter is not specified, the zone where the cluster compute resides is used by default.
	//
	// example:
	//
	// cn-shanghai-l
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateAgenticDBProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBProjectRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgenticDBProjectRequest) GetDefaultBranchName() *string {
	return s.DefaultBranchName
}

func (s *CreateAgenticDBProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgenticDBProjectRequest) GetHistoryRetentionSeconds() *int64 {
	return s.HistoryRetentionSeconds
}

func (s *CreateAgenticDBProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateAgenticDBProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgenticDBProjectRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAgenticDBProjectRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateAgenticDBProjectRequest) SetDBClusterId(v string) *CreateAgenticDBProjectRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetDefaultBranchName(v string) *CreateAgenticDBProjectRequest {
	s.DefaultBranchName = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetDescription(v string) *CreateAgenticDBProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetHistoryRetentionSeconds(v int64) *CreateAgenticDBProjectRequest {
	s.HistoryRetentionSeconds = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetProjectName(v string) *CreateAgenticDBProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetRegionId(v string) *CreateAgenticDBProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetTenantId(v string) *CreateAgenticDBProjectRequest {
	s.TenantId = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetZoneId(v string) *CreateAgenticDBProjectRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) Validate() error {
	return dara.Validate(s)
}
