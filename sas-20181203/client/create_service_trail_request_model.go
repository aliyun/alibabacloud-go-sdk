// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTrailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateServiceTrailRequest
	GetRegionId() *string
}

type CreateServiceTrailRequest struct {
	// The region ID of the instance. Valid values:
	//
	// 	- **cn-hangzhou**: International
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceTrailRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTrailRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceTrailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceTrailRequest) SetRegionId(v string) *CreateServiceTrailRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceTrailRequest) Validate() error {
	return dara.Validate(s)
}
