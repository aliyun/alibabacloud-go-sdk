// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsPath(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetAsPath() *string
	SetClientToken(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetClientToken() *string
	SetCommunity(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetCommunity() *string
	SetDestinationCidrBlock(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *DescribeExpressConnectRouterRouteEntriesRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetEcrId() *string
	SetMaxResults(v int32) *DescribeExpressConnectRouterRouteEntriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetNextToken() *string
	SetNexthopInstanceId(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetNexthopInstanceId() *string
	SetQueryRegionId(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetQueryRegionId() *string
}

type DescribeExpressConnectRouterRouteEntriesRequest struct {
	// The Autonomous System (AS) path of the route.
	//
	// example:
	//
	// [64993,64512]
	AsPath *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The community value that is carried in the Border Gateway Protocol (BGP) route.
	//
	// example:
	//
	// 9001:9263
	Community *string `json:"Community,omitempty" xml:"Community,omitempty"`
	// The destination CIDR block of the route that you want to query.
	//
	// example:
	//
	// 172.20.47.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the next-hop instance.
	//
	// example:
	//
	// br-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	// The region ID of the ECR.
	//
	// example:
	//
	// cn-hangzhou
	QueryRegionId *string `json:"QueryRegionId,omitempty" xml:"QueryRegionId,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetAsPath() *string {
	return s.AsPath
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetCommunity() *string {
	return s.Community
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetNexthopInstanceId() *string {
	return s.NexthopInstanceId
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetQueryRegionId() *string {
	return s.QueryRegionId
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetAsPath(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.AsPath = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetCommunity(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.Community = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetNextToken(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetQueryRegionId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.QueryRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
