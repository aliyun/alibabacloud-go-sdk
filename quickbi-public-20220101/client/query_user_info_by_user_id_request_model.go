// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QueryUserInfoByUserIdRequest
	GetUserId() *string
}

type QueryUserInfoByUserIdRequest struct {
	// The ID of the user. The UserID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserInfoByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserInfoByUserIdRequest) SetUserId(v string) *QueryUserInfoByUserIdRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserInfoByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
