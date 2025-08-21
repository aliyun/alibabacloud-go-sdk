// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *GetLoginProfileRequest
	GetUserPrincipalName() *string
}

type GetLoginProfileRequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetLoginProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *GetLoginProfileRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetLoginProfileRequest) SetUserPrincipalName(v string) *GetLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetLoginProfileRequest) Validate() error {
	return dara.Validate(s)
}
