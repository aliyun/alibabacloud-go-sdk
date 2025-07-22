// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppRecordTemplateRequest
	GetAppId() *string
	SetClientToken(v string) *ModifyAppRecordTemplateRequest
	GetClientToken() *string
	SetRecordTemplate(v *ModifyAppRecordTemplateRequestRecordTemplate) *ModifyAppRecordTemplateRequest
	GetRecordTemplate() *ModifyAppRecordTemplateRequestRecordTemplate
}

type ModifyAppRecordTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	RecordTemplate *ModifyAppRecordTemplateRequestRecordTemplate `json:"RecordTemplate,omitempty" xml:"RecordTemplate,omitempty" type:"Struct"`
}

func (s ModifyAppRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRecordTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppRecordTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppRecordTemplateRequest) GetRecordTemplate() *ModifyAppRecordTemplateRequestRecordTemplate {
	return s.RecordTemplate
}

func (s *ModifyAppRecordTemplateRequest) SetAppId(v string) *ModifyAppRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRecordTemplateRequest) SetClientToken(v string) *ModifyAppRecordTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppRecordTemplateRequest) SetRecordTemplate(v *ModifyAppRecordTemplateRequestRecordTemplate) *ModifyAppRecordTemplateRequest {
	s.RecordTemplate = v
	return s
}

func (s *ModifyAppRecordTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyAppRecordTemplateRequestRecordTemplate struct {
	// example:
	//
	// 180
	DelayStopTime *int32 `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// record/{AppId}/{ChannelId}_{TaskId}/{EscapedStartTime}_{EscapedEndTime}
	FilePrefix *string `json:"FilePrefix,omitempty" xml:"FilePrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1800
	FileSplitInterval *int32 `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	// This parameter is required.
	Formats []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// This parameter is required.
	LayoutIds []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 模版
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2xh6****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyAppRecordTemplateRequestRecordTemplate) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRecordTemplateRequestRecordTemplate) GoString() string {
	return s.String()
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) GetDelayStopTime() *int32 {
	return s.DelayStopTime
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) GetFilePrefix() *string {
	return s.FilePrefix
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) GetFileSplitInterval() *int32 {
	return s.FileSplitInterval
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) GetFormats() []*string {
	return s.Formats
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) GetName() *string {
	return s.Name
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) SetDelayStopTime(v int32) *ModifyAppRecordTemplateRequestRecordTemplate {
	s.DelayStopTime = &v
	return s
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) SetFilePrefix(v string) *ModifyAppRecordTemplateRequestRecordTemplate {
	s.FilePrefix = &v
	return s
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) SetFileSplitInterval(v int32) *ModifyAppRecordTemplateRequestRecordTemplate {
	s.FileSplitInterval = &v
	return s
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) SetFormats(v []*string) *ModifyAppRecordTemplateRequestRecordTemplate {
	s.Formats = v
	return s
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) SetLayoutIds(v []*string) *ModifyAppRecordTemplateRequestRecordTemplate {
	s.LayoutIds = v
	return s
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) SetMediaEncode(v int32) *ModifyAppRecordTemplateRequestRecordTemplate {
	s.MediaEncode = &v
	return s
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) SetName(v string) *ModifyAppRecordTemplateRequestRecordTemplate {
	s.Name = &v
	return s
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) SetTemplateId(v string) *ModifyAppRecordTemplateRequestRecordTemplate {
	s.TemplateId = &v
	return s
}

func (s *ModifyAppRecordTemplateRequestRecordTemplate) Validate() error {
	return dara.Validate(s)
}
