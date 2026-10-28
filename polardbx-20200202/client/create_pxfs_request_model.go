// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassCode(v string) *CreatePxfsRequest
	GetClassCode() *string
	SetDBInstanceName(v string) *CreatePxfsRequest
	GetDBInstanceName() *string
	SetNodeCount(v int64) *CreatePxfsRequest
	GetNodeCount() *int64
	SetRegionId(v string) *CreatePxfsRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *CreatePxfsRequest
	GetSecurityIPList() *string
	SetStorageSize(v int64) *CreatePxfsRequest
	GetStorageSize() *int64
}

type CreatePxfsRequest struct {
	// The node specifications.
	//
	// example:
	//
	// polarx.pxfs.x4.xlarge
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 3
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP whitelist.
	//
	// example:
	//
	// 192.168.1.0/24,10.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The storage capacity.
	//
	// example:
	//
	// 500
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s CreatePxfsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsRequest) GoString() string {
	return s.String()
}

func (s *CreatePxfsRequest) GetClassCode() *string {
	return s.ClassCode
}

func (s *CreatePxfsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreatePxfsRequest) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *CreatePxfsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePxfsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreatePxfsRequest) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *CreatePxfsRequest) SetClassCode(v string) *CreatePxfsRequest {
	s.ClassCode = &v
	return s
}

func (s *CreatePxfsRequest) SetDBInstanceName(v string) *CreatePxfsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePxfsRequest) SetNodeCount(v int64) *CreatePxfsRequest {
	s.NodeCount = &v
	return s
}

func (s *CreatePxfsRequest) SetRegionId(v string) *CreatePxfsRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePxfsRequest) SetSecurityIPList(v string) *CreatePxfsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreatePxfsRequest) SetStorageSize(v int64) *CreatePxfsRequest {
	s.StorageSize = &v
	return s
}

func (s *CreatePxfsRequest) Validate() error {
	return dara.Validate(s)
}
