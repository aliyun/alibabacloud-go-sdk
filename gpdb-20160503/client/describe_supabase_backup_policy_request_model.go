// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *DescribeSupabaseBackupPolicyRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeSupabaseBackupPolicyRequest
	GetRegionId() *string
}

type DescribeSupabaseBackupPolicyRequest struct {
	// Instance ID of the Supabase instance. You can obtain instance ID from the Supabase page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-263****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSupabaseBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseBackupPolicyRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeSupabaseBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSupabaseBackupPolicyRequest) SetProjectId(v string) *DescribeSupabaseBackupPolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyRequest) SetRegionId(v string) *DescribeSupabaseBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
