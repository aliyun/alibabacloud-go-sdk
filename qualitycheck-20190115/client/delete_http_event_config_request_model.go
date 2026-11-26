// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpEventConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteHttpEventConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteHttpEventConfigRequest
	GetJsonStr() *string
}

type DeleteHttpEventConfigRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The input JSON string.
	//
	// example:
	//
	// {\"id\":\"7211\"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteHttpEventConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpEventConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpEventConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteHttpEventConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteHttpEventConfigRequest) SetBaseMeAgentId(v int64) *DeleteHttpEventConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteHttpEventConfigRequest) SetJsonStr(v string) *DeleteHttpEventConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteHttpEventConfigRequest) Validate() error {
	return dara.Validate(s)
}
