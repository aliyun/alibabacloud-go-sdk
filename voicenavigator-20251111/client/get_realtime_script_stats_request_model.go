// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeScriptStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRealtimeScriptStatsRequest
	GetInstanceId() *string
	SetScriptId(v string) *GetRealtimeScriptStatsRequest
	GetScriptId() *string
}

type GetRealtimeScriptStatsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd6fc91bc13445c2af7f2e3e31418520
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s GetRealtimeScriptStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeScriptStatsRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeScriptStatsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRealtimeScriptStatsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetRealtimeScriptStatsRequest) SetInstanceId(v string) *GetRealtimeScriptStatsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRealtimeScriptStatsRequest) SetScriptId(v string) *GetRealtimeScriptStatsRequest {
	s.ScriptId = &v
	return s
}

func (s *GetRealtimeScriptStatsRequest) Validate() error {
	return dara.Validate(s)
}
