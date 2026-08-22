// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0SecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) *DescribeContext0SecurityIpsResponseBody
	GetAccessDeniedDetail() *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *DescribeContext0SecurityIpsResponseBodyData) *DescribeContext0SecurityIpsResponseBody
	GetData() *DescribeContext0SecurityIpsResponseBodyData
	SetRequestId(v string) *DescribeContext0SecurityIpsResponseBody
	GetRequestId() *string
}

type DescribeContext0SecurityIpsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The monitoring data.
	Data *DescribeContext0SecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContext0SecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0SecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContext0SecurityIpsResponseBody) GetAccessDeniedDetail() *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeContext0SecurityIpsResponseBody) GetData() *DescribeContext0SecurityIpsResponseBodyData {
	return s.Data
}

func (s *DescribeContext0SecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContext0SecurityIpsResponseBody) SetAccessDeniedDetail(v *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) *DescribeContext0SecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBody) SetData(v *DescribeContext0SecurityIpsResponseBodyData) *DescribeContext0SecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBody) SetRequestId(v string) *DescribeContext0SecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBody) Validate() error {
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

type DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail struct {
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
	// NoPermissionType
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

func (s DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeContext0SecurityIpsResponseBodyData struct {
	// The name of the context service instance.
	//
	// example:
	//
	// context0-example
	Context0InstanceName *string `json:"Context0InstanceName,omitempty" xml:"Context0InstanceName,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The list of whitelist groups.
	GroupItems []*DescribeContext0SecurityIpsResponseBodyDataGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
}

func (s DescribeContext0SecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0SecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContext0SecurityIpsResponseBodyData) GetContext0InstanceName() *string {
	return s.Context0InstanceName
}

func (s *DescribeContext0SecurityIpsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContext0SecurityIpsResponseBodyData) GetGroupItems() []*DescribeContext0SecurityIpsResponseBodyDataGroupItems {
	return s.GroupItems
}

func (s *DescribeContext0SecurityIpsResponseBodyData) SetContext0InstanceName(v string) *DescribeContext0SecurityIpsResponseBodyData {
	s.Context0InstanceName = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyData) SetDBInstanceName(v string) *DescribeContext0SecurityIpsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyData) SetGroupItems(v []*DescribeContext0SecurityIpsResponseBodyDataGroupItems) *DescribeContext0SecurityIpsResponseBodyData {
	s.GroupItems = v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyData) Validate() error {
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

type DescribeContext0SecurityIpsResponseBodyDataGroupItems struct {
	// The name of the whitelist group.
	//
	// example:
	//
	// defaultGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the group.
	//
	// example:
	//
	// group1
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The details of the whitelist group.
	//
	// example:
	//
	// 127.0.0.1,172.168.0.0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeContext0SecurityIpsResponseBodyDataGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0SecurityIpsResponseBodyDataGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeContext0SecurityIpsResponseBodyDataGroupItems) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeContext0SecurityIpsResponseBodyDataGroupItems) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribeContext0SecurityIpsResponseBodyDataGroupItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeContext0SecurityIpsResponseBodyDataGroupItems) SetGroupName(v string) *DescribeContext0SecurityIpsResponseBodyDataGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyDataGroupItems) SetGroupTag(v string) *DescribeContext0SecurityIpsResponseBodyDataGroupItems {
	s.GroupTag = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyDataGroupItems) SetSecurityIPList(v string) *DescribeContext0SecurityIpsResponseBodyDataGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponseBodyDataGroupItems) Validate() error {
	return dara.Validate(s)
}
