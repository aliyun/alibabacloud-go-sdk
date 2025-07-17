// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenarioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArmsScenarios(v []*ListScenarioResponseBodyArmsScenarios) *ListScenarioResponseBody
	GetArmsScenarios() []*ListScenarioResponseBodyArmsScenarios
	SetRequestId(v string) *ListScenarioResponseBody
	GetRequestId() *string
}

type ListScenarioResponseBody struct {
	// The detailed information of the business monitoring job.
	ArmsScenarios []*ListScenarioResponseBodyArmsScenarios `json:"ArmsScenarios,omitempty" xml:"ArmsScenarios,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 98027D1F-3AEB-492C-A4AA-E9217992****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListScenarioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenarioResponseBody) GetArmsScenarios() []*ListScenarioResponseBodyArmsScenarios {
	return s.ArmsScenarios
}

func (s *ListScenarioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScenarioResponseBody) SetArmsScenarios(v []*ListScenarioResponseBodyArmsScenarios) *ListScenarioResponseBody {
	s.ArmsScenarios = v
	return s
}

func (s *ListScenarioResponseBody) SetRequestId(v string) *ListScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScenarioResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListScenarioResponseBodyArmsScenarios struct {
	// The ID of the application.
	//
	// example:
	//
	// b590lhguqs@28f515462******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the business monitoring job was created.
	//
	// example:
	//
	// 1585214916000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extended information. The value is a JSON string.
	//
	// example:
	//
	// {"_MODE": "CUSTOM-TRANSACTION","_SCENARIO": "USER-DEFINED"}
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// The ID of the business monitoring job.
	//
	// example:
	//
	// 132
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the business monitoring job.
	//
	// example:
	//
	// k8s_deployment_css-guns-vip-main-prod_silence
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The code of the business monitoring job.
	//
	// example:
	//
	// a9f8****
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// The time when the business monitoring job was updated.
	//
	// example:
	//
	// 1585214916000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 113197164949****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListScenarioResponseBodyArmsScenarios) String() string {
	return dara.Prettify(s)
}

func (s ListScenarioResponseBodyArmsScenarios) GoString() string {
	return s.String()
}

func (s *ListScenarioResponseBodyArmsScenarios) GetAppId() *string {
	return s.AppId
}

func (s *ListScenarioResponseBodyArmsScenarios) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListScenarioResponseBodyArmsScenarios) GetExtensions() *string {
	return s.Extensions
}

func (s *ListScenarioResponseBodyArmsScenarios) GetId() *int64 {
	return s.Id
}

func (s *ListScenarioResponseBodyArmsScenarios) GetName() *string {
	return s.Name
}

func (s *ListScenarioResponseBodyArmsScenarios) GetRegionId() *string {
	return s.RegionId
}

func (s *ListScenarioResponseBodyArmsScenarios) GetSign() *string {
	return s.Sign
}

func (s *ListScenarioResponseBodyArmsScenarios) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListScenarioResponseBodyArmsScenarios) GetUserId() *string {
	return s.UserId
}

func (s *ListScenarioResponseBodyArmsScenarios) SetAppId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.AppId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetCreateTime(v string) *ListScenarioResponseBodyArmsScenarios {
	s.CreateTime = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetExtensions(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Extensions = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetId(v int64) *ListScenarioResponseBodyArmsScenarios {
	s.Id = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetName(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Name = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetRegionId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.RegionId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetSign(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Sign = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetUpdateTime(v string) *ListScenarioResponseBodyArmsScenarios {
	s.UpdateTime = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetUserId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.UserId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) Validate() error {
	return dara.Validate(s)
}
