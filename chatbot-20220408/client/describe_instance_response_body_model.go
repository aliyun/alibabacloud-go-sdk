// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *DescribeInstanceResponseBody
	GetAvatar() *string
	SetCategories(v []*DescribeInstanceResponseBodyCategories) *DescribeInstanceResponseBody
	GetCategories() []*DescribeInstanceResponseBodyCategories
	SetCreateTime(v string) *DescribeInstanceResponseBody
	GetCreateTime() *string
	SetEditStatus(v string) *DescribeInstanceResponseBody
	GetEditStatus() *string
	SetInstanceId(v string) *DescribeInstanceResponseBody
	GetInstanceId() *string
	SetIntroduction(v string) *DescribeInstanceResponseBody
	GetIntroduction() *string
	SetLanguageCode(v string) *DescribeInstanceResponseBody
	GetLanguageCode() *string
	SetName(v string) *DescribeInstanceResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetRobotType(v string) *DescribeInstanceResponseBody
	GetRobotType() *string
	SetTimeZone(v string) *DescribeInstanceResponseBody
	GetTimeZone() *string
}

type DescribeInstanceResponseBody struct {
	// example:
	//
	// /alimefe/meebot/robot/0.0.5/img/xxx-90-97.png
	Avatar     *string                                   `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Categories []*DescribeInstanceResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// PUBLISHED
	EditStatus *string `json:"EditStatus,omitempty" xml:"EditStatus,omitempty"`
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
	// 907AA5F2-0521-49AB-80AB-1ADEFAB2B901
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// scenario_im
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
	// example:
	//
	// Asia/Chongqing
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetAvatar() *string {
	return s.Avatar
}

func (s *DescribeInstanceResponseBody) GetCategories() []*DescribeInstanceResponseBodyCategories {
	return s.Categories
}

func (s *DescribeInstanceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInstanceResponseBody) GetEditStatus() *string {
	return s.EditStatus
}

func (s *DescribeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBody) GetIntroduction() *string {
	return s.Introduction
}

func (s *DescribeInstanceResponseBody) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *DescribeInstanceResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetRobotType() *string {
	return s.RobotType
}

func (s *DescribeInstanceResponseBody) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeInstanceResponseBody) SetAvatar(v string) *DescribeInstanceResponseBody {
	s.Avatar = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCategories(v []*DescribeInstanceResponseBodyCategories) *DescribeInstanceResponseBody {
	s.Categories = v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreateTime(v string) *DescribeInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEditStatus(v string) *DescribeInstanceResponseBody {
	s.EditStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIntroduction(v string) *DescribeInstanceResponseBody {
	s.Introduction = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetLanguageCode(v string) *DescribeInstanceResponseBody {
	s.LanguageCode = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetName(v string) *DescribeInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRobotType(v string) *DescribeInstanceResponseBody {
	s.RobotType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTimeZone(v string) *DescribeInstanceResponseBody {
	s.TimeZone = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	if s.Categories != nil {
		for _, item := range s.Categories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceResponseBodyCategories struct {
	AbilityType *string `json:"AbilityType,omitempty" xml:"AbilityType,omitempty"`
	// example:
	//
	// 30000066832
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 杭州市防疫政策
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s DescribeInstanceResponseBodyCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyCategories) GetAbilityType() *string {
	return s.AbilityType
}

func (s *DescribeInstanceResponseBodyCategories) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DescribeInstanceResponseBodyCategories) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBodyCategories) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *DescribeInstanceResponseBodyCategories) SetAbilityType(v string) *DescribeInstanceResponseBodyCategories {
	s.AbilityType = &v
	return s
}

func (s *DescribeInstanceResponseBodyCategories) SetCategoryId(v int64) *DescribeInstanceResponseBodyCategories {
	s.CategoryId = &v
	return s
}

func (s *DescribeInstanceResponseBodyCategories) SetName(v string) *DescribeInstanceResponseBodyCategories {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyCategories) SetParentCategoryId(v int64) *DescribeInstanceResponseBodyCategories {
	s.ParentCategoryId = &v
	return s
}

func (s *DescribeInstanceResponseBodyCategories) Validate() error {
	return dara.Validate(s)
}
