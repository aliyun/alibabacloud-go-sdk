// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcVersionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeCdcVersionListRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeCdcVersionListRequest
	GetRegionId() *string
}

type DescribeCdcVersionListRequest struct {
	// example:
	//
	// pxc-bjrl7****k2vp7
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCdcVersionListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcVersionListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdcVersionListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCdcVersionListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCdcVersionListRequest) SetDBInstanceName(v string) *DescribeCdcVersionListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCdcVersionListRequest) SetRegionId(v string) *DescribeCdcVersionListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCdcVersionListRequest) Validate() error {
	return dara.Validate(s)
}
