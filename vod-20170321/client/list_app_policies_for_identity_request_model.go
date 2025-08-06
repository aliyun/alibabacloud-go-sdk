// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPoliciesForIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppPoliciesForIdentityRequest
	GetAppId() *string
	SetIdentityName(v string) *ListAppPoliciesForIdentityRequest
	GetIdentityName() *string
	SetIdentityType(v string) *ListAppPoliciesForIdentityRequest
	GetIdentityType() *string
}

type ListAppPoliciesForIdentityRequest struct {
	// The ID of the application. Default value: **app-1000000**. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the identity.
	//
	// 	- Specifies the ID of the RAM user when the IdentityType parameter is set to RamUser.
	//
	// 	- Specifies the name of the RAM role when the IdentityType parameter is set to RamRole.
	//
	// example:
	//
	// test****name
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The type of the identity. Valid values:
	//
	// 	- **RamUser**: a RAM user.
	//
	// 	- **RamRole**: a RAM role.
	//
	// example:
	//
	// RamUser
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
}

func (s ListAppPoliciesForIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppPoliciesForIdentityRequest) GoString() string {
	return s.String()
}

func (s *ListAppPoliciesForIdentityRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppPoliciesForIdentityRequest) GetIdentityName() *string {
	return s.IdentityName
}

func (s *ListAppPoliciesForIdentityRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *ListAppPoliciesForIdentityRequest) SetAppId(v string) *ListAppPoliciesForIdentityRequest {
	s.AppId = &v
	return s
}

func (s *ListAppPoliciesForIdentityRequest) SetIdentityName(v string) *ListAppPoliciesForIdentityRequest {
	s.IdentityName = &v
	return s
}

func (s *ListAppPoliciesForIdentityRequest) SetIdentityType(v string) *ListAppPoliciesForIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *ListAppPoliciesForIdentityRequest) Validate() error {
	return dara.Validate(s)
}
