// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAcceleratePoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v []*ListEnterpriseAcceleratePoliciesResponseBodyPolicies) *ListEnterpriseAcceleratePoliciesResponseBody
	GetPolicies() []*ListEnterpriseAcceleratePoliciesResponseBodyPolicies
	SetRequestId(v string) *ListEnterpriseAcceleratePoliciesResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListEnterpriseAcceleratePoliciesResponseBody
	GetTotal() *int32
}

type ListEnterpriseAcceleratePoliciesResponseBody struct {
	Policies []*ListEnterpriseAcceleratePoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// DB0471D0-C05C-556D-9F40-0325D890036F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEnterpriseAcceleratePoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAcceleratePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAcceleratePoliciesResponseBody) GetPolicies() []*ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListEnterpriseAcceleratePoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnterpriseAcceleratePoliciesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListEnterpriseAcceleratePoliciesResponseBody) SetPolicies(v []*ListEnterpriseAcceleratePoliciesResponseBodyPolicies) *ListEnterpriseAcceleratePoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBody) SetRequestId(v string) *ListEnterpriseAcceleratePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBody) SetTotal(v int32) *ListEnterpriseAcceleratePoliciesResponseBody {
	s.Total = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseAcceleratePoliciesResponseBodyPolicies struct {
	// example:
	//
	// whitelist
	AccelerationType *string `json:"AccelerationType,omitempty" xml:"AccelerationType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// eap-eec34d4b12fcca61
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
	// example:
	//
	// 0
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	OnTls *int32 `json:"OnTls,omitempty" xml:"OnTls,omitempty"`
	// example:
	//
	// 99
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	ShowInClient *int32 `json:"ShowInClient,omitempty" xml:"ShowInClient,omitempty"`
	// example:
	//
	// 12.34.56.XX
	UpstreamHost *string `json:"UpstreamHost,omitempty" xml:"UpstreamHost,omitempty"`
	// example:
	//
	// 1000
	UpstreamPort *int32 `json:"UpstreamPort,omitempty" xml:"UpstreamPort,omitempty"`
	// example:
	//
	// connector
	UpstreamType       *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	UserAttributeGroup *string `json:"UserAttributeGroup,omitempty" xml:"UserAttributeGroup,omitempty"`
}

func (s ListEnterpriseAcceleratePoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetEapId() *string {
	return s.EapId
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetEnabled() *int32 {
	return s.Enabled
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetName() *string {
	return s.Name
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetOnTls() *int32 {
	return s.OnTls
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetShowInClient() *int32 {
	return s.ShowInClient
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetUpstreamHost() *string {
	return s.UpstreamHost
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetUpstreamPort() *int32 {
	return s.UpstreamPort
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) GetUserAttributeGroup() *string {
	return s.UserAttributeGroup
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetAccelerationType(v string) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.AccelerationType = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetDescription(v string) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.Description = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetEapId(v string) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.EapId = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetEnabled(v int32) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.Enabled = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetName(v string) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetOnTls(v int32) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.OnTls = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetPriority(v int32) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.Priority = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetShowInClient(v int32) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.ShowInClient = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetUpstreamHost(v string) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.UpstreamHost = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetUpstreamPort(v int32) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.UpstreamPort = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetUpstreamType(v string) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.UpstreamType = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) SetUserAttributeGroup(v string) *ListEnterpriseAcceleratePoliciesResponseBodyPolicies {
	s.UserAttributeGroup = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponseBodyPolicies) Validate() error {
	return dara.Validate(s)
}
