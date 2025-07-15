// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailoverTestJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeFailoverTestJobsRequest
	GetClientToken() *string
	SetFilter(v []*DescribeFailoverTestJobsRequestFilter) *DescribeFailoverTestJobsRequest
	GetFilter() []*DescribeFailoverTestJobsRequestFilter
	SetMaxResults(v int32) *DescribeFailoverTestJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFailoverTestJobsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeFailoverTestJobsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeFailoverTestJobsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeFailoverTestJobsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeFailoverTestJobsRequest
	GetResourceOwnerAccount() *string
}

type DescribeFailoverTestJobsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The filter information.
	Filter []*DescribeFailoverTestJobsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If the value of **NextToken*	- is not returned, it indicates that no next query is to be sent.
	//
	// 	- If a value of **NextToken*	- is returned, the value is the token that is used for the subsequent query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where you want to perform the failover test.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeFailoverTestJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeFailoverTestJobsRequest) GetFilter() []*DescribeFailoverTestJobsRequestFilter {
	return s.Filter
}

func (s *DescribeFailoverTestJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFailoverTestJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFailoverTestJobsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeFailoverTestJobsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeFailoverTestJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFailoverTestJobsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeFailoverTestJobsRequest) SetClientToken(v string) *DescribeFailoverTestJobsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFailoverTestJobsRequest) SetFilter(v []*DescribeFailoverTestJobsRequestFilter) *DescribeFailoverTestJobsRequest {
	s.Filter = v
	return s
}

func (s *DescribeFailoverTestJobsRequest) SetMaxResults(v int32) *DescribeFailoverTestJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFailoverTestJobsRequest) SetNextToken(v string) *DescribeFailoverTestJobsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFailoverTestJobsRequest) SetOwnerAccount(v string) *DescribeFailoverTestJobsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFailoverTestJobsRequest) SetOwnerId(v int64) *DescribeFailoverTestJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFailoverTestJobsRequest) SetRegionId(v string) *DescribeFailoverTestJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFailoverTestJobsRequest) SetResourceOwnerAccount(v string) *DescribeFailoverTestJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeFailoverTestJobsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeFailoverTestJobsRequestFilter struct {
	// The filter key. Valid values:
	//
	// 	- **JobId**
	//
	// 	- **JobName**
	//
	// 	- **JobStatus**
	//
	// 	- **ResourceId**
	//
	// 	- **ResourceName**
	//
	// 	- **ResourceType**
	//
	// > You can specify at most five different filter keys. If you specify ResourceId or ResourceName, you must also specify ResourceType. The logical operator among the filter keys is AND. Results that meet all specified filter keys are returned.
	//
	// example:
	//
	// JobId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter key.
	//
	// > You can specify at most five filter values for each filter key. The logical operator among filter values is OR. If a filter value is matched, the filter key is considered matched.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeFailoverTestJobsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeFailoverTestJobsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeFailoverTestJobsRequestFilter) SetKey(v string) *DescribeFailoverTestJobsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeFailoverTestJobsRequestFilter) SetValue(v []*string) *DescribeFailoverTestJobsRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeFailoverTestJobsRequestFilter) Validate() error {
	return dara.Validate(s)
}
