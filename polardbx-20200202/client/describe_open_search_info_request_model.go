// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenSearchInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeOpenSearchInfoRequest
	GetRegionId() *string
}

type DescribeOpenSearchInfoRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region in which the instance resides. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOpenSearchInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenSearchInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchInfoRequest) SetDBInstanceName(v string) *DescribeOpenSearchInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenSearchInfoRequest) SetRegionId(v string) *DescribeOpenSearchInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchInfoRequest) Validate() error {
	return dara.Validate(s)
}
