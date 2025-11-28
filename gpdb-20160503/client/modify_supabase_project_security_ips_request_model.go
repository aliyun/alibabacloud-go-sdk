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
}

type ModifySupabaseProjectSecurityIpsRequest struct {
	// The Supabase project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-407****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// A comma-separated list of IP addresses and CIDR blocks to set as the whitelist. You can specify up to 1,000 entries. Supported formats:
	//
	// 	- Single IP: 10.23.12.24
	//
	// 	- CIDR Block: 10.23.12.0/24 (the prefix`/24` indicates the length must be between 1 and 32)``
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
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

func (s *ModifySupabaseProjectSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
