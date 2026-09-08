// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionList(v []*ListRegionsResponseBodyRegionList) *ListRegionsResponseBody
	GetRegionList() []*ListRegionsResponseBodyRegionList
	SetRequestId(v string) *ListRegionsResponseBody
	GetRequestId() *string
}

type ListRegionsResponseBody struct {
	RegionList []*ListRegionsResponseBodyRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) GetRegionList() []*ListRegionsResponseBodyRegionList {
	return s.RegionList
}

func (s *ListRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionsResponseBody) SetRegionList(v []*ListRegionsResponseBodyRegionList) *ListRegionsResponseBody {
	s.RegionList = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) Validate() error {
	if s.RegionList != nil {
		for _, item := range s.RegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRegionsResponseBodyRegionList struct {
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 华北 3（张家口）
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s ListRegionsResponseBodyRegionList) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBodyRegionList) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegionList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRegionsResponseBodyRegionList) GetRegionName() *string {
	return s.RegionName
}

func (s *ListRegionsResponseBodyRegionList) SetRegionId(v string) *ListRegionsResponseBodyRegionList {
	s.RegionId = &v
	return s
}

func (s *ListRegionsResponseBodyRegionList) SetRegionName(v string) *ListRegionsResponseBodyRegionList {
	s.RegionName = &v
	return s
}

func (s *ListRegionsResponseBodyRegionList) Validate() error {
	return dara.Validate(s)
}
