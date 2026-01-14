// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableBusiRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*ListAvailableBusiRegionsResponseBodyRegions) *ListAvailableBusiRegionsResponseBody
	GetRegions() []*ListAvailableBusiRegionsResponseBodyRegions
	SetRequestId(v string) *ListAvailableBusiRegionsResponseBody
	GetRequestId() *string
}

type ListAvailableBusiRegionsResponseBody struct {
	// The information about the regions.
	Regions []*ListAvailableBusiRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAvailableBusiRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableBusiRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsResponseBody) GetRegions() []*ListAvailableBusiRegionsResponseBodyRegions {
	return s.Regions
}

func (s *ListAvailableBusiRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableBusiRegionsResponseBody) SetRegions(v []*ListAvailableBusiRegionsResponseBodyRegions) *ListAvailableBusiRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListAvailableBusiRegionsResponseBody) SetRequestId(v string) *ListAvailableBusiRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableBusiRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvailableBusiRegionsResponseBodyRegions struct {
	// Indicates whether the region is in the Chinese mainland. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ChinaMainland *bool `json:"ChinaMainland,omitempty" xml:"ChinaMainland,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China (Qingdao)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Indicates whether it is a point of presence (PoP) of Alibaba Cloud. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Pop *bool `json:"Pop,omitempty" xml:"Pop,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAvailableBusiRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableBusiRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) GetChinaMainland() *bool {
	return s.ChinaMainland
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) GetPop() *bool {
	return s.Pop
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) SetChinaMainland(v bool) *ListAvailableBusiRegionsResponseBodyRegions {
	s.ChinaMainland = &v
	return s
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) SetLocalName(v string) *ListAvailableBusiRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) SetPop(v bool) *ListAvailableBusiRegionsResponseBodyRegions {
	s.Pop = &v
	return s
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) SetRegionId(v string) *ListAvailableBusiRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
