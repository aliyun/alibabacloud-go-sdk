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
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// b3d5ac22-9643-49c6-aa84-777f6656f9f5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ID of the script to delete
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c589560-6b9d-4415-a3e0-049c4ff05f56
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
