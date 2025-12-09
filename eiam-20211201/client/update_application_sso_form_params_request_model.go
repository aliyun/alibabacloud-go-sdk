// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationSsoFormParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationSsoFormParamsRequest
	GetApplicationId() *string
	SetApplicationTemplateParams(v []*UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) *UpdateApplicationSsoFormParamsRequest
	GetApplicationTemplateParams() []*UpdateApplicationSsoFormParamsRequestApplicationTemplateParams
	SetInstanceId(v string) *UpdateApplicationSsoFormParamsRequest
	GetInstanceId() *string
}

type UpdateApplicationSsoFormParamsRequest struct {
	// IDaaS的应用主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// app_11111
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用模板创建参数，应用创建来源为模板时才可以指定
	//
	// This parameter is required.
	ApplicationTemplateParams []*UpdateApplicationSsoFormParamsRequestApplicationTemplateParams `json:"ApplicationTemplateParams,omitempty" xml:"ApplicationTemplateParams,omitempty" type:"Repeated"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateApplicationSsoFormParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationSsoFormParamsRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationSsoFormParamsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationSsoFormParamsRequest) GetApplicationTemplateParams() []*UpdateApplicationSsoFormParamsRequestApplicationTemplateParams {
	return s.ApplicationTemplateParams
}

func (s *UpdateApplicationSsoFormParamsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationSsoFormParamsRequest) SetApplicationId(v string) *UpdateApplicationSsoFormParamsRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationSsoFormParamsRequest) SetApplicationTemplateParams(v []*UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) *UpdateApplicationSsoFormParamsRequest {
	s.ApplicationTemplateParams = v
	return s
}

func (s *UpdateApplicationSsoFormParamsRequest) SetInstanceId(v string) *UpdateApplicationSsoFormParamsRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationSsoFormParamsRequest) Validate() error {
	if s.ApplicationTemplateParams != nil {
		for _, item := range s.ApplicationTemplateParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationSsoFormParamsRequestApplicationTemplateParams struct {
	// 应用模板创建参数具体名称
	//
	// example:
	//
	// aliyunUid
	TemplateParamName *string `json:"TemplateParamName,omitempty" xml:"TemplateParamName,omitempty"`
	// 应用模板创建参数真实的取值
	//
	// example:
	//
	// 123456789
	TemplateParamValue *string `json:"TemplateParamValue,omitempty" xml:"TemplateParamValue,omitempty"`
}

func (s UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) GoString() string {
	return s.String()
}

func (s *UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) GetTemplateParamName() *string {
	return s.TemplateParamName
}

func (s *UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) GetTemplateParamValue() *string {
	return s.TemplateParamValue
}

func (s *UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) SetTemplateParamName(v string) *UpdateApplicationSsoFormParamsRequestApplicationTemplateParams {
	s.TemplateParamName = &v
	return s
}

func (s *UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) SetTemplateParamValue(v string) *UpdateApplicationSsoFormParamsRequestApplicationTemplateParams {
	s.TemplateParamValue = &v
	return s
}

func (s *UpdateApplicationSsoFormParamsRequestApplicationTemplateParams) Validate() error {
	return dara.Validate(s)
}
