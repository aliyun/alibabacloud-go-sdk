// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseMem0PublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentConnectionString(v string) *ReleaseMem0PublicConnectionRequest
	GetCurrentConnectionString() *string
	SetDBInstanceName(v string) *ReleaseMem0PublicConnectionRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ReleaseMem0PublicConnectionRequest
	GetRegionId() *string
}

type ReleaseMem0PublicConnectionRequest struct {
	// The public endpoint to be released. This parameter is optional. If this parameter is not specified, the public endpoint of Mem0 is automatically released.
	//
	// example:
	//
	// pxc-hzjasdyuoo.polarx.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0****r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID of the instance. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the details of regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseMem0PublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseMem0PublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseMem0PublicConnectionRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ReleaseMem0PublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleaseMem0PublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseMem0PublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseMem0PublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseMem0PublicConnectionRequest) SetDBInstanceName(v string) *ReleaseMem0PublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseMem0PublicConnectionRequest) SetRegionId(v string) *ReleaseMem0PublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseMem0PublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
