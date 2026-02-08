// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalQuestionId(v string) *DeleteGlobalQuestionRequest
	GetGlobalQuestionId() *string
	SetInstanceId(v string) *DeleteGlobalQuestionRequest
	GetInstanceId() *string
	SetScriptId(v string) *DeleteGlobalQuestionRequest
	GetScriptId() *string
}

type DeleteGlobalQuestionRequest struct {
	// Global question ID
	//
	// > Obtain it by calling the ListGlobalQuestions API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35f1361e-4377-494c-9f10-4274bda0317f
	GlobalQuestionId *string `json:"GlobalQuestionId,omitempty" xml:"GlobalQuestionId,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0fe7f71c-8771-42ef-9bb1-19aa16ae7120
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DeleteGlobalQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalQuestionRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalQuestionRequest) GetGlobalQuestionId() *string {
	return s.GlobalQuestionId
}

func (s *DeleteGlobalQuestionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteGlobalQuestionRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteGlobalQuestionRequest) SetGlobalQuestionId(v string) *DeleteGlobalQuestionRequest {
	s.GlobalQuestionId = &v
	return s
}

func (s *DeleteGlobalQuestionRequest) SetInstanceId(v string) *DeleteGlobalQuestionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteGlobalQuestionRequest) SetScriptId(v string) *DeleteGlobalQuestionRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteGlobalQuestionRequest) Validate() error {
	return dara.Validate(s)
}
