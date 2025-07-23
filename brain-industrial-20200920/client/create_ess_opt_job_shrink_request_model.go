// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEssOptJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreateEssOptJobShrinkRequest
	GetBusinessKey() *string
	SetDuration(v int32) *CreateEssOptJobShrinkRequest
	GetDuration() *int32
	SetElecPriceShrink(v string) *CreateEssOptJobShrinkRequest
	GetElecPriceShrink() *string
	SetFreq(v string) *CreateEssOptJobShrinkRequest
	GetFreq() *string
	SetGenPriceShrink(v string) *CreateEssOptJobShrinkRequest
	GetGenPriceShrink() *string
	SetLocationShrink(v string) *CreateEssOptJobShrinkRequest
	GetLocationShrink() *string
	SetModelVersion(v string) *CreateEssOptJobShrinkRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreateEssOptJobShrinkRequest
	GetRunDate() *string
	SetSystemDataShrink(v string) *CreateEssOptJobShrinkRequest
	GetSystemDataShrink() *string
	SetTimeZone(v string) *CreateEssOptJobShrinkRequest
	GetTimeZone() *string
	SetTopoType(v string) *CreateEssOptJobShrinkRequest
	GetTopoType() *string
}

type CreateEssOptJobShrinkRequest struct {
	BusinessKey *string `json:"BusinessKey,omitempty" xml:"BusinessKey,omitempty"`
	// example:
	//
	// 1
	Duration        *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ElecPriceShrink *string `json:"ElecPrice,omitempty" xml:"ElecPrice,omitempty"`
	// example:
	//
	// FIFTEEN_MIN
	Freq           *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	GenPriceShrink *string `json:"GenPrice,omitempty" xml:"GenPrice,omitempty"`
	LocationShrink *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-02-12
	RunDate          *string `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	SystemDataShrink *string `json:"SystemData,omitempty" xml:"SystemData,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// example:
	//
	// LOAD_ESS_SOLAR
	TopoType *string `json:"TopoType,omitempty" xml:"TopoType,omitempty"`
}

func (s CreateEssOptJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobShrinkRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreateEssOptJobShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateEssOptJobShrinkRequest) GetElecPriceShrink() *string {
	return s.ElecPriceShrink
}

func (s *CreateEssOptJobShrinkRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreateEssOptJobShrinkRequest) GetGenPriceShrink() *string {
	return s.GenPriceShrink
}

func (s *CreateEssOptJobShrinkRequest) GetLocationShrink() *string {
	return s.LocationShrink
}

func (s *CreateEssOptJobShrinkRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreateEssOptJobShrinkRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreateEssOptJobShrinkRequest) GetSystemDataShrink() *string {
	return s.SystemDataShrink
}

func (s *CreateEssOptJobShrinkRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateEssOptJobShrinkRequest) GetTopoType() *string {
	return s.TopoType
}

func (s *CreateEssOptJobShrinkRequest) SetBusinessKey(v string) *CreateEssOptJobShrinkRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetDuration(v int32) *CreateEssOptJobShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetElecPriceShrink(v string) *CreateEssOptJobShrinkRequest {
	s.ElecPriceShrink = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetFreq(v string) *CreateEssOptJobShrinkRequest {
	s.Freq = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetGenPriceShrink(v string) *CreateEssOptJobShrinkRequest {
	s.GenPriceShrink = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetLocationShrink(v string) *CreateEssOptJobShrinkRequest {
	s.LocationShrink = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetModelVersion(v string) *CreateEssOptJobShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetRunDate(v string) *CreateEssOptJobShrinkRequest {
	s.RunDate = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetSystemDataShrink(v string) *CreateEssOptJobShrinkRequest {
	s.SystemDataShrink = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetTimeZone(v string) *CreateEssOptJobShrinkRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetTopoType(v string) *CreateEssOptJobShrinkRequest {
	s.TopoType = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
