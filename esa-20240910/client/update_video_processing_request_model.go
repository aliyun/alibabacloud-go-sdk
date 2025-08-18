// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoProcessingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateVideoProcessingRequest
	GetConfigId() *int64
	SetFlvSeekEnd(v string) *UpdateVideoProcessingRequest
	GetFlvSeekEnd() *string
	SetFlvSeekStart(v string) *UpdateVideoProcessingRequest
	GetFlvSeekStart() *string
	SetFlvVideoSeekMode(v string) *UpdateVideoProcessingRequest
	GetFlvVideoSeekMode() *string
	SetMp4SeekEnd(v string) *UpdateVideoProcessingRequest
	GetMp4SeekEnd() *string
	SetMp4SeekStart(v string) *UpdateVideoProcessingRequest
	GetMp4SeekStart() *string
	SetRule(v string) *UpdateVideoProcessingRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateVideoProcessingRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateVideoProcessingRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateVideoProcessingRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateVideoProcessingRequest
	GetSiteId() *int64
	SetVideoSeekEnable(v string) *UpdateVideoProcessingRequest
	GetVideoSeekEnable() *string
}

type UpdateVideoProcessingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// end
	FlvSeekEnd *string `json:"FlvSeekEnd,omitempty" xml:"FlvSeekEnd,omitempty"`
	// example:
	//
	// start
	FlvSeekStart *string `json:"FlvSeekStart,omitempty" xml:"FlvSeekStart,omitempty"`
	// example:
	//
	// by_byte
	FlvVideoSeekMode *string `json:"FlvVideoSeekMode,omitempty" xml:"FlvVideoSeekMode,omitempty"`
	// example:
	//
	// end
	Mp4SeekEnd *string `json:"Mp4SeekEnd,omitempty" xml:"Mp4SeekEnd,omitempty"`
	// example:
	//
	// start
	Mp4SeekStart *string `json:"Mp4SeekStart,omitempty" xml:"Mp4SeekStart,omitempty"`
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// on
	VideoSeekEnable *string `json:"VideoSeekEnable,omitempty" xml:"VideoSeekEnable,omitempty"`
}

func (s UpdateVideoProcessingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoProcessingRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoProcessingRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateVideoProcessingRequest) GetFlvSeekEnd() *string {
	return s.FlvSeekEnd
}

func (s *UpdateVideoProcessingRequest) GetFlvSeekStart() *string {
	return s.FlvSeekStart
}

func (s *UpdateVideoProcessingRequest) GetFlvVideoSeekMode() *string {
	return s.FlvVideoSeekMode
}

func (s *UpdateVideoProcessingRequest) GetMp4SeekEnd() *string {
	return s.Mp4SeekEnd
}

func (s *UpdateVideoProcessingRequest) GetMp4SeekStart() *string {
	return s.Mp4SeekStart
}

func (s *UpdateVideoProcessingRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateVideoProcessingRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateVideoProcessingRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateVideoProcessingRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateVideoProcessingRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateVideoProcessingRequest) GetVideoSeekEnable() *string {
	return s.VideoSeekEnable
}

func (s *UpdateVideoProcessingRequest) SetConfigId(v int64) *UpdateVideoProcessingRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetFlvSeekEnd(v string) *UpdateVideoProcessingRequest {
	s.FlvSeekEnd = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetFlvSeekStart(v string) *UpdateVideoProcessingRequest {
	s.FlvSeekStart = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetFlvVideoSeekMode(v string) *UpdateVideoProcessingRequest {
	s.FlvVideoSeekMode = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetMp4SeekEnd(v string) *UpdateVideoProcessingRequest {
	s.Mp4SeekEnd = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetMp4SeekStart(v string) *UpdateVideoProcessingRequest {
	s.Mp4SeekStart = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetRule(v string) *UpdateVideoProcessingRequest {
	s.Rule = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetRuleEnable(v string) *UpdateVideoProcessingRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetRuleName(v string) *UpdateVideoProcessingRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetSequence(v int32) *UpdateVideoProcessingRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetSiteId(v int64) *UpdateVideoProcessingRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateVideoProcessingRequest) SetVideoSeekEnable(v string) *UpdateVideoProcessingRequest {
	s.VideoSeekEnable = &v
	return s
}

func (s *UpdateVideoProcessingRequest) Validate() error {
	return dara.Validate(s)
}
