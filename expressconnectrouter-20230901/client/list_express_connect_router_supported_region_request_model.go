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
}

type ListExpressConnectRouterSupportedRegionRequest struct {
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
	// The type of the network instance. Valid values:
	//
	// 	- **TR**
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
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

func (s *ListExpressConnectRouterSupportedRegionRequest) SetClientToken(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.ClientToken = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionRequest) SetNodeType(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.NodeType = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionRequest) Validate() error {
	return dara.Validate(s)
}
