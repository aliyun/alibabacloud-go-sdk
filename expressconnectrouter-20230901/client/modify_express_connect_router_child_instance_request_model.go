// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterChildInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildInstanceId(v string) *ModifyExpressConnectRouterChildInstanceRequest
	GetChildInstanceId() *string
	SetChildInstanceType(v string) *ModifyExpressConnectRouterChildInstanceRequest
	GetChildInstanceType() *string
	SetClientToken(v string) *ModifyExpressConnectRouterChildInstanceRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyExpressConnectRouterChildInstanceRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyExpressConnectRouterChildInstanceRequest
	GetDryRun() *bool
	SetEcrId(v string) *ModifyExpressConnectRouterChildInstanceRequest
	GetEcrId() *string
	SetVersion(v string) *ModifyExpressConnectRouterChildInstanceRequest
	GetVersion() *string
}

type ModifyExpressConnectRouterChildInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vbr-t4n6xu2d5q0iaad1yl4le
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecr-gny3gqp41n7vrrn5iz
	EcrId   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ModifyExpressConnectRouterChildInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) GetVersion() *string {
	return s.Version
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *ModifyExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *ModifyExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *ModifyExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) SetDescription(v string) *ModifyExpressConnectRouterChildInstanceRequest {
	s.Description = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *ModifyExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *ModifyExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) SetVersion(v string) *ModifyExpressConnectRouterChildInstanceRequest {
	s.Version = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceRequest) Validate() error {
	return dara.Validate(s)
}
