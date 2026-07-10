// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLangfuseEndpointsRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeLangfuseEndpointsRequest
	GetRegionId() *string
}

type DescribeLangfuseEndpointsRequest struct {
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

func (s DescribeLangfuseEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseEndpointsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLangfuseEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLangfuseEndpointsRequest) SetDBInstanceId(v string) *DescribeLangfuseEndpointsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLangfuseEndpointsRequest) SetRegionId(v string) *DescribeLangfuseEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLangfuseEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
