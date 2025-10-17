// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *BeginSessionRequest
	GetAgentKey() *string
	SetInstanceId(v string) *BeginSessionRequest
	GetInstanceId() *string
	SetSandBox(v bool) *BeginSessionRequest
	GetSandBox() *bool
	SetVendorParam(v string) *BeginSessionRequest
	GetVendorParam() *string
}

type BeginSessionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SandBox     *bool   `json:"SandBox,omitempty" xml:"SandBox,omitempty"`
	VendorParam *string `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
}

func (s BeginSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s BeginSessionRequest) GoString() string {
	return s.String()
}

func (s *BeginSessionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *BeginSessionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BeginSessionRequest) GetSandBox() *bool {
	return s.SandBox
}

func (s *BeginSessionRequest) GetVendorParam() *string {
	return s.VendorParam
}

func (s *BeginSessionRequest) SetAgentKey(v string) *BeginSessionRequest {
	s.AgentKey = &v
	return s
}

func (s *BeginSessionRequest) SetInstanceId(v string) *BeginSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *BeginSessionRequest) SetSandBox(v bool) *BeginSessionRequest {
	s.SandBox = &v
	return s
}

func (s *BeginSessionRequest) SetVendorParam(v string) *BeginSessionRequest {
	s.VendorParam = &v
	return s
}

func (s *BeginSessionRequest) Validate() error {
	return dara.Validate(s)
}
