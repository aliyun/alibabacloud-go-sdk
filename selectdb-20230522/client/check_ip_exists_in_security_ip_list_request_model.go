// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckIpExistsInSecurityIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CheckIpExistsInSecurityIpListRequest
	GetDBInstanceId() *string
	SetIp(v string) *CheckIpExistsInSecurityIpListRequest
	GetIp() *string
	SetRegionId(v string) *CheckIpExistsInSecurityIpListRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *CheckIpExistsInSecurityIpListRequest
	GetResourceOwnerId() *int64
}

type CheckIpExistsInSecurityIpListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-2bl4dj****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.239
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckIpExistsInSecurityIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckIpExistsInSecurityIpListRequest) GoString() string {
	return s.String()
}

func (s *CheckIpExistsInSecurityIpListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckIpExistsInSecurityIpListRequest) GetIp() *string {
	return s.Ip
}

func (s *CheckIpExistsInSecurityIpListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckIpExistsInSecurityIpListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckIpExistsInSecurityIpListRequest) SetDBInstanceId(v string) *CheckIpExistsInSecurityIpListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckIpExistsInSecurityIpListRequest) SetIp(v string) *CheckIpExistsInSecurityIpListRequest {
	s.Ip = &v
	return s
}

func (s *CheckIpExistsInSecurityIpListRequest) SetRegionId(v string) *CheckIpExistsInSecurityIpListRequest {
	s.RegionId = &v
	return s
}

func (s *CheckIpExistsInSecurityIpListRequest) SetResourceOwnerId(v int64) *CheckIpExistsInSecurityIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckIpExistsInSecurityIpListRequest) Validate() error {
	return dara.Validate(s)
}
