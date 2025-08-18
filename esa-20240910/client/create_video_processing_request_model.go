// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoProcessingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlvSeekEnd(v string) *CreateVideoProcessingRequest
	GetFlvSeekEnd() *string
	SetFlvSeekStart(v string) *CreateVideoProcessingRequest
	GetFlvSeekStart() *string
	SetFlvVideoSeekMode(v string) *CreateVideoProcessingRequest
	GetFlvVideoSeekMode() *string
	SetMp4SeekEnd(v string) *CreateVideoProcessingRequest
	GetMp4SeekEnd() *string
	SetMp4SeekStart(v string) *CreateVideoProcessingRequest
	GetMp4SeekStart() *string
	SetRule(v string) *CreateVideoProcessingRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateVideoProcessingRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateVideoProcessingRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateVideoProcessingRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateVideoProcessingRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateVideoProcessingRequest
	GetSiteVersion() *int32
	SetVideoSeekEnable(v string) *CreateVideoProcessingRequest
	GetVideoSeekEnable() *string
}

type CreateVideoProcessingRequest struct {
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
	// (http.host eq \\"video.example.com\\")
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
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// example:
	//
	// on
	VideoSeekEnable *string `json:"VideoSeekEnable,omitempty" xml:"VideoSeekEnable,omitempty"`
}

func (s CreateVideoProcessingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoProcessingRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoProcessingRequest) GetFlvSeekEnd() *string {
	return s.FlvSeekEnd
}

func (s *CreateVideoProcessingRequest) GetFlvSeekStart() *string {
	return s.FlvSeekStart
}

func (s *CreateVideoProcessingRequest) GetFlvVideoSeekMode() *string {
	return s.FlvVideoSeekMode
}

func (s *CreateVideoProcessingRequest) GetMp4SeekEnd() *string {
	return s.Mp4SeekEnd
}

func (s *CreateVideoProcessingRequest) GetMp4SeekStart() *string {
	return s.Mp4SeekStart
}

func (s *CreateVideoProcessingRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateVideoProcessingRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateVideoProcessingRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateVideoProcessingRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateVideoProcessingRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateVideoProcessingRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateVideoProcessingRequest) GetVideoSeekEnable() *string {
	return s.VideoSeekEnable
}

func (s *CreateVideoProcessingRequest) SetFlvSeekEnd(v string) *CreateVideoProcessingRequest {
	s.FlvSeekEnd = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetFlvSeekStart(v string) *CreateVideoProcessingRequest {
	s.FlvSeekStart = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetFlvVideoSeekMode(v string) *CreateVideoProcessingRequest {
	s.FlvVideoSeekMode = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetMp4SeekEnd(v string) *CreateVideoProcessingRequest {
	s.Mp4SeekEnd = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetMp4SeekStart(v string) *CreateVideoProcessingRequest {
	s.Mp4SeekStart = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetRule(v string) *CreateVideoProcessingRequest {
	s.Rule = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetRuleEnable(v string) *CreateVideoProcessingRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetRuleName(v string) *CreateVideoProcessingRequest {
	s.RuleName = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetSequence(v int32) *CreateVideoProcessingRequest {
	s.Sequence = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetSiteId(v int64) *CreateVideoProcessingRequest {
	s.SiteId = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetSiteVersion(v int32) *CreateVideoProcessingRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateVideoProcessingRequest) SetVideoSeekEnable(v string) *CreateVideoProcessingRequest {
	s.VideoSeekEnable = &v
	return s
}

func (s *CreateVideoProcessingRequest) Validate() error {
	return dara.Validate(s)
}
