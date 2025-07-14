// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *ListUserGroupsByUserIdRequest
	GetUserId() *string
}

type ListUserGroupsByUserIdRequest struct {
	// The ID of the user. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserGroupsByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsByUserIdRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsByUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListUserGroupsByUserIdRequest) SetUserId(v string) *ListUserGroupsByUserIdRequest {
	s.UserId = &v
	return s
}

func (s *ListUserGroupsByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
