// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForPrivateAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolices(v []*ListApplicationsForPrivateAccessPolicyResponseBodyPolices) *ListApplicationsForPrivateAccessPolicyResponseBody
	GetPolices() []*ListApplicationsForPrivateAccessPolicyResponseBodyPolices
	SetRequestId(v string) *ListApplicationsForPrivateAccessPolicyResponseBody
	GetRequestId() *string
}

type ListApplicationsForPrivateAccessPolicyResponseBody struct {
	Polices []*ListApplicationsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBody) GetPolices() []*ListApplicationsForPrivateAccessPolicyResponseBodyPolices {
	return s.Polices
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBody) SetPolices(v []*ListApplicationsForPrivateAccessPolicyResponseBodyPolices) *ListApplicationsForPrivateAccessPolicyResponseBody {
	s.Polices = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBody) SetRequestId(v string) *ListApplicationsForPrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBody) Validate() error {
	if s.Polices != nil {
		for _, item := range s.Polices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsForPrivateAccessPolicyResponseBodyPolices struct {
	Applications []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolices) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolices) GetApplications() []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	return s.Applications
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolices) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolices) SetApplications(v []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) *ListApplicationsForPrivateAccessPolicyResponseBodyPolices {
	s.Applications = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolices) SetPolicyId(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolices) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-7a9243dd02f4****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// application_name
	Name       *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GetAddresses() []*string {
	return s.Addresses
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GetName() *string {
	return s.Name
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GetPortRanges() []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges {
	return s.PortRanges
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GetProtocol() *string {
	return s.Protocol
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetAddresses(v []*string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Addresses = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetApplicationId(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetCreateTime(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetDescription(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetName(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Name = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetPortRanges(v []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.PortRanges = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetProtocol(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Protocol = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetStatus(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Status = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) Validate() error {
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) SetBegin(v int32) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) SetEnd(v int32) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges {
	s.End = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) Validate() error {
	return dara.Validate(s)
}
