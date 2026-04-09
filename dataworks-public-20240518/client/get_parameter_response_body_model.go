// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameter(v *GetParameterResponseBodyParameter) *GetParameterResponseBody
	GetParameter() *GetParameterResponseBodyParameter
	SetRequestId(v string) *GetParameterResponseBody
	GetRequestId() *string
}

type GetParameterResponseBody struct {
	Parameter *GetParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParameterResponseBody) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBody) GetParameter() *GetParameterResponseBodyParameter {
	return s.Parameter
}

func (s *GetParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParameterResponseBody) SetParameter(v *GetParameterResponseBodyParameter) *GetParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *GetParameterResponseBody) SetRequestId(v string) *GetParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParameterResponseBody) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetParameterResponseBodyParameter struct {
	// example:
	//
	// 1640000000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123456789
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 这是一个测试参数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1640000000000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 123456789
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// workspace.para
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1000
	ProjectId  *int64                                         `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Properties []*GetParameterResponseBodyParameterProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// example:
	//
	// Project
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// PlainConstant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetParameterResponseBodyParameter) String() string {
	return dara.Prettify(s)
}

func (s GetParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBodyParameter) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetParameterResponseBodyParameter) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetParameterResponseBodyParameter) GetDescription() *string {
	return s.Description
}

func (s *GetParameterResponseBodyParameter) GetId() *int64 {
	return s.Id
}

func (s *GetParameterResponseBodyParameter) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetParameterResponseBodyParameter) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetParameterResponseBodyParameter) GetName() *string {
	return s.Name
}

func (s *GetParameterResponseBodyParameter) GetOwner() *string {
	return s.Owner
}

func (s *GetParameterResponseBodyParameter) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetParameterResponseBodyParameter) GetProperties() []*GetParameterResponseBodyParameterProperties {
	return s.Properties
}

func (s *GetParameterResponseBodyParameter) GetScope() *string {
	return s.Scope
}

func (s *GetParameterResponseBodyParameter) GetType() *string {
	return s.Type
}

func (s *GetParameterResponseBodyParameter) GetVersion() *int32 {
	return s.Version
}

func (s *GetParameterResponseBodyParameter) SetCreateTime(v int64) *GetParameterResponseBodyParameter {
	s.CreateTime = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetCreateUser(v string) *GetParameterResponseBodyParameter {
	s.CreateUser = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetDescription(v string) *GetParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetId(v int64) *GetParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetModifyTime(v int64) *GetParameterResponseBodyParameter {
	s.ModifyTime = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetModifyUser(v string) *GetParameterResponseBodyParameter {
	s.ModifyUser = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetName(v string) *GetParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetOwner(v string) *GetParameterResponseBodyParameter {
	s.Owner = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetProjectId(v int64) *GetParameterResponseBodyParameter {
	s.ProjectId = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetProperties(v []*GetParameterResponseBodyParameterProperties) *GetParameterResponseBodyParameter {
	s.Properties = v
	return s
}

func (s *GetParameterResponseBodyParameter) SetScope(v string) *GetParameterResponseBodyParameter {
	s.Scope = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetType(v string) *GetParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetVersion(v int32) *GetParameterResponseBodyParameter {
	s.Version = &v
	return s
}

func (s *GetParameterResponseBodyParameter) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetParameterResponseBodyParameterProperties struct {
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// value123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetParameterResponseBodyParameterProperties) String() string {
	return dara.Prettify(s)
}

func (s GetParameterResponseBodyParameterProperties) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBodyParameterProperties) GetEnvType() *string {
	return s.EnvType
}

func (s *GetParameterResponseBodyParameterProperties) GetValue() *string {
	return s.Value
}

func (s *GetParameterResponseBodyParameterProperties) SetEnvType(v string) *GetParameterResponseBodyParameterProperties {
	s.EnvType = &v
	return s
}

func (s *GetParameterResponseBodyParameterProperties) SetValue(v string) *GetParameterResponseBodyParameterProperties {
	s.Value = &v
	return s
}

func (s *GetParameterResponseBodyParameterProperties) Validate() error {
	return dara.Validate(s)
}
