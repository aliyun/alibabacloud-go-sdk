// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppRecordTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeAppRecordTemplatesResponseBodyTemplates) *DescribeAppRecordTemplatesResponseBody
	GetTemplates() []*DescribeAppRecordTemplatesResponseBodyTemplates
	SetTotalNum(v int64) *DescribeAppRecordTemplatesResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeAppRecordTemplatesResponseBody
	GetTotalPage() *int64
}

type DescribeAppRecordTemplatesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2A7497D0-BEAE-58E7-B13A-751BD8EAE4C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of templates.
	Templates []*DescribeAppRecordTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppRecordTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppRecordTemplatesResponseBody) GetTemplates() []*DescribeAppRecordTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeAppRecordTemplatesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeAppRecordTemplatesResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAppRecordTemplatesResponseBody) SetRequestId(v string) *DescribeAppRecordTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBody) SetTemplates(v []*DescribeAppRecordTemplatesResponseBodyTemplates) *DescribeAppRecordTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBody) SetTotalNum(v int64) *DescribeAppRecordTemplatesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBody) SetTotalPage(v int64) *DescribeAppRecordTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppRecordTemplatesResponseBodyTemplates struct {
	// The creation time. The time is displayed in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC.
	//
	// example:
	//
	// 2020-09-04T06:22:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The delay time before stopping the recording. Unit: seconds.
	//
	// example:
	//
	// 180
	DelayStopTime *int64 `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	// The naming rule for recording files.
	//
	// > To ensure the uniqueness of recording file names, the current naming rule is record/{AppId}/{ChannelId}_{TaskId}/{EscapedStartTime}_{EscapedEndTime}.
	//
	// example:
	//
	// record/{AppId}/{ChannelId_TaskId}/{EscapedStartTime}_{EscapedEndTime}
	FilePrefix *string `json:"FilePrefix,omitempty" xml:"FilePrefix,omitempty"`
	// The splitting duration of the recording file. The value must be greater than or equal to 1800. Unit: seconds.
	//
	// example:
	//
	// 1800
	FileSplitInterval *int64 `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	// The formats of the recording file. Valid values: .m3u8, .mp4, and .flv.
	//
	// > Only lowercase letters are supported for file formats.
	Formats []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The list of layout IDs.
	LayoutIds []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// The encoding options.
	//
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// The template name.
	//
	// example:
	//
	// Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// wv7N****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeAppRecordTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetDelayStopTime() *int64 {
	return s.DelayStopTime
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetFilePrefix() *string {
	return s.FilePrefix
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetFileSplitInterval() *int64 {
	return s.FileSplitInterval
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetFormats() []*string {
	return s.Formats
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetCreateTime(v string) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetDelayStopTime(v int64) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.DelayStopTime = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetFilePrefix(v string) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.FilePrefix = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetFileSplitInterval(v int64) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.FileSplitInterval = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetFormats(v []*string) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.Formats = v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetLayoutIds(v []*string) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.LayoutIds = v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetMediaEncode(v int32) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.MediaEncode = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetName(v string) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) SetTemplateId(v string) *DescribeAppRecordTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
