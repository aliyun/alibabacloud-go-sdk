// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSchemeTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateSchemeTaskConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateSchemeTaskConfigRequest
	GetJsonStr() *string
}

type UpdateSchemeTaskConfigRequest struct {
	// Workspace ID
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// Input parameter JSON. For details, see the request parameters section.
	//
	// example:
	//
	// {"schemeTaskConfigId":368,"status":1,"name":"检测任务 2022-09-21 16:59:50"}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateSchemeTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateSchemeTaskConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *UpdateSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateSchemeTaskConfigRequest) SetJsonStr(v string) *UpdateSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateSchemeTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
