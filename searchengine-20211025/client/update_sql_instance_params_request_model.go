// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCombineParam(v map[string]interface{}) *UpdateSqlInstanceParamsRequest
	GetCombineParam() map[string]interface{}
	SetDynamicParam(v map[string]interface{}) *UpdateSqlInstanceParamsRequest
	GetDynamicParam() map[string]interface{}
	SetKvpair(v map[string]interface{}) *UpdateSqlInstanceParamsRequest
	GetKvpair() map[string]interface{}
	SetParams(v map[string]interface{}) *UpdateSqlInstanceParamsRequest
	GetParams() map[string]interface{}
	SetStaticParam(v map[string]interface{}) *UpdateSqlInstanceParamsRequest
	GetStaticParam() map[string]interface{}
}

type UpdateSqlInstanceParamsRequest struct {
	CombineParam map[string]interface{} `json:"combineParam,omitempty" xml:"combineParam,omitempty"`
	DynamicParam map[string]interface{} `json:"dynamicParam,omitempty" xml:"dynamicParam,omitempty"`
	Kvpair       map[string]interface{} `json:"kvpair,omitempty" xml:"kvpair,omitempty"`
	Params       map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	StaticParam  map[string]interface{} `json:"staticParam,omitempty" xml:"staticParam,omitempty"`
}

func (s UpdateSqlInstanceParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceParamsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceParamsRequest) GetCombineParam() map[string]interface{} {
	return s.CombineParam
}

func (s *UpdateSqlInstanceParamsRequest) GetDynamicParam() map[string]interface{} {
	return s.DynamicParam
}

func (s *UpdateSqlInstanceParamsRequest) GetKvpair() map[string]interface{} {
	return s.Kvpair
}

func (s *UpdateSqlInstanceParamsRequest) GetParams() map[string]interface{} {
	return s.Params
}

func (s *UpdateSqlInstanceParamsRequest) GetStaticParam() map[string]interface{} {
	return s.StaticParam
}

func (s *UpdateSqlInstanceParamsRequest) SetCombineParam(v map[string]interface{}) *UpdateSqlInstanceParamsRequest {
	s.CombineParam = v
	return s
}

func (s *UpdateSqlInstanceParamsRequest) SetDynamicParam(v map[string]interface{}) *UpdateSqlInstanceParamsRequest {
	s.DynamicParam = v
	return s
}

func (s *UpdateSqlInstanceParamsRequest) SetKvpair(v map[string]interface{}) *UpdateSqlInstanceParamsRequest {
	s.Kvpair = v
	return s
}

func (s *UpdateSqlInstanceParamsRequest) SetParams(v map[string]interface{}) *UpdateSqlInstanceParamsRequest {
	s.Params = v
	return s
}

func (s *UpdateSqlInstanceParamsRequest) SetStaticParam(v map[string]interface{}) *UpdateSqlInstanceParamsRequest {
	s.StaticParam = v
	return s
}

func (s *UpdateSqlInstanceParamsRequest) Validate() error {
	return dara.Validate(s)
}
