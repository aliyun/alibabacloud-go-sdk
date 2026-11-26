// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SetSecurityInfoRequest
	GetDirectoryId() *string
	SetSecurityEmail(v string) *SetSecurityInfoRequest
	GetSecurityEmail() *string
	SetUserId(v string) *SetSecurityInfoRequest
	GetUserId() *string
	SetVerifyType(v string) *SetSecurityInfoRequest
	GetVerifyType() *string
}

type SetSecurityInfoRequest struct {
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The security email address of the user.
	//
	// The value can be up to 64 characters in length, must conform to the standard email format, and cannot contain whitespace characters. The security email address of each user in the same directory must be unique. The uniqueness check is case-insensitive, but the original case is preserved when the address is saved.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alice@example.com
	SecurityEmail *string `json:"SecurityEmail,omitempty" xml:"SecurityEmail,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The verification type for the security information. Valid values:
	//
	// - SecurityEmail: security email address.
	//
	// This parameter is required.
	//
	// example:
	//
	// SecurityEmail
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s SetSecurityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityInfoRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityInfoRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetSecurityInfoRequest) GetSecurityEmail() *string {
	return s.SecurityEmail
}

func (s *SetSecurityInfoRequest) GetUserId() *string {
	return s.UserId
}

func (s *SetSecurityInfoRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *SetSecurityInfoRequest) SetDirectoryId(v string) *SetSecurityInfoRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetSecurityInfoRequest) SetSecurityEmail(v string) *SetSecurityInfoRequest {
	s.SecurityEmail = &v
	return s
}

func (s *SetSecurityInfoRequest) SetUserId(v string) *SetSecurityInfoRequest {
	s.UserId = &v
	return s
}

func (s *SetSecurityInfoRequest) SetVerifyType(v string) *SetSecurityInfoRequest {
	s.VerifyType = &v
	return s
}

func (s *SetSecurityInfoRequest) Validate() error {
	return dara.Validate(s)
}
