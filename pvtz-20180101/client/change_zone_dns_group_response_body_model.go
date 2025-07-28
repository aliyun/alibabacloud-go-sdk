// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeZoneDnsGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeZoneDnsGroupResponseBody
	GetRequestId() *string
	SetZoneId(v string) *ChangeZoneDnsGroupResponseBody
	GetZoneId() *string
}

type ChangeZoneDnsGroupResponseBody struct {
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
	// e0cff188756b1d4579b25e54b66cb830
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ChangeZoneDnsGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeZoneDnsGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeZoneDnsGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeZoneDnsGroupResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *ChangeZoneDnsGroupResponseBody) SetRequestId(v string) *ChangeZoneDnsGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeZoneDnsGroupResponseBody) SetZoneId(v string) *ChangeZoneDnsGroupResponseBody {
	s.ZoneId = &v
	return s
}

func (s *ChangeZoneDnsGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
