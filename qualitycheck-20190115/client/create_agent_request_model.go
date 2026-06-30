// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateAgentRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateAgentRequest
	GetJsonStr() *string
}

type CreateAgentRequest struct {
	// The business workspace ID.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete **JSON string*	- information. For more information, see the following details.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateAgentRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateAgentRequest) SetBaseMeAgentId(v int64) *CreateAgentRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateAgentRequest) SetJsonStr(v string) *CreateAgentRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateAgentRequest) Validate() error {
	return dara.Validate(s)
}
