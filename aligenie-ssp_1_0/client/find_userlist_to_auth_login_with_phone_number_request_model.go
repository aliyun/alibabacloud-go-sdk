// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindUserlistToAuthLoginWithPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest
	GetCode() *string
	SetPhoneNumber(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest
	GetPhoneNumber() *string
	SetRegion(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest
	GetRegion() *string
	SetSessionId(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest
	GetSessionId() *string
}

type FindUserlistToAuthLoginWithPhoneNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18612345678
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// +86
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbe2eb4458302b9246c6da17fbc95f4b
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) GetCode() *string {
	return s.Code
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) GetRegion() *string {
	return s.Region
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) SetCode(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest {
	s.Code = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) SetPhoneNumber(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) SetRegion(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest {
	s.Region = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) SetSessionId(v string) *FindUserlistToAuthLoginWithPhoneNumberRequest {
	s.SessionId = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
