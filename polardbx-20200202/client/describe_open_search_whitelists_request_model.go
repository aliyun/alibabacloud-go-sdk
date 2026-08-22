// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchWhitelistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenSearchWhitelistsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeOpenSearchWhitelistsRequest
	GetRegionId() *string
}

type DescribeOpenSearchWhitelistsRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0****r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOpenSearchWhitelistsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchWhitelistsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchWhitelistsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenSearchWhitelistsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchWhitelistsRequest) SetDBInstanceName(v string) *DescribeOpenSearchWhitelistsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsRequest) SetRegionId(v string) *DescribeOpenSearchWhitelistsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsRequest) Validate() error {
	return dara.Validate(s)
}
