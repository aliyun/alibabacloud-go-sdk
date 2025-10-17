// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstanceResponseBodyInstances) *ListInstanceResponseBody
	GetInstances() []*ListInstanceResponseBodyInstances
	SetPageNumber(v int64) *ListInstanceResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListInstanceResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListInstanceResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListInstanceResponseBody
	GetTotalCount() *int64
}

type ListInstanceResponseBody struct {
	Instances []*ListInstanceResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 92B81548-42B9-4B34-924B-4E778AEB412B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) GetInstances() []*ListInstanceResponseBodyInstances {
	return s.Instances
}

func (s *ListInstanceResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstanceResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceResponseBody) SetInstances(v []*ListInstanceResponseBodyInstances) *ListInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstanceResponseBody) SetPageNumber(v int64) *ListInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceResponseBody) SetPageSize(v int64) *ListInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetTotalCount(v int64) *ListInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResponseBodyInstances struct {
	// example:
	//
	// /alimefe/meebot/robot/0.0.5/img/xxx-90-97.png
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 用于C端问答的机器人
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// example:
	//
	// 智能客服-小C
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// scenario_im
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s ListInstanceResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyInstances) GetAvatar() *string {
	return s.Avatar
}

func (s *ListInstanceResponseBodyInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstanceResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceResponseBodyInstances) GetIntroduction() *string {
	return s.Introduction
}

func (s *ListInstanceResponseBodyInstances) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *ListInstanceResponseBodyInstances) GetName() *string {
	return s.Name
}

func (s *ListInstanceResponseBodyInstances) GetRobotType() *string {
	return s.RobotType
}

func (s *ListInstanceResponseBodyInstances) SetAvatar(v string) *ListInstanceResponseBodyInstances {
	s.Avatar = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetCreateTime(v string) *ListInstanceResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceId(v string) *ListInstanceResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetIntroduction(v string) *ListInstanceResponseBodyInstances {
	s.Introduction = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetLanguageCode(v string) *ListInstanceResponseBodyInstances {
	s.LanguageCode = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetName(v string) *ListInstanceResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetRobotType(v string) *ListInstanceResponseBodyInstances {
	s.RobotType = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
