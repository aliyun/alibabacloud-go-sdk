// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgenticDBProjectHistoryRetentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyAgenticDBProjectHistoryRetentionRequest
	GetDBClusterId() *string
	SetHistoryRetentionSeconds(v int64) *ModifyAgenticDBProjectHistoryRetentionRequest
	GetHistoryRetentionSeconds() *int64
	SetProjectId(v string) *ModifyAgenticDBProjectHistoryRetentionRequest
	GetProjectId() *string
	SetRegionId(v string) *ModifyAgenticDBProjectHistoryRetentionRequest
	GetRegionId() *string
	SetTenantId(v string) *ModifyAgenticDBProjectHistoryRetentionRequest
	GetTenantId() *string
}

type ModifyAgenticDBProjectHistoryRetentionRequest struct {
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The default valid range is 3,600 to 2,592,000 seconds (1 hour to 30 days), inclusive, subject to service configurations. Shortening the duration may clear historical data. Extending the duration does not restore cleared data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 604800
	HistoryRetentionSeconds *int64 `json:"HistoryRetentionSeconds,omitempty" xml:"HistoryRetentionSeconds,omitempty"`
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
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyAgenticDBProjectHistoryRetentionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgenticDBProjectHistoryRetentionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) GetHistoryRetentionSeconds() *int64 {
	return s.HistoryRetentionSeconds
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) SetDBClusterId(v string) *ModifyAgenticDBProjectHistoryRetentionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) SetHistoryRetentionSeconds(v int64) *ModifyAgenticDBProjectHistoryRetentionRequest {
	s.HistoryRetentionSeconds = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) SetProjectId(v string) *ModifyAgenticDBProjectHistoryRetentionRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) SetRegionId(v string) *ModifyAgenticDBProjectHistoryRetentionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) SetTenantId(v string) *ModifyAgenticDBProjectHistoryRetentionRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionRequest) Validate() error {
	return dara.Validate(s)
}
