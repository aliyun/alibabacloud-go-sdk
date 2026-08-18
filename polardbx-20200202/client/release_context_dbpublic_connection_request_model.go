// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseContextDBPublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentConnectionString(v string) *ReleaseContextDBPublicConnectionRequest
	GetCurrentConnectionString() *string
	SetDBInstanceName(v string) *ReleaseContextDBPublicConnectionRequest
	GetDBInstanceName() *string
	SetNodeType(v string) *ReleaseContextDBPublicConnectionRequest
	GetNodeType() *string
	SetRegionId(v string) *ReleaseContextDBPublicConnectionRequest
	GetRegionId() *string
}

type ReleaseContextDBPublicConnectionRequest struct {
	// The public network connection string to release. This parameter is optional. If you do not specify this parameter, the Mem0 public endpoint is subject to automatic release.
	//
	// example:
	//
	// pxc-hzjasdyuoo.polarx.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0****r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The target node type: service or dashboard.
	//
	// example:
	//
	// service
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseContextDBPublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContextDBPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseContextDBPublicConnectionRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ReleaseContextDBPublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleaseContextDBPublicConnectionRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ReleaseContextDBPublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseContextDBPublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseContextDBPublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionRequest) SetDBInstanceName(v string) *ReleaseContextDBPublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionRequest) SetNodeType(v string) *ReleaseContextDBPublicConnectionRequest {
	s.NodeType = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionRequest) SetRegionId(v string) *ReleaseContextDBPublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
