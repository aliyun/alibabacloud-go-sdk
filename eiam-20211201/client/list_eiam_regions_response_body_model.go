// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEiamRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*ListEiamRegionsResponseBodyRegions) *ListEiamRegionsResponseBody
	GetRegions() []*ListEiamRegionsResponseBodyRegions
	SetRequestId(v string) *ListEiamRegionsResponseBody
	GetRequestId() *string
}

type ListEiamRegionsResponseBody struct {
	// The region list.
	Regions []*ListEiamRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEiamRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEiamRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEiamRegionsResponseBody) GetRegions() []*ListEiamRegionsResponseBodyRegions {
	return s.Regions
}

func (s *ListEiamRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEiamRegionsResponseBody) SetRegions(v []*ListEiamRegionsResponseBodyRegions) *ListEiamRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListEiamRegionsResponseBody) SetRequestId(v string) *ListEiamRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEiamRegionsResponseBody) Validate() error {
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

type ListEiamRegionsResponseBodyRegions struct {
	// The region name.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEiamRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListEiamRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListEiamRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListEiamRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEiamRegionsResponseBodyRegions) SetLocalName(v string) *ListEiamRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListEiamRegionsResponseBodyRegions) SetRegionId(v string) *ListEiamRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListEiamRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
