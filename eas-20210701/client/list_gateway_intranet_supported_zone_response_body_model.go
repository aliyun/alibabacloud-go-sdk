// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetSupportedZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListGatewayIntranetSupportedZoneResponseBody
	GetRequestId() *string
	SetZones(v []*string) *ListGatewayIntranetSupportedZoneResponseBody
	GetZones() []*string
}

type ListGatewayIntranetSupportedZoneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zones that are supported by the region.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListGatewayIntranetSupportedZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetSupportedZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetSupportedZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayIntranetSupportedZoneResponseBody) GetZones() []*string {
	return s.Zones
}

func (s *ListGatewayIntranetSupportedZoneResponseBody) SetRequestId(v string) *ListGatewayIntranetSupportedZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayIntranetSupportedZoneResponseBody) SetZones(v []*string) *ListGatewayIntranetSupportedZoneResponseBody {
	s.Zones = v
	return s
}

func (s *ListGatewayIntranetSupportedZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
