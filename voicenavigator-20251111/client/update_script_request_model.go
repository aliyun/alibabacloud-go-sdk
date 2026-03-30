// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *UpdateScriptRequest
	GetConcurrency() *int32
	SetDescription(v string) *UpdateScriptRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateScriptRequest
	GetInstanceId() *string
	SetName(v string) *UpdateScriptRequest
	GetName() *string
	SetScriptId(v string) *UpdateScriptRequest
	GetScriptId() *string
}

type UpdateScriptRequest struct {
	// example:
	//
	// 10
	Concurrency *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 19ac2375-53e3-477f-abe9-6cd334227981
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s UpdateScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptRequest) GoString() string {
	return s.String()
}

func (s *UpdateScriptRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *UpdateScriptRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateScriptRequest) GetName() *string {
	return s.Name
}

func (s *UpdateScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *UpdateScriptRequest) SetConcurrency(v int32) *UpdateScriptRequest {
	s.Concurrency = &v
	return s
}

func (s *UpdateScriptRequest) SetDescription(v string) *UpdateScriptRequest {
	s.Description = &v
	return s
}

func (s *UpdateScriptRequest) SetInstanceId(v string) *UpdateScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScriptRequest) SetName(v string) *UpdateScriptRequest {
	s.Name = &v
	return s
}

func (s *UpdateScriptRequest) SetScriptId(v string) *UpdateScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *UpdateScriptRequest) Validate() error {
	return dara.Validate(s)
}
