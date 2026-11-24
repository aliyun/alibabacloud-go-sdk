// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterAllowedPrefixHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest
	GetAssociationId() *string
	SetClientToken(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest
	GetEcrId() *string
	SetInstanceId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest
	GetInstanceType() *string
	SetVersion(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest
	GetVersion() *string
}

type DescribeExpressConnectRouterAllowedPrefixHistoryRequest struct {
	// The instance ID of the VPC or transit router (TR) associated with the Express Connect Router (ECR) for forward routing.
	//
	// > You must specify either the **InstanceId*	- parameter or the **AssociationId*	- parameter.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client to ensure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether it is only a dry run for this request. Valid values:
	//
	// - **true**: Performs a dry run for the request without modifying EPS resources.
	//
	// - **false*	- (default): Sends a normal request.
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
	// The instance ID to load into the Express Connect Router (ECR).
	//
	// > You must specify either the **InstanceId*	- parameter or the **AssociationId*	- parameter.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// - **VPC**: VPC instance.
	//
	// - **TR**: Transit router instance.
	//
	// example:
	//
	// VPC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetAssociationId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetClientToken(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetDryRun(v bool) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetEcrId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetInstanceId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetInstanceType(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetVersion(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.Version = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) Validate() error {
	return dara.Validate(s)
}
