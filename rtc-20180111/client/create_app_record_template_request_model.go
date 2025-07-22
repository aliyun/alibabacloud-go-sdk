// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppRecordTemplateRequest
	GetAppId() *string
	SetClientToken(v string) *CreateAppRecordTemplateRequest
	GetClientToken() *string
	SetRecordTemplate(v *CreateAppRecordTemplateRequestRecordTemplate) *CreateAppRecordTemplateRequest
	GetRecordTemplate() *CreateAppRecordTemplateRequestRecordTemplate
}

type CreateAppRecordTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	RecordTemplate *CreateAppRecordTemplateRequestRecordTemplate `json:"RecordTemplate,omitempty" xml:"RecordTemplate,omitempty" type:"Struct"`
}

func (s CreateAppRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRecordTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppRecordTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppRecordTemplateRequest) GetRecordTemplate() *CreateAppRecordTemplateRequestRecordTemplate {
	return s.RecordTemplate
}

func (s *CreateAppRecordTemplateRequest) SetAppId(v string) *CreateAppRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppRecordTemplateRequest) SetClientToken(v string) *CreateAppRecordTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppRecordTemplateRequest) SetRecordTemplate(v *CreateAppRecordTemplateRequestRecordTemplate) *CreateAppRecordTemplateRequest {
	s.RecordTemplate = v
	return s
}

func (s *CreateAppRecordTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAppRecordTemplateRequestRecordTemplate struct {
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
	Formats   []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
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
}

func (s CreateAppRecordTemplateRequestRecordTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRecordTemplateRequestRecordTemplate) GoString() string {
	return s.String()
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) GetDelayStopTime() *int32 {
	return s.DelayStopTime
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) GetFilePrefix() *string {
	return s.FilePrefix
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) GetFileSplitInterval() *int32 {
	return s.FileSplitInterval
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) GetFormats() []*string {
	return s.Formats
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) GetName() *string {
	return s.Name
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) SetDelayStopTime(v int32) *CreateAppRecordTemplateRequestRecordTemplate {
	s.DelayStopTime = &v
	return s
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) SetFilePrefix(v string) *CreateAppRecordTemplateRequestRecordTemplate {
	s.FilePrefix = &v
	return s
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) SetFileSplitInterval(v int32) *CreateAppRecordTemplateRequestRecordTemplate {
	s.FileSplitInterval = &v
	return s
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) SetFormats(v []*string) *CreateAppRecordTemplateRequestRecordTemplate {
	s.Formats = v
	return s
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) SetLayoutIds(v []*string) *CreateAppRecordTemplateRequestRecordTemplate {
	s.LayoutIds = v
	return s
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) SetMediaEncode(v int32) *CreateAppRecordTemplateRequestRecordTemplate {
	s.MediaEncode = &v
	return s
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) SetName(v string) *CreateAppRecordTemplateRequestRecordTemplate {
	s.Name = &v
	return s
}

func (s *CreateAppRecordTemplateRequestRecordTemplate) Validate() error {
	return dara.Validate(s)
}
