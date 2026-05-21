// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscardUpgradeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DiscardUpgradeRequest
	GetRegionId() *string
}

type DiscardUpgradeRequest struct {
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DiscardUpgradeRequest) String() string {
	return dara.Prettify(s)
}

func (s DiscardUpgradeRequest) GoString() string {
	return s.String()
}

func (s *DiscardUpgradeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DiscardUpgradeRequest) SetRegionId(v string) *DiscardUpgradeRequest {
	s.RegionId = &v
	return s
}

func (s *DiscardUpgradeRequest) Validate() error {
	return dara.Validate(s)
}
