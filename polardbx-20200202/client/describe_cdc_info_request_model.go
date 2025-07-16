// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeCdcInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeCdcInfoRequest
	GetRegionId() *string
}

type DescribeCdcInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCdcInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdcInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCdcInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCdcInfoRequest) SetDBInstanceName(v string) *DescribeCdcInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCdcInfoRequest) SetRegionId(v string) *DescribeCdcInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCdcInfoRequest) Validate() error {
	return dara.Validate(s)
}
