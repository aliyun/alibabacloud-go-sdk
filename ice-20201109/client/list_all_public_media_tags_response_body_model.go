// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPublicMediaTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaTagList(v []*ListAllPublicMediaTagsResponseBodyMediaTagList) *ListAllPublicMediaTagsResponseBody
	GetMediaTagList() []*ListAllPublicMediaTagsResponseBodyMediaTagList
	SetRequestId(v string) *ListAllPublicMediaTagsResponseBody
	GetRequestId() *string
}

type ListAllPublicMediaTagsResponseBody struct {
	// The tags of media assets in the public media library.
	MediaTagList []*ListAllPublicMediaTagsResponseBodyMediaTagList `json:"MediaTagList,omitempty" xml:"MediaTagList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B45F83B7-7F87-4792-BFE9-63CD2137CAF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAllPublicMediaTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllPublicMediaTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponseBody) GetMediaTagList() []*ListAllPublicMediaTagsResponseBodyMediaTagList {
	return s.MediaTagList
}

func (s *ListAllPublicMediaTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllPublicMediaTagsResponseBody) SetMediaTagList(v []*ListAllPublicMediaTagsResponseBodyMediaTagList) *ListAllPublicMediaTagsResponseBody {
	s.MediaTagList = v
	return s
}

func (s *ListAllPublicMediaTagsResponseBody) SetRequestId(v string) *ListAllPublicMediaTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBody) Validate() error {
	if s.MediaTagList != nil {
		for _, item := range s.MediaTagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllPublicMediaTagsResponseBodyMediaTagList struct {
	// The ID of the media tag.
	//
	// example:
	//
	// sticker-gif
	MediaTagId *string `json:"MediaTagId,omitempty" xml:"MediaTagId,omitempty"`
	// The name of the media tag in Chinese.
	//
	// example:
	//
	// Gif
	MediaTagNameChinese *string `json:"MediaTagNameChinese,omitempty" xml:"MediaTagNameChinese,omitempty"`
	// The name of the material tag in English.
	MediaTagNameEnglish *string `json:"MediaTagNameEnglish,omitempty" xml:"MediaTagNameEnglish,omitempty"`
	// The options.
	Options []*ListAllPublicMediaTagsResponseBodyMediaTagListOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
}

func (s ListAllPublicMediaTagsResponseBodyMediaTagList) String() string {
	return dara.Prettify(s)
}

func (s ListAllPublicMediaTagsResponseBodyMediaTagList) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) GetMediaTagId() *string {
	return s.MediaTagId
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) GetMediaTagNameChinese() *string {
	return s.MediaTagNameChinese
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) GetMediaTagNameEnglish() *string {
	return s.MediaTagNameEnglish
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) GetOptions() []*ListAllPublicMediaTagsResponseBodyMediaTagListOptions {
	return s.Options
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) SetMediaTagId(v string) *ListAllPublicMediaTagsResponseBodyMediaTagList {
	s.MediaTagId = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) SetMediaTagNameChinese(v string) *ListAllPublicMediaTagsResponseBodyMediaTagList {
	s.MediaTagNameChinese = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) SetMediaTagNameEnglish(v string) *ListAllPublicMediaTagsResponseBodyMediaTagList {
	s.MediaTagNameEnglish = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) SetOptions(v []*ListAllPublicMediaTagsResponseBodyMediaTagListOptions) *ListAllPublicMediaTagsResponseBodyMediaTagList {
	s.Options = v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) Validate() error {
	if s.Options != nil {
		for _, item := range s.Options {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllPublicMediaTagsResponseBodyMediaTagListOptions struct {
	// The option name in Chinese.
	OptionChineseName *string `json:"OptionChineseName,omitempty" xml:"OptionChineseName,omitempty"`
	// The option name in English.
	//
	// example:
	//
	// Angry
	OptionEnglishName *string `json:"OptionEnglishName,omitempty" xml:"OptionEnglishName,omitempty"`
	// The option ID.
	//
	// example:
	//
	// Angry
	OptionId *string `json:"OptionId,omitempty" xml:"OptionId,omitempty"`
}

func (s ListAllPublicMediaTagsResponseBodyMediaTagListOptions) String() string {
	return dara.Prettify(s)
}

func (s ListAllPublicMediaTagsResponseBodyMediaTagListOptions) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagListOptions) GetOptionChineseName() *string {
	return s.OptionChineseName
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagListOptions) GetOptionEnglishName() *string {
	return s.OptionEnglishName
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagListOptions) GetOptionId() *string {
	return s.OptionId
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagListOptions) SetOptionChineseName(v string) *ListAllPublicMediaTagsResponseBodyMediaTagListOptions {
	s.OptionChineseName = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagListOptions) SetOptionEnglishName(v string) *ListAllPublicMediaTagsResponseBodyMediaTagListOptions {
	s.OptionEnglishName = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagListOptions) SetOptionId(v string) *ListAllPublicMediaTagsResponseBodyMediaTagListOptions {
	s.OptionId = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagListOptions) Validate() error {
	return dara.Validate(s)
}
