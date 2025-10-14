// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicTagRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDynamicTagRuleListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeDynamicTagRuleListResponseBody
	GetMessage() *string
	SetPageNumber(v string) *DescribeDynamicTagRuleListResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeDynamicTagRuleListResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeDynamicTagRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDynamicTagRuleListResponseBody
	GetSuccess() *bool
	SetTagGroupList(v *DescribeDynamicTagRuleListResponseBodyTagGroupList) *DescribeDynamicTagRuleListResponseBody
	GetTagGroupList() *DescribeDynamicTagRuleListResponseBodyTagGroupList
	SetTotal(v int32) *DescribeDynamicTagRuleListResponseBody
	GetTotal() *int32
}

type DescribeDynamicTagRuleListResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0D50523D-8D59-4A61-B58E-E2286ECFB3A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The tag rules of application groups.
	TagGroupList *DescribeDynamicTagRuleListResponseBodyTagGroupList `json:"TagGroupList,omitempty" xml:"TagGroupList,omitempty" type:"Struct"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDynamicTagRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDynamicTagRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDynamicTagRuleListResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDynamicTagRuleListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDynamicTagRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDynamicTagRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDynamicTagRuleListResponseBody) GetTagGroupList() *DescribeDynamicTagRuleListResponseBodyTagGroupList {
	return s.TagGroupList
}

func (s *DescribeDynamicTagRuleListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeDynamicTagRuleListResponseBody) SetCode(v string) *DescribeDynamicTagRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetMessage(v string) *DescribeDynamicTagRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetPageNumber(v string) *DescribeDynamicTagRuleListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetPageSize(v string) *DescribeDynamicTagRuleListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetRequestId(v string) *DescribeDynamicTagRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetSuccess(v bool) *DescribeDynamicTagRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetTagGroupList(v *DescribeDynamicTagRuleListResponseBodyTagGroupList) *DescribeDynamicTagRuleListResponseBody {
	s.TagGroupList = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetTotal(v int32) *DescribeDynamicTagRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) Validate() error {
	if s.TagGroupList != nil {
		if err := s.TagGroupList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDynamicTagRuleListResponseBodyTagGroupList struct {
	TagGroup []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup `json:"TagGroup,omitempty" xml:"TagGroup,omitempty" type:"Repeated"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupList) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupList) GetTagGroup() []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	return s.TagGroup
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupList) SetTagGroup(v []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) *DescribeDynamicTagRuleListResponseBodyTagGroupList {
	s.TagGroup = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupList) Validate() error {
	if s.TagGroup != nil {
		for _, item := range s.TagGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup struct {
	// The alert contact group.
	ContactGroupList *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Struct"`
	// The ID of the tag rule.
	//
	// example:
	//
	// 1536df65-a719-429d-8813-73cc40d7****
	DynamicTagRuleId *string `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
	// The conditional expressions used to create an application group based on the tag.
	MatchExpress *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Struct"`
	// The logical operator that is used between conditional expressions. Valid values:
	//
	// 	- `and`
	//
	// 	- `or`
	//
	// >  Only one logical operator can be used in a request.
	//
	// example:
	//
	// or
	MatchExpressFilterRelation *string `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	// The ID of the region to which the tags belong.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of adding instances that meet the tag rule to the application group. Valid values:
	//
	// 	- `RUNNING`
	//
	// 	- `FINISH`
	//
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag key.
	//
	// example:
	//
	// tagkey1
	TagKey            *string                                                                      `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValueBlacklist *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist `json:"TagValueBlacklist,omitempty" xml:"TagValueBlacklist,omitempty" type:"Struct"`
	// The IDs of the alert templates.
	TemplateIdList *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Struct"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetContactGroupList() *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList {
	return s.ContactGroupList
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetDynamicTagRuleId() *string {
	return s.DynamicTagRuleId
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetMatchExpress() *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress {
	return s.MatchExpress
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetMatchExpressFilterRelation() *string {
	return s.MatchExpressFilterRelation
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetStatus() *string {
	return s.Status
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetTagValueBlacklist() *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist {
	return s.TagValueBlacklist
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GetTemplateIdList() *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList {
	return s.TemplateIdList
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetContactGroupList(v *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.ContactGroupList = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetDynamicTagRuleId(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.DynamicTagRuleId = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetMatchExpress(v *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.MatchExpress = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetMatchExpressFilterRelation(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetRegionId(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.RegionId = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetStatus(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.Status = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetTagKey(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.TagKey = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetTagValueBlacklist(v *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.TagValueBlacklist = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetTemplateIdList(v *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.TemplateIdList = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) Validate() error {
	if s.ContactGroupList != nil {
		if err := s.ContactGroupList.Validate(); err != nil {
			return err
		}
	}
	if s.MatchExpress != nil {
		if err := s.MatchExpress.Validate(); err != nil {
			return err
		}
	}
	if s.TagValueBlacklist != nil {
		if err := s.TagValueBlacklist.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateIdList != nil {
		if err := s.TemplateIdList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList struct {
	ContactGroupList []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList) GetContactGroupList() []*string {
	return s.ContactGroupList
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList) SetContactGroupList(v []*string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList {
	s.ContactGroupList = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupContactGroupList) Validate() error {
	return dara.Validate(s)
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress struct {
	MatchExpress []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) GetMatchExpress() []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress {
	return s.MatchExpress
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) SetMatchExpress(v []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress {
	s.MatchExpress = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) Validate() error {
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

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress struct {
	// The tag key.
	//
	// example:
	//
	// azone-version
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// The `TagValue` and `TagValueMatchFunction` parameters must be used in pairs.
	//
	// example:
	//
	// *
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The method that is used to match tag values. Valid values:
	//
	// 	- all: includes all
	//
	// 	- startWith: starts with a prefix
	//
	// 	- endWith: ends with a suffix
	//
	// 	- contains: contains
	//
	// 	- notContains: does not contain
	//
	// 	- equals: equals
	//
	// example:
	//
	// all
	TagValueMatchFunction *string `json:"TagValueMatchFunction,omitempty" xml:"TagValueMatchFunction,omitempty"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) GetTagValueMatchFunction() *string {
	return s.TagValueMatchFunction
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) SetTagKey(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress {
	s.TagKey = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) SetTagValue(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress {
	s.TagValue = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) SetTagValueMatchFunction(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress {
	s.TagValueMatchFunction = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) Validate() error {
	return dara.Validate(s)
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist struct {
	TagValueBlacklist []*string `json:"TagValueBlacklist,omitempty" xml:"TagValueBlacklist,omitempty" type:"Repeated"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist) GetTagValueBlacklist() []*string {
	return s.TagValueBlacklist
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist) SetTagValueBlacklist(v []*string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist {
	s.TagValueBlacklist = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTagValueBlacklist) Validate() error {
	return dara.Validate(s)
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList struct {
	TemplateIdList []*string `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) GetTemplateIdList() []*string {
	return s.TemplateIdList
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) SetTemplateIdList(v []*string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList {
	s.TemplateIdList = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) Validate() error {
	return dara.Validate(s)
}
