// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateUserRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateUserRequest
	GetJsonStr() *string
}

type CreateUserRequest struct {
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateUserRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateUserRequest) SetBaseMeAgentId(v int64) *CreateUserRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateUserRequest) SetJsonStr(v string) *CreateUserRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}
