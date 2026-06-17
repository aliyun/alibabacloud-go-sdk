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
	// The alert contact groups. The value of N can be from 1 to 100. Alert notifications for the application group are sent to the alert contacts in these alert contact groups.
	//
	// An alert contact group can contain one or more alert contacts. For more information about how to create alert contacts and alert contact groups, see [PutContact](https://help.aliyun.com/document_detail/114923.html) and [PutContactGroup](https://help.aliyun.com/document_detail/114929.html). For more information about how to obtain alert contact groups, see [DescribeContactGroupList](https://help.aliyun.com/document_detail/114922.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroupList []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
	// Specifies whether to automatically install the CloudMonitor agent for the application group. CloudMonitor automatically installs the agent on the hosts in the application group. Valid values:
	//
	// - true: enabled.
	//
	// - false (default): disabled.
	//
	// example:
	//
	// true
	EnableInstallAgent *bool `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	// Specifies whether to automatically subscribe to event notifications for the application group. When a critical or warning event occurs on a resource in the application group, CloudMonitor sends an alert notification. Valid values:
	//
	// - true: enabled.
	//
	// - false (default): disabled.
	//
	// example:
	//
	// true
	EnableSubscribeEvent *bool `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	// The match expressions that are used to create an application group from tags.
	//
	// This parameter is required.
	MatchExpress []*CreateDynamicTagGroupRequestMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
	// The relationship between the conditional expressions for the tag values. Valid values:
	//
	// - and (default)
	//
	// - or
	//
	// example:
	//
	// and
	MatchExpressFilterRelation *string `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag key of the resource.
	//
	// For more information about how to query the tag keys of resources, see [DescribeTagKeyList](https://help.aliyun.com/document_detail/145558.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs_instance
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The ID of the region to which the tag belongs.
	//
	// example:
	//
	// cn-hangzhou
	TagRegionId *string `json:"TagRegionId,omitempty" xml:"TagRegionId,omitempty"`
	// The ID of the alert template.
	//
	// For more information about how to query the IDs of alert templates, see [DescribeMetricRuleTemplateList](https://help.aliyun.com/document_detail/114982.html).
	//
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
	if s.MatchExpress != nil {
		for _, item := range s.MatchExpress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDynamicTagGroupRequestMatchExpress struct {
	// The key of the tag that is used to create the group. If multiple resources have this tag key, the resources that meet the filter conditions are added to the same group based on the same key-value pair.
	//
	// example:
	//
	// appname
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The value of the resource tag. The value of N is 1.
	//
	// > You must specify both the `MatchExpress.N.TagValueMatchFunction` and `MatchExpress.N.TagValue` parameters.
	//
	// example:
	//
	// instance
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The method that is used to match the values of resource tags. The value of N is 1. Valid values:
	//
	// - contains: includes.
	//
	// - startWith: prefix.
	//
	// - endWith: suffix.
	//
	// - notContains: does not include.
	//
	// - equals: equals.
	//
	// - all: all.
	//
	// > You must specify both the `MatchExpress.N.TagValueMatchFunction` and `MatchExpress.N.TagValue` parameters.
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
