// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SendVerificationMessageRequest
	GetDirectoryId() *string
	SetUserId(v string) *SendVerificationMessageRequest
	GetUserId() *string
	SetVerifyType(v string) *SendVerificationMessageRequest
	GetVerifyType() *string
}

type SendVerificationMessageRequest struct {
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
	// The verification type for the security information. Valid values:
	//
	// - SecurityEmail: security email address
	//
	// This parameter is required.
	//
	// example:
	//
	// SecurityEmail
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s SendVerificationMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationMessageRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationMessageRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SendVerificationMessageRequest) GetUserId() *string {
	return s.UserId
}

func (s *SendVerificationMessageRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *SendVerificationMessageRequest) SetDirectoryId(v string) *SendVerificationMessageRequest {
	s.DirectoryId = &v
	return s
}

func (s *SendVerificationMessageRequest) SetUserId(v string) *SendVerificationMessageRequest {
	s.UserId = &v
	return s
}

func (s *SendVerificationMessageRequest) SetVerifyType(v string) *SendVerificationMessageRequest {
	s.VerifyType = &v
	return s
}

func (s *SendVerificationMessageRequest) Validate() error {
	return dara.Validate(s)
}
