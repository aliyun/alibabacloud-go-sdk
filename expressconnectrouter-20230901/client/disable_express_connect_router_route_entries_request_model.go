// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableExpressConnectRouterRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableExpressConnectRouterRouteEntriesRequest
	GetClientToken() *string
	SetDestinationCidrBlock(v string) *DisableExpressConnectRouterRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *DisableExpressConnectRouterRouteEntriesRequest
	GetDryRun() *bool
	SetEcrId(v string) *DisableExpressConnectRouterRouteEntriesRequest
	GetEcrId() *string
	SetNexthopInstanceId(v string) *DisableExpressConnectRouterRouteEntriesRequest
	GetNexthopInstanceId() *string
}

type DisableExpressConnectRouterRouteEntriesRequest struct {
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
	// The destination CIDR block of the route entry in the route table of the ECR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.153.32.16/28
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
	// The ID of the next-hop instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
}

func (s DisableExpressConnectRouterRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) GetNexthopInstanceId() *string {
	return s.NexthopInstanceId
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DisableExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.NexthopInstanceId = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
