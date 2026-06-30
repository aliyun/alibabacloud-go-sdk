// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateAgentRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateAgentRequest
	GetJsonStr() *string
}

type UpdateAgentRequest struct {
	// The business space ID.
	//
	// example:
	//
	// 12345
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following details.
	//
	// example:
	//
	// xxx
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateAgentRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateAgentRequest) SetBaseMeAgentId(v int64) *UpdateAgentRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateAgentRequest) SetJsonStr(v string) *UpdateAgentRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateAgentRequest) Validate() error {
	return dara.Validate(s)
}
