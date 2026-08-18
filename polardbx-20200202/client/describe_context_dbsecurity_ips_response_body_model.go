// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) *DescribeContextDBSecurityIpsResponseBody
	GetAccessDeniedDetail() *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *DescribeContextDBSecurityIpsResponseBodyData) *DescribeContextDBSecurityIpsResponseBody
	GetData() *DescribeContextDBSecurityIpsResponseBodyData
	SetRequestId(v string) *DescribeContextDBSecurityIpsResponseBody
	GetRequestId() *string
}

type DescribeContextDBSecurityIpsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The paginated result of the instance list.
	Data *DescribeContextDBSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D6A4256F-7B83-5BD7-9AC0-72E1FAC05330
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContextDBSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContextDBSecurityIpsResponseBody) GetAccessDeniedDetail() *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeContextDBSecurityIpsResponseBody) GetData() *DescribeContextDBSecurityIpsResponseBodyData {
	return s.Data
}

func (s *DescribeContextDBSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContextDBSecurityIpsResponseBody) SetAccessDeniedDetail(v *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) *DescribeContextDBSecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBody) SetData(v *DescribeContextDBSecurityIpsResponseBodyData) *DescribeContextDBSecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBody) SetRequestId(v string) *DescribeContextDBSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBody) Validate() error {
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

type DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail struct {
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
	// The authentication principal type.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The diagnostic information.
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

func (s DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeContextDBSecurityIpsResponseBodyData struct {
	// The context service instance name.
	//
	// example:
	//
	// pxt-*********
	ContextDBInstanceName *string `json:"ContextDBInstanceName,omitempty" xml:"ContextDBInstanceName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The whitelist group list.
	GroupItems []*DescribeContextDBSecurityIpsResponseBodyDataGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
}

func (s DescribeContextDBSecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBSecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContextDBSecurityIpsResponseBodyData) GetContextDBInstanceName() *string {
	return s.ContextDBInstanceName
}

func (s *DescribeContextDBSecurityIpsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContextDBSecurityIpsResponseBodyData) GetGroupItems() []*DescribeContextDBSecurityIpsResponseBodyDataGroupItems {
	return s.GroupItems
}

func (s *DescribeContextDBSecurityIpsResponseBodyData) SetContextDBInstanceName(v string) *DescribeContextDBSecurityIpsResponseBodyData {
	s.ContextDBInstanceName = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyData) SetDBInstanceName(v string) *DescribeContextDBSecurityIpsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyData) SetGroupItems(v []*DescribeContextDBSecurityIpsResponseBodyDataGroupItems) *DescribeContextDBSecurityIpsResponseBodyData {
	s.GroupItems = v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyData) Validate() error {
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

type DescribeContextDBSecurityIpsResponseBodyDataGroupItems struct {
	// The whitelist group name.
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
	// The details of the whitelist group.
	//
	// example:
	//
	// 127.0.0.1,172.168.0.0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeContextDBSecurityIpsResponseBodyDataGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBSecurityIpsResponseBodyDataGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeContextDBSecurityIpsResponseBodyDataGroupItems) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeContextDBSecurityIpsResponseBodyDataGroupItems) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribeContextDBSecurityIpsResponseBodyDataGroupItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeContextDBSecurityIpsResponseBodyDataGroupItems) SetGroupName(v string) *DescribeContextDBSecurityIpsResponseBodyDataGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyDataGroupItems) SetGroupTag(v string) *DescribeContextDBSecurityIpsResponseBodyDataGroupItems {
	s.GroupTag = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyDataGroupItems) SetSecurityIPList(v string) *DescribeContextDBSecurityIpsResponseBodyDataGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponseBodyDataGroupItems) Validate() error {
	return dara.Validate(s)
}
