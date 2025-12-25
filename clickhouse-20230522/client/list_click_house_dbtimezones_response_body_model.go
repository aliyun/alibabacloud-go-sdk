// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClickHouseDBTimezonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListClickHouseDBTimezonesResponseBody
	GetRequestId() *string
	SetTimeZones(v []*ListClickHouseDBTimezonesResponseBodyTimeZones) *ListClickHouseDBTimezonesResponseBody
	GetTimeZones() []*ListClickHouseDBTimezonesResponseBodyTimeZones
}

type ListClickHouseDBTimezonesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeZones []*ListClickHouseDBTimezonesResponseBodyTimeZones `json:"TimeZones,omitempty" xml:"TimeZones,omitempty" type:"Repeated"`
}

func (s ListClickHouseDBTimezonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClickHouseDBTimezonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClickHouseDBTimezonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClickHouseDBTimezonesResponseBody) GetTimeZones() []*ListClickHouseDBTimezonesResponseBodyTimeZones {
	return s.TimeZones
}

func (s *ListClickHouseDBTimezonesResponseBody) SetRequestId(v string) *ListClickHouseDBTimezonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClickHouseDBTimezonesResponseBody) SetTimeZones(v []*ListClickHouseDBTimezonesResponseBodyTimeZones) *ListClickHouseDBTimezonesResponseBody {
	s.TimeZones = v
	return s
}

func (s *ListClickHouseDBTimezonesResponseBody) Validate() error {
	if s.TimeZones != nil {
		for _, item := range s.TimeZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClickHouseDBTimezonesResponseBodyTimeZones struct {
	// example:
	//
	// Asia/Shanghai
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClickHouseDBTimezonesResponseBodyTimeZones) String() string {
	return dara.Prettify(s)
}

func (s ListClickHouseDBTimezonesResponseBodyTimeZones) GoString() string {
	return s.String()
}

func (s *ListClickHouseDBTimezonesResponseBodyTimeZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListClickHouseDBTimezonesResponseBodyTimeZones) SetZoneId(v string) *ListClickHouseDBTimezonesResponseBodyTimeZones {
	s.ZoneId = &v
	return s
}

func (s *ListClickHouseDBTimezonesResponseBodyTimeZones) Validate() error {
	return dara.Validate(s)
}
