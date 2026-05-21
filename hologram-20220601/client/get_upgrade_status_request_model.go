// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpgradeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetUpgradeStatusRequest
	GetRegionId() *string
}

type GetUpgradeStatusRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUpgradeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUpgradeStatusRequest) SetRegionId(v string) *GetUpgradeStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetUpgradeStatusRequest) Validate() error {
	return dara.Validate(s)
}
