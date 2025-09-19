// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceTrailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteServiceTrailRequest
	GetRegionId() *string
}

type DeleteServiceTrailRequest struct {
	// The region in which your Security Center service is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: center.
	//
	// 	- **ap-southeast-1**: Singapore.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteServiceTrailRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceTrailRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceTrailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServiceTrailRequest) SetRegionId(v string) *DeleteServiceTrailRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceTrailRequest) Validate() error {
	return dara.Validate(s)
}
