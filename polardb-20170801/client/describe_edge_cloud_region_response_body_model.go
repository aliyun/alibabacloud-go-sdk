// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeCloudRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeCloudRegions(v []*DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) *DescribeEdgeCloudRegionResponseBody
	GetEdgeCloudRegions() []*DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions
	SetRequestId(v string) *DescribeEdgeCloudRegionResponseBody
	GetRequestId() *string
}

type DescribeEdgeCloudRegionResponseBody struct {
	// The list of available edge cloud regions. When a whitelist is configured, only the regions available to the current account are returned.
	EdgeCloudRegions []*DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions `json:"EdgeCloudRegions,omitempty" xml:"EdgeCloudRegions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEdgeCloudRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeCloudRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeCloudRegionResponseBody) GetEdgeCloudRegions() []*DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions {
	return s.EdgeCloudRegions
}

func (s *DescribeEdgeCloudRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEdgeCloudRegionResponseBody) SetEdgeCloudRegions(v []*DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) *DescribeEdgeCloudRegionResponseBody {
	s.EdgeCloudRegions = v
	return s
}

func (s *DescribeEdgeCloudRegionResponseBody) SetRequestId(v string) *DescribeEdgeCloudRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdgeCloudRegionResponseBody) Validate() error {
	if s.EdgeCloudRegions != nil {
		for _, item := range s.EdgeCloudRegions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions struct {
	// The ID of the central control region corresponding to the edge cloud region.
	//
	// example:
	//
	// cn-beijing
	CentralControlRegionId *string `json:"CentralControlRegionId,omitempty" xml:"CentralControlRegionId,omitempty"`
	// The ID of the edge cloud region.
	//
	// example:
	//
	// sg-singapore
	EdgeCloudRegionId *string `json:"EdgeCloudRegionId,omitempty" xml:"EdgeCloudRegionId,omitempty"`
	// The name of the edge cloud region.
	//
	// example:
	//
	// edgeCloudRegionName
	EdgeCloudRegionName *string `json:"EdgeCloudRegionName,omitempty" xml:"EdgeCloudRegionName,omitempty"`
	// The list of standby edge cloud regions.
	StandbyEdgeCloudRegionList []*DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList `json:"StandbyEdgeCloudRegionList,omitempty" xml:"StandbyEdgeCloudRegionList,omitempty" type:"Repeated"`
}

func (s DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) GoString() string {
	return s.String()
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) GetCentralControlRegionId() *string {
	return s.CentralControlRegionId
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) GetEdgeCloudRegionId() *string {
	return s.EdgeCloudRegionId
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) GetEdgeCloudRegionName() *string {
	return s.EdgeCloudRegionName
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) GetStandbyEdgeCloudRegionList() []*DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList {
	return s.StandbyEdgeCloudRegionList
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) SetCentralControlRegionId(v string) *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions {
	s.CentralControlRegionId = &v
	return s
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) SetEdgeCloudRegionId(v string) *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions {
	s.EdgeCloudRegionId = &v
	return s
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) SetEdgeCloudRegionName(v string) *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions {
	s.EdgeCloudRegionName = &v
	return s
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) SetStandbyEdgeCloudRegionList(v []*DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList) *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions {
	s.StandbyEdgeCloudRegionList = v
	return s
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegions) Validate() error {
	if s.StandbyEdgeCloudRegionList != nil {
		for _, item := range s.StandbyEdgeCloudRegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList struct {
	// The ID of the standby edge cloud region.
	//
	// example:
	//
	// standbyEdgeCloudRegionId
	StandbyEdgeCloudRegionId *string `json:"StandbyEdgeCloudRegionId,omitempty" xml:"StandbyEdgeCloudRegionId,omitempty"`
	// The name of the standby edge cloud region.
	//
	// example:
	//
	// standbyEdgeCloudRegionName
	StandbyEdgeCloudRegionName *string `json:"StandbyEdgeCloudRegionName,omitempty" xml:"StandbyEdgeCloudRegionName,omitempty"`
}

func (s DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList) GoString() string {
	return s.String()
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList) GetStandbyEdgeCloudRegionId() *string {
	return s.StandbyEdgeCloudRegionId
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList) GetStandbyEdgeCloudRegionName() *string {
	return s.StandbyEdgeCloudRegionName
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList) SetStandbyEdgeCloudRegionId(v string) *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList {
	s.StandbyEdgeCloudRegionId = &v
	return s
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList) SetStandbyEdgeCloudRegionName(v string) *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList {
	s.StandbyEdgeCloudRegionName = &v
	return s
}

func (s *DescribeEdgeCloudRegionResponseBodyEdgeCloudRegionsStandbyEdgeCloudRegionList) Validate() error {
	return dara.Validate(s)
}
