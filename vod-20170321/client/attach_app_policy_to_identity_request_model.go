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
	// The ID of the application. Default value: **app-1000000**. For more information, see [Multi-application service](https://help.aliyun.com/document_detail/113600.html).
	//
	// > This parameter is optional only if you set the policy name to VODAppAdministratorAccess.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the RAM user or the name of the RAM role.
	//
	// 	- Specify the ID of the RAM user when the IdentityType parameter is set to RamUser.
	//
	// 	- Specify the name of the RAM role when the IdentityType parameter is set to RamRole.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The type of the identity. Valid values:
	//
	// 	- **RamUser**: a RAM user
	//
	// 	- **RamRole**: a RAM role
	//
	// This parameter is required.
	//
	// example:
	//
	// RamRole
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The name of the policy. Only system policies are supported. Separate multiple policy names with commas (,). Valid values:
	//
	// 	- **VODAppFullAccess**: permissions to manage all resources in an application.
	//
	// 	- **VODAppReadOnlyAccess**: permissions to read all resources in an application.
	//
	// 	- **VODAppAdministratorAccess**: permissions of the application administrator.
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
