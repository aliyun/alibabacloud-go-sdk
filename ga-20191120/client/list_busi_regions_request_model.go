// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusiRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListBusiRegionsRequest
	GetRegionId() *string
}

type ListBusiRegionsRequest struct {
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBusiRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBusiRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBusiRegionsRequest) SetRegionId(v string) *ListBusiRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBusiRegionsRequest) Validate() error {
	return dara.Validate(s)
}
