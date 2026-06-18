// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMem0SecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) *DescribeMem0SecurityIpsResponseBody
	GetAccessDeniedDetail() *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *DescribeMem0SecurityIpsResponseBodyData) *DescribeMem0SecurityIpsResponseBody
	GetData() *DescribeMem0SecurityIpsResponseBodyData
	SetRequestId(v string) *DescribeMem0SecurityIpsResponseBody
	GetRequestId() *string
}

type DescribeMem0SecurityIpsResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data list.
	Data *DescribeMem0SecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMem0SecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0SecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMem0SecurityIpsResponseBody) GetAccessDeniedDetail() *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeMem0SecurityIpsResponseBody) GetData() *DescribeMem0SecurityIpsResponseBodyData {
	return s.Data
}

func (s *DescribeMem0SecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMem0SecurityIpsResponseBody) SetAccessDeniedDetail(v *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) *DescribeMem0SecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBody) SetData(v *DescribeMem0SecurityIpsResponseBodyData) *DescribeMem0SecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBody) SetRequestId(v string) *DescribeMem0SecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBody) Validate() error {
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

type DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail struct {
	// Same as above.
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
	// Same as above.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// Same as above.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAacnceTlBNjlFODgyLTlBMDUtNUUyRC04MzE5LUQwMzZDRjJFQTM3NA==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of no-permission error.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeMem0SecurityIpsResponseBodyData struct {
	// The name of the memory engine instance.
	//
	// example:
	//
	// pxc-***-mem
	CustinsName *string `json:"CustinsName,omitempty" xml:"CustinsName,omitempty"`
	// The groups corresponding to the consumed service.
	Groups []*DescribeMem0SecurityIpsResponseBodyDataGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
}

func (s DescribeMem0SecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0SecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMem0SecurityIpsResponseBodyData) GetCustinsName() *string {
	return s.CustinsName
}

func (s *DescribeMem0SecurityIpsResponseBodyData) GetGroups() []*DescribeMem0SecurityIpsResponseBodyDataGroups {
	return s.Groups
}

func (s *DescribeMem0SecurityIpsResponseBodyData) SetCustinsName(v string) *DescribeMem0SecurityIpsResponseBodyData {
	s.CustinsName = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyData) SetGroups(v []*DescribeMem0SecurityIpsResponseBodyDataGroups) *DescribeMem0SecurityIpsResponseBodyData {
	s.Groups = v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyData) Validate() error {
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

type DescribeMem0SecurityIpsResponseBodyDataGroups struct {
	// The name of the whitelist group.
	//
	// example:
	//
	// bigdata
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the group.
	//
	// example:
	//
	// test
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// Indicates whether the IP addresses that are already used for DNAT entries can also be used for SNAT entries. Valid values:
	//
	// - **true**: The IP addresses can also be used for SNAT entries.
	//
	// - **false**: The IP addresses cannot be used for SNAT entries.
	//
	// example:
	//
	// 127.0.0.1
	IpLists *string `json:"IpLists,omitempty" xml:"IpLists,omitempty"`
}

func (s DescribeMem0SecurityIpsResponseBodyDataGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0SecurityIpsResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *DescribeMem0SecurityIpsResponseBodyDataGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeMem0SecurityIpsResponseBodyDataGroups) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribeMem0SecurityIpsResponseBodyDataGroups) GetIpLists() *string {
	return s.IpLists
}

func (s *DescribeMem0SecurityIpsResponseBodyDataGroups) SetGroupName(v string) *DescribeMem0SecurityIpsResponseBodyDataGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyDataGroups) SetGroupTag(v string) *DescribeMem0SecurityIpsResponseBodyDataGroups {
	s.GroupTag = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyDataGroups) SetIpLists(v string) *DescribeMem0SecurityIpsResponseBodyDataGroups {
	s.IpLists = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponseBodyDataGroups) Validate() error {
	return dara.Validate(s)
}
