// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) *DescribeSupabaseIpWhitelistResponseBody
	GetAccessDeniedDetail() *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail
	SetData(v *DescribeSupabaseIpWhitelistResponseBodyData) *DescribeSupabaseIpWhitelistResponseBody
	GetData() *DescribeSupabaseIpWhitelistResponseBodyData
	SetRequestId(v string) *DescribeSupabaseIpWhitelistResponseBody
	GetRequestId() *string
}

type DescribeSupabaseIpWhitelistResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The whitelist information.
	Data *DescribeSupabaseIpWhitelistResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSupabaseIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseIpWhitelistResponseBody) GetAccessDeniedDetail() *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeSupabaseIpWhitelistResponseBody) GetData() *DescribeSupabaseIpWhitelistResponseBodyData {
	return s.Data
}

func (s *DescribeSupabaseIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupabaseIpWhitelistResponseBody) SetAccessDeniedDetail(v *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) *DescribeSupabaseIpWhitelistResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBody) SetData(v *DescribeSupabaseIpWhitelistResponseBodyData) *DescribeSupabaseIpWhitelistResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBody) SetRequestId(v string) *DescribeSupabaseIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBody) Validate() error {
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

type DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail struct {
	// The authentication action.
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
	// The type of the permission denial.
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

func (s DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeSupabaseIpWhitelistResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// pxc-unrf5ssig0ecg8
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The list of whitelist groups.
	GroupItems []*DescribeSupabaseIpWhitelistResponseBodyDataGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
}

func (s DescribeSupabaseIpWhitelistResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseIpWhitelistResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseIpWhitelistResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSupabaseIpWhitelistResponseBodyData) GetGroupItems() []*DescribeSupabaseIpWhitelistResponseBodyDataGroupItems {
	return s.GroupItems
}

func (s *DescribeSupabaseIpWhitelistResponseBodyData) SetDBInstanceName(v string) *DescribeSupabaseIpWhitelistResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyData) SetGroupItems(v []*DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) *DescribeSupabaseIpWhitelistResponseBodyData {
	s.GroupItems = v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyData) Validate() error {
	if s.GroupItems != nil {
		for _, item := range s.GroupItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSupabaseIpWhitelistResponseBodyDataGroupItems struct {
	// The group name.
	//
	// example:
	//
	// defaultGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The group tag.
	//
	// example:
	//
	// group1
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The IP whitelist.
	//
	// example:
	//
	// 127.0.0.1,172.168.0.0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) SetGroupName(v string) *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) SetGroupTag(v string) *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems {
	s.GroupTag = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) SetSecurityIPList(v string) *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponseBodyDataGroupItems) Validate() error {
	return dara.Validate(s)
}
