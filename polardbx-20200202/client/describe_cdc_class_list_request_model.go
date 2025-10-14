// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcClassListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeCdcClassListRequest
	GetDBInstanceName() *string
	SetInstanceName(v string) *DescribeCdcClassListRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeCdcClassListRequest
	GetRegionId() *string
}

type DescribeCdcClassListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// pxc-hzrh51fze****pon-cdc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCdcClassListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcClassListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdcClassListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCdcClassListRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCdcClassListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCdcClassListRequest) SetDBInstanceName(v string) *DescribeCdcClassListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCdcClassListRequest) SetInstanceName(v string) *DescribeCdcClassListRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeCdcClassListRequest) SetRegionId(v string) *DescribeCdcClassListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCdcClassListRequest) Validate() error {
	return dara.Validate(s)
}
