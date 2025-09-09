// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShardTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeShardTaskInfoRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeShardTaskInfoRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *DescribeShardTaskInfoRequest
	GetRegionId() *string
	SetSourceTableName(v string) *DescribeShardTaskInfoRequest
	GetSourceTableName() *string
	SetTargetTableName(v string) *DescribeShardTaskInfoRequest
	GetTargetTableName() *string
}

type DescribeShardTaskInfoRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// jjjjjj_ppppp
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table that you want to convert or shard.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the table that is generated after you convert or shard the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s DescribeShardTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeShardTaskInfoRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeShardTaskInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeShardTaskInfoRequest) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *DescribeShardTaskInfoRequest) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *DescribeShardTaskInfoRequest) SetDbName(v string) *DescribeShardTaskInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetDrdsInstanceId(v string) *DescribeShardTaskInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetRegionId(v string) *DescribeShardTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetSourceTableName(v string) *DescribeShardTaskInfoRequest {
	s.SourceTableName = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) SetTargetTableName(v string) *DescribeShardTaskInfoRequest {
	s.TargetTableName = &v
	return s
}

func (s *DescribeShardTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
