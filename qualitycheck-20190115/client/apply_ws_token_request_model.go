// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyWsTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ApplyWsTokenRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ApplyWsTokenRequest
	GetJsonStr() *string
}

type ApplyWsTokenRequest struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {
	//
	//     "business": "test",
	//
	//     "callType": 1,
	//
	//     "callee": "13111111111",
	//
	//     "caller": "13800000000",
	//
	//     "skillGroupId": 1,
	//
	//     "skillGroupName": "test",
	//
	//     "taskConfigId": 399,
	//
	//     "tid": "2025012412cb129e-1579-46b5-9326-1b2ececf8f30"
	//
	// }
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ApplyWsTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyWsTokenRequest) GoString() string {
	return s.String()
}

func (s *ApplyWsTokenRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ApplyWsTokenRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ApplyWsTokenRequest) SetBaseMeAgentId(v int64) *ApplyWsTokenRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ApplyWsTokenRequest) SetJsonStr(v string) *ApplyWsTokenRequest {
	s.JsonStr = &v
	return s
}

func (s *ApplyWsTokenRequest) Validate() error {
	return dara.Validate(s)
}
