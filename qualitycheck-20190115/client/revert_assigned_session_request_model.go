// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAssignedSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *RevertAssignedSessionRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *RevertAssignedSessionRequest
	GetJsonStr() *string
}

type RevertAssignedSessionRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"isSchemeData":1,"searchParam":{"schemeTaskConfigId":1,"sourceDataType":1,"startTime":"2022-09-20 00:00:00","endTime":"2022-09-26 23:59:59"}}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s RevertAssignedSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertAssignedSessionRequest) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *RevertAssignedSessionRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *RevertAssignedSessionRequest) SetBaseMeAgentId(v int64) *RevertAssignedSessionRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *RevertAssignedSessionRequest) SetJsonStr(v string) *RevertAssignedSessionRequest {
	s.JsonStr = &v
	return s
}

func (s *RevertAssignedSessionRequest) Validate() error {
	return dara.Validate(s)
}
