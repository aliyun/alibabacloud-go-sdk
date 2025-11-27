// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCollationTimeZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollationTimeZones(v *DescribeCollationTimeZonesResponseBodyCollationTimeZones) *DescribeCollationTimeZonesResponseBody
	GetCollationTimeZones() *DescribeCollationTimeZonesResponseBodyCollationTimeZones
	SetRequestId(v string) *DescribeCollationTimeZonesResponseBody
	GetRequestId() *string
}

type DescribeCollationTimeZonesResponseBody struct {
	// The list of the character set collations and time zones that are available.
	CollationTimeZones *DescribeCollationTimeZonesResponseBodyCollationTimeZones `json:"CollationTimeZones,omitempty" xml:"CollationTimeZones,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4EAED246-DB18-4C8D-9EB5-C319626F2A77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCollationTimeZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollationTimeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCollationTimeZonesResponseBody) GetCollationTimeZones() *DescribeCollationTimeZonesResponseBodyCollationTimeZones {
	return s.CollationTimeZones
}

func (s *DescribeCollationTimeZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCollationTimeZonesResponseBody) SetCollationTimeZones(v *DescribeCollationTimeZonesResponseBodyCollationTimeZones) *DescribeCollationTimeZonesResponseBody {
	s.CollationTimeZones = v
	return s
}

func (s *DescribeCollationTimeZonesResponseBody) SetRequestId(v string) *DescribeCollationTimeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCollationTimeZonesResponseBody) Validate() error {
	if s.CollationTimeZones != nil {
		if err := s.CollationTimeZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCollationTimeZonesResponseBodyCollationTimeZones struct {
	CollationTimeZone []*DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone `json:"CollationTimeZone,omitempty" xml:"CollationTimeZone,omitempty" type:"Repeated"`
}

func (s DescribeCollationTimeZonesResponseBodyCollationTimeZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollationTimeZonesResponseBodyCollationTimeZones) GoString() string {
	return s.String()
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZones) GetCollationTimeZone() []*DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone {
	return s.CollationTimeZone
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZones) SetCollationTimeZone(v []*DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) *DescribeCollationTimeZonesResponseBodyCollationTimeZones {
	s.CollationTimeZone = v
	return s
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZones) Validate() error {
	if s.CollationTimeZone != nil {
		for _, item := range s.CollationTimeZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone struct {
	// The description.
	//
	// example:
	//
	// Kabul
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The offset of the UTC time. The offset is in the following format: (UTC+*HH:mm*).
	//
	// example:
	//
	// (UTC+04:30)
	StandardTimeOffset *string `json:"StandardTimeOffset,omitempty" xml:"StandardTimeOffset,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Afghanistan Standard Time
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) GoString() string {
	return s.String()
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) GetDescription() *string {
	return s.Description
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) GetStandardTimeOffset() *string {
	return s.StandardTimeOffset
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) SetDescription(v string) *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone {
	s.Description = &v
	return s
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) SetStandardTimeOffset(v string) *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone {
	s.StandardTimeOffset = &v
	return s
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) SetTimeZone(v string) *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone {
	s.TimeZone = &v
	return s
}

func (s *DescribeCollationTimeZonesResponseBodyCollationTimeZonesCollationTimeZone) Validate() error {
	return dara.Validate(s)
}
