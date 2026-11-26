// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *DeleteSecurityInfoRequest
	GetDirectoryId() *string
	SetUserId(v string) *DeleteSecurityInfoRequest
	GetUserId() *string
	SetVerifyType(v string) *DeleteSecurityInfoRequest
	GetVerifyType() *string
}

type DeleteSecurityInfoRequest struct {
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The verification type of the security information. Valid values:
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

func (s DeleteSecurityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityInfoRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteSecurityInfoRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteSecurityInfoRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *DeleteSecurityInfoRequest) SetDirectoryId(v string) *DeleteSecurityInfoRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteSecurityInfoRequest) SetUserId(v string) *DeleteSecurityInfoRequest {
	s.UserId = &v
	return s
}

func (s *DeleteSecurityInfoRequest) SetVerifyType(v string) *DeleteSecurityInfoRequest {
	s.VerifyType = &v
	return s
}

func (s *DeleteSecurityInfoRequest) Validate() error {
	return dara.Validate(s)
}
