// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateRelayPollerScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlatform(v string) *GenerateRelayPollerScriptRequest
	GetPlatform() *string
	SetTargetId(v string) *GenerateRelayPollerScriptRequest
	GetTargetId() *string
}

type GenerateRelayPollerScriptRequest struct {
	// The target platform in the "operating system-architecture" format. Only linux-amd64 and linux-arm64 are supported. Compatible architecture values include amd64, x86_64, x86, arm64, and aarch64. If only the architecture is specified, the operating system defaults to linux. Other operating systems such as macOS and Windows return HTTP status code 400. If this parameter is not specified, the default value is linux-amd64.
	//
	// example:
	//
	// linux-amd64
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The unique identifier of the scan target. The target must use the enterprise_relay connection method (see CreateAttackTarget). Otherwise, HTTP status code 400 is returned. If the target does not exist or belongs to another tenant, HTTP status code 400 is returned without exposing whether the resource exists. This parameter is registered as optional but is required in practice.
	//
	// example:
	//
	// target-abc123def4567
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s GenerateRelayPollerScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateRelayPollerScriptRequest) GoString() string {
	return s.String()
}

func (s *GenerateRelayPollerScriptRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GenerateRelayPollerScriptRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *GenerateRelayPollerScriptRequest) SetPlatform(v string) *GenerateRelayPollerScriptRequest {
	s.Platform = &v
	return s
}

func (s *GenerateRelayPollerScriptRequest) SetTargetId(v string) *GenerateRelayPollerScriptRequest {
	s.TargetId = &v
	return s
}

func (s *GenerateRelayPollerScriptRequest) Validate() error {
	return dara.Validate(s)
}
