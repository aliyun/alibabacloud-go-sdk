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
	// The configuration of the data event selector. This parameter is a JSON array that can contain a maximum of 20 elements.
	//
	// Each element in the JSON array includes the following fields:
	//
	// - `ServiceName`: The name of the Alibaba Cloud service that supports data events.
	//
	// - `ReadWriteType`: The type of data event. Valid values: Read, Write, and All.
	//
	// - `EventName`: This field contains the `Equals` and `NotEquals` subfields.
	//
	//   For example, the following configuration specifies that only `GetObject`, `CopyObject`, and `AppendObject`events are delivered:
	//
	//   `{"EventName":{"Equals":["GetObject","CopyObject","AppendObject"]}}`
	//
	//   If you specify `NotEquals`, events other than `GetObject`, `CopyObject`, and `AppendObject` are delivered.
	//
	// - `ResourceArn`: This field also contains the `Equals` and `NotEquals` subfields, similar to `EventName`. For example:
	//
	//   `{"ResourceArn":{"Equals":[arn1,...,arnx]}}`
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"EventName":{"Equals":["GetObject","CopyObject","AppendObject"]},"ReadWriteType":"All","ServiceName":"Oss"}]
	EventSelectors *string `json:"EventSelectors,omitempty" xml:"EventSelectors,omitempty"`
	// Specifies whether the trail tracks data events in all regions.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	IsTrailAllRegion *bool `json:"IsTrailAllRegion,omitempty" xml:"IsTrailAllRegion,omitempty"`
	// The name of the trail.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
	// The regions where the trail tracks data events. Separate multiple region IDs with a comma (`,`).
	//
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
