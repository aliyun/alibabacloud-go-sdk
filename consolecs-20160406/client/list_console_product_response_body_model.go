// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsoleProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsoleProductResponseBody
	GetCode() *string
	SetData(v []*ListConsoleProductResponseBodyData) *ListConsoleProductResponseBody
	GetData() []*ListConsoleProductResponseBodyData
	SetRequestId(v string) *ListConsoleProductResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConsoleProductResponseBody
	GetSuccess() *bool
}

type ListConsoleProductResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListConsoleProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListConsoleProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsoleProductResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsoleProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsoleProductResponseBody) GetData() []*ListConsoleProductResponseBodyData {
	return s.Data
}

func (s *ListConsoleProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsoleProductResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConsoleProductResponseBody) SetCode(v string) *ListConsoleProductResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsoleProductResponseBody) SetData(v []*ListConsoleProductResponseBodyData) *ListConsoleProductResponseBody {
	s.Data = v
	return s
}

func (s *ListConsoleProductResponseBody) SetRequestId(v string) *ListConsoleProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsoleProductResponseBody) SetSuccess(v bool) *ListConsoleProductResponseBody {
	s.Success = &v
	return s
}

func (s *ListConsoleProductResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConsoleProductResponseBodyData struct {
	BelongedCategory    *string   `json:"BelongedCategory,omitempty" xml:"BelongedCategory,omitempty"`
	Categories          []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	ChannelLinks        []*string `json:"ChannelLinks,omitempty" xml:"ChannelLinks,omitempty" type:"Repeated"`
	DocId               *string   `json:"DocId,omitempty" xml:"DocId,omitempty"`
	Keywords            []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Names               *string   `json:"Names,omitempty" xml:"Names,omitempty"`
	Pinyin              *string   `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	ProductId           *string   `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RelatedConsoleAppId *string   `json:"RelatedConsoleAppId,omitempty" xml:"RelatedConsoleAppId,omitempty"`
	RelatedPipId        *string   `json:"RelatedPipId,omitempty" xml:"RelatedPipId,omitempty"`
	ShowInNav           *bool     `json:"ShowInNav,omitempty" xml:"ShowInNav,omitempty"`
	SupportedAccounts   []*string `json:"SupportedAccounts,omitempty" xml:"SupportedAccounts,omitempty" type:"Repeated"`
	SupportedChannels   []*string `json:"SupportedChannels,omitempty" xml:"SupportedChannels,omitempty" type:"Repeated"`
	Tag                 *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TagExpireTime       *string   `json:"TagExpireTime,omitempty" xml:"TagExpireTime,omitempty"`
}

func (s ListConsoleProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsoleProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsoleProductResponseBodyData) GetBelongedCategory() *string {
	return s.BelongedCategory
}

func (s *ListConsoleProductResponseBodyData) GetCategories() []*string {
	return s.Categories
}

func (s *ListConsoleProductResponseBodyData) GetChannelLinks() []*string {
	return s.ChannelLinks
}

func (s *ListConsoleProductResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *ListConsoleProductResponseBodyData) GetKeywords() []*string {
	return s.Keywords
}

func (s *ListConsoleProductResponseBodyData) GetNames() *string {
	return s.Names
}

func (s *ListConsoleProductResponseBodyData) GetPinyin() *string {
	return s.Pinyin
}

func (s *ListConsoleProductResponseBodyData) GetProductId() *string {
	return s.ProductId
}

func (s *ListConsoleProductResponseBodyData) GetRelatedConsoleAppId() *string {
	return s.RelatedConsoleAppId
}

func (s *ListConsoleProductResponseBodyData) GetRelatedPipId() *string {
	return s.RelatedPipId
}

func (s *ListConsoleProductResponseBodyData) GetShowInNav() *bool {
	return s.ShowInNav
}

func (s *ListConsoleProductResponseBodyData) GetSupportedAccounts() []*string {
	return s.SupportedAccounts
}

func (s *ListConsoleProductResponseBodyData) GetSupportedChannels() []*string {
	return s.SupportedChannels
}

func (s *ListConsoleProductResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ListConsoleProductResponseBodyData) GetTagExpireTime() *string {
	return s.TagExpireTime
}

func (s *ListConsoleProductResponseBodyData) SetBelongedCategory(v string) *ListConsoleProductResponseBodyData {
	s.BelongedCategory = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetCategories(v []*string) *ListConsoleProductResponseBodyData {
	s.Categories = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetChannelLinks(v []*string) *ListConsoleProductResponseBodyData {
	s.ChannelLinks = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetDocId(v string) *ListConsoleProductResponseBodyData {
	s.DocId = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetKeywords(v []*string) *ListConsoleProductResponseBodyData {
	s.Keywords = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetNames(v string) *ListConsoleProductResponseBodyData {
	s.Names = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetPinyin(v string) *ListConsoleProductResponseBodyData {
	s.Pinyin = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetProductId(v string) *ListConsoleProductResponseBodyData {
	s.ProductId = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetRelatedConsoleAppId(v string) *ListConsoleProductResponseBodyData {
	s.RelatedConsoleAppId = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetRelatedPipId(v string) *ListConsoleProductResponseBodyData {
	s.RelatedPipId = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetShowInNav(v bool) *ListConsoleProductResponseBodyData {
	s.ShowInNav = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetSupportedAccounts(v []*string) *ListConsoleProductResponseBodyData {
	s.SupportedAccounts = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetSupportedChannels(v []*string) *ListConsoleProductResponseBodyData {
	s.SupportedChannels = v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetTag(v string) *ListConsoleProductResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) SetTagExpireTime(v string) *ListConsoleProductResponseBodyData {
	s.TagExpireTime = &v
	return s
}

func (s *ListConsoleProductResponseBodyData) Validate() error {
	return dara.Validate(s)
}
