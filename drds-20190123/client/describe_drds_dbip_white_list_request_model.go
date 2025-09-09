// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBIpWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsDBIpWhiteListRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsDBIpWhiteListRequest
	GetDrdsInstanceId() *string
	SetGroupName(v string) *DescribeDrdsDBIpWhiteListRequest
	GetGroupName() *string
	SetRegionId(v string) *DescribeDrdsDBIpWhiteListRequest
	GetRegionId() *string
}

type DescribeDrdsDBIpWhiteListRequest struct {
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The name of the whitelist group.
	//
	// example:
	//
	// group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsDBIpWhiteListRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsDBIpWhiteListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDrdsDBIpWhiteListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetDbName(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetGroupName(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) SetRegionId(v string) *DescribeDrdsDBIpWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
