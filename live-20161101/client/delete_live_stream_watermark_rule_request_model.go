// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamWatermarkRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DeleteLiveStreamWatermarkRuleRequest
	GetApp() *string
	SetDomain(v string) *DeleteLiveStreamWatermarkRuleRequest
	GetDomain() *string
	SetOwnerId(v int64) *DeleteLiveStreamWatermarkRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveStreamWatermarkRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *DeleteLiveStreamWatermarkRuleRequest
	GetRuleId() *string
	SetStream(v string) *DeleteLiveStreamWatermarkRuleRequest
	GetStream() *string
}

type DeleteLiveStreamWatermarkRuleRequest struct {
	// The AppName of the live stream.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the watermark rule.
	//
	// > Get this ID from the response of the [AddLiveStreamWatermarkRule](https://help.aliyun.com/document_detail/2848100.html) operation.
	//
	// example:
	//
	// 445409ec-7eaa-461d -8f29-4bec2eb9****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The stream name. The following rules apply:
	//
	// - To match a specific stream, enter the full stream name. Example: liveStreamA.
	//
	// - Use a wildcard for matching. The asterisk (\\*) matches all streams.
	//
	// - You can match by prefix or suffix.
	//
	// > 	- For wildcard matching, use only one asterisk (\\*) at the beginning or end of the string. Enclose matching items in parentheses. Separate multiple matching items with a vertical bar (|).
	//
	// >
	//
	// > 	- Example: `*(t1|t2)` matches all streams ending with `t1` or `t2`. `(abc|123)*` matches all streams starting with `abc` or `123`.
	//
	// example:
	//
	// liveStreamA
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s DeleteLiveStreamWatermarkRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamWatermarkRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamWatermarkRuleRequest) GetApp() *string {
	return s.App
}

func (s *DeleteLiveStreamWatermarkRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteLiveStreamWatermarkRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveStreamWatermarkRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveStreamWatermarkRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteLiveStreamWatermarkRuleRequest) GetStream() *string {
	return s.Stream
}

func (s *DeleteLiveStreamWatermarkRuleRequest) SetApp(v string) *DeleteLiveStreamWatermarkRuleRequest {
	s.App = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleRequest) SetDomain(v string) *DeleteLiveStreamWatermarkRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleRequest) SetOwnerId(v int64) *DeleteLiveStreamWatermarkRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleRequest) SetRegionId(v string) *DeleteLiveStreamWatermarkRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleRequest) SetRuleId(v string) *DeleteLiveStreamWatermarkRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleRequest) SetStream(v string) *DeleteLiveStreamWatermarkRuleRequest {
	s.Stream = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleRequest) Validate() error {
	return dara.Validate(s)
}
