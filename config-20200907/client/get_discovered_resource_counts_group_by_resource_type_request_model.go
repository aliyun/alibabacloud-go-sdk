// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceCountsGroupByResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *GetDiscoveredResourceCountsGroupByResourceTypeRequest
	GetRegion() *string
}

type GetDiscoveredResourceCountsGroupByResourceTypeRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeRequest) GetRegion() *string {
	return s.Region
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeRequest) SetRegion(v string) *GetDiscoveredResourceCountsGroupByResourceTypeRequest {
	s.Region = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
