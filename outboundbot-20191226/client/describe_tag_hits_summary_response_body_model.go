// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagHitsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTagHitsSummaryResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeTagHitsSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeTagHitsSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTagHitsSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTagHitsSummaryResponseBody
	GetSuccess() *bool
	SetTagGroups(v []*DescribeTagHitsSummaryResponseBodyTagGroups) *DescribeTagHitsSummaryResponseBody
	GetTagGroups() []*DescribeTagHitsSummaryResponseBodyTagGroups
	SetTagHitsList(v []*DescribeTagHitsSummaryResponseBodyTagHitsList) *DescribeTagHitsSummaryResponseBody
	GetTagHitsList() []*DescribeTagHitsSummaryResponseBodyTagHitsList
}

type DescribeTagHitsSummaryResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5391EB13-A0E7-402D-A407-B99D4ABAF22A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success     *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TagGroups   []*DescribeTagHitsSummaryResponseBodyTagGroups   `json:"TagGroups,omitempty" xml:"TagGroups,omitempty" type:"Repeated"`
	TagHitsList []*DescribeTagHitsSummaryResponseBodyTagHitsList `json:"TagHitsList,omitempty" xml:"TagHitsList,omitempty" type:"Repeated"`
}

func (s DescribeTagHitsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagHitsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagHitsSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTagHitsSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeTagHitsSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTagHitsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagHitsSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTagHitsSummaryResponseBody) GetTagGroups() []*DescribeTagHitsSummaryResponseBodyTagGroups {
	return s.TagGroups
}

func (s *DescribeTagHitsSummaryResponseBody) GetTagHitsList() []*DescribeTagHitsSummaryResponseBodyTagHitsList {
	return s.TagHitsList
}

func (s *DescribeTagHitsSummaryResponseBody) SetCode(v string) *DescribeTagHitsSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBody) SetHttpStatusCode(v int32) *DescribeTagHitsSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBody) SetMessage(v string) *DescribeTagHitsSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBody) SetRequestId(v string) *DescribeTagHitsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBody) SetSuccess(v bool) *DescribeTagHitsSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBody) SetTagGroups(v []*DescribeTagHitsSummaryResponseBodyTagGroups) *DescribeTagHitsSummaryResponseBody {
	s.TagGroups = v
	return s
}

func (s *DescribeTagHitsSummaryResponseBody) SetTagHitsList(v []*DescribeTagHitsSummaryResponseBodyTagHitsList) *DescribeTagHitsSummaryResponseBody {
	s.TagHitsList = v
	return s
}

func (s *DescribeTagHitsSummaryResponseBody) Validate() error {
	if s.TagGroups != nil {
		for _, item := range s.TagGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagHitsList != nil {
		for _, item := range s.TagHitsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagHitsSummaryResponseBodyTagGroups struct {
	// ID
	//
	// example:
	//
	// 8bb6f8ca-85a3-49f8-86a5-3127902a2156
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 84fc7c41-f918-4a47-b742-a439b35a8567
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	// example:
	//
	// 0
	TagGroupIndex *int32 `json:"TagGroupIndex,omitempty" xml:"TagGroupIndex,omitempty"`
}

func (s DescribeTagHitsSummaryResponseBodyTagGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagHitsSummaryResponseBodyTagGroups) GoString() string {
	return s.String()
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) GetId() *string {
	return s.Id
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) GetTagGroup() *string {
	return s.TagGroup
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) GetTagGroupIndex() *int32 {
	return s.TagGroupIndex
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) SetId(v string) *DescribeTagHitsSummaryResponseBodyTagGroups {
	s.Id = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) SetScriptId(v string) *DescribeTagHitsSummaryResponseBodyTagGroups {
	s.ScriptId = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) SetTagGroup(v string) *DescribeTagHitsSummaryResponseBodyTagGroups {
	s.TagGroup = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) SetTagGroupIndex(v int32) *DescribeTagHitsSummaryResponseBodyTagGroups {
	s.TagGroupIndex = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBodyTagGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeTagHitsSummaryResponseBodyTagHitsList struct {
	// example:
	//
	// 1
	HitCount *int32  `json:"HitCount,omitempty" xml:"HitCount,omitempty"`
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	TagName  *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeTagHitsSummaryResponseBodyTagHitsList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagHitsSummaryResponseBodyTagHitsList) GoString() string {
	return s.String()
}

func (s *DescribeTagHitsSummaryResponseBodyTagHitsList) GetHitCount() *int32 {
	return s.HitCount
}

func (s *DescribeTagHitsSummaryResponseBodyTagHitsList) GetTagGroup() *string {
	return s.TagGroup
}

func (s *DescribeTagHitsSummaryResponseBodyTagHitsList) GetTagName() *string {
	return s.TagName
}

func (s *DescribeTagHitsSummaryResponseBodyTagHitsList) SetHitCount(v int32) *DescribeTagHitsSummaryResponseBodyTagHitsList {
	s.HitCount = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBodyTagHitsList) SetTagGroup(v string) *DescribeTagHitsSummaryResponseBodyTagHitsList {
	s.TagGroup = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBodyTagHitsList) SetTagName(v string) *DescribeTagHitsSummaryResponseBodyTagHitsList {
	s.TagName = &v
	return s
}

func (s *DescribeTagHitsSummaryResponseBodyTagHitsList) Validate() error {
	return dara.Validate(s)
}
