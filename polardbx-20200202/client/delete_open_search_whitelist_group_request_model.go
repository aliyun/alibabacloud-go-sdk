// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchWhitelistGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteOpenSearchWhitelistGroupRequest
	GetDBInstanceName() *string
	SetGroupId(v string) *DeleteOpenSearchWhitelistGroupRequest
	GetGroupId() *string
	SetRegionId(v string) *DeleteOpenSearchWhitelistGroupRequest
	GetRegionId() *string
}

type DeleteOpenSearchWhitelistGroupRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the service group.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_y6sqhtr6jyu52w4oublg3674du
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the region where the instance resides. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteOpenSearchWhitelistGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchWhitelistGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchWhitelistGroupRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteOpenSearchWhitelistGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteOpenSearchWhitelistGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteOpenSearchWhitelistGroupRequest) SetDBInstanceName(v string) *DeleteOpenSearchWhitelistGroupRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupRequest) SetGroupId(v string) *DeleteOpenSearchWhitelistGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupRequest) SetRegionId(v string) *DeleteOpenSearchWhitelistGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupRequest) Validate() error {
	return dara.Validate(s)
}
