// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiStageVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateApiStageVariableRequest
	GetGroupId() *string
	SetSecurityToken(v string) *CreateApiStageVariableRequest
	GetSecurityToken() *string
	SetStageId(v string) *CreateApiStageVariableRequest
	GetStageId() *string
	SetStageRouteModel(v string) *CreateApiStageVariableRequest
	GetStageRouteModel() *string
	SetSupportRoute(v bool) *CreateApiStageVariableRequest
	GetSupportRoute() *bool
	SetVariableName(v string) *CreateApiStageVariableRequest
	GetVariableName() *string
	SetVariableValue(v string) *CreateApiStageVariableRequest
	GetVariableValue() *string
}

type CreateApiStageVariableRequest struct {
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 523e8dc7bbe04613b5b1d726c2a7889d
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the runtime environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// Deprecated
	//
	// The routing model of the environment.
	//
	// example:
	//
	// {
	//
	//     "location": "HEAD",
	//
	//     "parameterCatalog": "CUSTOM",
	//
	//     "parameterType": "String",
	//
	//     "serviceParameterName": "TestConstant",
	//
	//     "routeMatchSymbol": "IN",
	//
	//     "routeRules": [
	//
	//         {
	//
	//             "conditionValue": "aaa,bbb",
	//
	//             "resultValue": "apigateway-test.com"
	//
	//         }
	//
	//     ]
	//
	// }
	StageRouteModel *string `json:"StageRouteModel,omitempty" xml:"StageRouteModel,omitempty"`
	// Specifies whether routing is supported.
	//
	// example:
	//
	// true
	SupportRoute *bool `json:"SupportRoute,omitempty" xml:"SupportRoute,omitempty"`
	// The name of the variable to be added. This parameter is case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// serverName
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// api.domain.com
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s CreateApiStageVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiStageVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateApiStageVariableRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateApiStageVariableRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateApiStageVariableRequest) GetStageId() *string {
	return s.StageId
}

func (s *CreateApiStageVariableRequest) GetStageRouteModel() *string {
	return s.StageRouteModel
}

func (s *CreateApiStageVariableRequest) GetSupportRoute() *bool {
	return s.SupportRoute
}

func (s *CreateApiStageVariableRequest) GetVariableName() *string {
	return s.VariableName
}

func (s *CreateApiStageVariableRequest) GetVariableValue() *string {
	return s.VariableValue
}

func (s *CreateApiStageVariableRequest) SetGroupId(v string) *CreateApiStageVariableRequest {
	s.GroupId = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetSecurityToken(v string) *CreateApiStageVariableRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetStageId(v string) *CreateApiStageVariableRequest {
	s.StageId = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetStageRouteModel(v string) *CreateApiStageVariableRequest {
	s.StageRouteModel = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetSupportRoute(v bool) *CreateApiStageVariableRequest {
	s.SupportRoute = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetVariableName(v string) *CreateApiStageVariableRequest {
	s.VariableName = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetVariableValue(v string) *CreateApiStageVariableRequest {
	s.VariableValue = &v
	return s
}

func (s *CreateApiStageVariableRequest) Validate() error {
	return dara.Validate(s)
}
