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
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The client token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The recording template.
	//
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
	if s.RecordTemplate != nil {
		if err := s.RecordTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppRecordTemplateRequestRecordTemplate struct {
	// The delay time before stopping the recording. Unit: seconds.
	//
	// example:
	//
	// 180
	DelayStopTime *int32 `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	// The naming rule for recorded files.
	//
	// > To ensure the uniqueness of recorded file names, the current naming rule for recorded files is record/{AppId}/{ChannelId}_{TaskId}/{EscapedStartTime}_{EscapedEndTime}.
	//
	// This parameter is required.
	//
	// example:
	//
	// record/{AppId}/{ChannelId}_{TaskId}/{EscapedStartTime}_{EscapedEndTime}
	FilePrefix *string `json:"FilePrefix,omitempty" xml:"FilePrefix,omitempty"`
	// The splitting duration of the recorded file. The value must be greater than or equal to 1800. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1800
	FileSplitInterval *int32 `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	// The format of the recorded file. Only .m3u8, .mp4, and .flv are supported.
	//
	// This parameter is required.
	Formats []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The list of layout IDs. A maximum of 16 layout IDs can be specified.
	//
	// Note: The layouts in the list must correspond to different numbers of data sources.
	LayoutIds []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// The encoding options.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Template
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
