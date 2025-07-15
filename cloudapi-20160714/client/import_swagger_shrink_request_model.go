// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportSwaggerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ImportSwaggerShrinkRequest
	GetData() *string
	SetDataFormat(v string) *ImportSwaggerShrinkRequest
	GetDataFormat() *string
	SetDryRun(v bool) *ImportSwaggerShrinkRequest
	GetDryRun() *bool
	SetGlobalConditionShrink(v string) *ImportSwaggerShrinkRequest
	GetGlobalConditionShrink() *string
	SetGroupId(v string) *ImportSwaggerShrinkRequest
	GetGroupId() *string
	SetOverwrite(v bool) *ImportSwaggerShrinkRequest
	GetOverwrite() *bool
	SetSecurityToken(v string) *ImportSwaggerShrinkRequest
	GetSecurityToken() *string
}

type ImportSwaggerShrinkRequest struct {
	// The Swagger text content.
	//
	// This parameter is required.
	//
	// example:
	//
	// "A Swagger API definition in YAML"
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The Swagger text format:
	//
	// 	- json
	//
	// 	- yaml
	//
	// This parameter is required.
	//
	// example:
	//
	// yaml
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	// The pre-inspection.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The global conditions.
	//
	// example:
	//
	// {}
	GlobalConditionShrink *string `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
	// The ID of the API group to which the Swagger is imported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether to overwrite the existing API.
	//
	// APIs with the same HTTP request type and backend request path are considered the same.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Overwrite     *bool   `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ImportSwaggerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportSwaggerShrinkRequest) GetData() *string {
	return s.Data
}

func (s *ImportSwaggerShrinkRequest) GetDataFormat() *string {
	return s.DataFormat
}

func (s *ImportSwaggerShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ImportSwaggerShrinkRequest) GetGlobalConditionShrink() *string {
	return s.GlobalConditionShrink
}

func (s *ImportSwaggerShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ImportSwaggerShrinkRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *ImportSwaggerShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ImportSwaggerShrinkRequest) SetData(v string) *ImportSwaggerShrinkRequest {
	s.Data = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetDataFormat(v string) *ImportSwaggerShrinkRequest {
	s.DataFormat = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetDryRun(v bool) *ImportSwaggerShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetGlobalConditionShrink(v string) *ImportSwaggerShrinkRequest {
	s.GlobalConditionShrink = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetGroupId(v string) *ImportSwaggerShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetOverwrite(v bool) *ImportSwaggerShrinkRequest {
	s.Overwrite = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetSecurityToken(v string) *ImportSwaggerShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
