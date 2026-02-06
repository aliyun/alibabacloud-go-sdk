// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentProfilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIp(v string) *ListAgentProfilesRequest
	GetAppIp() *string
	SetInstanceId(v string) *ListAgentProfilesRequest
	GetInstanceId() *string
	SetScriptId(v string) *ListAgentProfilesRequest
	GetScriptId() *string
}

type ListAgentProfilesRequest struct {
	// example:
	//
	// 127.0.0.1
	AppIp *string `json:"AppIp,omitempty" xml:"AppIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListAgentProfilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentProfilesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentProfilesRequest) GetAppIp() *string {
	return s.AppIp
}

func (s *ListAgentProfilesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentProfilesRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListAgentProfilesRequest) SetAppIp(v string) *ListAgentProfilesRequest {
	s.AppIp = &v
	return s
}

func (s *ListAgentProfilesRequest) SetInstanceId(v string) *ListAgentProfilesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentProfilesRequest) SetScriptId(v string) *ListAgentProfilesRequest {
	s.ScriptId = &v
	return s
}

func (s *ListAgentProfilesRequest) Validate() error {
	return dara.Validate(s)
}
