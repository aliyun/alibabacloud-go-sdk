// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchWhitelistGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateOpenSearchWhitelistGroupRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *CreateOpenSearchWhitelistGroupRequest
	GetGroupName() *string
	SetIPs(v string) *CreateOpenSearchWhitelistGroupRequest
	GetIPs() *string
	SetRegionId(v string) *CreateOpenSearchWhitelistGroupRequest
	GetRegionId() *string
	SetRemark(v string) *CreateOpenSearchWhitelistGroupRequest
	GetRemark() *string
}

type CreateOpenSearchWhitelistGroupRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzravgpt8q****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the whitelist group.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The list of allowed source IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	IPs *string `json:"IPs,omitempty" xml:"IPs,omitempty"`
	// The ID of the region in which the instance resides. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the group ID.
	//
	// example:
	//
	// MAIN
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateOpenSearchWhitelistGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchWhitelistGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchWhitelistGroupRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateOpenSearchWhitelistGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateOpenSearchWhitelistGroupRequest) GetIPs() *string {
	return s.IPs
}

func (s *CreateOpenSearchWhitelistGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOpenSearchWhitelistGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateOpenSearchWhitelistGroupRequest) SetDBInstanceName(v string) *CreateOpenSearchWhitelistGroupRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupRequest) SetGroupName(v string) *CreateOpenSearchWhitelistGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupRequest) SetIPs(v string) *CreateOpenSearchWhitelistGroupRequest {
	s.IPs = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupRequest) SetRegionId(v string) *CreateOpenSearchWhitelistGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupRequest) SetRemark(v string) *CreateOpenSearchWhitelistGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupRequest) Validate() error {
	return dara.Validate(s)
}
