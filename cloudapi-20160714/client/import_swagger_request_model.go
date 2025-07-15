// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportSwaggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ImportSwaggerRequest
	GetData() *string
	SetDataFormat(v string) *ImportSwaggerRequest
	GetDataFormat() *string
	SetDryRun(v bool) *ImportSwaggerRequest
	GetDryRun() *bool
	SetGlobalCondition(v map[string]interface{}) *ImportSwaggerRequest
	GetGlobalCondition() map[string]interface{}
	SetGroupId(v string) *ImportSwaggerRequest
	GetGroupId() *string
	SetOverwrite(v bool) *ImportSwaggerRequest
	GetOverwrite() *bool
	SetSecurityToken(v string) *ImportSwaggerRequest
	GetSecurityToken() *string
}

type ImportSwaggerRequest struct {
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
	GlobalCondition map[string]interface{} `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
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

func (s ImportSwaggerRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerRequest) GoString() string {
	return s.String()
}

func (s *ImportSwaggerRequest) GetData() *string {
	return s.Data
}

func (s *ImportSwaggerRequest) GetDataFormat() *string {
	return s.DataFormat
}

func (s *ImportSwaggerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ImportSwaggerRequest) GetGlobalCondition() map[string]interface{} {
	return s.GlobalCondition
}

func (s *ImportSwaggerRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ImportSwaggerRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *ImportSwaggerRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ImportSwaggerRequest) SetData(v string) *ImportSwaggerRequest {
	s.Data = &v
	return s
}

func (s *ImportSwaggerRequest) SetDataFormat(v string) *ImportSwaggerRequest {
	s.DataFormat = &v
	return s
}

func (s *ImportSwaggerRequest) SetDryRun(v bool) *ImportSwaggerRequest {
	s.DryRun = &v
	return s
}

func (s *ImportSwaggerRequest) SetGlobalCondition(v map[string]interface{}) *ImportSwaggerRequest {
	s.GlobalCondition = v
	return s
}

func (s *ImportSwaggerRequest) SetGroupId(v string) *ImportSwaggerRequest {
	s.GroupId = &v
	return s
}

func (s *ImportSwaggerRequest) SetOverwrite(v bool) *ImportSwaggerRequest {
	s.Overwrite = &v
	return s
}

func (s *ImportSwaggerRequest) SetSecurityToken(v string) *ImportSwaggerRequest {
	s.SecurityToken = &v
	return s
}

func (s *ImportSwaggerRequest) Validate() error {
	return dara.Validate(s)
}
