// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteIntentRequest
	GetInstanceId() *string
	SetIntentId(v string) *DeleteIntentRequest
	GetIntentId() *string
	SetScriptId(v string) *DeleteIntentRequest
	GetScriptId() *string
}

type DeleteIntentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c21fb0ec-fb5e-476f-a6bf-81a892739c8d
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aa279896-64a6-4182-864c-4f2b04ec8d17
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DeleteIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteIntentRequest) GetIntentId() *string {
	return s.IntentId
}

func (s *DeleteIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteIntentRequest) SetInstanceId(v string) *DeleteIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIntentRequest) SetIntentId(v string) *DeleteIntentRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteIntentRequest) SetScriptId(v string) *DeleteIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteIntentRequest) Validate() error {
	return dara.Validate(s)
}
