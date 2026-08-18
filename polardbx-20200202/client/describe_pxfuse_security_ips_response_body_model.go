// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfuseSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) *DescribePxfuseSecurityIpsResponseBody
	GetAccessDeniedDetail() *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *DescribePxfuseSecurityIpsResponseBodyData) *DescribePxfuseSecurityIpsResponseBody
	GetData() *DescribePxfuseSecurityIpsResponseBodyData
	SetRequestId(v string) *DescribePxfuseSecurityIpsResponseBody
	GetRequestId() *string
}

type DescribePxfuseSecurityIpsResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The instance details.
	Data *DescribePxfuseSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePxfuseSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePxfuseSecurityIpsResponseBody) GetAccessDeniedDetail() *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribePxfuseSecurityIpsResponseBody) GetData() *DescribePxfuseSecurityIpsResponseBodyData {
	return s.Data
}

func (s *DescribePxfuseSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePxfuseSecurityIpsResponseBody) SetAccessDeniedDetail(v *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) *DescribePxfuseSecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBody) SetData(v *DescribePxfuseSecurityIpsResponseBodyData) *DescribePxfuseSecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBody) SetRequestId(v string) *DescribePxfuseSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail struct {
	// The description is the same as above.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authentication principal.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the no-permission error.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePxfuseSecurityIpsResponseBodyData struct {
	// The name of the memory engine instance.
	//
	// example:
	//
	// pxc-***-mem
	CustinsName *string `json:"CustinsName,omitempty" xml:"CustinsName,omitempty"`
	// The groups corresponding to the consumed service.
	Groups []*DescribePxfuseSecurityIpsResponseBodyDataGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
}

func (s DescribePxfuseSecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseSecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePxfuseSecurityIpsResponseBodyData) GetCustinsName() *string {
	return s.CustinsName
}

func (s *DescribePxfuseSecurityIpsResponseBodyData) GetGroups() []*DescribePxfuseSecurityIpsResponseBodyDataGroups {
	return s.Groups
}

func (s *DescribePxfuseSecurityIpsResponseBodyData) SetCustinsName(v string) *DescribePxfuseSecurityIpsResponseBodyData {
	s.CustinsName = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyData) SetGroups(v []*DescribePxfuseSecurityIpsResponseBodyDataGroups) *DescribePxfuseSecurityIpsResponseBodyData {
	s.Groups = v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyData) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePxfuseSecurityIpsResponseBodyDataGroups struct {
	// The name of the whitelist group.
	//
	// example:
	//
	// bigdata
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The group tag.
	//
	// example:
	//
	// test
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// Specifies whether an IP address that is already used for a DNAT entry can also be used for an SNAT entry. Valid values:
	//
	// - **true**: Destination IP address can also be used for an SNAT entry.
	//
	// - **false**: Destination IP address cannot be used for an SNAT entry.
	//
	// example:
	//
	// 127.0.0.1
	IpLists *string `json:"IpLists,omitempty" xml:"IpLists,omitempty"`
}

func (s DescribePxfuseSecurityIpsResponseBodyDataGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseSecurityIpsResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *DescribePxfuseSecurityIpsResponseBodyDataGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePxfuseSecurityIpsResponseBodyDataGroups) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribePxfuseSecurityIpsResponseBodyDataGroups) GetIpLists() *string {
	return s.IpLists
}

func (s *DescribePxfuseSecurityIpsResponseBodyDataGroups) SetGroupName(v string) *DescribePxfuseSecurityIpsResponseBodyDataGroups {
	s.GroupName = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyDataGroups) SetGroupTag(v string) *DescribePxfuseSecurityIpsResponseBodyDataGroups {
	s.GroupTag = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyDataGroups) SetIpLists(v string) *DescribePxfuseSecurityIpsResponseBodyDataGroups {
	s.IpLists = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponseBodyDataGroups) Validate() error {
	return dara.Validate(s)
}
