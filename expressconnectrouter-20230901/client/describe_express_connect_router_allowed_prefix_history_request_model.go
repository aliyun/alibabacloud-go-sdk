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
}

type DescribeExpressConnectRouterAllowedPrefixHistoryRequest struct {
	// The ID of the association between the ECR and the virtual private cloud (VPC) or transit router (TR).
	//
	// >  You must specify either **InstanceId*	- or **AssociationId**.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
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
	// The ID of the network instance that is associated with the ECR.
	//
	// >  You must specify either **InstanceId*	- or **AssociationId**.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **TR**
	//
	// example:
	//
	// VPC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
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

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) Validate() error {
	return dara.Validate(s)
}
