// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataEventSelectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventSelectors(v string) *PutDataEventSelectorRequest
	GetEventSelectors() *string
	SetIsTrailAllRegion(v bool) *PutDataEventSelectorRequest
	GetIsTrailAllRegion() *bool
	SetTrailName(v string) *PutDataEventSelectorRequest
	GetTrailName() *string
	SetTrailRegionIds(v string) *PutDataEventSelectorRequest
	GetTrailRegionIds() *string
}

type PutDataEventSelectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"EventName":{"Equals":["GetObject","CopyObject","AppendObject"]},"ReadWriteType":"All","ServiceName":"Oss"}]
	EventSelectors *string `json:"EventSelectors,omitempty" xml:"EventSelectors,omitempty"`
	// example:
	//
	// true
	IsTrailAllRegion *bool `json:"IsTrailAllRegion,omitempty" xml:"IsTrailAllRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
	// example:
	//
	// cn-shanghai,cn-hangzhou
	TrailRegionIds *string `json:"TrailRegionIds,omitempty" xml:"TrailRegionIds,omitempty"`
}

func (s PutDataEventSelectorRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDataEventSelectorRequest) GoString() string {
	return s.String()
}

func (s *PutDataEventSelectorRequest) GetEventSelectors() *string {
	return s.EventSelectors
}

func (s *PutDataEventSelectorRequest) GetIsTrailAllRegion() *bool {
	return s.IsTrailAllRegion
}

func (s *PutDataEventSelectorRequest) GetTrailName() *string {
	return s.TrailName
}

func (s *PutDataEventSelectorRequest) GetTrailRegionIds() *string {
	return s.TrailRegionIds
}

func (s *PutDataEventSelectorRequest) SetEventSelectors(v string) *PutDataEventSelectorRequest {
	s.EventSelectors = &v
	return s
}

func (s *PutDataEventSelectorRequest) SetIsTrailAllRegion(v bool) *PutDataEventSelectorRequest {
	s.IsTrailAllRegion = &v
	return s
}

func (s *PutDataEventSelectorRequest) SetTrailName(v string) *PutDataEventSelectorRequest {
	s.TrailName = &v
	return s
}

func (s *PutDataEventSelectorRequest) SetTrailRegionIds(v string) *PutDataEventSelectorRequest {
	s.TrailRegionIds = &v
	return s
}

func (s *PutDataEventSelectorRequest) Validate() error {
	return dara.Validate(s)
}
