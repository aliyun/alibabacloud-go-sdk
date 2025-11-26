// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateScriptRequest
	GetClusterId() *string
	SetRegionId(v string) *UpdateScriptRequest
	GetRegionId() *string
	SetScript(v *Script) *UpdateScriptRequest
	GetScript() *Script
	SetScriptId(v string) *UpdateScriptRequest
	GetScriptId() *string
	SetScriptType(v string) *UpdateScriptRequest
	GetScriptType() *string
}

type UpdateScriptRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The script.
	//
	// This parameter is required.
	Script *Script `json:"Script,omitempty" xml:"Script,omitempty"`
	// The script ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cs-da7476a7679a4d4c9cede62ebe09****
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The type of the script. Valid values:
	//
	// 	- BOOTSTRAP: indicates a bootstrap action of the Elastic Compute Service (ECS) instance.
	//
	// 	- NORMAL: indicates a common script.
	//
	// This parameter is required.
	//
	// example:
	//
	// BOOTSTRAP
	ScriptType *string `json:"ScriptType,omitempty" xml:"ScriptType,omitempty"`
}

func (s UpdateScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptRequest) GoString() string {
	return s.String()
}

func (s *UpdateScriptRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateScriptRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateScriptRequest) GetScript() *Script {
	return s.Script
}

func (s *UpdateScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *UpdateScriptRequest) GetScriptType() *string {
	return s.ScriptType
}

func (s *UpdateScriptRequest) SetClusterId(v string) *UpdateScriptRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateScriptRequest) SetRegionId(v string) *UpdateScriptRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateScriptRequest) SetScript(v *Script) *UpdateScriptRequest {
	s.Script = v
	return s
}

func (s *UpdateScriptRequest) SetScriptId(v string) *UpdateScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *UpdateScriptRequest) SetScriptType(v string) *UpdateScriptRequest {
	s.ScriptType = &v
	return s
}

func (s *UpdateScriptRequest) Validate() error {
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	return nil
}
