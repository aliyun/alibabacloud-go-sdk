// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetScriptRequest
	GetInstanceId() *string
	SetScriptId(v string) *GetScriptRequest
	GetScriptId() *string
}

type GetScriptRequest struct {
	// example:
	//
	// 8a503680-815d-473e-a9b0-e010f47a64d2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ac6db76a-bbe7-4a2c-b7cc-7f62da7bb4c6
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s GetScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScriptRequest) GoString() string {
	return s.String()
}

func (s *GetScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetScriptRequest) SetInstanceId(v string) *GetScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *GetScriptRequest) SetScriptId(v string) *GetScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *GetScriptRequest) Validate() error {
	return dara.Validate(s)
}
