// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilters interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdRegex(v string) *Filters
	GetAppIdRegex() *string
	SetAppNameRegex(v string) *Filters
	GetAppNameRegex() *string
	SetAppState(v string) *Filters
	GetAppState() *string
	SetAppType(v string) *Filters
	GetAppType() *string
	SetExecutionTimeRange(v *FiltersExecutionTimeRange) *Filters
	GetExecutionTimeRange() *FiltersExecutionTimeRange
	SetSubmitTimeRange(v *FiltersSubmitTimeRange) *Filters
	GetSubmitTimeRange() *FiltersSubmitTimeRange
	SetTermiatedTimeRange(v *FiltersTermiatedTimeRange) *Filters
	GetTermiatedTimeRange() *FiltersTermiatedTimeRange
}

type Filters struct {
	AppIdRegex         *string                    `json:"AppIdRegex,omitempty" xml:"AppIdRegex,omitempty"`
	AppNameRegex       *string                    `json:"AppNameRegex,omitempty" xml:"AppNameRegex,omitempty"`
	AppState           *string                    `json:"AppState,omitempty" xml:"AppState,omitempty"`
	AppType            *string                    `json:"AppType,omitempty" xml:"AppType,omitempty"`
	ExecutionTimeRange *FiltersExecutionTimeRange `json:"ExecutionTimeRange,omitempty" xml:"ExecutionTimeRange,omitempty" type:"Struct"`
	SubmitTimeRange    *FiltersSubmitTimeRange    `json:"SubmitTimeRange,omitempty" xml:"SubmitTimeRange,omitempty" type:"Struct"`
	TermiatedTimeRange *FiltersTermiatedTimeRange `json:"TermiatedTimeRange,omitempty" xml:"TermiatedTimeRange,omitempty" type:"Struct"`
}

func (s Filters) String() string {
	return dara.Prettify(s)
}

func (s Filters) GoString() string {
	return s.String()
}

func (s *Filters) GetAppIdRegex() *string {
	return s.AppIdRegex
}

func (s *Filters) GetAppNameRegex() *string {
	return s.AppNameRegex
}

func (s *Filters) GetAppState() *string {
	return s.AppState
}

func (s *Filters) GetAppType() *string {
	return s.AppType
}

func (s *Filters) GetExecutionTimeRange() *FiltersExecutionTimeRange {
	return s.ExecutionTimeRange
}

func (s *Filters) GetSubmitTimeRange() *FiltersSubmitTimeRange {
	return s.SubmitTimeRange
}

func (s *Filters) GetTermiatedTimeRange() *FiltersTermiatedTimeRange {
	return s.TermiatedTimeRange
}

func (s *Filters) SetAppIdRegex(v string) *Filters {
	s.AppIdRegex = &v
	return s
}

func (s *Filters) SetAppNameRegex(v string) *Filters {
	s.AppNameRegex = &v
	return s
}

func (s *Filters) SetAppState(v string) *Filters {
	s.AppState = &v
	return s
}

func (s *Filters) SetAppType(v string) *Filters {
	s.AppType = &v
	return s
}

func (s *Filters) SetExecutionTimeRange(v *FiltersExecutionTimeRange) *Filters {
	s.ExecutionTimeRange = v
	return s
}

func (s *Filters) SetSubmitTimeRange(v *FiltersSubmitTimeRange) *Filters {
	s.SubmitTimeRange = v
	return s
}

func (s *Filters) SetTermiatedTimeRange(v *FiltersTermiatedTimeRange) *Filters {
	s.TermiatedTimeRange = v
	return s
}

func (s *Filters) Validate() error {
	return dara.Validate(s)
}

type FiltersExecutionTimeRange struct {
	MaxTimeInSeconds *int64 `json:"MaxTimeInSeconds,omitempty" xml:"MaxTimeInSeconds,omitempty"`
	MinTimeInSeconds *int64 `json:"MinTimeInSeconds,omitempty" xml:"MinTimeInSeconds,omitempty"`
}

func (s FiltersExecutionTimeRange) String() string {
	return dara.Prettify(s)
}

func (s FiltersExecutionTimeRange) GoString() string {
	return s.String()
}

func (s *FiltersExecutionTimeRange) GetMaxTimeInSeconds() *int64 {
	return s.MaxTimeInSeconds
}

func (s *FiltersExecutionTimeRange) GetMinTimeInSeconds() *int64 {
	return s.MinTimeInSeconds
}

func (s *FiltersExecutionTimeRange) SetMaxTimeInSeconds(v int64) *FiltersExecutionTimeRange {
	s.MaxTimeInSeconds = &v
	return s
}

func (s *FiltersExecutionTimeRange) SetMinTimeInSeconds(v int64) *FiltersExecutionTimeRange {
	s.MinTimeInSeconds = &v
	return s
}

func (s *FiltersExecutionTimeRange) Validate() error {
	return dara.Validate(s)
}

type FiltersSubmitTimeRange struct {
	MaxTimeInMills *int64 `json:"MaxTimeInMills,omitempty" xml:"MaxTimeInMills,omitempty"`
	MinTimeInMills *int64 `json:"MinTimeInMills,omitempty" xml:"MinTimeInMills,omitempty"`
}

func (s FiltersSubmitTimeRange) String() string {
	return dara.Prettify(s)
}

func (s FiltersSubmitTimeRange) GoString() string {
	return s.String()
}

func (s *FiltersSubmitTimeRange) GetMaxTimeInMills() *int64 {
	return s.MaxTimeInMills
}

func (s *FiltersSubmitTimeRange) GetMinTimeInMills() *int64 {
	return s.MinTimeInMills
}

func (s *FiltersSubmitTimeRange) SetMaxTimeInMills(v int64) *FiltersSubmitTimeRange {
	s.MaxTimeInMills = &v
	return s
}

func (s *FiltersSubmitTimeRange) SetMinTimeInMills(v int64) *FiltersSubmitTimeRange {
	s.MinTimeInMills = &v
	return s
}

func (s *FiltersSubmitTimeRange) Validate() error {
	return dara.Validate(s)
}

type FiltersTermiatedTimeRange struct {
	MaxTimeInMills *int64 `json:"MaxTimeInMills,omitempty" xml:"MaxTimeInMills,omitempty"`
	MinTimeInMills *int64 `json:"MinTimeInMills,omitempty" xml:"MinTimeInMills,omitempty"`
}

func (s FiltersTermiatedTimeRange) String() string {
	return dara.Prettify(s)
}

func (s FiltersTermiatedTimeRange) GoString() string {
	return s.String()
}

func (s *FiltersTermiatedTimeRange) GetMaxTimeInMills() *int64 {
	return s.MaxTimeInMills
}

func (s *FiltersTermiatedTimeRange) GetMinTimeInMills() *int64 {
	return s.MinTimeInMills
}

func (s *FiltersTermiatedTimeRange) SetMaxTimeInMills(v int64) *FiltersTermiatedTimeRange {
	s.MaxTimeInMills = &v
	return s
}

func (s *FiltersTermiatedTimeRange) SetMinTimeInMills(v int64) *FiltersTermiatedTimeRange {
	s.MinTimeInMills = &v
	return s
}

func (s *FiltersTermiatedTimeRange) Validate() error {
	return dara.Validate(s)
}
