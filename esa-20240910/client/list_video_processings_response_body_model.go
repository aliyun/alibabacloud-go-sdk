// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVideoProcessingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListVideoProcessingsResponseBodyConfigs) *ListVideoProcessingsResponseBody
	GetConfigs() []*ListVideoProcessingsResponseBodyConfigs
	SetPageNumber(v int32) *ListVideoProcessingsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVideoProcessingsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListVideoProcessingsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVideoProcessingsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListVideoProcessingsResponseBody
	GetTotalPage() *int32
}

type ListVideoProcessingsResponseBody struct {
	Configs    []*ListVideoProcessingsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                     `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListVideoProcessingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVideoProcessingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVideoProcessingsResponseBody) GetConfigs() []*ListVideoProcessingsResponseBodyConfigs {
	return s.Configs
}

func (s *ListVideoProcessingsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVideoProcessingsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVideoProcessingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVideoProcessingsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVideoProcessingsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListVideoProcessingsResponseBody) SetConfigs(v []*ListVideoProcessingsResponseBodyConfigs) *ListVideoProcessingsResponseBody {
	s.Configs = v
	return s
}

func (s *ListVideoProcessingsResponseBody) SetPageNumber(v int32) *ListVideoProcessingsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVideoProcessingsResponseBody) SetPageSize(v int32) *ListVideoProcessingsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVideoProcessingsResponseBody) SetRequestId(v string) *ListVideoProcessingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVideoProcessingsResponseBody) SetTotalCount(v int32) *ListVideoProcessingsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVideoProcessingsResponseBody) SetTotalPage(v int32) *ListVideoProcessingsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListVideoProcessingsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVideoProcessingsResponseBodyConfigs struct {
	ConfigId         *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType       *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	FlvSeekEnd       *string `json:"FlvSeekEnd,omitempty" xml:"FlvSeekEnd,omitempty"`
	FlvSeekStart     *string `json:"FlvSeekStart,omitempty" xml:"FlvSeekStart,omitempty"`
	FlvVideoSeekMode *string `json:"FlvVideoSeekMode,omitempty" xml:"FlvVideoSeekMode,omitempty"`
	Mp4SeekEnd       *string `json:"Mp4SeekEnd,omitempty" xml:"Mp4SeekEnd,omitempty"`
	Mp4SeekStart     *string `json:"Mp4SeekStart,omitempty" xml:"Mp4SeekStart,omitempty"`
	Rule             *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable       *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName         *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence         *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion      *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	VideoSeekEnable  *string `json:"VideoSeekEnable,omitempty" xml:"VideoSeekEnable,omitempty"`
}

func (s ListVideoProcessingsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListVideoProcessingsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetFlvSeekEnd() *string {
	return s.FlvSeekEnd
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetFlvSeekStart() *string {
	return s.FlvSeekStart
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetFlvVideoSeekMode() *string {
	return s.FlvVideoSeekMode
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetMp4SeekEnd() *string {
	return s.Mp4SeekEnd
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetMp4SeekStart() *string {
	return s.Mp4SeekStart
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListVideoProcessingsResponseBodyConfigs) GetVideoSeekEnable() *string {
	return s.VideoSeekEnable
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetConfigId(v int64) *ListVideoProcessingsResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetConfigType(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetFlvSeekEnd(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.FlvSeekEnd = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetFlvSeekStart(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.FlvSeekStart = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetFlvVideoSeekMode(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.FlvVideoSeekMode = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetMp4SeekEnd(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.Mp4SeekEnd = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetMp4SeekStart(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.Mp4SeekStart = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetRule(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetRuleEnable(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetRuleName(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetSequence(v int32) *ListVideoProcessingsResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetSiteVersion(v int32) *ListVideoProcessingsResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) SetVideoSeekEnable(v string) *ListVideoProcessingsResponseBodyConfigs {
	s.VideoSeekEnable = &v
	return s
}

func (s *ListVideoProcessingsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
