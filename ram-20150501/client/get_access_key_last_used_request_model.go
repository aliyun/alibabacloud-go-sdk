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
	SetUserName(v string) *GetAccessKeyLastUsedRequest
	GetUserName() *string
}

type GetAccessKeyLastUsedRequest struct {
	// example:
	//
	// LTAI4GFTgcR8m8cZQDTH****
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *GetAccessKeyLastUsedRequest) GetUserName() *string {
	return s.UserName
}

func (s *GetAccessKeyLastUsedRequest) SetUserAccessKeyId(v string) *GetAccessKeyLastUsedRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetAccessKeyLastUsedRequest) SetUserName(v string) *GetAccessKeyLastUsedRequest {
	s.UserName = &v
	return s
}

func (s *GetAccessKeyLastUsedRequest) Validate() error {
	return dara.Validate(s)
}
