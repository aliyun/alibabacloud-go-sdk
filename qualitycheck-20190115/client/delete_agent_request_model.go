// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteAgentRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteAgentRequest
	GetJsonStr() *string
}

type DeleteAgentRequest struct {
	// baseMeAgentId
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// A JSON-formatted string. For more information about the complete JSON string, see the following details.
	//
	// example:
	//
	// {\\"id\\":486}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteAgentRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteAgentRequest) SetBaseMeAgentId(v int64) *DeleteAgentRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteAgentRequest) SetJsonStr(v string) *DeleteAgentRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteAgentRequest) Validate() error {
	return dara.Validate(s)
}
