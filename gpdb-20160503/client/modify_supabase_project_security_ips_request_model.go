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
	// This parameter is required.
	//
	// example:
	//
	// sbp-407****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
