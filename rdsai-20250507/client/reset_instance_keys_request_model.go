// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstanceKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ResetInstanceKeysRequest
	GetInstanceName() *string
	SetRegionId(v string) *ResetInstanceKeysRequest
	GetRegionId() *string
}

type ResetInstanceKeysRequest struct {
	// The instance ID of the AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetInstanceKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetInstanceKeysRequest) GoString() string {
	return s.String()
}

func (s *ResetInstanceKeysRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ResetInstanceKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetInstanceKeysRequest) SetInstanceName(v string) *ResetInstanceKeysRequest {
	s.InstanceName = &v
	return s
}

func (s *ResetInstanceKeysRequest) SetRegionId(v string) *ResetInstanceKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ResetInstanceKeysRequest) Validate() error {
	return dara.Validate(s)
}
