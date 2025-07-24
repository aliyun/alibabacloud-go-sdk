// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateScriptShrinkRequest
	GetClusterId() *string
	SetRegionId(v string) *UpdateScriptShrinkRequest
	GetRegionId() *string
	SetScriptShrink(v string) *UpdateScriptShrinkRequest
	GetScriptShrink() *string
	SetScriptId(v string) *UpdateScriptShrinkRequest
	GetScriptId() *string
	SetScriptType(v string) *UpdateScriptShrinkRequest
	GetScriptType() *string
}

type UpdateScriptShrinkRequest struct {
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
	ScriptShrink *string `json:"Script,omitempty" xml:"Script,omitempty"`
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

func (s UpdateScriptShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScriptShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateScriptShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateScriptShrinkRequest) GetScriptShrink() *string {
	return s.ScriptShrink
}

func (s *UpdateScriptShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *UpdateScriptShrinkRequest) GetScriptType() *string {
	return s.ScriptType
}

func (s *UpdateScriptShrinkRequest) SetClusterId(v string) *UpdateScriptShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateScriptShrinkRequest) SetRegionId(v string) *UpdateScriptShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateScriptShrinkRequest) SetScriptShrink(v string) *UpdateScriptShrinkRequest {
	s.ScriptShrink = &v
	return s
}

func (s *UpdateScriptShrinkRequest) SetScriptId(v string) *UpdateScriptShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *UpdateScriptShrinkRequest) SetScriptType(v string) *UpdateScriptShrinkRequest {
	s.ScriptType = &v
	return s
}

func (s *UpdateScriptShrinkRequest) Validate() error {
	return dara.Validate(s)
}
