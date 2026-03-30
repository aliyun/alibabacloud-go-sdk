// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *GetUserMFAInfoRequest
	GetUserPrincipalName() *string
}

type GetUserMFAInfoRequest struct {
	// The logon name of the RAM user. This parameter is differently set in the following scenarios:
	//
	// 	- If you use a RAM user to call this operation, this parameter can be left empty. If you do not specify this parameter, information about the MFA device that is bound to the RAM user is queried.
	//
	// 	- If you use an Alibaba Cloud account to call this operation, you must set this parameter to the logon name of the RAM user that you want to query.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserMFAInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetUserMFAInfoRequest) SetUserPrincipalName(v string) *GetUserMFAInfoRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetUserMFAInfoRequest) Validate() error {
	return dara.Validate(s)
}
