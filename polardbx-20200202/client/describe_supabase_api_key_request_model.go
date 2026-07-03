// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeSupabaseApiKeyRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeSupabaseApiKeyRequest
	GetRegionId() *string
}

type DescribeSupabaseApiKeyRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSupabaseApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseApiKeyRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSupabaseApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSupabaseApiKeyRequest) SetDBInstanceName(v string) *DescribeSupabaseApiKeyRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSupabaseApiKeyRequest) SetRegionId(v string) *DescribeSupabaseApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSupabaseApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
