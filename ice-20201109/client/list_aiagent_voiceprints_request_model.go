// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentVoiceprintsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAIAgentVoiceprintsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAIAgentVoiceprintsRequest
	GetPageSize() *int32
	SetRegistrationMode(v string) *ListAIAgentVoiceprintsRequest
	GetRegistrationMode() *string
	SetVoiceprintId(v string) *ListAIAgentVoiceprintsRequest
	GetVoiceprintId() *string
}

type ListAIAgentVoiceprintsRequest struct {
	// The page number to return. Must be 1 or greater.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The registration mode for the voiceprint. The default value is `Explicit`.
	//
	// example:
	//
	// Explicit
	RegistrationMode *string `json:"RegistrationMode,omitempty" xml:"RegistrationMode,omitempty"`
	// The unique voiceprint ID. If specified, this operation returns the details of a single voiceprint. If omitted, it returns a paginated list of all voiceprints under your account.
	//
	// example:
	//
	// vp_1699123456_8527
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
}

func (s ListAIAgentVoiceprintsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentVoiceprintsRequest) GoString() string {
	return s.String()
}

func (s *ListAIAgentVoiceprintsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAIAgentVoiceprintsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAIAgentVoiceprintsRequest) GetRegistrationMode() *string {
	return s.RegistrationMode
}

func (s *ListAIAgentVoiceprintsRequest) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *ListAIAgentVoiceprintsRequest) SetPageNumber(v int32) *ListAIAgentVoiceprintsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAIAgentVoiceprintsRequest) SetPageSize(v int32) *ListAIAgentVoiceprintsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAIAgentVoiceprintsRequest) SetRegistrationMode(v string) *ListAIAgentVoiceprintsRequest {
	s.RegistrationMode = &v
	return s
}

func (s *ListAIAgentVoiceprintsRequest) SetVoiceprintId(v string) *ListAIAgentVoiceprintsRequest {
	s.VoiceprintId = &v
	return s
}

func (s *ListAIAgentVoiceprintsRequest) Validate() error {
	return dara.Validate(s)
}
