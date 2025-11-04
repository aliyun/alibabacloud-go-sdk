// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAIAgentVoiceprintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVoiceprintId(v string) *ClearAIAgentVoiceprintRequest
	GetVoiceprintId() *string
}

type ClearAIAgentVoiceprintRequest struct {
	// The unique identifier for the voiceprint.
	//
	// example:
	//
	// vp_1699123456_8527
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
}

func (s ClearAIAgentVoiceprintRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearAIAgentVoiceprintRequest) GoString() string {
	return s.String()
}

func (s *ClearAIAgentVoiceprintRequest) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *ClearAIAgentVoiceprintRequest) SetVoiceprintId(v string) *ClearAIAgentVoiceprintRequest {
	s.VoiceprintId = &v
	return s
}

func (s *ClearAIAgentVoiceprintRequest) Validate() error {
	return dara.Validate(s)
}
