// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAIAgentVoiceprintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAIAgentVoiceprintResponseBody
	GetRequestId() *string
	SetVoiceprintId(v string) *SetAIAgentVoiceprintResponseBody
	GetVoiceprintId() *string
}

type SetAIAgentVoiceprintResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 550e8400********55440000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the registered voiceprint.
	//
	// example:
	//
	// vp_1699123456_8527
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
}

func (s SetAIAgentVoiceprintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAIAgentVoiceprintResponseBody) GoString() string {
	return s.String()
}

func (s *SetAIAgentVoiceprintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAIAgentVoiceprintResponseBody) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *SetAIAgentVoiceprintResponseBody) SetRequestId(v string) *SetAIAgentVoiceprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAIAgentVoiceprintResponseBody) SetVoiceprintId(v string) *SetAIAgentVoiceprintResponseBody {
	s.VoiceprintId = &v
	return s
}

func (s *SetAIAgentVoiceprintResponseBody) Validate() error {
	return dara.Validate(s)
}
