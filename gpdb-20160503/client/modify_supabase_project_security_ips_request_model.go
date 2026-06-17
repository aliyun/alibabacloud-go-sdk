// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *ModifySupabaseProjectSecurityIpsRequest
	GetProjectId() *string
	SetRegionId(v string) *ModifySupabaseProjectSecurityIpsRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifySupabaseProjectSecurityIpsRequest
	GetSecurityIPList() *string
	SetUpdateDb(v bool) *ModifySupabaseProjectSecurityIpsRequest
	GetUpdateDb() *bool
	SetUpdateWeb(v bool) *ModifySupabaseProjectSecurityIpsRequest
	GetUpdateWeb() *bool
}

type ModifySupabaseProjectSecurityIpsRequest struct {
	// The Supabase instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-407****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// > For more information, see [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) to view available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of IP addresses for the whitelist. Up to 1,000 IP addresses are supported. Separate multiple IP addresses with commas. The following formats are supported:
	//
	// - 10.23.12.24 (IP address)
	//
	// - 10.23.12.24/24 (A CIDR block, where `/24` indicates the prefix length. The prefix length must be an integer in the range `[1,32]`.)
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// Specifies whether to modify the whitelist for database port 5432. The default value is true.
	UpdateDb *bool `json:"UpdateDb,omitempty" xml:"UpdateDb,omitempty"`
	// Specifies whether to modify the whitelist for HTTP port 80 and HTTPS port 443. The default value is true.
	UpdateWeb *bool `json:"UpdateWeb,omitempty" xml:"UpdateWeb,omitempty"`
}

func (s ModifySupabaseProjectSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectSecurityIpsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifySupabaseProjectSecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySupabaseProjectSecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySupabaseProjectSecurityIpsRequest) GetUpdateDb() *bool {
	return s.UpdateDb
}

func (s *ModifySupabaseProjectSecurityIpsRequest) GetUpdateWeb() *bool {
	return s.UpdateWeb
}

func (s *ModifySupabaseProjectSecurityIpsRequest) SetProjectId(v string) *ModifySupabaseProjectSecurityIpsRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsRequest) SetRegionId(v string) *ModifySupabaseProjectSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsRequest) SetSecurityIPList(v string) *ModifySupabaseProjectSecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsRequest) SetUpdateDb(v bool) *ModifySupabaseProjectSecurityIpsRequest {
	s.UpdateDb = &v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsRequest) SetUpdateWeb(v bool) *ModifySupabaseProjectSecurityIpsRequest {
	s.UpdateWeb = &v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
