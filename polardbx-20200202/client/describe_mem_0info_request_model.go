// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMem0InfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeMem0InfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeMem0InfoRequest
	GetRegionId() *string
}

type DescribeMem0InfoRequest struct {
	// Instance Name
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-spsil01pww4hfz
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Region
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMem0InfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0InfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMem0InfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeMem0InfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMem0InfoRequest) SetDBInstanceName(v string) *DescribeMem0InfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeMem0InfoRequest) SetRegionId(v string) *DescribeMem0InfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMem0InfoRequest) Validate() error {
	return dara.Validate(s)
}
