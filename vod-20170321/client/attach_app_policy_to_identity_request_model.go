// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAppPolicyToIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AttachAppPolicyToIdentityRequest
	GetAppId() *string
	SetIdentityName(v string) *AttachAppPolicyToIdentityRequest
	GetIdentityName() *string
	SetIdentityType(v string) *AttachAppPolicyToIdentityRequest
	GetIdentityType() *string
	SetPolicyNames(v string) *AttachAppPolicyToIdentityRequest
	GetPolicyNames() *string
}

type AttachAppPolicyToIdentityRequest struct {
	// The application ID. Default value: **app-1000000**. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// > If the policy name is VODAppAdministratorAccess, this parameter is optional. For other policies, this parameter is required.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The identity name.
	//
	// - If the type is RamUser, specify the Resource Access Management (RAM) user ID.
	//
	// - If the type is RamRole, specify the role name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The identity type. Valid values:
	//
	// - **RamUser**: Resource Access Management (RAM) user.
	//
	// - **RamRole**: RAM role.
	//
	// This parameter is required.
	//
	// example:
	//
	// RamRole
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The policy names. Separate multiple names with commas (,). Only system policies are supported. Valid values:
	//
	// - **VODAppFullAccess**: permissions to manage and operate all resources in the application.
	//
	// - **VODAppReadOnlyAccess**: read-only permissions on all resources in the application.
	//
	// - **VODAppAdministratorAccess**: application administrator permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// VODAppFullAccess
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
}

func (s AttachAppPolicyToIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachAppPolicyToIdentityRequest) GoString() string {
	return s.String()
}

func (s *AttachAppPolicyToIdentityRequest) GetAppId() *string {
	return s.AppId
}

func (s *AttachAppPolicyToIdentityRequest) GetIdentityName() *string {
	return s.IdentityName
}

func (s *AttachAppPolicyToIdentityRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *AttachAppPolicyToIdentityRequest) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *AttachAppPolicyToIdentityRequest) SetAppId(v string) *AttachAppPolicyToIdentityRequest {
	s.AppId = &v
	return s
}

func (s *AttachAppPolicyToIdentityRequest) SetIdentityName(v string) *AttachAppPolicyToIdentityRequest {
	s.IdentityName = &v
	return s
}

func (s *AttachAppPolicyToIdentityRequest) SetIdentityType(v string) *AttachAppPolicyToIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *AttachAppPolicyToIdentityRequest) SetPolicyNames(v string) *AttachAppPolicyToIdentityRequest {
	s.PolicyNames = &v
	return s
}

func (s *AttachAppPolicyToIdentityRequest) Validate() error {
	return dara.Validate(s)
}
