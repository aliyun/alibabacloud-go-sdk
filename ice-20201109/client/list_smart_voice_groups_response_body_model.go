// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartVoiceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSmartVoiceGroupsResponseBody
	GetRequestId() *string
	SetVoiceGroups(v []*ListSmartVoiceGroupsResponseBodyVoiceGroups) *ListSmartVoiceGroupsResponseBody
	GetVoiceGroups() []*ListSmartVoiceGroupsResponseBodyVoiceGroups
}

type ListSmartVoiceGroupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 627B30EB-1D0A-5C6D-8467-431626E0FA10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried speaker groups.
	VoiceGroups []*ListSmartVoiceGroupsResponseBodyVoiceGroups `json:"VoiceGroups,omitempty" xml:"VoiceGroups,omitempty" type:"Repeated"`
}

func (s ListSmartVoiceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSmartVoiceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmartVoiceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSmartVoiceGroupsResponseBody) GetVoiceGroups() []*ListSmartVoiceGroupsResponseBodyVoiceGroups {
	return s.VoiceGroups
}

func (s *ListSmartVoiceGroupsResponseBody) SetRequestId(v string) *ListSmartVoiceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBody) SetVoiceGroups(v []*ListSmartVoiceGroupsResponseBodyVoiceGroups) *ListSmartVoiceGroupsResponseBody {
	s.VoiceGroups = v
	return s
}

func (s *ListSmartVoiceGroupsResponseBody) Validate() error {
	if s.VoiceGroups != nil {
		for _, item := range s.VoiceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSmartVoiceGroupsResponseBodyVoiceGroups struct {
	// The name of the speaker group.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The speakers.
	VoiceList []*ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList `json:"VoiceList,omitempty" xml:"VoiceList,omitempty" type:"Repeated"`
}

func (s ListSmartVoiceGroupsResponseBodyVoiceGroups) String() string {
	return dara.Prettify(s)
}

func (s ListSmartVoiceGroupsResponseBodyVoiceGroups) GoString() string {
	return s.String()
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroups) GetType() *string {
	return s.Type
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroups) GetVoiceList() []*ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	return s.VoiceList
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroups) SetType(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroups {
	s.Type = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroups) SetVoiceList(v []*ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) *ListSmartVoiceGroupsResponseBodyVoiceGroups {
	s.VoiceList = v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroups) Validate() error {
	if s.VoiceList != nil {
		for _, item := range s.VoiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList struct {
	// The speaker description.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The speaker name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks of the speaker.
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SupportSampleRate *string `json:"SupportSampleRate,omitempty" xml:"SupportSampleRate,omitempty"`
	// The tag of the speaker type.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The speaker ID.
	//
	// example:
	//
	// zhitian
	Voice       *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	VoiceSource *string `json:"VoiceSource,omitempty" xml:"VoiceSource,omitempty"`
	// The speaker type.
	//
	// Valid values:
	//
	// 	- Male
	//
	// 	- Female
	//
	// 	- Boy
	//
	// 	- Girl
	//
	// example:
	//
	// Female
	VoiceType *string `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
	// The URL of the sample audio file.
	//
	// example:
	//
	// https://***.com/zhiqing.mp3
	VoiceUrl *string `json:"VoiceUrl,omitempty" xml:"VoiceUrl,omitempty"`
}

func (s ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) String() string {
	return dara.Prettify(s)
}

func (s ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GoString() string {
	return s.String()
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetDesc() *string {
	return s.Desc
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetName() *string {
	return s.Name
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetRemark() *string {
	return s.Remark
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetSupportSampleRate() *string {
	return s.SupportSampleRate
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetTag() *string {
	return s.Tag
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetVoice() *string {
	return s.Voice
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetVoiceSource() *string {
	return s.VoiceSource
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetVoiceType() *string {
	return s.VoiceType
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) GetVoiceUrl() *string {
	return s.VoiceUrl
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetDesc(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.Desc = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetName(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.Name = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetRemark(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.Remark = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetSupportSampleRate(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.SupportSampleRate = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetTag(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.Tag = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetVoice(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.Voice = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetVoiceSource(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.VoiceSource = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetVoiceType(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.VoiceType = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) SetVoiceUrl(v string) *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList {
	s.VoiceUrl = &v
	return s
}

func (s *ListSmartVoiceGroupsResponseBodyVoiceGroupsVoiceList) Validate() error {
	return dara.Validate(s)
}
