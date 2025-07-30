// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAssignedSessionGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *RevertAssignedSessionGroupRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *RevertAssignedSessionGroupRequest
	GetJsonStr() *string
}

type RevertAssignedSessionGroupRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"isSchemeData":1,"forceRevertSessionGroup":true,"sessionGroupIdList":["1"]}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s RevertAssignedSessionGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertAssignedSessionGroupRequest) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionGroupRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *RevertAssignedSessionGroupRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *RevertAssignedSessionGroupRequest) SetBaseMeAgentId(v int64) *RevertAssignedSessionGroupRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *RevertAssignedSessionGroupRequest) SetJsonStr(v string) *RevertAssignedSessionGroupRequest {
	s.JsonStr = &v
	return s
}

func (s *RevertAssignedSessionGroupRequest) Validate() error {
	return dara.Validate(s)
}
