// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteZoneResponseBody
	GetRequestId() *string
	SetZoneId(v string) *DeleteZoneResponseBody
	GetZoneId() *string
}

type DeleteZoneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E246E023-F2EB-4034-83F7-B13FCF31459C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// example:
	//
	// 0e41496f12da01311d314f17b801****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteZoneResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteZoneResponseBody) SetRequestId(v string) *DeleteZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteZoneResponseBody) SetZoneId(v string) *DeleteZoneResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DeleteZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
