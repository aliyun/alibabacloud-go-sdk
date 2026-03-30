// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *GetAccessKeyLastUsedRequest
	GetUserAccessKeyId() *string
	SetUserPrincipalName(v string) *GetAccessKeyLastUsedRequest
	GetUserPrincipalName() *string
}

type GetAccessKeyLastUsedRequest struct {
	// The ID of the AccessKey pair that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// If you do not specify this parameter, the AccessKey pair of the current user is queried.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetAccessKeyLastUsedRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetAccessKeyLastUsedRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetAccessKeyLastUsedRequest) SetUserAccessKeyId(v string) *GetAccessKeyLastUsedRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetAccessKeyLastUsedRequest) SetUserPrincipalName(v string) *GetAccessKeyLastUsedRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetAccessKeyLastUsedRequest) Validate() error {
	return dara.Validate(s)
}
