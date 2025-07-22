// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *DescribeExpressConnectRouterAssociationRequest
	GetAssociationId() *string
	SetAssociationNodeType(v string) *DescribeExpressConnectRouterAssociationRequest
	GetAssociationNodeType() *string
	SetAssociationRegionId(v string) *DescribeExpressConnectRouterAssociationRequest
	GetAssociationRegionId() *string
	SetCenId(v string) *DescribeExpressConnectRouterAssociationRequest
	GetCenId() *string
	SetClientToken(v string) *DescribeExpressConnectRouterAssociationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeExpressConnectRouterAssociationRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeExpressConnectRouterAssociationRequest
	GetEcrId() *string
	SetMaxResults(v int32) *DescribeExpressConnectRouterAssociationRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeExpressConnectRouterAssociationRequest
	GetNextToken() *string
	SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationRequest
	GetTransitRouterId() *string
	SetVpcId(v string) *DescribeExpressConnectRouterAssociationRequest
	GetVpcId() *string
}

type DescribeExpressConnectRouterAssociationRequest struct {
	// The ID of the association between the ECR and the VPC or TR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The type of the associated resource. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **TR**
	//
	// example:
	//
	// VPC
	AssociationNodeType *string `json:"AssociationNodeType,omitempty" xml:"AssociationNodeType,omitempty"`
	// The region ID of the VPC or TR.
	//
	// example:
	//
	// cn-hangzhou
	AssociationRegionId *string `json:"AssociationRegionId,omitempty" xml:"AssociationRegionId,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-of3o1the3i4gwb****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 20.
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
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The TR ID.
	//
	// example:
	//
	// tr-2ze4i71c6be454e2l****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1h37fchc6jmfyln****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetAssociationNodeType() *string {
	return s.AssociationNodeType
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetAssociationRegionId() *string {
	return s.AssociationRegionId
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationNodeType(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationNodeType = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationRegionId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetCenId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.CenId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetClientToken(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetDryRun(v bool) *DescribeExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetEcrId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterAssociationRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetNextToken(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetVpcId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) Validate() error {
	return dara.Validate(s)
}
