// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogShipperRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogShipperRegionList(v []*ListLogShipperRegionsResponseBodyLogShipperRegionList) *ListLogShipperRegionsResponseBody
	GetLogShipperRegionList() []*ListLogShipperRegionsResponseBodyLogShipperRegionList
	SetRequestId(v string) *ListLogShipperRegionsResponseBody
	GetRequestId() *string
}

type ListLogShipperRegionsResponseBody struct {
	// The regions supported by the log delivery feature.
	LogShipperRegionList []*ListLogShipperRegionsResponseBodyLogShipperRegionList `json:"LogShipperRegionList,omitempty" xml:"LogShipperRegionList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F9C4DE22-D242-5ABA-87EC-325ECBDC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLogShipperRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogShipperRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogShipperRegionsResponseBody) GetLogShipperRegionList() []*ListLogShipperRegionsResponseBodyLogShipperRegionList {
	return s.LogShipperRegionList
}

func (s *ListLogShipperRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogShipperRegionsResponseBody) SetLogShipperRegionList(v []*ListLogShipperRegionsResponseBodyLogShipperRegionList) *ListLogShipperRegionsResponseBody {
	s.LogShipperRegionList = v
	return s
}

func (s *ListLogShipperRegionsResponseBody) SetRequestId(v string) *ListLogShipperRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogShipperRegionsResponseBody) Validate() error {
	if s.LogShipperRegionList != nil {
		for _, item := range s.LogShipperRegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogShipperRegionsResponseBodyLogShipperRegionList struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLogShipperRegionsResponseBodyLogShipperRegionList) String() string {
	return dara.Prettify(s)
}

func (s ListLogShipperRegionsResponseBodyLogShipperRegionList) GoString() string {
	return s.String()
}

func (s *ListLogShipperRegionsResponseBodyLogShipperRegionList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLogShipperRegionsResponseBodyLogShipperRegionList) SetRegionId(v string) *ListLogShipperRegionsResponseBodyLogShipperRegionList {
	s.RegionId = &v
	return s
}

func (s *ListLogShipperRegionsResponseBodyLogShipperRegionList) Validate() error {
	return dara.Validate(s)
}
