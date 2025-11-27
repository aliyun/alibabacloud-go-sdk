// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptRCInquiredSystemEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *AcceptRCInquiredSystemEventRequest
	GetEventId() *string
	SetRegionId(v string) *AcceptRCInquiredSystemEventRequest
	GetRegionId() *string
}

type AcceptRCInquiredSystemEventRequest struct {
	// The ID of the system event.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// e-2zeielxl1qzq8slb****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The region ID of the system event.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AcceptRCInquiredSystemEventRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptRCInquiredSystemEventRequest) GoString() string {
	return s.String()
}

func (s *AcceptRCInquiredSystemEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *AcceptRCInquiredSystemEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AcceptRCInquiredSystemEventRequest) SetEventId(v string) *AcceptRCInquiredSystemEventRequest {
	s.EventId = &v
	return s
}

func (s *AcceptRCInquiredSystemEventRequest) SetRegionId(v string) *AcceptRCInquiredSystemEventRequest {
	s.RegionId = &v
	return s
}

func (s *AcceptRCInquiredSystemEventRequest) Validate() error {
	return dara.Validate(s)
}
