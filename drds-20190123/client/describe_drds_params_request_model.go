// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsParamsRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsParamsRequest
	GetDrdsInstanceId() *string
	SetParamLevel(v string) *DescribeDrdsParamsRequest
	GetParamLevel() *string
	SetRegionId(v string) *DescribeDrdsParamsRequest
	GetRegionId() *string
}

type DescribeDrdsParamsRequest struct {
	// The name of the database.
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
	// DescribeDrdsParams
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The type of nodes whose parameters you want to query. Valid values:
	//
	// 	- **INSTANCE: the instance level.**
	//
	// 	- **DB**: the database level.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// The ID of the region where the PolarDB-X 1.0 instance is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsParamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsParamsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsParamsRequest) GetParamLevel() *string {
	return s.ParamLevel
}

func (s *DescribeDrdsParamsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsParamsRequest) SetDbName(v string) *DescribeDrdsParamsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetDrdsInstanceId(v string) *DescribeDrdsParamsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetParamLevel(v string) *DescribeDrdsParamsRequest {
	s.ParamLevel = &v
	return s
}

func (s *DescribeDrdsParamsRequest) SetRegionId(v string) *DescribeDrdsParamsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsParamsRequest) Validate() error {
	return dara.Validate(s)
}
