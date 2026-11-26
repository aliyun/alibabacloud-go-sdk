// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpEventConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetHttpEventConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetHttpEventConfigRequest
	GetJsonStr() *string
}

type GetHttpEventConfigRequest struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"id":"配置ID"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetHttpEventConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpEventConfigRequest) GoString() string {
	return s.String()
}

func (s *GetHttpEventConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetHttpEventConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetHttpEventConfigRequest) SetBaseMeAgentId(v int64) *GetHttpEventConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetHttpEventConfigRequest) SetJsonStr(v string) *GetHttpEventConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *GetHttpEventConfigRequest) Validate() error {
	return dara.Validate(s)
}
