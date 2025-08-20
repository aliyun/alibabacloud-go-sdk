// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclEnable(v bool) *GetInstanceEndpointResponseBody
	GetAclEnable() *bool
	SetAclEntries(v []*GetInstanceEndpointResponseBodyAclEntries) *GetInstanceEndpointResponseBody
	GetAclEntries() []*GetInstanceEndpointResponseBodyAclEntries
	SetCode(v string) *GetInstanceEndpointResponseBody
	GetCode() *string
	SetDomains(v []*GetInstanceEndpointResponseBodyDomains) *GetInstanceEndpointResponseBody
	GetDomains() []*GetInstanceEndpointResponseBodyDomains
	SetEnable(v bool) *GetInstanceEndpointResponseBody
	GetEnable() *bool
	SetIsSuccess(v bool) *GetInstanceEndpointResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetInstanceEndpointResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetInstanceEndpointResponseBody
	GetStatus() *string
}

type GetInstanceEndpointResponseBody struct {
	// Indicates whether the access control list (ACL) feature is enabled.
	//
	// example:
	//
	// true
	AclEnable *bool `json:"AclEnable,omitempty" xml:"AclEnable,omitempty"`
	// The ACLs.
	AclEntries []*GetInstanceEndpointResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Domain names.
	Domains []*GetInstanceEndpointResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the ACL feature is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8F3D5EC5-39D1-4C53-A198-48C54C658FA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointResponseBody) GetAclEnable() *bool {
	return s.AclEnable
}

func (s *GetInstanceEndpointResponseBody) GetAclEntries() []*GetInstanceEndpointResponseBodyAclEntries {
	return s.AclEntries
}

func (s *GetInstanceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceEndpointResponseBody) GetDomains() []*GetInstanceEndpointResponseBodyDomains {
	return s.Domains
}

func (s *GetInstanceEndpointResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *GetInstanceEndpointResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetInstanceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceEndpointResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceEndpointResponseBody) SetAclEnable(v bool) *GetInstanceEndpointResponseBody {
	s.AclEnable = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetAclEntries(v []*GetInstanceEndpointResponseBodyAclEntries) *GetInstanceEndpointResponseBody {
	s.AclEntries = v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetCode(v string) *GetInstanceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetDomains(v []*GetInstanceEndpointResponseBodyDomains) *GetInstanceEndpointResponseBody {
	s.Domains = v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetEnable(v bool) *GetInstanceEndpointResponseBody {
	s.Enable = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetIsSuccess(v bool) *GetInstanceEndpointResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetRequestId(v string) *GetInstanceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetStatus(v string) *GetInstanceEndpointResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceEndpointResponseBodyAclEntries struct {
	// Remarks for public IP address whitelists.
	//
	// example:
	//
	// 1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The public IP address whitelist.
	//
	// example:
	//
	// 192.168.1.0/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s GetInstanceEndpointResponseBodyAclEntries) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceEndpointResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointResponseBodyAclEntries) GetComment() *string {
	return s.Comment
}

func (s *GetInstanceEndpointResponseBodyAclEntries) GetEntry() *string {
	return s.Entry
}

func (s *GetInstanceEndpointResponseBodyAclEntries) SetComment(v string) *GetInstanceEndpointResponseBodyAclEntries {
	s.Comment = &v
	return s
}

func (s *GetInstanceEndpointResponseBodyAclEntries) SetEntry(v string) *GetInstanceEndpointResponseBodyAclEntries {
	s.Entry = &v
	return s
}

func (s *GetInstanceEndpointResponseBodyAclEntries) Validate() error {
	return dara.Validate(s)
}

type GetInstanceEndpointResponseBodyDomains struct {
	// The domain name that is used to access the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// shanghai-instance1-registry.cn-shanghai.cr.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- `SYSTEM`: a system domain name.
	//
	// 	- `USER`: a user domain name.
	//
	// example:
	//
	// SYSTEM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceEndpointResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceEndpointResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointResponseBodyDomains) GetDomain() *string {
	return s.Domain
}

func (s *GetInstanceEndpointResponseBodyDomains) GetType() *string {
	return s.Type
}

func (s *GetInstanceEndpointResponseBodyDomains) SetDomain(v string) *GetInstanceEndpointResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *GetInstanceEndpointResponseBodyDomains) SetType(v string) *GetInstanceEndpointResponseBodyDomains {
	s.Type = &v
	return s
}

func (s *GetInstanceEndpointResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}
