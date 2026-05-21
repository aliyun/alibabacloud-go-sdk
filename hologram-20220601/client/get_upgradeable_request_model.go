// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpgradeableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetUpgradeableRequest
	GetRegionId() *string
}

type GetUpgradeableRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUpgradeableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeableRequest) GoString() string {
	return s.String()
}

func (s *GetUpgradeableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUpgradeableRequest) SetRegionId(v string) *GetUpgradeableRequest {
	s.RegionId = &v
	return s
}

func (s *GetUpgradeableRequest) Validate() error {
	return dara.Validate(s)
}
