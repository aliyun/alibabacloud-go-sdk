// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTrailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceTrailRequest
	GetRegionId() *string
}

type GetServiceTrailRequest struct {
	// The ID of the region in which the instance resides. Valid value:
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

func (s GetServiceTrailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTrailRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTrailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceTrailRequest) SetRegionId(v string) *GetServiceTrailRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTrailRequest) Validate() error {
	return dara.Validate(s)
}
