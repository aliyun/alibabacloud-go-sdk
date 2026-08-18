// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateContextDBRequest
	GetDBInstanceName() *string
	SetOpenSearchInstanceName(v string) *CreateContextDBRequest
	GetOpenSearchInstanceName() *string
	SetRegionId(v string) *CreateContextDBRequest
	GetRegionId() *string
}

type CreateContextDBRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the PolarDB-X Search instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxs-********
	OpenSearchInstanceName *string `json:"OpenSearchInstanceName,omitempty" xml:"OpenSearchInstanceName,omitempty"`
	// The ID of the region where the instance resides. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateContextDBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDBRequest) GoString() string {
	return s.String()
}

func (s *CreateContextDBRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateContextDBRequest) GetOpenSearchInstanceName() *string {
	return s.OpenSearchInstanceName
}

func (s *CreateContextDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateContextDBRequest) SetDBInstanceName(v string) *CreateContextDBRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateContextDBRequest) SetOpenSearchInstanceName(v string) *CreateContextDBRequest {
	s.OpenSearchInstanceName = &v
	return s
}

func (s *CreateContextDBRequest) SetRegionId(v string) *CreateContextDBRequest {
	s.RegionId = &v
	return s
}

func (s *CreateContextDBRequest) Validate() error {
	return dara.Validate(s)
}
