// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListUsersRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListUsersRequest
	GetJsonStr() *string
}

type ListUsersRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListUsersRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListUsersRequest) SetBaseMeAgentId(v int64) *ListUsersRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListUsersRequest) SetJsonStr(v string) *ListUsersRequest {
	s.JsonStr = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
