// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamWatermarkRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *AddLiveStreamWatermarkRuleRequest
	GetApp() *string
	SetDescription(v string) *AddLiveStreamWatermarkRuleRequest
	GetDescription() *string
	SetDomain(v string) *AddLiveStreamWatermarkRuleRequest
	GetDomain() *string
	SetName(v string) *AddLiveStreamWatermarkRuleRequest
	GetName() *string
	SetOwnerId(v int64) *AddLiveStreamWatermarkRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLiveStreamWatermarkRuleRequest
	GetRegionId() *string
	SetStream(v string) *AddLiveStreamWatermarkRuleRequest
	GetStream() *string
	SetTemplateId(v string) *AddLiveStreamWatermarkRuleRequest
	GetTemplateId() *string
}

type AddLiveStreamWatermarkRuleRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The description of the custom rule.
	//
	// example:
	//
	// my rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the custom rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// WatermarkRule****
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. The following rules apply:
	//
	// 	- A stream name can be exactly matched. Example: liveStreamA.
	//
	// 	- Fuzzy match is also supported. The use of an asterisk (`*`) allows all approximate matches to be found.
	//
	// 	- You can place the asterisk before or after an approximate string.
	//
	//
	//
	// >	- Fuzzy match: Only one asterisk (`*`) before or after an approximate string is allowed. The approximate string must be enclosed in `()`. Separate multiple strings with vertical bars (`|`).
	//
	// >	- For example, `*(t1|t2)` matches all streams whose name has the `t1` or `t2` suffix, and `(abc|123)*` matches all streams whose name has the `abc` or `123` prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStreamA
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// The ID of the watermark template.
	//
	// >  You can obtain the template ID by checking the value of the TemplateId parameter that is returned by the [AddLiveStreamWatermark](https://help.aliyun.com/document_detail/410759.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddLiveStreamWatermarkRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamWatermarkRuleRequest) GoString() string {
	return s.String()
}

func (s *AddLiveStreamWatermarkRuleRequest) GetApp() *string {
	return s.App
}

func (s *AddLiveStreamWatermarkRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *AddLiveStreamWatermarkRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddLiveStreamWatermarkRuleRequest) GetName() *string {
	return s.Name
}

func (s *AddLiveStreamWatermarkRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveStreamWatermarkRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveStreamWatermarkRuleRequest) GetStream() *string {
	return s.Stream
}

func (s *AddLiveStreamWatermarkRuleRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddLiveStreamWatermarkRuleRequest) SetApp(v string) *AddLiveStreamWatermarkRuleRequest {
	s.App = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleRequest) SetDescription(v string) *AddLiveStreamWatermarkRuleRequest {
	s.Description = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleRequest) SetDomain(v string) *AddLiveStreamWatermarkRuleRequest {
	s.Domain = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleRequest) SetName(v string) *AddLiveStreamWatermarkRuleRequest {
	s.Name = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleRequest) SetOwnerId(v int64) *AddLiveStreamWatermarkRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleRequest) SetRegionId(v string) *AddLiveStreamWatermarkRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleRequest) SetStream(v string) *AddLiveStreamWatermarkRuleRequest {
	s.Stream = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleRequest) SetTemplateId(v string) *AddLiveStreamWatermarkRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleRequest) Validate() error {
	return dara.Validate(s)
}
