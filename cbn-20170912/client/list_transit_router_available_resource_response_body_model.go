// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v []*string) *ListTransitRouterAvailableResourceResponseBody
	GetAvailableZones() []*string
	SetMasterZones(v []*string) *ListTransitRouterAvailableResourceResponseBody
	GetMasterZones() []*string
	SetRequestId(v string) *ListTransitRouterAvailableResourceResponseBody
	GetRequestId() *string
	SetSlaveZones(v []*string) *ListTransitRouterAvailableResourceResponseBody
	GetSlaveZones() []*string
	SetSupportMulticast(v bool) *ListTransitRouterAvailableResourceResponseBody
	GetSupportMulticast() *bool
}

type ListTransitRouterAvailableResourceResponseBody struct {
	// A list of zone IDs.
	AvailableZones []*string `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// A list of primary zones.
	MasterZones []*string `json:"MasterZones,omitempty" xml:"MasterZones,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B4F480E0-4E76-5E43-9966-8322C28A158A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of secondary zone IDs.
	SlaveZones []*string `json:"SlaveZones,omitempty" xml:"SlaveZones,omitempty" type:"Repeated"`
	// Indicates whether the zone supports the multicast feature.
	//
	// example:
	//
	// false
	SupportMulticast *bool `json:"SupportMulticast,omitempty" xml:"SupportMulticast,omitempty"`
}

func (s ListTransitRouterAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterAvailableResourceResponseBody) GetAvailableZones() []*string {
	return s.AvailableZones
}

func (s *ListTransitRouterAvailableResourceResponseBody) GetMasterZones() []*string {
	return s.MasterZones
}

func (s *ListTransitRouterAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterAvailableResourceResponseBody) GetSlaveZones() []*string {
	return s.SlaveZones
}

func (s *ListTransitRouterAvailableResourceResponseBody) GetSupportMulticast() *bool {
	return s.SupportMulticast
}

func (s *ListTransitRouterAvailableResourceResponseBody) SetAvailableZones(v []*string) *ListTransitRouterAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *ListTransitRouterAvailableResourceResponseBody) SetMasterZones(v []*string) *ListTransitRouterAvailableResourceResponseBody {
	s.MasterZones = v
	return s
}

func (s *ListTransitRouterAvailableResourceResponseBody) SetRequestId(v string) *ListTransitRouterAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterAvailableResourceResponseBody) SetSlaveZones(v []*string) *ListTransitRouterAvailableResourceResponseBody {
	s.SlaveZones = v
	return s
}

func (s *ListTransitRouterAvailableResourceResponseBody) SetSupportMulticast(v bool) *ListTransitRouterAvailableResourceResponseBody {
	s.SupportMulticast = &v
	return s
}

func (s *ListTransitRouterAvailableResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
