// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBindRamUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CheckBindRamUserRequest
	GetDBClusterId() *string
	SetRegionId(v string) *CheckBindRamUserRequest
	GetRegionId() *string
}

type CheckBindRamUserRequest struct {
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the cluster ID.
	//
	// example:
	//
	// amv-wz9842849v6****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckBindRamUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckBindRamUserRequest) GoString() string {
	return s.String()
}

func (s *CheckBindRamUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckBindRamUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckBindRamUserRequest) SetDBClusterId(v string) *CheckBindRamUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckBindRamUserRequest) SetRegionId(v string) *CheckBindRamUserRequest {
	s.RegionId = &v
	return s
}

func (s *CheckBindRamUserRequest) Validate() error {
	return dara.Validate(s)
}
