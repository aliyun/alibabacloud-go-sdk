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
	ArmsScenarios []*ListScenarioResponseBodyArmsScenarios `json:"ArmsScenarios,omitempty" xml:"ArmsScenarios,omitempty" type:"Repeated"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.ArmsScenarios != nil {
		for _, item := range s.ArmsScenarios {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScenarioResponseBodyArmsScenarios struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sign       *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
