// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupList(v []*string) *CreateDynamicTagGroupRequest
	GetContactGroupList() []*string
	SetEnableInstallAgent(v bool) *CreateDynamicTagGroupRequest
	GetEnableInstallAgent() *bool
	SetEnableSubscribeEvent(v bool) *CreateDynamicTagGroupRequest
	GetEnableSubscribeEvent() *bool
	SetMatchExpress(v []*CreateDynamicTagGroupRequestMatchExpress) *CreateDynamicTagGroupRequest
	GetMatchExpress() []*CreateDynamicTagGroupRequestMatchExpress
	SetMatchExpressFilterRelation(v string) *CreateDynamicTagGroupRequest
	GetMatchExpressFilterRelation() *string
	SetRegionId(v string) *CreateDynamicTagGroupRequest
	GetRegionId() *string
	SetTagKey(v string) *CreateDynamicTagGroupRequest
	GetTagKey() *string
	SetTagRegionId(v string) *CreateDynamicTagGroupRequest
	GetTagRegionId() *string
	SetTemplateIdList(v []*string) *CreateDynamicTagGroupRequest
	GetTemplateIdList() []*string
}

type CreateDynamicTagGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroupList []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
	// Specifies whether the CloudMonitor agent is automatically installed for the application group. CloudMonitor determines whether to automatically install the CloudMonitor agent for the hosts in an application group based on the value of this parameter. Valid values:
	//
	// 	- true: The CloudMonitor agent is automatically installed.
	//
	// 	- false (default value): The CloudMonitor agent is not automatically installed.
	//
	// example:
	//
	// true
	EnableInstallAgent *bool `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	// Specifies whether the application group automatically subscribes to event notifications. If events whose severity level is critical or warning occur on resources in an application group, CloudMonitor sends alert notifications. Valid values:
	//
	// 	- true: The application group automatically subscribes to event notifications.
	//
	// 	- false (default value): The application group does not automatically subscribe to event notifications.
	//
	// example:
	//
	// true
	EnableSubscribeEvent *bool `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	// The conditional expressions used to create an application group based on the tag.
	//
	// This parameter is required.
	MatchExpress []*CreateDynamicTagGroupRequestMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
	// The relationship between the conditional expressions for the tag values of the cloud resources. Valid values:
	//
	// 	- and (default)
	//
	// 	- or
	//
	// example:
	//
	// and
	MatchExpressFilterRelation *string `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag keys of the cloud resources.
	//
	// For more information about how to obtain tag keys, see [DescribeTagKeyList](https://help.aliyun.com/document_detail/145558.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs_instance
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The ID of the region to which the tags belong.
	//
	// example:
	//
	// cn-hangzhou
	TagRegionId *string `json:"TagRegionId,omitempty" xml:"TagRegionId,omitempty"`
	// example:
	//
	// 85****
	TemplateIdList []*string `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
}

func (s CreateDynamicTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicTagGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDynamicTagGroupRequest) GetContactGroupList() []*string {
	return s.ContactGroupList
}

func (s *CreateDynamicTagGroupRequest) GetEnableInstallAgent() *bool {
	return s.EnableInstallAgent
}

func (s *CreateDynamicTagGroupRequest) GetEnableSubscribeEvent() *bool {
	return s.EnableSubscribeEvent
}

func (s *CreateDynamicTagGroupRequest) GetMatchExpress() []*CreateDynamicTagGroupRequestMatchExpress {
	return s.MatchExpress
}

func (s *CreateDynamicTagGroupRequest) GetMatchExpressFilterRelation() *string {
	return s.MatchExpressFilterRelation
}

func (s *CreateDynamicTagGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDynamicTagGroupRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateDynamicTagGroupRequest) GetTagRegionId() *string {
	return s.TagRegionId
}

func (s *CreateDynamicTagGroupRequest) GetTemplateIdList() []*string {
	return s.TemplateIdList
}

func (s *CreateDynamicTagGroupRequest) SetContactGroupList(v []*string) *CreateDynamicTagGroupRequest {
	s.ContactGroupList = v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetEnableInstallAgent(v bool) *CreateDynamicTagGroupRequest {
	s.EnableInstallAgent = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetEnableSubscribeEvent(v bool) *CreateDynamicTagGroupRequest {
	s.EnableSubscribeEvent = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetMatchExpress(v []*CreateDynamicTagGroupRequestMatchExpress) *CreateDynamicTagGroupRequest {
	s.MatchExpress = v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetMatchExpressFilterRelation(v string) *CreateDynamicTagGroupRequest {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetRegionId(v string) *CreateDynamicTagGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetTagKey(v string) *CreateDynamicTagGroupRequest {
	s.TagKey = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetTagRegionId(v string) *CreateDynamicTagGroupRequest {
	s.TagRegionId = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetTemplateIdList(v []*string) *CreateDynamicTagGroupRequest {
	s.TemplateIdList = v
	return s
}

func (s *CreateDynamicTagGroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDynamicTagGroupRequestMatchExpress struct {
	// The keys of the tags that are used to create the application group. If a specified key is attached to multiple resources, the resources that have the same key-value pair are added to the same group.
	//
	// example:
	//
	// appname
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The tag values of the cloud resources. Set the value of N to 1.
	//
	// >  If you set the `MatchExpress.N.TagValueMatchFunction` parameter, you must also set the `MatchExpress.N.TagValue` parameter.
	//
	// example:
	//
	// instance
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The method that is used to match the tag values of the cloud resources. Set the value of N to 1. Valid values:
	//
	// 	- contains: contains
	//
	// 	- startWith: starts with a prefix
	//
	// 	- endWith: ends with a suffix
	//
	// 	- notContains: does not contain
	//
	// 	- equals: equals
	//
	// 	- all: matches all
	//
	// >  If you set the `MatchExpress.N.TagValueMatchFunction` parameter, you must also set the `MatchExpress.N.TagValue` parameter.
	//
	// example:
	//
	// contains
	TagValueMatchFunction *string `json:"TagValueMatchFunction,omitempty" xml:"TagValueMatchFunction,omitempty"`
}

func (s CreateDynamicTagGroupRequestMatchExpress) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicTagGroupRequestMatchExpress) GoString() string {
	return s.String()
}

func (s *CreateDynamicTagGroupRequestMatchExpress) GetTagName() *string {
	return s.TagName
}

func (s *CreateDynamicTagGroupRequestMatchExpress) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateDynamicTagGroupRequestMatchExpress) GetTagValueMatchFunction() *string {
	return s.TagValueMatchFunction
}

func (s *CreateDynamicTagGroupRequestMatchExpress) SetTagName(v string) *CreateDynamicTagGroupRequestMatchExpress {
	s.TagName = &v
	return s
}

func (s *CreateDynamicTagGroupRequestMatchExpress) SetTagValue(v string) *CreateDynamicTagGroupRequestMatchExpress {
	s.TagValue = &v
	return s
}

func (s *CreateDynamicTagGroupRequestMatchExpress) SetTagValueMatchFunction(v string) *CreateDynamicTagGroupRequestMatchExpress {
	s.TagValueMatchFunction = &v
	return s
}

func (s *CreateDynamicTagGroupRequestMatchExpress) Validate() error {
	return dara.Validate(s)
}
