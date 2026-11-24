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
	// The ID used to associate the Express Connect Router (ECR) with the VBR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The instance ID of the network instance to be queried.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The region where the network instance resides.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values: **VBR**: virtual border router.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client and ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether this request is only a dry run. Valid values:
	//
	// - **true**: sends a check request without modifying the EPS resource.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Instance ID of the Express Connect Router (ECR) instance.
	//
	// > You can call the DescribeExpressConnectRouter API operation to obtain instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If this is the first query or no next query is required, leave this parameter empty.
	//
	// - If a next query is required, set the value to the NextToken value returned from the previous API call.
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
