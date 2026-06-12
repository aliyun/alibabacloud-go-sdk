// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceUpgradeHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceInstanceUpgradeHistoryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceUpgradeHistoryRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceInstanceUpgradeHistoryRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *ListServiceInstanceUpgradeHistoryRequest
	GetServiceInstanceId() *string
}

type ListServiceInstanceUpgradeHistoryRequest struct {
	// The number of entries to return on each page for a paged query. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Set this to the NextToken value from a previous response.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceUpgradeHistoryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceUpgradeHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceInstanceUpgradeHistoryRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetMaxResults(v int32) *ListServiceInstanceUpgradeHistoryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetNextToken(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetRegionId(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetServiceInstanceId(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) Validate() error {
	return dara.Validate(s)
}
