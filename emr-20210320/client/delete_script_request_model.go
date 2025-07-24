// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteScriptRequest
	GetClusterId() *string
	SetRegionId(v string) *DeleteScriptRequest
	GetRegionId() *string
	SetScriptId(v string) *DeleteScriptRequest
	GetScriptId() *string
	SetScriptType(v string) *DeleteScriptRequest
	GetScriptType() *string
}

type DeleteScriptRequest struct {
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

func (s DeleteScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptRequest) GoString() string {
	return s.String()
}

func (s *DeleteScriptRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteScriptRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteScriptRequest) GetScriptType() *string {
	return s.ScriptType
}

func (s *DeleteScriptRequest) SetClusterId(v string) *DeleteScriptRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteScriptRequest) SetRegionId(v string) *DeleteScriptRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScriptRequest) SetScriptId(v string) *DeleteScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteScriptRequest) SetScriptType(v string) *DeleteScriptRequest {
	s.ScriptType = &v
	return s
}

func (s *DeleteScriptRequest) Validate() error {
	return dara.Validate(s)
}
