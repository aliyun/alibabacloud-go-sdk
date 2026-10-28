// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) *DescribePxfsSecurityIpsResponseBody
	GetAccessDeniedDetail() *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *DescribePxfsSecurityIpsResponseBodyData) *DescribePxfsSecurityIpsResponseBody
	GetData() *DescribePxfsSecurityIpsResponseBodyData
	SetRequestId(v string) *DescribePxfsSecurityIpsResponseBody
	GetRequestId() *string
}

type DescribePxfsSecurityIpsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *DescribePxfsSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePxfsSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePxfsSecurityIpsResponseBody) GetAccessDeniedDetail() *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribePxfsSecurityIpsResponseBody) GetData() *DescribePxfsSecurityIpsResponseBodyData {
	return s.Data
}

func (s *DescribePxfsSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePxfsSecurityIpsResponseBody) SetAccessDeniedDetail(v *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) *DescribePxfsSecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBody) SetData(v *DescribePxfsSecurityIpsResponseBodyData) *DescribePxfsSecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBody) SetRequestId(v string) *DescribePxfsSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBody) Validate() error {
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

type DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity used for authentication in the request.
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
	// The description is as described above.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The no-permission type.
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

func (s DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePxfsSecurityIpsResponseBodyData struct {
	// The name of the managed instance.
	//
	// example:
	//
	// pxfs-xxxxxxxx
	CustinsName *string `json:"CustinsName,omitempty" xml:"CustinsName,omitempty"`
	// The whitelist groups.
	Groups []*DescribePxfsSecurityIpsResponseBodyDataGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
}

func (s DescribePxfsSecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsSecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePxfsSecurityIpsResponseBodyData) GetCustinsName() *string {
	return s.CustinsName
}

func (s *DescribePxfsSecurityIpsResponseBodyData) GetGroups() []*DescribePxfsSecurityIpsResponseBodyDataGroups {
	return s.Groups
}

func (s *DescribePxfsSecurityIpsResponseBodyData) SetCustinsName(v string) *DescribePxfsSecurityIpsResponseBodyData {
	s.CustinsName = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyData) SetGroups(v []*DescribePxfsSecurityIpsResponseBodyDataGroups) *DescribePxfsSecurityIpsResponseBodyData {
	s.Groups = v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyData) Validate() error {
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

type DescribePxfsSecurityIpsResponseBodyDataGroups struct {
	// The name of the whitelist group.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the whitelist group.
	//
	// example:
	//
	// default
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The IP addresses.
	//
	// example:
	//
	// 192.168.1.0/24,10.0.0.1
	IpLists *string `json:"IpLists,omitempty" xml:"IpLists,omitempty"`
}

func (s DescribePxfsSecurityIpsResponseBodyDataGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsSecurityIpsResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *DescribePxfsSecurityIpsResponseBodyDataGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePxfsSecurityIpsResponseBodyDataGroups) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribePxfsSecurityIpsResponseBodyDataGroups) GetIpLists() *string {
	return s.IpLists
}

func (s *DescribePxfsSecurityIpsResponseBodyDataGroups) SetGroupName(v string) *DescribePxfsSecurityIpsResponseBodyDataGroups {
	s.GroupName = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyDataGroups) SetGroupTag(v string) *DescribePxfsSecurityIpsResponseBodyDataGroups {
	s.GroupTag = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyDataGroups) SetIpLists(v string) *DescribePxfsSecurityIpsResponseBodyDataGroups {
	s.IpLists = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponseBodyDataGroups) Validate() error {
	return dara.Validate(s)
}
