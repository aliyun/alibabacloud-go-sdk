// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAppPolicyFromIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DetachAppPolicyFromIdentityRequest
	GetAppId() *string
	SetIdentityName(v string) *DetachAppPolicyFromIdentityRequest
	GetIdentityName() *string
	SetIdentityType(v string) *DetachAppPolicyFromIdentityRequest
	GetIdentityType() *string
	SetPolicyNames(v string) *DetachAppPolicyFromIdentityRequest
	GetPolicyNames() *string
}

type DetachAppPolicyFromIdentityRequest struct {
	// The application ID. If the policy name is VODAppAdministratorAccess, this parameter is optional. For other policies, this parameter is required.
	//
	// - Value (default): **app-1000000**.
	//
	// - For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
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
	// test****name
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
	// RamUser
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The policy names. Separate multiple names with commas (,). Only system policies are supported. Valid values:
	//
	// - **VODAppFullAccess**: permissions to manage and operate all resources in the application.
	//
	// - **VODAppReadOnlyAccess**: read-only permissions for all resources in the application.
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

func (s DetachAppPolicyFromIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachAppPolicyFromIdentityRequest) GoString() string {
	return s.String()
}

func (s *DetachAppPolicyFromIdentityRequest) GetAppId() *string {
	return s.AppId
}

func (s *DetachAppPolicyFromIdentityRequest) GetIdentityName() *string {
	return s.IdentityName
}

func (s *DetachAppPolicyFromIdentityRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *DetachAppPolicyFromIdentityRequest) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *DetachAppPolicyFromIdentityRequest) SetAppId(v string) *DetachAppPolicyFromIdentityRequest {
	s.AppId = &v
	return s
}

func (s *DetachAppPolicyFromIdentityRequest) SetIdentityName(v string) *DetachAppPolicyFromIdentityRequest {
	s.IdentityName = &v
	return s
}

func (s *DetachAppPolicyFromIdentityRequest) SetIdentityType(v string) *DetachAppPolicyFromIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *DetachAppPolicyFromIdentityRequest) SetPolicyNames(v string) *DetachAppPolicyFromIdentityRequest {
	s.PolicyNames = &v
	return s
}

func (s *DetachAppPolicyFromIdentityRequest) Validate() error {
	return dara.Validate(s)
}
