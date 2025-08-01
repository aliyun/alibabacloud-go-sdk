// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceFuncStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *GetServiceFuncStatusRequest
	GetChannel() *string
	SetParams(v *GetServiceFuncStatusRequestParams) *GetServiceFuncStatusRequest
	GetParams() *GetServiceFuncStatusRequestParams
	SetServiceName(v string) *GetServiceFuncStatusRequest
	GetServiceName() *string
}

type GetServiceFuncStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	Params *GetServiceFuncStatusRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// livetrace
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s GetServiceFuncStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceFuncStatusRequest) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetServiceFuncStatusRequest) GetParams() *GetServiceFuncStatusRequestParams {
	return s.Params
}

func (s *GetServiceFuncStatusRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceFuncStatusRequest) SetChannel(v string) *GetServiceFuncStatusRequest {
	s.Channel = &v
	return s
}

func (s *GetServiceFuncStatusRequest) SetParams(v *GetServiceFuncStatusRequestParams) *GetServiceFuncStatusRequest {
	s.Params = v
	return s
}

func (s *GetServiceFuncStatusRequest) SetServiceName(v string) *GetServiceFuncStatusRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceFuncStatusRequest) Validate() error {
	return dara.Validate(s)
}

type GetServiceFuncStatusRequestParams struct {
	// This parameter is required.
	//
	// example:
	//
	// mullprof
	FunctionName *string `json:"function_name,omitempty" xml:"function_name,omitempty"`
	// example:
	//
	// i-2zei55fwj8nnu31h3z46
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 1338904783509062
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetServiceFuncStatusRequestParams) String() string {
	return dara.Prettify(s)
}

func (s GetServiceFuncStatusRequestParams) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusRequestParams) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetServiceFuncStatusRequestParams) GetInstance() *string {
	return s.Instance
}

func (s *GetServiceFuncStatusRequestParams) GetUid() *string {
	return s.Uid
}

func (s *GetServiceFuncStatusRequestParams) SetFunctionName(v string) *GetServiceFuncStatusRequestParams {
	s.FunctionName = &v
	return s
}

func (s *GetServiceFuncStatusRequestParams) SetInstance(v string) *GetServiceFuncStatusRequestParams {
	s.Instance = &v
	return s
}

func (s *GetServiceFuncStatusRequestParams) SetUid(v string) *GetServiceFuncStatusRequestParams {
	s.Uid = &v
	return s
}

func (s *GetServiceFuncStatusRequestParams) Validate() error {
	return dara.Validate(s)
}
