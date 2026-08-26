// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrustedOriginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTrustedOriginRequest
	GetClientToken() *string
	SetInstanceId(v string) *CreateTrustedOriginRequest
	GetInstanceId() *string
	SetOrigin(v string) *CreateTrustedOriginRequest
	GetOrigin() *string
	SetTrustOriginName(v string) *CreateTrustedOriginRequest
	GetTrustOriginName() *string
	SetTrustedOriginScene(v []*string) *CreateTrustedOriginRequest
	GetTrustedOriginScene() []*string
}

type CreateTrustedOriginRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The browser origin in the format of scheme://host[:port]. This value cannot be modified after creation.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://console.qoder.com
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The name of the trusted origin used for management, display, and auditing. If not specified, the value is empty.
	//
	// example:
	//
	// Qoder Production Console
	TrustOriginName *string `json:"TrustOriginName,omitempty" xml:"TrustOriginName,omitempty"`
	// The trusted origin scenes. Only iframe_embed and cors are supported. At least one value is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// iframe_embed
	TrustedOriginScene []*string `json:"TrustedOriginScene,omitempty" xml:"TrustedOriginScene,omitempty" type:"Repeated"`
}

func (s CreateTrustedOriginRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrustedOriginRequest) GoString() string {
	return s.String()
}

func (s *CreateTrustedOriginRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTrustedOriginRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTrustedOriginRequest) GetOrigin() *string {
	return s.Origin
}

func (s *CreateTrustedOriginRequest) GetTrustOriginName() *string {
	return s.TrustOriginName
}

func (s *CreateTrustedOriginRequest) GetTrustedOriginScene() []*string {
	return s.TrustedOriginScene
}

func (s *CreateTrustedOriginRequest) SetClientToken(v string) *CreateTrustedOriginRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTrustedOriginRequest) SetInstanceId(v string) *CreateTrustedOriginRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTrustedOriginRequest) SetOrigin(v string) *CreateTrustedOriginRequest {
	s.Origin = &v
	return s
}

func (s *CreateTrustedOriginRequest) SetTrustOriginName(v string) *CreateTrustedOriginRequest {
	s.TrustOriginName = &v
	return s
}

func (s *CreateTrustedOriginRequest) SetTrustedOriginScene(v []*string) *CreateTrustedOriginRequest {
	s.TrustedOriginScene = v
	return s
}

func (s *CreateTrustedOriginRequest) Validate() error {
	return dara.Validate(s)
}
