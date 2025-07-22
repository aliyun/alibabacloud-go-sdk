// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachExpressConnectRouterChildInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildInstanceId(v string) *AttachExpressConnectRouterChildInstanceRequest
	GetChildInstanceId() *string
	SetChildInstanceOwnerId(v int64) *AttachExpressConnectRouterChildInstanceRequest
	GetChildInstanceOwnerId() *int64
	SetChildInstanceRegionId(v string) *AttachExpressConnectRouterChildInstanceRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *AttachExpressConnectRouterChildInstanceRequest
	GetChildInstanceType() *string
	SetClientToken(v string) *AttachExpressConnectRouterChildInstanceRequest
	GetClientToken() *string
	SetDescription(v string) *AttachExpressConnectRouterChildInstanceRequest
	GetDescription() *string
	SetDryRun(v bool) *AttachExpressConnectRouterChildInstanceRequest
	GetDryRun() *bool
	SetEcrId(v string) *AttachExpressConnectRouterChildInstanceRequest
	GetEcrId() *string
}

type AttachExpressConnectRouterChildInstanceRequest struct {
	// The VBR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VBR belongs.
	//
	// >  If you want to connect to a network instance that belongs to a different account, this parameter is required.
	//
	// example:
	//
	// 190550144868****
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The region ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Set the value to **VBR**.
	//
	// This parameter is required.
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
	// The description of the sub-instance. It must be 0 to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// ecr-a5xqrgbeidz1w8****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s AttachExpressConnectRouterChildInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *AttachExpressConnectRouterChildInstanceRequest) GetChildInstanceOwnerId() *int64 {
	return s.ChildInstanceOwnerId
}

func (s *AttachExpressConnectRouterChildInstanceRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *AttachExpressConnectRouterChildInstanceRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *AttachExpressConnectRouterChildInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachExpressConnectRouterChildInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *AttachExpressConnectRouterChildInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AttachExpressConnectRouterChildInstanceRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceOwnerId(v int64) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceRegionId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetDescription(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.Description = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *AttachExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) Validate() error {
	return dara.Validate(s)
}
