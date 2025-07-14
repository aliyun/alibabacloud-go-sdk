// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListByUserGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupIds(v string) *ListByUserGroupIdRequest
	GetUserGroupIds() *string
}

type ListByUserGroupIdRequest struct {
	// The ID of the user group that you want to query. Separate multiple user groups with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 34fe-***-6dcb,84q9-****-4a274
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
}

func (s ListByUserGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListByUserGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdRequest) GetUserGroupIds() *string {
	return s.UserGroupIds
}

func (s *ListByUserGroupIdRequest) SetUserGroupIds(v string) *ListByUserGroupIdRequest {
	s.UserGroupIds = &v
	return s
}

func (s *ListByUserGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
