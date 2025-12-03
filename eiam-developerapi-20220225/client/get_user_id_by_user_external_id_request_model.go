// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByUserExternalIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserExternalId(v string) *GetUserIdByUserExternalIdRequest
	GetUserExternalId() *string
	SetUserSourceId(v string) *GetUserIdByUserExternalIdRequest
	GetUserSourceId() *string
	SetUserSourceType(v string) *GetUserIdByUserExternalIdRequest
	GetUserSourceType() *string
}

type GetUserIdByUserExternalIdRequest struct {
	// The external ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx001
	UserExternalId *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	// The source ID of the account. If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"userSourceId,omitempty" xml:"userSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in Identity as a Service (IDaaS).
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// This parameter is required.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"userSourceType,omitempty" xml:"userSourceType,omitempty"`
}

func (s GetUserIdByUserExternalIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByUserExternalIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdRequest) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *GetUserIdByUserExternalIdRequest) GetUserSourceId() *string {
	return s.UserSourceId
}

func (s *GetUserIdByUserExternalIdRequest) GetUserSourceType() *string {
	return s.UserSourceType
}

func (s *GetUserIdByUserExternalIdRequest) SetUserExternalId(v string) *GetUserIdByUserExternalIdRequest {
	s.UserExternalId = &v
	return s
}

func (s *GetUserIdByUserExternalIdRequest) SetUserSourceId(v string) *GetUserIdByUserExternalIdRequest {
	s.UserSourceId = &v
	return s
}

func (s *GetUserIdByUserExternalIdRequest) SetUserSourceType(v string) *GetUserIdByUserExternalIdRequest {
	s.UserSourceType = &v
	return s
}

func (s *GetUserIdByUserExternalIdRequest) Validate() error {
	return dara.Validate(s)
}
