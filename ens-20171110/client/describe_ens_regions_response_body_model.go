// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnsRegionsResponseBody
	GetCode() *int32
	SetEnsRegions(v *DescribeEnsRegionsResponseBodyEnsRegions) *DescribeEnsRegionsResponseBody
	GetEnsRegions() *DescribeEnsRegionsResponseBodyEnsRegions
	SetRequestId(v string) *DescribeEnsRegionsResponseBody
	GetRequestId() *string
}

type DescribeEnsRegionsResponseBody struct {
	// The service code. 0 is returned for a successful request. An error code is returned for a failed request.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the regions.
	EnsRegions *DescribeEnsRegionsResponseBodyEnsRegions `json:"EnsRegions,omitempty" xml:"EnsRegions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnsRegionsResponseBody) GetEnsRegions() *DescribeEnsRegionsResponseBodyEnsRegions {
	return s.EnsRegions
}

func (s *DescribeEnsRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsRegionsResponseBody) SetCode(v int32) *DescribeEnsRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnsRegionsResponseBody) SetEnsRegions(v *DescribeEnsRegionsResponseBodyEnsRegions) *DescribeEnsRegionsResponseBody {
	s.EnsRegions = v
	return s
}

func (s *DescribeEnsRegionsResponseBody) SetRequestId(v string) *DescribeEnsRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsRegionsResponseBodyEnsRegions struct {
	EnsRegions []*DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions `json:"EnsRegions,omitempty" xml:"EnsRegions,omitempty" type:"Repeated"`
}

func (s DescribeEnsRegionsResponseBodyEnsRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionsResponseBodyEnsRegions) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponseBodyEnsRegions) GetEnsRegions() []*DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	return s.EnsRegions
}

func (s *DescribeEnsRegionsResponseBodyEnsRegions) SetEnsRegions(v []*DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) *DescribeEnsRegionsResponseBodyEnsRegions {
	s.EnsRegions = v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions struct {
	// The code of the region.
	//
	// example:
	//
	// NorthEastChina
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// NorthChina
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-dalian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Dalian Unicom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The province where the node is deployed.
	//
	// example:
	//
	// Liaoning Province
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) GetArea() *string {
	return s.Area
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) GetEnName() *string {
	return s.EnName
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) GetName() *string {
	return s.Name
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) GetProvince() *string {
	return s.Province
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetArea(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.Area = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetEnName(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.EnName = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetEnsRegionId(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetName(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.Name = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) SetProvince(v string) *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions {
	s.Province = &v
	return s
}

func (s *DescribeEnsRegionsResponseBodyEnsRegionsEnsRegions) Validate() error {
	return dara.Validate(s)
}
