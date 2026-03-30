// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteScriptRequest
	GetInstanceId() *string
	SetScriptId(v string) *DeleteScriptRequest
	GetScriptId() *string
}

type DeleteScriptRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DeleteScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptRequest) GoString() string {
	return s.String()
}

func (s *DeleteScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteScriptRequest) SetInstanceId(v string) *DeleteScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScriptRequest) SetScriptId(v string) *DeleteScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteScriptRequest) Validate() error {
	return dara.Validate(s)
}
