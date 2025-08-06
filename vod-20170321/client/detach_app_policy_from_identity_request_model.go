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
	// The ID of the application. This parameter is optional if you set PolicyNames to VODAppAdministratorAccess. This parameter is required if you set PolicyNames to a value other than VODAppAdministratorAccess.
	//
	// 	- Default value: **app-1000000**.
	//
	// 	- For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the RAM user or the name of the RAM role.
	//
	// 	- Specifies the ID of the RAM user for this parameter if you set IdentityType to RamUser.
	//
	// 	- Specifies the name of the RAM role for this parameter if you set IdentityType to RamRole.
	//
	// This parameter is required.
	//
	// example:
	//
	// test****name
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The type of the identity. Valid values:
	//
	// 	- **RamUser**: RAM user
	//
	// 	- **RamRole**: RAM role
	//
	// This parameter is required.
	//
	// example:
	//
	// RamUser
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The name of the policy. Separate multiple names with commas (,). Only system policies are supported.
	//
	// 	- **VODAppFullAccess**: permissions to manage all resources in an application
	//
	// 	- **VODAppReadOnlyAccess**: permissions to read all resources in an application
	//
	// 	- **VODAppAdministratorAccess**: permissions of the application administrator
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
