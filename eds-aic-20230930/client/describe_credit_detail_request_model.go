// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeCreditDetailRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribeCreditDetailRequest
	GetInstanceIds() []*string
	SetPackageIds(v []*string) *DescribeCreditDetailRequest
	GetPackageIds() []*string
	SetPageNum(v string) *DescribeCreditDetailRequest
	GetPageNum() *string
	SetPageSize(v string) *DescribeCreditDetailRequest
	GetPageSize() *string
	SetStartTime(v int64) *DescribeCreditDetailRequest
	GetStartTime() *int64
}

type DescribeCreditDetailRequest struct {
	// The end time.
	//
	// example:
	//
	// 1782906240000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The list of resource plan or credit booster pack IDs.
	PackageIds []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1782819840000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCreditDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreditDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeCreditDetailRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeCreditDetailRequest) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *DescribeCreditDetailRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *DescribeCreditDetailRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCreditDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeCreditDetailRequest) SetEndTime(v int64) *DescribeCreditDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCreditDetailRequest) SetInstanceIds(v []*string) *DescribeCreditDetailRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCreditDetailRequest) SetPackageIds(v []*string) *DescribeCreditDetailRequest {
	s.PackageIds = v
	return s
}

func (s *DescribeCreditDetailRequest) SetPageNum(v string) *DescribeCreditDetailRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCreditDetailRequest) SetPageSize(v string) *DescribeCreditDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCreditDetailRequest) SetStartTime(v int64) *DescribeCreditDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCreditDetailRequest) Validate() error {
	return dara.Validate(s)
}
