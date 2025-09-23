// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeZonesResponseBody
	GetRequestId() *string
	SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody
	GetZones() []*DescribeZonesResponseBodyZones
}

type DescribeZonesResponseBody struct {
	// example:
	//
	// E2FDB684-751D-424D-98B9-704BEA******
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZonesResponseBody) GetZones() []*DescribeZonesResponseBodyZones {
	return s.Zones
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZones struct {
	// example:
	//
	// UnSet
	ModeCode *string `json:"ModeCode,omitempty" xml:"ModeCode,omitempty"`
	// example:
	//
	// 50
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) GetModeCode() *string {
	return s.ModeCode
}

func (s *DescribeZonesResponseBodyZones) GetPriority() *string {
	return s.Priority
}

func (s *DescribeZonesResponseBodyZones) GetStatus() *string {
	return s.Status
}

func (s *DescribeZonesResponseBodyZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZones) SetModeCode(v string) *DescribeZonesResponseBodyZones {
	s.ModeCode = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetPriority(v string) *DescribeZonesResponseBodyZones {
	s.Priority = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetStatus(v string) *DescribeZonesResponseBodyZones {
	s.Status = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}
