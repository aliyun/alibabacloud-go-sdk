// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBMiniEngineVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBMiniEngineVersionsRequest
	GetDBInstanceId() *string
	SetDedicatedHostGroupId(v string) *DescribeDBMiniEngineVersionsRequest
	GetDedicatedHostGroupId() *string
	SetEngine(v string) *DescribeDBMiniEngineVersionsRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeDBMiniEngineVersionsRequest
	GetEngineVersion() *string
	SetMinorVersionTag(v string) *DescribeDBMiniEngineVersionsRequest
	GetMinorVersionTag() *string
	SetNodeType(v string) *DescribeDBMiniEngineVersionsRequest
	GetNodeType() *string
	SetRegionId(v string) *DescribeDBMiniEngineVersionsRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeDBMiniEngineVersionsRequest
	GetResourceOwnerId() *int64
	SetStorageType(v string) *DescribeDBMiniEngineVersionsRequest
	GetStorageType() *string
}

type DescribeDBMiniEngineVersionsRequest struct {
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	MinorVersionTag      *string `json:"MinorVersionTag,omitempty" xml:"MinorVersionTag,omitempty"`
	NodeType             *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// This parameter is required.
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StorageType     *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeDBMiniEngineVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBMiniEngineVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBMiniEngineVersionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBMiniEngineVersionsRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBMiniEngineVersionsRequest) GetMinorVersionTag() *string {
	return s.MinorVersionTag
}

func (s *DescribeDBMiniEngineVersionsRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBMiniEngineVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBMiniEngineVersionsRequest) SetDBInstanceId(v string) *DescribeDBMiniEngineVersionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetDedicatedHostGroupId(v string) *DescribeDBMiniEngineVersionsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetEngine(v string) *DescribeDBMiniEngineVersionsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetEngineVersion(v string) *DescribeDBMiniEngineVersionsRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetMinorVersionTag(v string) *DescribeDBMiniEngineVersionsRequest {
	s.MinorVersionTag = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetNodeType(v string) *DescribeDBMiniEngineVersionsRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetRegionId(v string) *DescribeDBMiniEngineVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetResourceOwnerId(v int64) *DescribeDBMiniEngineVersionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetStorageType(v string) *DescribeDBMiniEngineVersionsRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) Validate() error {
	return dara.Validate(s)
}
