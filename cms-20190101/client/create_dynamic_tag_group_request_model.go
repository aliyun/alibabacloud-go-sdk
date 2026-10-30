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
	SetDryRun(v bool) *CreateDynamicTagGroupRequest
	GetDryRun() *bool
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
	// The alert contact groups. Valid values of N: 1 to 100. Alert notifications for the application group are sent to the alert contacts in the alert contact groups.
	//
	// An alert contact group contains one or more alert contacts. For more information about how to create alert contacts and alert contact groups, see [PutContact](https://help.aliyun.com/document_detail/114923.html) and [PutContactGroup](https://help.aliyun.com/document_detail/114929.html). For more information about how to query alert contact groups, see [DescribeContactGroupList](https://help.aliyun.com/document_detail/114922.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroupList []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: Performs a dry run without creating the application group. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false**: Performs a dry run and sends the actual request. If the request passes the dry run, the application group is created. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable automatic installation of the CloudMonitor agent for the application group. CloudMonitor automatically installs the CloudMonitor agent on the hosts in the application group. Valid values:
	//
	// - true: Enables automatic installation.
	//
	// - false: Disables automatic installation. This is the default value.
	//
	// example:
	//
	// true
	EnableInstallAgent *bool `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	// Specifies whether to enable automatic subscription to event notifications for the application group. When critical or warning events occur on resources in the application group, CloudMonitor sends alert notifications. Valid values:
	//
	// - true: Enables automatic subscription.
	//
	// - false: Disables automatic subscription. This is the default value.
	//
	// example:
	//
	// true
	EnableSubscribeEvent *bool `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	// The matching expression used to generate the application group based on tags.
	//
	// This parameter is required.
	MatchExpress []*CreateDynamicTagGroupRequestMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
	// The relationship between the conditional expressions of the resource tag values. Valid values:
	//
	// - and: Uses the AND logical operator. This is the default value.
	//
	// - or: Uses the OR logical operator.
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
	// The region ID of the tag.
	//
	// example:
	//
	// cn-hangzhou
	TagRegionId *string `json:"TagRegionId,omitempty" xml:"TagRegionId,omitempty"`
	// The IDs of the alert templates.
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

func (s *CreateDynamicTagGroupRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *CreateDynamicTagGroupRequest) SetDryRun(v bool) *CreateDynamicTagGroupRequest {
	s.DryRun = &v
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
	// The tag key used to create the group. If multiple resources contain this key, the resources that meet the filter conditions are added to the same group based on the same key-value pair.
	//
	// example:
	//
	// appname
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The tag value of the resource. Valid values of N: 1.
	//
	// > You must specify both the `MatchExpress.N.TagValueMatchFunction` and `MatchExpress.N.TagValue` parameters.
	//
	// example:
	//
	// instance
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The matching method for the tag value of the resource. Valid values of N: 1. Valid values:
	//
	// - contains: includes the specified value
	//
	// - startWith: starts with the specified value
	//
	// - endWith: ends with the specified value
	//
	// - notContains: does not include the specified value
	//
	// - equals: equals the specified value
	//
	// - all: all values
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
