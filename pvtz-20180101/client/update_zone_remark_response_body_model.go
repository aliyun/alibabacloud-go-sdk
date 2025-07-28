// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateZoneRemarkResponseBody
	GetRequestId() *string
	SetZoneId(v string) *UpdateZoneRemarkResponseBody
	GetZoneId() *string
}

type UpdateZoneRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C6F1D541-E7A6-447A-A2B5-9F7A20B2A8FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateZoneRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateZoneRemarkResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateZoneRemarkResponseBody) SetRequestId(v string) *UpdateZoneRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateZoneRemarkResponseBody) SetZoneId(v string) *UpdateZoneRemarkResponseBody {
	s.ZoneId = &v
	return s
}

func (s *UpdateZoneRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
