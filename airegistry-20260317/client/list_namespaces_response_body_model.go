// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody
	GetData() *ListNamespacesResponseBodyData
	SetRequestId(v string) *ListNamespacesResponseBody
	GetRequestId() *string
}

type ListNamespacesResponseBody struct {
	// The namespace query result.
	Data *ListNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D9E87E66-9EF0-5C10-A5E6-924020A0C9B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) GetData() *ListNamespacesResponseBodyData {
	return s.Data
}

func (s *ListNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespacesResponseBody) SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNamespacesResponseBodyData struct {
	// The returned data entries.
	Items []*ListNamespacesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyData) GetItems() []*ListNamespacesResponseBodyDataItems {
	return s.Items
}

func (s *ListNamespacesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNamespacesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespacesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNamespacesResponseBodyData) SetItems(v []*ListNamespacesResponseBodyDataItems) *ListNamespacesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListNamespacesResponseBodyData) SetPageNumber(v int32) *ListNamespacesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListNamespacesResponseBodyData) SetPageSize(v int32) *ListNamespacesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListNamespacesResponseBodyData) SetTotalCount(v int32) *ListNamespacesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListNamespacesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamespacesResponseBodyDataItems struct {
	// The time when the namespace was created.
	//
	// example:
	//
	// 2022-07-11T09:32:03+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// Default project, auto-created by EMR.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// magic:magic-cn-1us4sed5d01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// 76d4b6e4-31bf-475a-8710-6217ec049c1f
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The number of prompts in the namespace.
	//
	// example:
	//
	// 1
	PromptCount         *int32  `json:"PromptCount,omitempty" xml:"PromptCount,omitempty"`
	PublicAccessEnabled *bool   `json:"PublicAccessEnabled,omitempty" xml:"PublicAccessEnabled,omitempty"`
	PublicDomain        *string `json:"PublicDomain,omitempty" xml:"PublicDomain,omitempty"`
	// The number of skills in the namespace.
	//
	// example:
	//
	// 1
	SkillCount *int32 `json:"SkillCount,omitempty" xml:"SkillCount,omitempty"`
	// The source of the namespace.
	//
	// example:
	//
	// magic:magic-cn-fpi4secsq01
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source index of the namespace.
	//
	// example:
	//
	// 0
	SourceIndex *int32 `json:"SourceIndex,omitempty" xml:"SourceIndex,omitempty"`
	// The tags of the namespace.
	//
	// example:
	//
	// qa,test
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListNamespacesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyDataItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListNamespacesResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListNamespacesResponseBodyDataItems) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *ListNamespacesResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListNamespacesResponseBodyDataItems) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNamespacesResponseBodyDataItems) GetPromptCount() *int32 {
	return s.PromptCount
}

func (s *ListNamespacesResponseBodyDataItems) GetPublicAccessEnabled() *bool {
	return s.PublicAccessEnabled
}

func (s *ListNamespacesResponseBodyDataItems) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *ListNamespacesResponseBodyDataItems) GetSkillCount() *int32 {
	return s.SkillCount
}

func (s *ListNamespacesResponseBodyDataItems) GetSource() *string {
	return s.Source
}

func (s *ListNamespacesResponseBodyDataItems) GetSourceIndex() *int32 {
	return s.SourceIndex
}

func (s *ListNamespacesResponseBodyDataItems) GetTags() *string {
	return s.Tags
}

func (s *ListNamespacesResponseBodyDataItems) SetCreatedTime(v string) *ListNamespacesResponseBodyDataItems {
	s.CreatedTime = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetDescription(v string) *ListNamespacesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetIpWhitelist(v string) *ListNamespacesResponseBodyDataItems {
	s.IpWhitelist = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetName(v string) *ListNamespacesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetNamespaceId(v string) *ListNamespacesResponseBodyDataItems {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetPromptCount(v int32) *ListNamespacesResponseBodyDataItems {
	s.PromptCount = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetPublicAccessEnabled(v bool) *ListNamespacesResponseBodyDataItems {
	s.PublicAccessEnabled = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetPublicDomain(v string) *ListNamespacesResponseBodyDataItems {
	s.PublicDomain = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetSkillCount(v int32) *ListNamespacesResponseBodyDataItems {
	s.SkillCount = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetSource(v string) *ListNamespacesResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetSourceIndex(v int32) *ListNamespacesResponseBodyDataItems {
	s.SourceIndex = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetTags(v string) *ListNamespacesResponseBodyDataItems {
	s.Tags = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
