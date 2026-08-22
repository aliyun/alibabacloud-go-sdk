// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContext0Request interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteContext0Request
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteContext0Request
	GetRegionId() *string
}

type DeleteContext0Request struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
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

func (s DeleteContext0Request) String() string {
	return dara.Prettify(s)
}

func (s DeleteContext0Request) GoString() string {
	return s.String()
}

func (s *DeleteContext0Request) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteContext0Request) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteContext0Request) SetDBInstanceName(v string) *DeleteContext0Request {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteContext0Request) SetRegionId(v string) *DeleteContext0Request {
	s.RegionId = &v
	return s
}

func (s *DeleteContext0Request) Validate() error {
	return dara.Validate(s)
}
