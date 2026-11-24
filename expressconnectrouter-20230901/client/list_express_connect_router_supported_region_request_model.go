// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExpressConnectRouterSupportedRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListExpressConnectRouterSupportedRegionRequest
	GetClientToken() *string
	SetNodeType(v string) *ListExpressConnectRouterSupportedRegionRequest
	GetNodeType() *string
	SetVersion(v string) *ListExpressConnectRouterSupportedRegionRequest
	GetVersion() *string
}

type ListExpressConnectRouterSupportedRegionRequest struct {
	// A client token to ensure the idempotence of the request.
	//
	// Generate a value that is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- as the **ClientToken**. Each request has a unique **RequestId**.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The type of network instance. Valid values:
	//
	// - **TR**: transit router instance.
	//
	// - **VBR**: virtual border router instance.
	//
	// - **VPC**: virtual private cloud instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListExpressConnectRouterSupportedRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionRequest) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListExpressConnectRouterSupportedRegionRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ListExpressConnectRouterSupportedRegionRequest) GetVersion() *string {
	return s.Version
}

func (s *ListExpressConnectRouterSupportedRegionRequest) SetClientToken(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.ClientToken = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionRequest) SetNodeType(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.NodeType = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionRequest) SetVersion(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.Version = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionRequest) Validate() error {
	return dara.Validate(s)
}
