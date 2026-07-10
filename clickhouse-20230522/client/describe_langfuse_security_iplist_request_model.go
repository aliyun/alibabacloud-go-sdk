// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseSecurityIPListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLangfuseSecurityIPListRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeLangfuseSecurityIPListRequest
	GetRegionId() *string
}

type DescribeLangfuseSecurityIPListRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLangfuseSecurityIPListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseSecurityIPListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseSecurityIPListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLangfuseSecurityIPListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLangfuseSecurityIPListRequest) SetDBInstanceId(v string) *DescribeLangfuseSecurityIPListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListRequest) SetRegionId(v string) *DescribeLangfuseSecurityIPListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListRequest) Validate() error {
	return dara.Validate(s)
}
