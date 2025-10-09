// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoProcessingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetVideoProcessingResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetVideoProcessingResponseBody
	GetConfigType() *string
	SetFlvSeekEnd(v string) *GetVideoProcessingResponseBody
	GetFlvSeekEnd() *string
	SetFlvSeekStart(v string) *GetVideoProcessingResponseBody
	GetFlvSeekStart() *string
	SetFlvVideoSeekMode(v string) *GetVideoProcessingResponseBody
	GetFlvVideoSeekMode() *string
	SetMp4SeekEnd(v string) *GetVideoProcessingResponseBody
	GetMp4SeekEnd() *string
	SetMp4SeekStart(v string) *GetVideoProcessingResponseBody
	GetMp4SeekStart() *string
	SetRequestId(v string) *GetVideoProcessingResponseBody
	GetRequestId() *string
	SetRule(v string) *GetVideoProcessingResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetVideoProcessingResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetVideoProcessingResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetVideoProcessingResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetVideoProcessingResponseBody
	GetSiteVersion() *int32
	SetVideoSeekEnable(v string) *GetVideoProcessingResponseBody
	GetVideoSeekEnable() *string
}

type GetVideoProcessingResponseBody struct {
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The type of the configuration. Valid values:
	//
	// 	- global: global configuration.
	//
	// 	- rule: rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The custom end parameter for FLV files.
	//
	// example:
	//
	// end
	FlvSeekEnd *string `json:"FlvSeekEnd,omitempty" xml:"FlvSeekEnd,omitempty"`
	// The custom start parameter for FLV files.
	//
	// example:
	//
	// start
	FlvSeekStart *string `json:"FlvSeekStart,omitempty" xml:"FlvSeekStart,omitempty"`
	// FLV Seeking Valid values:
	//
	// 	- by_byte: Seek by byte.
	//
	// 	- by_time: Seek by time.
	//
	// example:
	//
	// by_byte
	FlvVideoSeekMode *string `json:"FlvVideoSeekMode,omitempty" xml:"FlvVideoSeekMode,omitempty"`
	// Customize the mp4 end parameter.
	//
	// example:
	//
	// end
	Mp4SeekEnd *string `json:"Mp4SeekEnd,omitempty" xml:"Mp4SeekEnd,omitempty"`
	// Customize the mp4 start parameter.
	//
	// example:
	//
	// start
	Mp4SeekStart *string `json:"Mp4SeekStart,omitempty" xml:"Mp4SeekStart,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Indicates whether the rule is enabled. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the scaling rule.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The order in which the rule is executed. A smaller value gives priority to the rule.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the website configurations.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Video seeking. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	VideoSeekEnable *string `json:"VideoSeekEnable,omitempty" xml:"VideoSeekEnable,omitempty"`
}

func (s GetVideoProcessingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoProcessingResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoProcessingResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetVideoProcessingResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetVideoProcessingResponseBody) GetFlvSeekEnd() *string {
	return s.FlvSeekEnd
}

func (s *GetVideoProcessingResponseBody) GetFlvSeekStart() *string {
	return s.FlvSeekStart
}

func (s *GetVideoProcessingResponseBody) GetFlvVideoSeekMode() *string {
	return s.FlvVideoSeekMode
}

func (s *GetVideoProcessingResponseBody) GetMp4SeekEnd() *string {
	return s.Mp4SeekEnd
}

func (s *GetVideoProcessingResponseBody) GetMp4SeekStart() *string {
	return s.Mp4SeekStart
}

func (s *GetVideoProcessingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoProcessingResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetVideoProcessingResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetVideoProcessingResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetVideoProcessingResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetVideoProcessingResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetVideoProcessingResponseBody) GetVideoSeekEnable() *string {
	return s.VideoSeekEnable
}

func (s *GetVideoProcessingResponseBody) SetConfigId(v int64) *GetVideoProcessingResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetConfigType(v string) *GetVideoProcessingResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetFlvSeekEnd(v string) *GetVideoProcessingResponseBody {
	s.FlvSeekEnd = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetFlvSeekStart(v string) *GetVideoProcessingResponseBody {
	s.FlvSeekStart = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetFlvVideoSeekMode(v string) *GetVideoProcessingResponseBody {
	s.FlvVideoSeekMode = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetMp4SeekEnd(v string) *GetVideoProcessingResponseBody {
	s.Mp4SeekEnd = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetMp4SeekStart(v string) *GetVideoProcessingResponseBody {
	s.Mp4SeekStart = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetRequestId(v string) *GetVideoProcessingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetRule(v string) *GetVideoProcessingResponseBody {
	s.Rule = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetRuleEnable(v string) *GetVideoProcessingResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetRuleName(v string) *GetVideoProcessingResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetSequence(v int32) *GetVideoProcessingResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetSiteVersion(v int32) *GetVideoProcessingResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetVideoProcessingResponseBody) SetVideoSeekEnable(v string) *GetVideoProcessingResponseBody {
	s.VideoSeekEnable = &v
	return s
}

func (s *GetVideoProcessingResponseBody) Validate() error {
	return dara.Validate(s)
}
