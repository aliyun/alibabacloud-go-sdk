// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterType(v string) *DescribeDBNodePerformanceRequest
	GetCharacterType() *string
	SetDBInstanceName(v string) *DescribeDBNodePerformanceRequest
	GetDBInstanceName() *string
	SetDBNodeIds(v string) *DescribeDBNodePerformanceRequest
	GetDBNodeIds() *string
	SetDBNodeRole(v string) *DescribeDBNodePerformanceRequest
	GetDBNodeRole() *string
	SetEndTime(v string) *DescribeDBNodePerformanceRequest
	GetEndTime() *string
	SetKey(v string) *DescribeDBNodePerformanceRequest
	GetKey() *string
	SetRegionId(v string) *DescribeDBNodePerformanceRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDBNodePerformanceRequest
	GetStartTime() *string
}

type DescribeDBNodePerformanceRequest struct {
	// The node type. Valid values: polarx_cn, polarx_dn, polarx_cdc, polarx_gms, and polarx_columnar.
	//
	// This parameter is required.
	//
	// example:
	//
	// polarx_cn
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The name of the PolarDB-X instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-*******
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The node names. Separate multiple node names with commas (,). You can specify up to 10 CDC nodes and compute nodes, or up to 1 data node at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-i-******,pxc-i-*******
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// The node role. Valid values:
	//
	// - master: primary node.
	//
	// - slave: secondary node. This value applies to DN and GMS nodes.
	//
	// - standby: secondary node. This value applies to CN nodes.
	//
	// example:
	//
	// master
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// The end of the time range to query. Specify the time in the YYYY-MM-ddTHH:mmZ format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2012-06-18T15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric names. Separate multiple metric names with commas (,). You can specify up to 6 metrics. For more information, see [Performance metrics](https://help.aliyun.com/document_detail/332726.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Cpu_Usage,Mem_Usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the YYYY-MM-ddTHH:mmZ format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2012-06-08T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBNodePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeDBNodePerformanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBNodePerformanceRequest) GetDBNodeIds() *string {
	return s.DBNodeIds
}

func (s *DescribeDBNodePerformanceRequest) GetDBNodeRole() *string {
	return s.DBNodeRole
}

func (s *DescribeDBNodePerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBNodePerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBNodePerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBNodePerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBNodePerformanceRequest) SetCharacterType(v string) *DescribeDBNodePerformanceRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBInstanceName(v string) *DescribeDBNodePerformanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBNodeIds(v string) *DescribeDBNodePerformanceRequest {
	s.DBNodeIds = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBNodeRole(v string) *DescribeDBNodePerformanceRequest {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetEndTime(v string) *DescribeDBNodePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetKey(v string) *DescribeDBNodePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetRegionId(v string) *DescribeDBNodePerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetStartTime(v string) *DescribeDBNodePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
