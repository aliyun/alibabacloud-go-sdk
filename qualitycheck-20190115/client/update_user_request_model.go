// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateUserRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateUserRequest
	GetJsonStr() *string
}

type UpdateUserRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"aliUid":"2951869706989****","roleName":"ADMIN"},{"aliUid":"2557461687048****","roleName":"ADMIN"}]
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateUserRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateUserRequest) SetBaseMeAgentId(v int64) *UpdateUserRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateUserRequest) SetJsonStr(v string) *UpdateUserRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
