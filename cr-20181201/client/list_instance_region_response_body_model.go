// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceRegionResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListInstanceRegionResponseBody
	GetIsSuccess() *bool
	SetRegions(v []*ListInstanceRegionResponseBodyRegions) *ListInstanceRegionResponseBody
	GetRegions() []*ListInstanceRegionResponseBodyRegions
	SetRequestId(v string) *ListInstanceRegionResponseBody
	GetRequestId() *string
}

type ListInstanceRegionResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The list of regions.
	Regions []*ListInstanceRegionResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 11F182E1-0F84-4F5B-8D3B-61E991482727
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceRegionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceRegionResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListInstanceRegionResponseBody) GetRegions() []*ListInstanceRegionResponseBodyRegions {
	return s.Regions
}

func (s *ListInstanceRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceRegionResponseBody) SetCode(v string) *ListInstanceRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceRegionResponseBody) SetIsSuccess(v bool) *ListInstanceRegionResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListInstanceRegionResponseBody) SetRegions(v []*ListInstanceRegionResponseBodyRegions) *ListInstanceRegionResponseBody {
	s.Regions = v
	return s
}

func (s *ListInstanceRegionResponseBody) SetRequestId(v string) *ListInstanceRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceRegionResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceRegionResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (Shenzhen)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceRegionResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRegionResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListInstanceRegionResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListInstanceRegionResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceRegionResponseBodyRegions) SetLocalName(v string) *ListInstanceRegionResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListInstanceRegionResponseBodyRegions) SetRegionId(v string) *ListInstanceRegionResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListInstanceRegionResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
