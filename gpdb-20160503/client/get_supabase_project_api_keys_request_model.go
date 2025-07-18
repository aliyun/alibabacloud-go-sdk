// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *GetSupabaseProjectApiKeysRequest
	GetProjectId() *string
	SetRegionId(v string) *GetSupabaseProjectApiKeysRequest
	GetRegionId() *string
}

type GetSupabaseProjectApiKeysRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sbp-481****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSupabaseProjectApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectApiKeysRequest) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectApiKeysRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSupabaseProjectApiKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSupabaseProjectApiKeysRequest) SetProjectId(v string) *GetSupabaseProjectApiKeysRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSupabaseProjectApiKeysRequest) SetRegionId(v string) *GetSupabaseProjectApiKeysRequest {
	s.RegionId = &v
	return s
}

func (s *GetSupabaseProjectApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
