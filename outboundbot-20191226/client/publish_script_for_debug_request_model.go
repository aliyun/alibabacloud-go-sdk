// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptForDebugRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PublishScriptForDebugRequest
	GetInstanceId() *string
	SetScriptId(v string) *PublishScriptForDebugRequest
	GetScriptId() *string
}

type PublishScriptForDebugRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0bfe34e5-a7fa-4aac-91d4-bd798518dffc
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s PublishScriptForDebugRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptForDebugRequest) GoString() string {
	return s.String()
}

func (s *PublishScriptForDebugRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PublishScriptForDebugRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *PublishScriptForDebugRequest) SetInstanceId(v string) *PublishScriptForDebugRequest {
	s.InstanceId = &v
	return s
}

func (s *PublishScriptForDebugRequest) SetScriptId(v string) *PublishScriptForDebugRequest {
	s.ScriptId = &v
	return s
}

func (s *PublishScriptForDebugRequest) Validate() error {
	return dara.Validate(s)
}
