// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetProxyPatternResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetProxyPatternResponseBody
	GetRequestId() *string
	SetZoneId(v string) *SetProxyPatternResponseBody
	GetZoneId() *string
}

type SetProxyPatternResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C6F1D541-E7A6-447A-A2B5-9F7A20B2A8FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The global ID of the zone.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SetProxyPatternResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetProxyPatternResponseBody) GoString() string {
	return s.String()
}

func (s *SetProxyPatternResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetProxyPatternResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *SetProxyPatternResponseBody) SetRequestId(v string) *SetProxyPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetProxyPatternResponseBody) SetZoneId(v string) *SetProxyPatternResponseBody {
	s.ZoneId = &v
	return s
}

func (s *SetProxyPatternResponseBody) Validate() error {
	return dara.Validate(s)
}
