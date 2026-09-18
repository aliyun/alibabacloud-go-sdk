// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateRelayPollerScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateRelayPollerScriptResponseBodyData) *GenerateRelayPollerScriptResponseBody
	GetData() *GenerateRelayPollerScriptResponseBodyData
	SetRequestId(v string) *GenerateRelayPollerScriptResponseBody
	GetRequestId() *string
}

type GenerateRelayPollerScriptResponseBody struct {
	// The generation result, which contains the target identifier, the normalized platform, and the installation script.
	Data *GenerateRelayPollerScriptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can use this ID for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateRelayPollerScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateRelayPollerScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateRelayPollerScriptResponseBody) GetData() *GenerateRelayPollerScriptResponseBodyData {
	return s.Data
}

func (s *GenerateRelayPollerScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateRelayPollerScriptResponseBody) SetData(v *GenerateRelayPollerScriptResponseBodyData) *GenerateRelayPollerScriptResponseBody {
	s.Data = v
	return s
}

func (s *GenerateRelayPollerScriptResponseBody) SetRequestId(v string) *GenerateRelayPollerScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateRelayPollerScriptResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateRelayPollerScriptResponseBodyData struct {
	// The normalized target platform in the operating system-architecture format.
	//
	// example:
	//
	// linux-amd64
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The installation script content (bash). The script contains a one-time access code, a temporary download link for the poller binary (signed and valid for 1 hour), and a checksum. The script does not contain the actual endpoint or credentials of the target, which are interactively entered during installation. The script carries access credential semantics. Transmit it through a trusted channel and re-download the script to obtain a new one after use.
	//
	// example:
	//
	// #!/bin/bash
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The scan target identifier echoed from the request.
	//
	// example:
	//
	// target-abc123def4567
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s GenerateRelayPollerScriptResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateRelayPollerScriptResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateRelayPollerScriptResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *GenerateRelayPollerScriptResponseBodyData) GetScript() *string {
	return s.Script
}

func (s *GenerateRelayPollerScriptResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *GenerateRelayPollerScriptResponseBodyData) SetPlatform(v string) *GenerateRelayPollerScriptResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GenerateRelayPollerScriptResponseBodyData) SetScript(v string) *GenerateRelayPollerScriptResponseBodyData {
	s.Script = &v
	return s
}

func (s *GenerateRelayPollerScriptResponseBodyData) SetTargetId(v string) *GenerateRelayPollerScriptResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *GenerateRelayPollerScriptResponseBodyData) Validate() error {
	return dara.Validate(s)
}
