// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceEndpointResponseBody
	GetCode() *string
	SetEndpoints(v []*ListInstanceEndpointResponseBodyEndpoints) *ListInstanceEndpointResponseBody
	GetEndpoints() []*ListInstanceEndpointResponseBodyEndpoints
	SetIsSuccess(v bool) *ListInstanceEndpointResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *ListInstanceEndpointResponseBody
	GetRequestId() *string
}

type ListInstanceEndpointResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The endpoints of the instance.
	Endpoints []*ListInstanceEndpointResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B21A877-66A2-4095-90EB-20A7781A4A67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceEndpointResponseBody) GetEndpoints() []*ListInstanceEndpointResponseBodyEndpoints {
	return s.Endpoints
}

func (s *ListInstanceEndpointResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListInstanceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceEndpointResponseBody) SetCode(v string) *ListInstanceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceEndpointResponseBody) SetEndpoints(v []*ListInstanceEndpointResponseBodyEndpoints) *ListInstanceEndpointResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListInstanceEndpointResponseBody) SetIsSuccess(v bool) *ListInstanceEndpointResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListInstanceEndpointResponseBody) SetRequestId(v string) *ListInstanceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceEndpointResponseBodyEndpoints struct {
	// Indicates whether the endpoint is added to an access control list (ACL).
	//
	// example:
	//
	// true
	AclEnable *bool `json:"AclEnable,omitempty" xml:"AclEnable,omitempty"`
	// The ACLs that are configured for the instance.
	AclEntries []*ListInstanceEndpointResponseBodyEndpointsAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The list of domain names of the Container Registry instance.
	Domains []*ListInstanceEndpointResponseBodyEndpointsDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the endpoint is added to an ACL.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the endpoint.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The VPCs associated with the instance.
	LinkedVpcs []*ListInstanceEndpointResponseBodyEndpointsLinkedVpcs `json:"LinkedVpcs,omitempty" xml:"LinkedVpcs,omitempty" type:"Repeated"`
	// The status of the endpoint.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceEndpointResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceEndpointResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBodyEndpoints) GetAclEnable() *bool {
	return s.AclEnable
}

func (s *ListInstanceEndpointResponseBodyEndpoints) GetAclEntries() []*ListInstanceEndpointResponseBodyEndpointsAclEntries {
	return s.AclEntries
}

func (s *ListInstanceEndpointResponseBodyEndpoints) GetDomains() []*ListInstanceEndpointResponseBodyEndpointsDomains {
	return s.Domains
}

func (s *ListInstanceEndpointResponseBodyEndpoints) GetEnable() *bool {
	return s.Enable
}

func (s *ListInstanceEndpointResponseBodyEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListInstanceEndpointResponseBodyEndpoints) GetLinkedVpcs() []*ListInstanceEndpointResponseBodyEndpointsLinkedVpcs {
	return s.LinkedVpcs
}

func (s *ListInstanceEndpointResponseBodyEndpoints) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetAclEnable(v bool) *ListInstanceEndpointResponseBodyEndpoints {
	s.AclEnable = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetAclEntries(v []*ListInstanceEndpointResponseBodyEndpointsAclEntries) *ListInstanceEndpointResponseBodyEndpoints {
	s.AclEntries = v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetDomains(v []*ListInstanceEndpointResponseBodyEndpointsDomains) *ListInstanceEndpointResponseBodyEndpoints {
	s.Domains = v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetEnable(v bool) *ListInstanceEndpointResponseBodyEndpoints {
	s.Enable = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetEndpointType(v string) *ListInstanceEndpointResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetLinkedVpcs(v []*ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) *ListInstanceEndpointResponseBodyEndpoints {
	s.LinkedVpcs = v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetStatus(v string) *ListInstanceEndpointResponseBodyEndpoints {
	s.Status = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}

type ListInstanceEndpointResponseBodyEndpointsAclEntries struct {
	// The information about the ACL.
	//
	// example:
	//
	// null
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s ListInstanceEndpointResponseBodyEndpointsAclEntries) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceEndpointResponseBodyEndpointsAclEntries) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBodyEndpointsAclEntries) GetEntry() *string {
	return s.Entry
}

func (s *ListInstanceEndpointResponseBodyEndpointsAclEntries) SetEntry(v string) *ListInstanceEndpointResponseBodyEndpointsAclEntries {
	s.Entry = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpointsAclEntries) Validate() error {
	return dara.Validate(s)
}

type ListInstanceEndpointResponseBodyEndpointsDomains struct {
	// The domain name of the Container Registry instance.
	//
	// example:
	//
	// t****-registry.cn-shanghai.cr.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- SYSTEM: system domain name.
	//
	// 	- USER: user domain name.
	//
	// example:
	//
	// SYSTEM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstanceEndpointResponseBodyEndpointsDomains) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceEndpointResponseBodyEndpointsDomains) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBodyEndpointsDomains) GetDomain() *string {
	return s.Domain
}

func (s *ListInstanceEndpointResponseBodyEndpointsDomains) GetType() *string {
	return s.Type
}

func (s *ListInstanceEndpointResponseBodyEndpointsDomains) SetDomain(v string) *ListInstanceEndpointResponseBodyEndpointsDomains {
	s.Domain = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpointsDomains) SetType(v string) *ListInstanceEndpointResponseBodyEndpointsDomains {
	s.Type = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpointsDomains) Validate() error {
	return dara.Validate(s)
}

type ListInstanceEndpointResponseBodyEndpointsLinkedVpcs struct {
	// VPC ID
	//
	// example:
	//
	// null
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) SetVpcId(v string) *ListInstanceEndpointResponseBodyEndpointsLinkedVpcs {
	s.VpcId = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) Validate() error {
	return dara.Validate(s)
}
