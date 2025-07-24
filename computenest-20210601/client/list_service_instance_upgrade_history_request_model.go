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
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-70a3b15bb62643xxxxxx
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
