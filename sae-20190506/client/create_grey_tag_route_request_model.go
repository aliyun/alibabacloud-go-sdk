// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGreyTagRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlbRules(v string) *CreateGreyTagRouteRequest
	GetAlbRules() *string
	SetAppId(v string) *CreateGreyTagRouteRequest
	GetAppId() *string
	SetDescription(v string) *CreateGreyTagRouteRequest
	GetDescription() *string
	SetDubboRules(v string) *CreateGreyTagRouteRequest
	GetDubboRules() *string
	SetName(v string) *CreateGreyTagRouteRequest
	GetName() *string
	SetScRules(v string) *CreateGreyTagRouteRequest
	GetScRules() *string
}

type CreateGreyTagRouteRequest struct {
	// The canary release rules created for an application that has ALB gateway routing configured.
	//
	// example:
	//
	// [{"condition":"AND","items":[{"cond":"==","name":"grey","operator":"rawvalue","type":"sourceIp","value":"127.0.0.1"},{"cond":"==","name":"grey","operator":"rawvalue","type":"cookie","value":"true"},{"cond":"==","name":"grey","operator":"rawvalue","type":"header","value":"true"}],"path":"/post-echo/hi"}]
	AlbRules *string `json:"AlbRules,omitempty" xml:"AlbRules,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7802c49a-67bc-4167-8369-9a9c003c****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The rule description. It can be up to 64 characters in length.
	//
	// example:
	//
	// Canary Release - Regions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The canary release rules created for a Dubbo application. This parameter is required when your application uses the Dubbo framework. You do not need to specify the **ScRules*	- parameter.
	//
	// example:
	//
	// [{"condition":"OR","group":"DUBBO","items":[{"cond":"==","expr":".key1","index":0,"operator":"rawvalue","value":"value1"},{"cond":"==","expr":".key2","index":0,"operator":"rawvalue","value":"value2"}],"methodName":"echo","serviceName":"com.alibaba.edas.boot.EchoService","version":"1.0.0"}]
	DubboRules *string `json:"DubboRules,omitempty" xml:"DubboRules,omitempty"`
	// The rule name. It must start with a lowercase letter, end with a digit or a lowercase letter, and can contain only lowercase letters, Chinese characters, digits, and hyphens (-). It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// dubbo-echo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The canary release rules created for a Spring Cloud application. This parameter is required when your application uses the Spring Cloud framework. You do not need to specify the **DubboRules*	- parameter.
	//
	// example:
	//
	// [{"condition":"OR","items":[{"cond":"==","name":"grey","operator":"rawvalue","type":"param","value":"true"},{"cond":"==","name":"grey","operator":"rawvalue","type":"cookie","value":"true"},{"cond":"==","name":"grey","operator":"rawvalue","type":"header","value":"true"}],"path":"/post-echo/hi"}]
	ScRules *string `json:"ScRules,omitempty" xml:"ScRules,omitempty"`
}

func (s CreateGreyTagRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateGreyTagRouteRequest) GetAlbRules() *string {
	return s.AlbRules
}

func (s *CreateGreyTagRouteRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateGreyTagRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGreyTagRouteRequest) GetDubboRules() *string {
	return s.DubboRules
}

func (s *CreateGreyTagRouteRequest) GetName() *string {
	return s.Name
}

func (s *CreateGreyTagRouteRequest) GetScRules() *string {
	return s.ScRules
}

func (s *CreateGreyTagRouteRequest) SetAlbRules(v string) *CreateGreyTagRouteRequest {
	s.AlbRules = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetAppId(v string) *CreateGreyTagRouteRequest {
	s.AppId = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetDescription(v string) *CreateGreyTagRouteRequest {
	s.Description = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetDubboRules(v string) *CreateGreyTagRouteRequest {
	s.DubboRules = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetName(v string) *CreateGreyTagRouteRequest {
	s.Name = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetScRules(v string) *CreateGreyTagRouteRequest {
	s.ScRules = &v
	return s
}

func (s *CreateGreyTagRouteRequest) Validate() error {
	return dara.Validate(s)
}
