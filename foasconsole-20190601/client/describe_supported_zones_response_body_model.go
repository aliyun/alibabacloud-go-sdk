// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportedZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSupportedZonesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSupportedZonesResponseBody
	GetSuccess() *bool
	SetZoneIds(v []*string) *DescribeSupportedZonesResponseBody
	GetZoneIds() []*string
}

type DescribeSupportedZonesResponseBody struct {
	// example:
	//
	// 23A9C718-DDAB-1696-B025-18FBC830F7C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	ZoneIds []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s DescribeSupportedZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportedZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportedZonesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSupportedZonesResponseBody) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *DescribeSupportedZonesResponseBody) SetRequestId(v string) *DescribeSupportedZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetSuccess(v bool) *DescribeSupportedZonesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetZoneIds(v []*string) *DescribeSupportedZonesResponseBody {
	s.ZoneIds = v
	return s
}

func (s *DescribeSupportedZonesResponseBody) Validate() error {
	return dara.Validate(s)
}
