// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptWaveformsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryScriptWaveformsRequest
	GetInstanceId() *string
	SetScriptContent(v string) *QueryScriptWaveformsRequest
	GetScriptContent() *string
	SetScriptId(v string) *QueryScriptWaveformsRequest
	GetScriptId() *string
}

type QueryScriptWaveformsRequest struct {
	// instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario audio Name
	//
	// This parameter is required.
	//
	// example:
	//
	// 抱歉打扰您了,稍后会有客户经理与您联系啊,再见!
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 5ab2d935-306c-478a-88bf-d08e4e25c1b7
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s QueryScriptWaveformsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptWaveformsRequest) GoString() string {
	return s.String()
}

func (s *QueryScriptWaveformsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryScriptWaveformsRequest) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *QueryScriptWaveformsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *QueryScriptWaveformsRequest) SetInstanceId(v string) *QueryScriptWaveformsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryScriptWaveformsRequest) SetScriptContent(v string) *QueryScriptWaveformsRequest {
	s.ScriptContent = &v
	return s
}

func (s *QueryScriptWaveformsRequest) SetScriptId(v string) *QueryScriptWaveformsRequest {
	s.ScriptId = &v
	return s
}

func (s *QueryScriptWaveformsRequest) Validate() error {
	return dara.Validate(s)
}
