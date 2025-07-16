// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest
	GetCurrentConnectionString() *string
	SetDBInstanceName(v string) *ReleaseInstancePublicConnectionRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ReleaseInstancePublicConnectionRequest
	GetRegionId() *string
}

type ReleaseInstancePublicConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasdyuoo.polarx.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ReleaseInstancePublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleaseInstancePublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceName(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetRegionId(v string) *ReleaseInstancePublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
