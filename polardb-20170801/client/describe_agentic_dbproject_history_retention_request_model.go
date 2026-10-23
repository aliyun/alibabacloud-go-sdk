// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectHistoryRetentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAgenticDBProjectHistoryRetentionRequest
	GetDBClusterId() *string
	SetProjectId(v string) *DescribeAgenticDBProjectHistoryRetentionRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeAgenticDBProjectHistoryRetentionRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribeAgenticDBProjectHistoryRetentionRequest
	GetTenantId() *string
}

type DescribeAgenticDBProjectHistoryRetentionRequest struct {
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The target project ID.
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

func (s DescribeAgenticDBProjectHistoryRetentionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectHistoryRetentionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) SetDBClusterId(v string) *DescribeAgenticDBProjectHistoryRetentionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) SetProjectId(v string) *DescribeAgenticDBProjectHistoryRetentionRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) SetRegionId(v string) *DescribeAgenticDBProjectHistoryRetentionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) SetTenantId(v string) *DescribeAgenticDBProjectHistoryRetentionRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionRequest) Validate() error {
	return dara.Validate(s)
}
