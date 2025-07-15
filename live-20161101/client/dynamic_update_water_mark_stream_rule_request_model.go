// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDynamicUpdateWaterMarkStreamRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DynamicUpdateWaterMarkStreamRuleRequest
	GetApp() *string
	SetDomain(v string) *DynamicUpdateWaterMarkStreamRuleRequest
	GetDomain() *string
	SetOwnerId(v int64) *DynamicUpdateWaterMarkStreamRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DynamicUpdateWaterMarkStreamRuleRequest
	GetRegionId() *string
	SetStream(v string) *DynamicUpdateWaterMarkStreamRuleRequest
	GetStream() *string
	SetTemplateId(v string) *DynamicUpdateWaterMarkStreamRuleRequest
	GetTemplateId() *string
}

type DynamicUpdateWaterMarkStreamRuleRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
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
	// pull.aliyundoc.com
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stream that contains the watermark.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// The watermark template ID. You can call the [DescribeLiveStreamWatermarks](https://help.aliyun.com/document_detail/2848102.html) operation to obtain available watermark template IDs.
	//
	// >  The TemplateId parameter is used to replace the watermark template ID during live streaming.
	//
	// This parameter is required.
	//
	// example:
	//
	// 749b7594-86d6-37b1-513b-e1e19845****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DynamicUpdateWaterMarkStreamRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DynamicUpdateWaterMarkStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) GetApp() *string {
	return s.App
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) GetStream() *string {
	return s.Stream
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) SetApp(v string) *DynamicUpdateWaterMarkStreamRuleRequest {
	s.App = &v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) SetDomain(v string) *DynamicUpdateWaterMarkStreamRuleRequest {
	s.Domain = &v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) SetOwnerId(v int64) *DynamicUpdateWaterMarkStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) SetRegionId(v string) *DynamicUpdateWaterMarkStreamRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) SetStream(v string) *DynamicUpdateWaterMarkStreamRuleRequest {
	s.Stream = &v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) SetTemplateId(v string) *DynamicUpdateWaterMarkStreamRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleRequest) Validate() error {
	return dara.Validate(s)
}
