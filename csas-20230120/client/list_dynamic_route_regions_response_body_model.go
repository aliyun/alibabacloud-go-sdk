// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicRouteRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*string) *ListDynamicRouteRegionsResponseBody
	GetRegions() []*string
	SetRequestId(v string) *ListDynamicRouteRegionsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListDynamicRouteRegionsResponseBody
	GetTotalNum() *int32
}

type ListDynamicRouteRegionsResponseBody struct {
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDynamicRouteRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicRouteRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicRouteRegionsResponseBody) GetRegions() []*string {
	return s.Regions
}

func (s *ListDynamicRouteRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDynamicRouteRegionsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListDynamicRouteRegionsResponseBody) SetRegions(v []*string) *ListDynamicRouteRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListDynamicRouteRegionsResponseBody) SetRequestId(v string) *ListDynamicRouteRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDynamicRouteRegionsResponseBody) SetTotalNum(v int32) *ListDynamicRouteRegionsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListDynamicRouteRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
