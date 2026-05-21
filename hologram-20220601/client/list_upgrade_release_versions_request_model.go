// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpgradeReleaseVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListUpgradeReleaseVersionsRequest
	GetRegionId() *string
}

type ListUpgradeReleaseVersionsRequest struct {
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUpgradeReleaseVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUpgradeReleaseVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListUpgradeReleaseVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUpgradeReleaseVersionsRequest) SetRegionId(v string) *ListUpgradeReleaseVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUpgradeReleaseVersionsRequest) Validate() error {
	return dara.Validate(s)
}
