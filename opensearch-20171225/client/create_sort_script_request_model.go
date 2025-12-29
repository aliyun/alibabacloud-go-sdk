// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSortScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSortScriptRequest
	GetDescription() *string
	SetScope(v string) *CreateSortScriptRequest
	GetScope() *string
	SetScriptName(v string) *CreateSortScriptRequest
	GetScriptName() *string
	SetType(v string) *CreateSortScriptRequest
	GetType() *string
}

type CreateSortScriptRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The sort phase to which the script applies.
	//
	// example:
	//
	// second_rank
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The script name.
	//
	// example:
	//
	// rank_cava_20230606_v7
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The script type. Set the value to cava_script.
	//
	// example:
	//
	// cava_script
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSortScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSortScriptRequest) GoString() string {
	return s.String()
}

func (s *CreateSortScriptRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSortScriptRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateSortScriptRequest) GetScriptName() *string {
	return s.ScriptName
}

func (s *CreateSortScriptRequest) GetType() *string {
	return s.Type
}

func (s *CreateSortScriptRequest) SetDescription(v string) *CreateSortScriptRequest {
	s.Description = &v
	return s
}

func (s *CreateSortScriptRequest) SetScope(v string) *CreateSortScriptRequest {
	s.Scope = &v
	return s
}

func (s *CreateSortScriptRequest) SetScriptName(v string) *CreateSortScriptRequest {
	s.ScriptName = &v
	return s
}

func (s *CreateSortScriptRequest) SetType(v string) *CreateSortScriptRequest {
	s.Type = &v
	return s
}

func (s *CreateSortScriptRequest) Validate() error {
	return dara.Validate(s)
}
