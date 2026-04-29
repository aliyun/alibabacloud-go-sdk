// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleUnknownThreatDetectEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventIdList(v []*string) *HandleUnknownThreatDetectEventRequest
	GetEventIdList() []*string
	SetStatus(v int32) *HandleUnknownThreatDetectEventRequest
	GetStatus() *int32
}

type HandleUnknownThreatDetectEventRequest struct {
	EventIdList []*string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s HandleUnknownThreatDetectEventRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleUnknownThreatDetectEventRequest) GoString() string {
	return s.String()
}

func (s *HandleUnknownThreatDetectEventRequest) GetEventIdList() []*string {
	return s.EventIdList
}

func (s *HandleUnknownThreatDetectEventRequest) GetStatus() *int32 {
	return s.Status
}

func (s *HandleUnknownThreatDetectEventRequest) SetEventIdList(v []*string) *HandleUnknownThreatDetectEventRequest {
	s.EventIdList = v
	return s
}

func (s *HandleUnknownThreatDetectEventRequest) SetStatus(v int32) *HandleUnknownThreatDetectEventRequest {
	s.Status = &v
	return s
}

func (s *HandleUnknownThreatDetectEventRequest) Validate() error {
	return dara.Validate(s)
}
