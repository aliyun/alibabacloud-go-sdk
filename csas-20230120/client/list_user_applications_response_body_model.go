// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListUserApplicationsResponseBodyApplications) *ListUserApplicationsResponseBody
	GetApplications() []*ListUserApplicationsResponseBodyApplications
	SetRequestId(v string) *ListUserApplicationsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListUserApplicationsResponseBody
	GetTotalNum() *int32
}

type ListUserApplicationsResponseBody struct {
	Applications []*ListUserApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListUserApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsResponseBody) GetApplications() []*ListUserApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListUserApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserApplicationsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListUserApplicationsResponseBody) SetApplications(v []*ListUserApplicationsResponseBodyApplications) *ListUserApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListUserApplicationsResponseBody) SetRequestId(v string) *ListUserApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserApplicationsResponseBody) SetTotalNum(v int32) *ListUserApplicationsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListUserApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserApplicationsResponseBodyApplications struct {
	// example:
	//
	// Block
	Action    *string   `json:"Action,omitempty" xml:"Action,omitempty"`
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-b927baf3e592****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// private_access_application_name
	Name       *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges []*ListUserApplicationsResponseBodyApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ListUserApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListUserApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsResponseBodyApplications) GetAction() *string {
	return s.Action
}

func (s *ListUserApplicationsResponseBodyApplications) GetAddresses() []*string {
	return s.Addresses
}

func (s *ListUserApplicationsResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListUserApplicationsResponseBodyApplications) GetName() *string {
	return s.Name
}

func (s *ListUserApplicationsResponseBodyApplications) GetPortRanges() []*ListUserApplicationsResponseBodyApplicationsPortRanges {
	return s.PortRanges
}

func (s *ListUserApplicationsResponseBodyApplications) GetProtocol() *string {
	return s.Protocol
}

func (s *ListUserApplicationsResponseBodyApplications) SetAction(v string) *ListUserApplicationsResponseBodyApplications {
	s.Action = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetAddresses(v []*string) *ListUserApplicationsResponseBodyApplications {
	s.Addresses = v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetApplicationId(v string) *ListUserApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetName(v string) *ListUserApplicationsResponseBodyApplications {
	s.Name = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetPortRanges(v []*ListUserApplicationsResponseBodyApplicationsPortRanges) *ListUserApplicationsResponseBodyApplications {
	s.PortRanges = v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetProtocol(v string) *ListUserApplicationsResponseBodyApplications {
	s.Protocol = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}

type ListUserApplicationsResponseBodyApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *string `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *string `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListUserApplicationsResponseBodyApplicationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s ListUserApplicationsResponseBodyApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsResponseBodyApplicationsPortRanges) GetBegin() *string {
	return s.Begin
}

func (s *ListUserApplicationsResponseBodyApplicationsPortRanges) GetEnd() *string {
	return s.End
}

func (s *ListUserApplicationsResponseBodyApplicationsPortRanges) SetBegin(v string) *ListUserApplicationsResponseBodyApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplicationsPortRanges) SetEnd(v string) *ListUserApplicationsResponseBodyApplicationsPortRanges {
	s.End = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplicationsPortRanges) Validate() error {
	return dara.Validate(s)
}
