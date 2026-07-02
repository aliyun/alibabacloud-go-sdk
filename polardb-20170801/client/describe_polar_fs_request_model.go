// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePolarFsRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribePolarFsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolarFsRequest
	GetPageSize() *int32
	SetPolarFsInstanceDescription(v string) *DescribePolarFsRequest
	GetPolarFsInstanceDescription() *string
	SetPolarFsInstanceIds(v string) *DescribePolarFsRequest
	GetPolarFsInstanceIds() *string
	SetPolarFsType(v string) *DescribePolarFsRequest
	GetPolarFsType() *string
	SetRegionId(v string) *DescribePolarFsRequest
	GetRegionId() *string
	SetRelativeDbClusterId(v string) *DescribePolarFsRequest
	GetRelativeDbClusterId() *string
	SetTag(v []*DescribePolarFsRequestTag) *DescribePolarFsRequest
	GetTag() []*DescribePolarFsRequestTag
}

type DescribePolarFsRequest struct {
	// example:
	//
	// pc-2ze8u21s03******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// pfs-ins1********
	PolarFsInstanceDescription *string `json:"PolarFsInstanceDescription,omitempty" xml:"PolarFsInstanceDescription,omitempty"`
	// example:
	//
	// pfs-ins1*******
	PolarFsInstanceIds *string `json:"PolarFsInstanceIds,omitempty" xml:"PolarFsInstanceIds,omitempty"`
	// example:
	//
	// all
	PolarFsType *string `json:"PolarFsType,omitempty" xml:"PolarFsType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// pc-2zejpr41d9xk*****
	RelativeDbClusterId *string                      `json:"RelativeDbClusterId,omitempty" xml:"RelativeDbClusterId,omitempty"`
	Tag                 []*DescribePolarFsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribePolarFsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarFsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarFsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarFsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolarFsRequest) GetPolarFsInstanceDescription() *string {
	return s.PolarFsInstanceDescription
}

func (s *DescribePolarFsRequest) GetPolarFsInstanceIds() *string {
	return s.PolarFsInstanceIds
}

func (s *DescribePolarFsRequest) GetPolarFsType() *string {
	return s.PolarFsType
}

func (s *DescribePolarFsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarFsRequest) GetRelativeDbClusterId() *string {
	return s.RelativeDbClusterId
}

func (s *DescribePolarFsRequest) GetTag() []*DescribePolarFsRequestTag {
	return s.Tag
}

func (s *DescribePolarFsRequest) SetDBClusterId(v string) *DescribePolarFsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarFsRequest) SetPageNumber(v int32) *DescribePolarFsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarFsRequest) SetPageSize(v int32) *DescribePolarFsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolarFsRequest) SetPolarFsInstanceDescription(v string) *DescribePolarFsRequest {
	s.PolarFsInstanceDescription = &v
	return s
}

func (s *DescribePolarFsRequest) SetPolarFsInstanceIds(v string) *DescribePolarFsRequest {
	s.PolarFsInstanceIds = &v
	return s
}

func (s *DescribePolarFsRequest) SetPolarFsType(v string) *DescribePolarFsRequest {
	s.PolarFsType = &v
	return s
}

func (s *DescribePolarFsRequest) SetRegionId(v string) *DescribePolarFsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarFsRequest) SetRelativeDbClusterId(v string) *DescribePolarFsRequest {
	s.RelativeDbClusterId = &v
	return s
}

func (s *DescribePolarFsRequest) SetTag(v []*DescribePolarFsRequestTag) *DescribePolarFsRequest {
	s.Tag = v
	return s
}

func (s *DescribePolarFsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsRequestTag struct {
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePolarFsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribePolarFsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribePolarFsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribePolarFsRequestTag) SetKey(v string) *DescribePolarFsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribePolarFsRequestTag) SetValue(v string) *DescribePolarFsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribePolarFsRequestTag) Validate() error {
	return dara.Validate(s)
}
