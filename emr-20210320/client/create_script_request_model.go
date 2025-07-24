// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateScriptRequest
	GetClusterId() *string
	SetRegionId(v string) *CreateScriptRequest
	GetRegionId() *string
	SetScriptType(v string) *CreateScriptRequest
	GetScriptType() *string
	SetScripts(v []*Script) *CreateScriptRequest
	GetScripts() []*Script
	SetTimeoutSecs(v string) *CreateScriptRequest
	GetTimeoutSecs() *string
}

type CreateScriptRequest struct {
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
	// The list of scripts.
	//
	// This parameter is required.
	Scripts []*Script `json:"Scripts,omitempty" xml:"Scripts,omitempty" type:"Repeated"`
	// The timeout period for a manual execution script. You cannot specify a timeout period for a bootstrap action.
	TimeoutSecs *string `json:"TimeoutSecs,omitempty" xml:"TimeoutSecs,omitempty"`
}

func (s CreateScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptRequest) GoString() string {
	return s.String()
}

func (s *CreateScriptRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateScriptRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateScriptRequest) GetScriptType() *string {
	return s.ScriptType
}

func (s *CreateScriptRequest) GetScripts() []*Script {
	return s.Scripts
}

func (s *CreateScriptRequest) GetTimeoutSecs() *string {
	return s.TimeoutSecs
}

func (s *CreateScriptRequest) SetClusterId(v string) *CreateScriptRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateScriptRequest) SetRegionId(v string) *CreateScriptRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScriptRequest) SetScriptType(v string) *CreateScriptRequest {
	s.ScriptType = &v
	return s
}

func (s *CreateScriptRequest) SetScripts(v []*Script) *CreateScriptRequest {
	s.Scripts = v
	return s
}

func (s *CreateScriptRequest) SetTimeoutSecs(v string) *CreateScriptRequest {
	s.TimeoutSecs = &v
	return s
}

func (s *CreateScriptRequest) Validate() error {
	return dara.Validate(s)
}
