// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentVoiceprintsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAIAgentVoiceprintsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAIAgentVoiceprintsResponseBody
	GetTotalCount() *int32
	SetVoiceprints(v []*ListAIAgentVoiceprintsResponseBodyVoiceprints) *ListAIAgentVoiceprintsResponseBody
	GetVoiceprints() []*ListAIAgentVoiceprintsResponseBodyVoiceprints
}

type ListAIAgentVoiceprintsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of voiceprints that match the query criteria.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The voiceprints.
	Voiceprints []*ListAIAgentVoiceprintsResponseBodyVoiceprints `json:"Voiceprints,omitempty" xml:"Voiceprints,omitempty" type:"Repeated"`
}

func (s ListAIAgentVoiceprintsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentVoiceprintsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIAgentVoiceprintsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIAgentVoiceprintsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAIAgentVoiceprintsResponseBody) GetVoiceprints() []*ListAIAgentVoiceprintsResponseBodyVoiceprints {
	return s.Voiceprints
}

func (s *ListAIAgentVoiceprintsResponseBody) SetRequestId(v string) *ListAIAgentVoiceprintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIAgentVoiceprintsResponseBody) SetTotalCount(v int32) *ListAIAgentVoiceprintsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAIAgentVoiceprintsResponseBody) SetVoiceprints(v []*ListAIAgentVoiceprintsResponseBodyVoiceprints) *ListAIAgentVoiceprintsResponseBody {
	s.Voiceprints = v
	return s
}

func (s *ListAIAgentVoiceprintsResponseBody) Validate() error {
	if s.Voiceprints != nil {
		for _, item := range s.Voiceprints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAIAgentVoiceprintsResponseBodyVoiceprints struct {
	// The creation time of the voiceprint.
	//
	// example:
	//
	// 2025-07-28T10:03:58.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time of the voiceprint.
	//
	// example:
	//
	// 2025-07-28T10:03:58.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The unique identifier for the voiceprint.
	//
	// example:
	//
	// vp_1699123456_8527
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
}

func (s ListAIAgentVoiceprintsResponseBodyVoiceprints) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentVoiceprintsResponseBodyVoiceprints) GoString() string {
	return s.String()
}

func (s *ListAIAgentVoiceprintsResponseBodyVoiceprints) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAIAgentVoiceprintsResponseBodyVoiceprints) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAIAgentVoiceprintsResponseBodyVoiceprints) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *ListAIAgentVoiceprintsResponseBodyVoiceprints) SetGmtCreate(v string) *ListAIAgentVoiceprintsResponseBodyVoiceprints {
	s.GmtCreate = &v
	return s
}

func (s *ListAIAgentVoiceprintsResponseBodyVoiceprints) SetGmtModified(v string) *ListAIAgentVoiceprintsResponseBodyVoiceprints {
	s.GmtModified = &v
	return s
}

func (s *ListAIAgentVoiceprintsResponseBodyVoiceprints) SetVoiceprintId(v string) *ListAIAgentVoiceprintsResponseBodyVoiceprints {
	s.VoiceprintId = &v
	return s
}

func (s *ListAIAgentVoiceprintsResponseBodyVoiceprints) Validate() error {
	return dara.Validate(s)
}
