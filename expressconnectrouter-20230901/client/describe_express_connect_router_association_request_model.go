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
	SetTag(v []*DescribeExpressConnectRouterAssociationRequestTag) *DescribeExpressConnectRouterAssociationRequest
	GetTag() []*DescribeExpressConnectRouterAssociationRequestTag
	SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationRequest
	GetTransitRouterId() *string
	SetVersion(v string) *DescribeExpressConnectRouterAssociationRequest
	GetVersion() *string
	SetVpcId(v string) *DescribeExpressConnectRouterAssociationRequest
	GetVpcId() *string
}

type DescribeExpressConnectRouterAssociationRequest struct {
	// The ID used to associate the Express Connect Router (ECR) with a VPC or TR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The type of the associated resource. Valid values:
	//
	// - **VPC**: VPC.
	//
	// - **TR**: transit router.
	//
	// example:
	//
	// VPC
	AssociationNodeType *string `json:"AssociationNodeType,omitempty" xml:"AssociationNodeType,omitempty"`
	// The region where the VPC or TR resides.
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
	// You can generate the value from your client. Make sure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether this is a dry run. Valid values:
	//
	// - **true**: sends a check request and does not modify service resources.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// > You can call the CreateExpressConnectRouter API operation to create an ECR instance and obtain the EcrId.
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
	// The list of tags.
	//
	// You can add up to 20 tags.
	Tag []*DescribeExpressConnectRouterAssociationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-2ze4i71c6be454e2l****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the VPC-connected instance.
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

func (s *DescribeExpressConnectRouterAssociationRequest) GetTag() []*DescribeExpressConnectRouterAssociationRequestTag {
	return s.Tag
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeExpressConnectRouterAssociationRequest) GetVersion() *string {
	return s.Version
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

func (s *DescribeExpressConnectRouterAssociationRequest) SetTag(v []*DescribeExpressConnectRouterAssociationRequestTag) *DescribeExpressConnectRouterAssociationRequest {
	s.Tag = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetVersion(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.Version = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetVpcId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExpressConnectRouterAssociationRequestTag struct {
	// The tag key of the resource. The tag key cannot be an empty string. It can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. The tag value can be an empty string. It can be up to 128 characters in length, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeExpressConnectRouterAssociationRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeExpressConnectRouterAssociationRequestTag) SetKey(v string) *DescribeExpressConnectRouterAssociationRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequestTag) SetValue(v string) *DescribeExpressConnectRouterAssociationRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequestTag) Validate() error {
	return dara.Validate(s)
}
