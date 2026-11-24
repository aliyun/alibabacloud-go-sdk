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
	SetVersion(v string) *DisableExpressConnectRouterRouteEntriesRequest
	GetVersion() *string
}

type DisableExpressConnectRouterRouteEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination CIDR block of the route entry in the gateway route table.
	//
	// This parameter is required.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether the request is a dry run. Valid values:
	//
	// - **true**: sends a check request and does not modify the EPS resource.
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
	// The next hop instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-hp3u4u5f03tfuljis****
	NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

func (s *DisableExpressConnectRouterRouteEntriesRequest) GetVersion() *string {
	return s.Version
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

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetVersion(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.Version = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
