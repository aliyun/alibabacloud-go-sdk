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
	SetVersion(v string) *DescribeExpressConnectRouterRouteEntriesRequest
	GetVersion() *string
}

type DescribeExpressConnectRouterRouteEntriesRequest struct {
	// The AS path traversed by the route entry.
	//
	// example:
	//
	// [64993,64512]
	AsPath *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value from your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The community value carried by the BGP route entry.
	//
	// example:
	//
	// 9001:9263
	Community *string `json:"Community,omitempty" xml:"Community,omitempty"`
	// The destination CIDR block of the route entry to query.
	//
	// example:
	//
	// 172.20.4.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether it is a dry run for this request. Valid values:
	//
	// - **true**: sends a check request and does not modify EPS resources.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - For the first query or when no subsequent query is required, leave this parameter empty.
	//
	// - If a subsequent query is required, set the value to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The next hop instance ID.
	//
	// example:
	//
	// vbr-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	// The region where the Express Connect Router (ECR) instance resides.
	//
	// example:
	//
	// cn-hangzhou
	QueryRegionId *string `json:"QueryRegionId,omitempty" xml:"QueryRegionId,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

func (s *DescribeExpressConnectRouterRouteEntriesRequest) GetVersion() *string {
	return s.Version
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

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetVersion(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.Version = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
