// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The description.
	//
	// example:
	//
	// Ask customers for their opinions and suggestions about the service
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name.
	//
	// example:
	//
	// Satisfaction survey
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s UpdateScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptRequest) GoString() string {
	return s.String()
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
