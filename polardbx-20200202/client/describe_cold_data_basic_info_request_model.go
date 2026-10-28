// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdDataBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeColdDataBasicInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeColdDataBasicInfoRequest
	GetRegionId() *string
}

type DescribeColdDataBasicInfoRequest struct {
	// The name of the PolarDB-X instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r****
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

func (s DescribeColdDataBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdDataBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeColdDataBasicInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColdDataBasicInfoRequest) SetDBInstanceName(v string) *DescribeColdDataBasicInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeColdDataBasicInfoRequest) SetRegionId(v string) *DescribeColdDataBasicInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColdDataBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
