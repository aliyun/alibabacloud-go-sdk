// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGreyTagRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlbRules(v string) *UpdateGreyTagRouteRequest
	GetAlbRules() *string
	SetDescription(v string) *UpdateGreyTagRouteRequest
	GetDescription() *string
	SetDubboRules(v string) *UpdateGreyTagRouteRequest
	GetDubboRules() *string
	SetGreyTagRouteId(v int64) *UpdateGreyTagRouteRequest
	GetGreyTagRouteId() *int64
	SetScRules(v string) *UpdateGreyTagRouteRequest
	GetScRules() *string
}

type UpdateGreyTagRouteRequest struct {
	// The canary release rule of the application for which ALB gateway routing is configured.
	//
	// example:
	//
	// [{"condition":"AND","items":[{"cond":"==","name":"grey","operator":"rawvalue","type":"sourceIp","value":"127.0.0.1"},{"cond":"==","name":"grey","operator":"rawvalue","type":"cookie","value":"true"},{"cond":"==","name":"grey","operator":"rawvalue","type":"header","value":"true"}],"path":"/post-echo/hi"}]
	AlbRules *string `json:"AlbRules,omitempty" xml:"AlbRules,omitempty"`
	// The description of the canary release rule.
	//
	// example:
	//
	// 灰度发布-地域灰度
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The canary release rule of the Dubbo application.
	//
	// example:
	//
	// [{"condition":"OR","group":"DUBBO","items":[{"cond":"==","expr":".key1","index":0,"operator":"rawvalue","value":"value1"},{"cond":"==","expr":".key2","index":0,"operator":"rawvalue","value":"value2"}],"methodName":"echo","serviceName":"com.alibaba.edas.boot.EchoService","version":"1.0.0"}]
	DubboRules *string `json:"DubboRules,omitempty" xml:"DubboRules,omitempty"`
	// The ID of the canary release rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
	// The canary release rule of the Spring Cloud application.
	//
	// example:
	//
	// [{"condition":"OR","items":[{"cond":"==","name":"grey","operator":"rawvalue","type":"param","value":"true"},{"cond":"==","name":"grey","operator":"rawvalue","type":"cookie","value":"true"},{"cond":"==","name":"grey","operator":"rawvalue","type":"header","value":"true"}],"path":"/post-echo/hi"}]
	ScRules *string `json:"ScRules,omitempty" xml:"ScRules,omitempty"`
}

func (s UpdateGreyTagRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateGreyTagRouteRequest) GetAlbRules() *string {
	return s.AlbRules
}

func (s *UpdateGreyTagRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGreyTagRouteRequest) GetDubboRules() *string {
	return s.DubboRules
}

func (s *UpdateGreyTagRouteRequest) GetGreyTagRouteId() *int64 {
	return s.GreyTagRouteId
}

func (s *UpdateGreyTagRouteRequest) GetScRules() *string {
	return s.ScRules
}

func (s *UpdateGreyTagRouteRequest) SetAlbRules(v string) *UpdateGreyTagRouteRequest {
	s.AlbRules = &v
	return s
}

func (s *UpdateGreyTagRouteRequest) SetDescription(v string) *UpdateGreyTagRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateGreyTagRouteRequest) SetDubboRules(v string) *UpdateGreyTagRouteRequest {
	s.DubboRules = &v
	return s
}

func (s *UpdateGreyTagRouteRequest) SetGreyTagRouteId(v int64) *UpdateGreyTagRouteRequest {
	s.GreyTagRouteId = &v
	return s
}

func (s *UpdateGreyTagRouteRequest) SetScRules(v string) *UpdateGreyTagRouteRequest {
	s.ScRules = &v
	return s
}

func (s *UpdateGreyTagRouteRequest) Validate() error {
	return dara.Validate(s)
}
