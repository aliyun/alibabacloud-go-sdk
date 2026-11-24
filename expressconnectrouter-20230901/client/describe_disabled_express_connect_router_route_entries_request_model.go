// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisabledExpressConnectRouterRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeDisabledExpressConnectRouterRouteEntriesRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest
	GetEcrId() *string
	SetMaxResults(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest
	GetNextToken() *string
	SetVersion(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest
	GetVersion() *string
}

type DescribeDisabledExpressConnectRouterRouteEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether the request is a dry run. Valid values:
	//
	// - **true**: performs a dry run on the request without modifying the EPS resources.
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
	// - If this is the first query or no next query is required, you do not need to specify this parameter.
	//
	// - If a next query is required, set the value to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// FFv4fzkNPW8Z+cZ+DBXXQ3Gmf3XlCgpBH43oaTYTAAcGc708Zb+pDyAGVJBo/MKsyrtZfPnX9Ztf02vgdIDyaNe8UuZdf/JJk069qxGKzqSKg=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetMaxResults(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetNextToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetVersion(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.Version = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
