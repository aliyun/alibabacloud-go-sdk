// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestGraphResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestDetails(v *DescribeRequestGraphResponseBodyRequestDetails) *DescribeRequestGraphResponseBody
	GetRequestDetails() *DescribeRequestGraphResponseBodyRequestDetails
	SetRequestId(v string) *DescribeRequestGraphResponseBody
	GetRequestId() *string
}

type DescribeRequestGraphResponseBody struct {
	// The details of the DNS requests.
	RequestDetails *DescribeRequestGraphResponseBodyRequestDetails `json:"RequestDetails,omitempty" xml:"RequestDetails,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EB71815-A421-4E51-8E8D-667F44ABE633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRequestGraphResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestGraphResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBody) GetRequestDetails() *DescribeRequestGraphResponseBodyRequestDetails {
	return s.RequestDetails
}

func (s *DescribeRequestGraphResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRequestGraphResponseBody) SetRequestDetails(v *DescribeRequestGraphResponseBodyRequestDetails) *DescribeRequestGraphResponseBody {
	s.RequestDetails = v
	return s
}

func (s *DescribeRequestGraphResponseBody) SetRequestId(v string) *DescribeRequestGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestGraphResponseBody) Validate() error {
	if s.RequestDetails != nil {
		if err := s.RequestDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRequestGraphResponseBodyRequestDetails struct {
	ZoneRequestTop []*DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop `json:"ZoneRequestTop,omitempty" xml:"ZoneRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeRequestGraphResponseBodyRequestDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestGraphResponseBodyRequestDetails) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBodyRequestDetails) GetZoneRequestTop() []*DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	return s.ZoneRequestTop
}

func (s *DescribeRequestGraphResponseBodyRequestDetails) SetZoneRequestTop(v []*DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) *DescribeRequestGraphResponseBodyRequestDetails {
	s.ZoneRequestTop = v
	return s
}

func (s *DescribeRequestGraphResponseBodyRequestDetails) Validate() error {
	if s.ZoneRequestTop != nil {
		for _, item := range s.ZoneRequestTop {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop struct {
	// The number of DNS requests.
	//
	// example:
	//
	// 103
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// The time when the data was collected. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-21T10:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The time when the data was collected. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1571652000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) GetTime() *string {
	return s.Time
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetRequestCount(v int64) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetTime(v string) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.Time = &v
	return s
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetTimestamp(v int64) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.Timestamp = &v
	return s
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) Validate() error {
	return dara.Validate(s)
}
