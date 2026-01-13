// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesPrivateRAGServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeZonesPrivateRAGServiceResponseBody
	GetRequestId() *string
	SetZoneIds(v []*string) *DescribeZonesPrivateRAGServiceResponseBody
	GetZoneIds() []*string
}

type DescribeZonesPrivateRAGServiceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of zones.
	ZoneIds []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s DescribeZonesPrivateRAGServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesPrivateRAGServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesPrivateRAGServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZonesPrivateRAGServiceResponseBody) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *DescribeZonesPrivateRAGServiceResponseBody) SetRequestId(v string) *DescribeZonesPrivateRAGServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesPrivateRAGServiceResponseBody) SetZoneIds(v []*string) *DescribeZonesPrivateRAGServiceResponseBody {
	s.ZoneIds = v
	return s
}

func (s *DescribeZonesPrivateRAGServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
