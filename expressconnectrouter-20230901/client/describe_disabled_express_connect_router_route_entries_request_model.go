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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
