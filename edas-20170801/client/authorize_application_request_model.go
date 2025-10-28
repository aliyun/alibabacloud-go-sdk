// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *AuthorizeApplicationRequest
	GetAppIds() *string
	SetTargetUserId(v string) *AuthorizeApplicationRequest
	GetTargetUserId() *string
}

type AuthorizeApplicationRequest struct {
	// The ID of the application. You can specify multiple IDs. Separate multiple IDs with semicolons (;). If you leave this parameter empty, the permissions on the application are revoked.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5fdf50e8-*****;696-******
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The ID of the RAM user to be authorized. The value of the parameter is in the `sub-account name@primary account UID` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@133434434****
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s AuthorizeApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *AuthorizeApplicationRequest) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *AuthorizeApplicationRequest) SetAppIds(v string) *AuthorizeApplicationRequest {
	s.AppIds = &v
	return s
}

func (s *AuthorizeApplicationRequest) SetTargetUserId(v string) *AuthorizeApplicationRequest {
	s.TargetUserId = &v
	return s
}

func (s *AuthorizeApplicationRequest) Validate() error {
	return dara.Validate(s)
}
