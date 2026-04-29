// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnknownThreatDetectStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDurationDaysAfterInit(v int32) *UpdateUnknownThreatDetectStrategyRequest
	GetDurationDaysAfterInit() *int32
	SetDurationDaysAfterStop(v int32) *UpdateUnknownThreatDetectStrategyRequest
	GetDurationDaysAfterStop() *int32
	SetId(v string) *UpdateUnknownThreatDetectStrategyRequest
	GetId() *string
	SetName(v string) *UpdateUnknownThreatDetectStrategyRequest
	GetName() *string
	SetStudyMode(v string) *UpdateUnknownThreatDetectStrategyRequest
	GetStudyMode() *string
}

type UpdateUnknownThreatDetectStrategyRequest struct {
	// example:
	//
	// 1
	DurationDaysAfterInit *int32 `json:"DurationDaysAfterInit,omitempty" xml:"DurationDaysAfterInit,omitempty"`
	// example:
	//
	// 1
	DurationDaysAfterStop *int32 `json:"DurationDaysAfterStop,omitempty" xml:"DurationDaysAfterStop,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// strategy****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// hash
	StudyMode *string `json:"StudyMode,omitempty" xml:"StudyMode,omitempty"`
}

func (s UpdateUnknownThreatDetectStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnknownThreatDetectStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateUnknownThreatDetectStrategyRequest) GetDurationDaysAfterInit() *int32 {
	return s.DurationDaysAfterInit
}

func (s *UpdateUnknownThreatDetectStrategyRequest) GetDurationDaysAfterStop() *int32 {
	return s.DurationDaysAfterStop
}

func (s *UpdateUnknownThreatDetectStrategyRequest) GetId() *string {
	return s.Id
}

func (s *UpdateUnknownThreatDetectStrategyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateUnknownThreatDetectStrategyRequest) GetStudyMode() *string {
	return s.StudyMode
}

func (s *UpdateUnknownThreatDetectStrategyRequest) SetDurationDaysAfterInit(v int32) *UpdateUnknownThreatDetectStrategyRequest {
	s.DurationDaysAfterInit = &v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyRequest) SetDurationDaysAfterStop(v int32) *UpdateUnknownThreatDetectStrategyRequest {
	s.DurationDaysAfterStop = &v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyRequest) SetId(v string) *UpdateUnknownThreatDetectStrategyRequest {
	s.Id = &v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyRequest) SetName(v string) *UpdateUnknownThreatDetectStrategyRequest {
	s.Name = &v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyRequest) SetStudyMode(v string) *UpdateUnknownThreatDetectStrategyRequest {
	s.StudyMode = &v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyRequest) Validate() error {
	return dara.Validate(s)
}
