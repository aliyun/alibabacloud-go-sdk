// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterChildInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *DescribeExpressConnectRouterChildInstanceRequest
	GetAssociationId() *string
	SetChildInstanceId(v string) *DescribeExpressConnectRouterChildInstanceRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *DescribeExpressConnectRouterChildInstanceRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *DescribeExpressConnectRouterChildInstanceRequest
	GetChildInstanceType() *string
	SetClientToken(v string) *DescribeExpressConnectRouterChildInstanceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeExpressConnectRouterChildInstanceRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeExpressConnectRouterChildInstanceRequest
	GetEcrId() *string
	SetMaxResults(v int32) *DescribeExpressConnectRouterChildInstanceRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeExpressConnectRouterChildInstanceRequest
	GetNextToken() *string
	SetVersion(v string) *DescribeExpressConnectRouterChildInstanceRequest
	GetVersion() *string
}

type DescribeExpressConnectRouterChildInstanceRequest struct {
	// The ID of the association between the ECR and the virtual private cloud (VPC) or transit router (TR).
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The ID of the network instance to be queried.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The region ID of the network instance.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of network instance. Set the value to VBR.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
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
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- If a value of NextToken is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Zn0J0/sCjivzwX9oBcwFnWaaas/kSG+J/WzLOxJHS4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetAssociationId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceRegionId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *DescribeExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterChildInstanceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetNextToken(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetVersion(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.Version = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) Validate() error {
	return dara.Validate(s)
}
