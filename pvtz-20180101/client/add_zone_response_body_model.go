// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddZoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddZoneResponseBody
	GetSuccess() *bool
	SetZoneId(v string) *AddZoneResponseBody
	GetZoneId() *string
	SetZoneName(v string) *AddZoneResponseBody
	GetZoneName() *string
}

type AddZoneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// example:
	//
	// 6fc186295683a131f63bb8b0cddc****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s AddZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddZoneResponseBody) GoString() string {
	return s.String()
}

func (s *AddZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddZoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddZoneResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddZoneResponseBody) GetZoneName() *string {
	return s.ZoneName
}

func (s *AddZoneResponseBody) SetRequestId(v string) *AddZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddZoneResponseBody) SetSuccess(v bool) *AddZoneResponseBody {
	s.Success = &v
	return s
}

func (s *AddZoneResponseBody) SetZoneId(v string) *AddZoneResponseBody {
	s.ZoneId = &v
	return s
}

func (s *AddZoneResponseBody) SetZoneName(v string) *AddZoneResponseBody {
	s.ZoneName = &v
	return s
}

func (s *AddZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
