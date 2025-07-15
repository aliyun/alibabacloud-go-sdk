// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunSwaggerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DryRunSwaggerShrinkRequest
	GetData() *string
	SetDataFormat(v string) *DryRunSwaggerShrinkRequest
	GetDataFormat() *string
	SetGlobalConditionShrink(v string) *DryRunSwaggerShrinkRequest
	GetGlobalConditionShrink() *string
	SetGroupId(v string) *DryRunSwaggerShrinkRequest
	GetGroupId() *string
	SetOverwrite(v bool) *DryRunSwaggerShrinkRequest
	GetOverwrite() *bool
	SetSecurityToken(v string) *DryRunSwaggerShrinkRequest
	GetSecurityToken() *string
}

type DryRunSwaggerShrinkRequest struct {
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
	// The global condition.
	//
	// example:
	//
	// {}
	GlobalConditionShrink *string `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// d633cf5524f841b9950e245b191bdabf
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

func (s DryRunSwaggerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerShrinkRequest) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerShrinkRequest) GetData() *string {
	return s.Data
}

func (s *DryRunSwaggerShrinkRequest) GetDataFormat() *string {
	return s.DataFormat
}

func (s *DryRunSwaggerShrinkRequest) GetGlobalConditionShrink() *string {
	return s.GlobalConditionShrink
}

func (s *DryRunSwaggerShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DryRunSwaggerShrinkRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *DryRunSwaggerShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DryRunSwaggerShrinkRequest) SetData(v string) *DryRunSwaggerShrinkRequest {
	s.Data = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetDataFormat(v string) *DryRunSwaggerShrinkRequest {
	s.DataFormat = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetGlobalConditionShrink(v string) *DryRunSwaggerShrinkRequest {
	s.GlobalConditionShrink = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetGroupId(v string) *DryRunSwaggerShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetOverwrite(v bool) *DryRunSwaggerShrinkRequest {
	s.Overwrite = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetSecurityToken(v string) *DryRunSwaggerShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
