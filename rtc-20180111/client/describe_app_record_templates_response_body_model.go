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
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeAppRecordTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	return dara.Validate(s)
}

type DescribeAppRecordTemplatesResponseBodyTemplates struct {
	// example:
	//
	// 2020-09-04T06:22:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 180
	DelayStopTime *int64 `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	// example:
	//
	// record/{AppId}/{ChannelId_TaskId}/{EscapedStartTime}_{EscapedEndTime}
	FilePrefix *string `json:"FilePrefix,omitempty" xml:"FilePrefix,omitempty"`
	// example:
	//
	// 1800
	FileSplitInterval *int64    `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	Formats           []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	LayoutIds         []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
