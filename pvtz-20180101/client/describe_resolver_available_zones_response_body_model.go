// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverAvailableZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v []*DescribeResolverAvailableZonesResponseBodyAvailableZones) *DescribeResolverAvailableZonesResponseBody
	GetAvailableZones() []*DescribeResolverAvailableZonesResponseBodyAvailableZones
	SetRequestId(v string) *DescribeResolverAvailableZonesResponseBody
	GetRequestId() *string
}

type DescribeResolverAvailableZonesResponseBody struct {
	// The queried zones.
	AvailableZones []*DescribeResolverAvailableZonesResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 23268E49-0C3E-4A2C-AB70-B4C7D092470B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResolverAvailableZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverAvailableZonesResponseBody) GetAvailableZones() []*DescribeResolverAvailableZonesResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeResolverAvailableZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResolverAvailableZonesResponseBody) SetAvailableZones(v []*DescribeResolverAvailableZonesResponseBodyAvailableZones) *DescribeResolverAvailableZonesResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeResolverAvailableZonesResponseBody) SetRequestId(v string) *DescribeResolverAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverAvailableZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResolverAvailableZonesResponseBodyAvailableZones struct {
	// The zone ID.
	//
	// example:
	//
	// cn-zhangjiakou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The state of resources in the zone. Valid values:
	//
	// 	- NORMAL: The resources are in the normal state.
	//
	// 	- SOLD_OUT: The resources are sold out.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResolverAvailableZonesResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverAvailableZonesResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeResolverAvailableZonesResponseBodyAvailableZones) GetAzId() *string {
	return s.AzId
}

func (s *DescribeResolverAvailableZonesResponseBodyAvailableZones) GetStatus() *string {
	return s.Status
}

func (s *DescribeResolverAvailableZonesResponseBodyAvailableZones) SetAzId(v string) *DescribeResolverAvailableZonesResponseBodyAvailableZones {
	s.AzId = &v
	return s
}

func (s *DescribeResolverAvailableZonesResponseBodyAvailableZones) SetStatus(v string) *DescribeResolverAvailableZonesResponseBodyAvailableZones {
	s.Status = &v
	return s
}

func (s *DescribeResolverAvailableZonesResponseBodyAvailableZones) Validate() error {
	return dara.Validate(s)
}
