// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessApplicationsForDynamicRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRoutes(v []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) *ListPrivateAccessApplicationsForDynamicRouteResponseBody
	GetDynamicRoutes() []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes
	SetRequestId(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBody
	GetRequestId() *string
}

type ListPrivateAccessApplicationsForDynamicRouteResponseBody struct {
	DynamicRoutes []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes `json:"DynamicRoutes,omitempty" xml:"DynamicRoutes,omitempty" type:"Repeated"`
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBody) GetDynamicRoutes() []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes {
	return s.DynamicRoutes
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBody) SetDynamicRoutes(v []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) *ListPrivateAccessApplicationsForDynamicRouteResponseBody {
	s.DynamicRoutes = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBody) SetRequestId(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes struct {
	Applications []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) GetApplications() []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	return s.Applications
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) GetDynamicRouteId() *string {
	return s.DynamicRouteId
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) SetApplications(v []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes {
	s.Applications = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) SetDynamicRouteId(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes {
	s.DynamicRouteId = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) Validate() error {
	return dara.Validate(s)
}

type ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-7a9243dd02f4****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 2022-04-13 13:33:24
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// application_name
	Name       *string                                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GetAddresses() []*string {
	return s.Addresses
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GetDescription() *string {
	return s.Description
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GetName() *string {
	return s.Name
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GetPortRanges() []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges {
	return s.PortRanges
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GetProtocol() *string {
	return s.Protocol
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GetStatus() *string {
	return s.Status
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetAddresses(v []*string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Addresses = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetApplicationId(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetCreateTime(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetDescription(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetName(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetPortRanges(v []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.PortRanges = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetProtocol(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Protocol = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetStatus(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) Validate() error {
	return dara.Validate(s)
}

type ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) SetBegin(v int32) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) SetEnd(v int32) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges {
	s.End = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) Validate() error {
	return dara.Validate(s)
}
