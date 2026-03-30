// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *CreateScriptRequest
	GetConcurrency() *int32
	SetDescription(v string) *CreateScriptRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateScriptRequest
	GetInstanceId() *string
	SetName(v string) *CreateScriptRequest
	GetName() *string
	SetNluEngine(v string) *CreateScriptRequest
	GetNluEngine() *string
}

type CreateScriptRequest struct {
	// example:
	//
	// 10
	Concurrency *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
}

func (s CreateScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptRequest) GoString() string {
	return s.String()
}

func (s *CreateScriptRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *CreateScriptRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScriptRequest) GetName() *string {
	return s.Name
}

func (s *CreateScriptRequest) GetNluEngine() *string {
	return s.NluEngine
}

func (s *CreateScriptRequest) SetConcurrency(v int32) *CreateScriptRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateScriptRequest) SetDescription(v string) *CreateScriptRequest {
	s.Description = &v
	return s
}

func (s *CreateScriptRequest) SetInstanceId(v string) *CreateScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScriptRequest) SetName(v string) *CreateScriptRequest {
	s.Name = &v
	return s
}

func (s *CreateScriptRequest) SetNluEngine(v string) *CreateScriptRequest {
	s.NluEngine = &v
	return s
}

func (s *CreateScriptRequest) Validate() error {
	return dara.Validate(s)
}
