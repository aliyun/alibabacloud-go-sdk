// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrustedOriginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTrustedOriginRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateTrustedOriginRequest
	GetInstanceId() *string
	SetTrustOriginName(v string) *UpdateTrustedOriginRequest
	GetTrustOriginName() *string
	SetTrustedOriginId(v string) *UpdateTrustedOriginRequest
	GetTrustedOriginId() *string
	SetTrustedOriginScene(v []*string) *UpdateTrustedOriginRequest
	GetTrustedOriginScene() []*string
}

type UpdateTrustedOriginRequest struct {
	// The idempotence token.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The trusted origin name.
	//
	// example:
	//
	// Qoder Production Console
	TrustOriginName *string `json:"TrustOriginName,omitempty" xml:"TrustOriginName,omitempty"`
	// The trusted origin ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// to_example
	TrustedOriginId *string `json:"TrustedOriginId,omitempty" xml:"TrustedOriginId,omitempty"`
	// The trusted origin scenario.
	//
	// example:
	//
	// iframe_embed
	TrustedOriginScene []*string `json:"TrustedOriginScene,omitempty" xml:"TrustedOriginScene,omitempty" type:"Repeated"`
}

func (s UpdateTrustedOriginRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrustedOriginRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrustedOriginRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTrustedOriginRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTrustedOriginRequest) GetTrustOriginName() *string {
	return s.TrustOriginName
}

func (s *UpdateTrustedOriginRequest) GetTrustedOriginId() *string {
	return s.TrustedOriginId
}

func (s *UpdateTrustedOriginRequest) GetTrustedOriginScene() []*string {
	return s.TrustedOriginScene
}

func (s *UpdateTrustedOriginRequest) SetClientToken(v string) *UpdateTrustedOriginRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTrustedOriginRequest) SetInstanceId(v string) *UpdateTrustedOriginRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTrustedOriginRequest) SetTrustOriginName(v string) *UpdateTrustedOriginRequest {
	s.TrustOriginName = &v
	return s
}

func (s *UpdateTrustedOriginRequest) SetTrustedOriginId(v string) *UpdateTrustedOriginRequest {
	s.TrustedOriginId = &v
	return s
}

func (s *UpdateTrustedOriginRequest) SetTrustedOriginScene(v []*string) *UpdateTrustedOriginRequest {
	s.TrustedOriginScene = v
	return s
}

func (s *UpdateTrustedOriginRequest) Validate() error {
	return dara.Validate(s)
}
