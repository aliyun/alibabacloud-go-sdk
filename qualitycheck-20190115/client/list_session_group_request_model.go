// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListSessionGroupRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListSessionGroupRequest
	GetJsonStr() *string
}

type ListSessionGroupRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"isSchemeData":1,"pageNumber":1,"pageSize":10,"callStartTime":"2022-09-17 00:00:00","callEndTime":"2022-09-23 23:59:59","schemeTaskConfigId":368}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s ListSessionGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSessionGroupRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListSessionGroupRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListSessionGroupRequest) SetBaseMeAgentId(v int64) *ListSessionGroupRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListSessionGroupRequest) SetJsonStr(v string) *ListSessionGroupRequest {
	s.JsonStr = &v
	return s
}

func (s *ListSessionGroupRequest) Validate() error {
	return dara.Validate(s)
}
