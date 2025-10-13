// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHiveAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DisableHiveAccessRequest
	GetRegionId() *string
}

type DisableHiveAccessRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableHiveAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableHiveAccessRequest) GoString() string {
	return s.String()
}

func (s *DisableHiveAccessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableHiveAccessRequest) SetRegionId(v string) *DisableHiveAccessRequest {
	s.RegionId = &v
	return s
}

func (s *DisableHiveAccessRequest) Validate() error {
	return dara.Validate(s)
}
