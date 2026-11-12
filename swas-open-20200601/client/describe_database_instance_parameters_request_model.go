// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstanceParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseInstanceId(v string) *DescribeDatabaseInstanceParametersRequest
	GetDatabaseInstanceId() *string
	SetRegionId(v string) *DescribeDatabaseInstanceParametersRequest
	GetRegionId() *string
}

type DescribeDatabaseInstanceParametersRequest struct {
	// The ID of the lightweight database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The ID of the region where the lightweight database instance is located. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDatabaseInstanceParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *DescribeDatabaseInstanceParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDatabaseInstanceParametersRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseInstanceParametersRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersRequest) SetRegionId(v string) *DescribeDatabaseInstanceParametersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersRequest) Validate() error {
	return dara.Validate(s)
}
