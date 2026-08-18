// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfuseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribePxfuseInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribePxfuseInfoRequest
	GetRegionId() *string
}

type DescribePxfuseInfoRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePxfuseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribePxfuseInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePxfuseInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfuseInfoRequest) SetDBInstanceName(v string) *DescribePxfuseInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePxfuseInfoRequest) SetRegionId(v string) *DescribePxfuseInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePxfuseInfoRequest) Validate() error {
	return dara.Validate(s)
}
