// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbRdsNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsDbRdsNameListRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsDbRdsNameListRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *DescribeDrdsDbRdsNameListRequest
	GetRegionId() *string
}

type DescribeDrdsDbRdsNameListRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdsxxxxxxxxxxx
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsDbRdsNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsDbRdsNameListRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsDbRdsNameListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsDbRdsNameListRequest) SetDbName(v string) *DescribeDrdsDbRdsNameListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbRdsNameListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListRequest) SetRegionId(v string) *DescribeDrdsDbRdsNameListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListRequest) Validate() error {
	return dara.Validate(s)
}
