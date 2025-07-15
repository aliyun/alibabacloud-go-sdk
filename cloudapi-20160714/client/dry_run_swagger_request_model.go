// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunSwaggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DryRunSwaggerRequest
	GetData() *string
	SetDataFormat(v string) *DryRunSwaggerRequest
	GetDataFormat() *string
	SetGlobalCondition(v map[string]interface{}) *DryRunSwaggerRequest
	GetGlobalCondition() map[string]interface{}
	SetGroupId(v string) *DryRunSwaggerRequest
	GetGroupId() *string
	SetOverwrite(v bool) *DryRunSwaggerRequest
	GetOverwrite() *bool
	SetSecurityToken(v string) *DryRunSwaggerRequest
	GetSecurityToken() *string
}

type DryRunSwaggerRequest struct {
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
	GlobalCondition map[string]interface{} `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
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

func (s DryRunSwaggerRequest) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerRequest) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerRequest) GetData() *string {
	return s.Data
}

func (s *DryRunSwaggerRequest) GetDataFormat() *string {
	return s.DataFormat
}

func (s *DryRunSwaggerRequest) GetGlobalCondition() map[string]interface{} {
	return s.GlobalCondition
}

func (s *DryRunSwaggerRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DryRunSwaggerRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *DryRunSwaggerRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DryRunSwaggerRequest) SetData(v string) *DryRunSwaggerRequest {
	s.Data = &v
	return s
}

func (s *DryRunSwaggerRequest) SetDataFormat(v string) *DryRunSwaggerRequest {
	s.DataFormat = &v
	return s
}

func (s *DryRunSwaggerRequest) SetGlobalCondition(v map[string]interface{}) *DryRunSwaggerRequest {
	s.GlobalCondition = v
	return s
}

func (s *DryRunSwaggerRequest) SetGroupId(v string) *DryRunSwaggerRequest {
	s.GroupId = &v
	return s
}

func (s *DryRunSwaggerRequest) SetOverwrite(v bool) *DryRunSwaggerRequest {
	s.Overwrite = &v
	return s
}

func (s *DryRunSwaggerRequest) SetSecurityToken(v string) *DryRunSwaggerRequest {
	s.SecurityToken = &v
	return s
}

func (s *DryRunSwaggerRequest) Validate() error {
	return dara.Validate(s)
}
