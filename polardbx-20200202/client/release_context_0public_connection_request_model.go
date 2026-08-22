// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseContext0PublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentConnectionString(v string) *ReleaseContext0PublicConnectionRequest
	GetCurrentConnectionString() *string
	SetDBInstanceName(v string) *ReleaseContext0PublicConnectionRequest
	GetDBInstanceName() *string
	SetNodeType(v string) *ReleaseContext0PublicConnectionRequest
	GetNodeType() *string
	SetRegionId(v string) *ReleaseContext0PublicConnectionRequest
	GetRegionId() *string
}

type ReleaseContext0PublicConnectionRequest struct {
	// The public network connection string to release. If this parameter is not specified, the Mem0 public address is subject to automatic release.
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
	// The target node type. Valid values:
	//
	// - service
	//
	// - dashboard
	//
	// example:
	//
	// dn
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseContext0PublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContext0PublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseContext0PublicConnectionRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ReleaseContext0PublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleaseContext0PublicConnectionRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ReleaseContext0PublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseContext0PublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseContext0PublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseContext0PublicConnectionRequest) SetDBInstanceName(v string) *ReleaseContext0PublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseContext0PublicConnectionRequest) SetNodeType(v string) *ReleaseContext0PublicConnectionRequest {
	s.NodeType = &v
	return s
}

func (s *ReleaseContext0PublicConnectionRequest) SetRegionId(v string) *ReleaseContext0PublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseContext0PublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
