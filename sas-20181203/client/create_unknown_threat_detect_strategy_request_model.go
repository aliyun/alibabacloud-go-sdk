// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUnknownThreatDetectStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetSelectionType(v string) *CreateUnknownThreatDetectStrategyRequest
	GetAssetSelectionType() *string
	SetDurationDaysAfterInit(v int32) *CreateUnknownThreatDetectStrategyRequest
	GetDurationDaysAfterInit() *int32
	SetDurationDaysAfterStop(v int32) *CreateUnknownThreatDetectStrategyRequest
	GetDurationDaysAfterStop() *int32
	SetName(v string) *CreateUnknownThreatDetectStrategyRequest
	GetName() *string
	SetStudyMode(v string) *CreateUnknownThreatDetectStrategyRequest
	GetStudyMode() *string
}

type CreateUnknownThreatDetectStrategyRequest struct {
	// example:
	//
	// UNKNOWN_THREAT_DETECT_CONFIG_****
	AssetSelectionType *string `json:"AssetSelectionType,omitempty" xml:"AssetSelectionType,omitempty"`
	// example:
	//
	// 1
	DurationDaysAfterInit *int32 `json:"DurationDaysAfterInit,omitempty" xml:"DurationDaysAfterInit,omitempty"`
	// example:
	//
	// 1
	DurationDaysAfterStop *int32 `json:"DurationDaysAfterStop,omitempty" xml:"DurationDaysAfterStop,omitempty"`
	// example:
	//
	// strategy****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// hash
	StudyMode *string `json:"StudyMode,omitempty" xml:"StudyMode,omitempty"`
}

func (s CreateUnknownThreatDetectStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUnknownThreatDetectStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateUnknownThreatDetectStrategyRequest) GetAssetSelectionType() *string {
	return s.AssetSelectionType
}

func (s *CreateUnknownThreatDetectStrategyRequest) GetDurationDaysAfterInit() *int32 {
	return s.DurationDaysAfterInit
}

func (s *CreateUnknownThreatDetectStrategyRequest) GetDurationDaysAfterStop() *int32 {
	return s.DurationDaysAfterStop
}

func (s *CreateUnknownThreatDetectStrategyRequest) GetName() *string {
	return s.Name
}

func (s *CreateUnknownThreatDetectStrategyRequest) GetStudyMode() *string {
	return s.StudyMode
}

func (s *CreateUnknownThreatDetectStrategyRequest) SetAssetSelectionType(v string) *CreateUnknownThreatDetectStrategyRequest {
	s.AssetSelectionType = &v
	return s
}

func (s *CreateUnknownThreatDetectStrategyRequest) SetDurationDaysAfterInit(v int32) *CreateUnknownThreatDetectStrategyRequest {
	s.DurationDaysAfterInit = &v
	return s
}

func (s *CreateUnknownThreatDetectStrategyRequest) SetDurationDaysAfterStop(v int32) *CreateUnknownThreatDetectStrategyRequest {
	s.DurationDaysAfterStop = &v
	return s
}

func (s *CreateUnknownThreatDetectStrategyRequest) SetName(v string) *CreateUnknownThreatDetectStrategyRequest {
	s.Name = &v
	return s
}

func (s *CreateUnknownThreatDetectStrategyRequest) SetStudyMode(v string) *CreateUnknownThreatDetectStrategyRequest {
	s.StudyMode = &v
	return s
}

func (s *CreateUnknownThreatDetectStrategyRequest) Validate() error {
	return dara.Validate(s)
}
