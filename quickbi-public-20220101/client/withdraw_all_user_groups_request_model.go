// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawAllUserGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *WithdrawAllUserGroupsRequest
	GetUserId() *string
}

type WithdrawAllUserGroupsRequest struct {
	// The ID of the user. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s WithdrawAllUserGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s WithdrawAllUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *WithdrawAllUserGroupsRequest) GetUserId() *string {
	return s.UserId
}

func (s *WithdrawAllUserGroupsRequest) SetUserId(v string) *WithdrawAllUserGroupsRequest {
	s.UserId = &v
	return s
}

func (s *WithdrawAllUserGroupsRequest) Validate() error {
	return dara.Validate(s)
}
